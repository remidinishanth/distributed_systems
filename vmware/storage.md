---
layout: page
title: "Storage"
category: "vmware"
---

vSphere presents guest **virtual disks (VMDKs)** on top of a **datastore**. The
datastore type decides the transport and the backing store — and what
data-integrity and redundancy guarantees exist *below* the guest OS.

![Datastore types: VMFS / NFS / vSAN / vVol](datastore_types.png)

## Datastore types

| Datastore type | Transport | Backing | Redundancy of its own | Bitrot detect + repair |
|---|---|---|---|---|
| **VMFS** | Direct-attached / FC / FCoE / iSCSI | Disk or LUN | None — inherits from the array/RAID below (or nothing on local disk) | No (depends entirely on the array) |
| **NFS** | Ethernet | Filer file system | Whatever the NAS provides | Not from VMware — depends on filer (NetApp/ZFS have it, others may not) |
| **vSAN** | Direct-attached (pooled local disks) | vSAN cluster | Built-in, policy-driven **FTT** (mirror / erasure across hosts) | Yes — end-to-end checksum on by default (since vSAN 6.2) + scrubber |
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

## Bitrot: detect vs. repair

* **Detection** = noticing a block is corrupt (checksums). **Repair/heal** =
  reconstructing the good data. Repair needs redundancy (a parity/replica copy);
  detection alone just flags the object as unreadable.
* On **plain VMFS / local disk**: neither, unless the array provides it.
* On **vSAN**: both (checksum + FTT replica).

## Thin provisioning gotcha

Datastore file size for a thin VMDK = **actual consumed**, not provisioned. A
257 GB thin disk with 60 GB used shows as a ~62 GB `.vmdk` file. When matching
VMDKs (datastore Files view) to guest disks (`lsblk`), compare *actual* usage,
not provisioned size, or they will look mismatched.
