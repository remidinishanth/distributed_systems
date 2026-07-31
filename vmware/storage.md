---
layout: page
title: "Storage"
category: "vmware"
---

vSphere presents guest **virtual disks (VMDKs)** on top of a **datastore**. The
datastore type decides the transport and the backing store — and, crucially,
what data-integrity and redundancy guarantees exist *below* the guest OS.

![Datastore types: VMFS / NFS / vSAN / vVol](datastore_types.png)

## Datastore types

| Datastore type | Transport | Backing | Redundancy of its own | Bitrot detect + repair |
|---|---|---|---|---|
| **VMFS** | Direct-attached / FC / FCoE / iSCSI | Disk or LUN | None — inherits from the array/RAID below (or nothing on local disk) | ❌ No (depends entirely on the array) |
| **NFS** | Ethernet | Filer file system | Whatever the NAS provides | ❌ Not from VMware — depends on filer (NetApp/ZFS have it, others may not) |
| **vSAN** | Direct-attached (pooled local disks) | vSAN cluster | Built-in, policy-driven **FTT** (mirror / erasure across hosts) | ✅ **Yes — end-to-end checksum on by default (since vSAN 6.2) + scrubber** |
| **vSphere Virtual Volume (vVol)** | FC / Ethernet | Storage container (array-managed) | Array-defined via VASA policy | Depends on array |

### Key distinction: VMFS vs vSAN

* **VMFS is a clustered *filesystem*, not a *protector*.** It stores VMDK blocks
  on one LUN/disk and verifies nothing over time. Any integrity/redundancy comes
  from the layer below (enterprise array scrubbing / T10 DIF-PI), which varies by
  vendor and is absent on local disk.
* **vSAN is a software-defined storage layer that owns integrity + redundancy.**
  Object checksum (CRC-32C) is enabled by default and a background scrubber
  repairs mismatches from the FTT copy. FTT=1 (mirror) costs 2x raw, FTT=2 costs
  3x raw.
* You cannot assume the hypervisor protects against bitrot — it only does on
  **vSAN**, which is a minority-to-substantial slice of deployments, not a given.

## Bitrot: detect vs. repair

* **Detection** = noticing a block is corrupt (checksums). **Repair/heal** =
  reconstructing the good data. Repair needs redundancy (a parity/replica copy);
  detection alone just flags the object as unreadable.
* On **plain VMFS / local disk**: neither, unless the array provides it.
* On **vSAN**: both (checksum + FTT replica).
* MinIO/AIStor provides its **own** end-to-end checksums, so it *detects* bitrot
  regardless of datastore; whether it can *heal* depends on its erasure-coding
  parity, not on the datastore.

## Thin provisioning gotcha

Datastore file size for a thin VMDK = **actual consumed**, not provisioned. A
257 GB thin disk with 60 GB used shows as a ~62 GB `.vmdk` file. When matching
VMDKs (datastore Files view) to guest disks (`lsblk`), compare *actual* usage,
not provisioned size, or they will look mismatched.

## MinIO / AIStor on vSphere — layering notes

* MinIO wants **JBOD + XFS**, no RAID/LVM/ZFS, because it does its own erasure
  coding and per-drive bitrot healing. LVM/RAID under it reduces performance,
  reliability, and predictability (per MinIO's storage requirements).
* **Multi-drive still protects even on a shared datastore.** Bit rot, single-VMDK
  loss, and per-mount filesystem corruption live *above* the physical disk, so
  erasure coding across several VMDKs heals them even when you cannot guarantee
  independent physical media. Single-drive detects but cannot heal — one bad block
  is a lost object, recoverable only from the VM backup.
* **Expansion without LVM or a new pool:** grow the VMDK, rescan the device in the
  guest, grow the partition, then `xfs_growfs` the mount — all online. Grow every
  drive by the same amount so the erasure set (sized to the smallest drive) gains
  usable capacity. You cannot add drives to an existing node's pool; a second pool
  on the same VM adds no fault isolation and must match the first pool's EC.
* **EC:M efficiency** = `(N - M) / N` usable. Parity is fixed per erasure set, so
  efficiency *improves* with more drives: EC:2 → 50% at 4 drives, 75% at 8, 83%
  at 12. If efficiency is the concern, the answer is more drives, not fewer.
* **The safe way to use LVM (if you must):** one physical disk → one LV → one MinIO
  drive (1:1). Never concatenate multiple VMDKs into one LV — losing any one VMDK
  then loses the whole volume, and MinIO sees a single drive so nothing heals.

### Observed in the Rubrik RSC-P appliance (`rscp-3.6.6-cdm-9.6-dev-63`)

MinIO runs on **3 dedicated VMDKs** (`sdd`/`sde`/`sdf`, 30 GB each), each its own
disk → own LV → **LUKS-encrypted** (`minio-minio1/2/3`). LVM is used *per disk*
for encryption, not for pooling (the safe 1:1 pattern above). Three drives implies
**EC:1** parity. This is visible from both inside the guest (`lsblk`) and from the
datastore Files view (three near-identical ~12 GiB VMDKs).
