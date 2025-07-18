---
layout: page
title: "Devices and Disk Layout"
category: "linux"
description: "Linux devices, disk layout, and filesystem configuration"
---

# /etc/fstab
# Created by anaconda on Mon Jan 20 15:59:48 2025
#
# Accessible filesystems, by reference, are maintained under '/dev/disk'
# See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info
#
/dev/mapper/root-os     /                       xfs     defaults        0 0
UUID=c2707e9e-2e27-43a5-91d1-0d25e1f08dd2 /boot                   ext4    defaults        1 2
/dev/mapper/root-home   /home                   xfs     defaults,nosuid        0 0
tmpfs /dev/shm tmpfs defaults,relatime,nodev,noexec,nosuid 0 0
/dev/mapper/opt-minio /opt/minio xfs defaults 0 2
/dev/mapper/opt-monitoring /opt/monitoring xfs defaults 0 2
/dev/mapper/opt-mq /opt/mq xfs defaults 0 2
/dev/mapper/opt-mysql /opt/mysql xfs defaults 0 2
/dev/mapper/opt-polaris /opt/polaris xfs defaults 0 2
/dev/mapper/opt-staging /opt/staging xfs defaults 0 2
```

Explanation 

<img width="844" alt="image" src="https://github.com/user-attachments/assets/69a6676c-fce7-4d1d-a716-da49ec5f5d25" />

```
# c entry details:
# /dev/mapper/root-os     /                       xfs     defaults        0 1
# * /dev/mapper/root-os <--- defines the storage device (i.e. /dev/sda2)
# * /                   <--- tells the mount command where it should mount the
#                            <file system> to.
# * xfs                 <--- defines the file system type of the device or
#                            partition to be mounted
# * defaults            <--- define particular options for filesystems. Some
#                            options relate only to the filesystem itself. Some
#                            of the more common options are:
#                            auto  - file system will mount automatically at boot
#                            exec  - allow the execution binaries that are on
#                                    that partition (default).
#                            ro    - mount the filesystem read only
#                            rw    - mount the filesystem read-write
#                            sync  - I/O should be done synchronously
#                            async - I/O should be done asynchronously
#                            nouser - only allow root to mount the filesystem
#                            defaults - default mount settings (equivalent to rw,
#                                       suid,dev,exec,auto,nouser,async).
#                            suid - allow the operation of suid, and sgid bits.
# * 0                   <--- is used by the dump utility to decide when to make
#                            a backup. When installed, dump checks the entry and
#                            uses the number to decide if a file system should
#                            be backed up. Possible entries are 0 and 1. If 0,
#                            dump will ignore the file system, if 1, dump will
#                            make a backup. Most users will not have dump
#                            installed, so they should put 0 for the <dump>
#                            entry
# * 1                       <--- fsck reads the <pass> number and determines in
#                            which order the file systems should be checked.
#                            Possible entries are 0, 1, and 2. The root file
#                            system should have the highest priority, 1, all
#                            other file systems you want to have checked should
#                            get a 2. File systems with a <pass> value 0 will
#                            not be checked by the fsck utility.
#
```

### Useful commands
<img width="775" alt="image" src="https://github.com/user-attachments/assets/a6343cab-f243-4109-ad0e-db2aff00deac" />"
category: "devices_and_disk_layout"
---

![image](https://github.com/user-attachments/assets/7df9cb19-2dcf-40c8-8edf-3b69fbd2fcc7)
Ref: The UNIX Time-Sharing System, Dennis M. Ritchie and Ken Thompson, Bell Laboratories


![image](https://github.com/user-attachments/assets/157b5629-38b1-415d-a8d9-a499aa379b16)

<img width="1638" alt="image" src="https://github.com/user-attachments/assets/338eac5f-e1b6-431f-a922-d1ff0a116e7d" />

<img width="1464" alt="image" src="https://github.com/user-attachments/assets/fb869e98-485d-4263-ac54-1d6c92607ceb" />

<img width="710" alt="image" src="https://github.com/user-attachments/assets/2872f4e1-7119-4cbb-8073-36972d632ada" />

Ref: https://tldp.org/LDP/Linux-Filesystem-Hierarchy/html/dev.html

For example:

The null device in Unix systems is /dev/null.
* Its purpose is to immediately discard anything sent to it.
* It’s also known as a bucket or a blackhole, like throwing something in a trash bucket or sending it to a blackhole never to be seen again.

![image](https://github.com/user-attachments/assets/94525d09-beda-4456-b43c-1018dbfba63d)

If we wanted to send a specific file descriptor output, we would use `1>` for stdout and we would use `2>` for stderr.

A `&>` sends both stdout and stderr file descriptors to `/dev/null`.

![image](https://github.com/user-attachments/assets/e7878018-f8c2-447c-90ec-ee052c535968)

### Devices

Linux splits all devices into three classes: block devices, character devices, and network devices.

<img width="780" alt="image" src="https://github.com/user-attachments/assets/c542150f-eb35-436c-8a2b-e9f549d1b921" />

A character-stream device transfers bytes one by one, whereas a block device transfers a block of bytes as a unit.


#### Block devices
* Block devices include all devices that allow random access to completely
independent, fixed-sized blocks of data, including hard disks and floppy disks,
CD-ROMs and Blu-ray discs, and flash memory. 
* Block devices are typically used to store file systems, but direct access to a block device is also allowed
so that programs can create and repair the file system that the device contains.
* Applications can also access these block devices directly if they wish.
For example, a database application may prefer to perform its own fine-tuned
layout of data onto a disk rather than using the general-purpose file system.

#### Character devices
Character devices include most other devices, such as mice and keyboards.

The fundamental difference between block and character devices is random
access—block devices are accessed randomly, while character devices are
accessed serially. For example, seeking to a certain position in a file might
be supported for a DVD but makes no sense for a pointing device such as a
mouse.

#### Network devices
Network devices are dealt with differently from block and character
devices. Users cannot directly transfer data to network devices. Instead,
they must communicate indirectly by opening a connection to the kernel’s
networking subsystem.

![image](https://github.com/user-attachments/assets/d2e80327-aaaa-4603-a0b7-b4ad8fa5ed01)

![image](https://github.com/user-attachments/assets/44ff55a4-0956-49e0-8f78-0a9befeb5214)

https://dev.to/shankarsurya035/how-to-create-lvm-partition-in-linux-dgo


> Special files for char drivers are identified by a “c” in the first column of the output of ls -l. Block devices appear in /dev as well, but they are identified by a “b.” The focus of this chapter is on char devices, but much of the following information applies to block devices as well.

## Logical Volume Management

LVM creates a logical volume that aggregates multiple disks to stripe reads and writes. 
It also serves as a scalable volume for maintenance.

* Logical Volume Manager (LVM) plays an important role in the Linux operating system by improving the availability, disk I/O, performance and capability of disk management.

* LVM is a widely used technique that is extremely flexible for disk management.

* This adds an extra layer between the physical disks and the file system, allowing you to create a logical volume instead of a physical disk.

* LVM allows you to easily resize, extend and decrease the logical volume when you need it.


![image](https://github.com/user-attachments/assets/8bc1e9ec-824b-4466-a098-de9888f56736)


<img width="594" alt="image" src="https://github.com/user-attachments/assets/42721ba9-0ca8-4267-a769-882bcc628f3a" />


![image](https://github.com/user-attachments/assets/2b17c6b2-287f-4106-b4ce-25aa65743f33)


```
[rksupport@VR-POLARIS-VW-D27E2E4 ~]$ lsblk
NAME              MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                 8:0    0   64G  0 disk
├─sda1              8:1    0    1G  0 part /boot
└─sda2              8:2    0   63G  0 part
  ├─root-os       253:0    0   24G  0 lvm  /
  ├─root-home     253:4    0   15G  0 lvm  /home
  └─root-reserved 253:5    0   24G  0 lvm
sdb                 8:16   0  100G  0 disk
└─opt-minio       253:1    0  102G  0 lvm  /opt/minio
sdc                 8:32   0  700G  0 disk
├─opt-polaris     253:2    0  501G  0 lvm  /opt/polaris
├─opt-monitoring  253:3    0  101G  0 lvm  /opt/monitoring
├─opt-staging     253:6    0   50G  0 lvm  /opt/staging
└─opt-reserved    253:7    0   49G  0 lvm
sdd                 8:48   0  450G  0 disk
├─opt-mysql       253:8    0  400G  0 lvm  /opt/mysql
└─opt-mq          253:9    0   50G  0 lvm  /opt/mq
sde                 8:64   0    1G  0 disk
└─opt-minio       253:1    0  102G  0 lvm  /opt/minio
sdf                 8:80   0    1G  0 disk
└─opt-minio       253:1    0  102G  0 lvm  /opt/minio
sdg                 8:96   0    1G  0 disk
└─opt-polaris     253:2    0  501G  0 lvm  /opt/polaris
sdh                 8:112  0    1G  0 disk
sdi                 8:128  0    1G  0 disk
```

This is basically from

```
Physical Device Layout (from lsblk)
================================================================================

sda (64G) ──┬─ sda1 (1G, partition) ─────────────────────────────────→ /boot
            │
            └─ sda2 (63G, partition) 
                  │
                  └─ [LVM PV: root-vg] ─┬─ root-os       (24G, LV)   → /
                                         ├─ root-home    (15G, LV)   → /home
                                         └─ root-reserved(24G, LV)   → (not mounted)

sdb (100G) ──┬─ sdb1 (entire disk as LVM PV)
               │
               └─ [LVM PV: opt-vg] ─────── opt-minio     (102G, LV)  → /opt/minio

sdc (700G) ──┬─ sdc1 (entire disk as LVM PV)
               │
               └─ [LVM PV: opt-vg] ─┬─ opt-polaris     (501G, LV)   → /opt/polaris
                                     ├─ opt-monitoring (101G, LV)   → /opt/monitoring
                                     ├─ opt-staging    (50G, LV)    → /opt/staging
                                     └─ opt-reserved   (49G, LV)    → (not mounted)
```

```
+------------------------+
|  Physical Disk (sda)   |
+------------------------+
           |
   +----------------+
   |   Partition    |  (e.g., sda2, marked for LVM)
   +----------------+
           |
   +-------------------------+                (Optional: several disks/partitions below)
   |   LVM Physical Volume   |
   +-------------------------+
           |
   +-------------------+
   |   Volume Group    |      (VG can use many PVs, combine them!)
   +-------------------+
           |
   +----------+---------+---------+---------+
   | Logical  | Logical | Logical | Logical |
   | Volume   | Volume  | Volume  | Volume  |
   +----------+---------+---------+---------+
     /         /home      /var
  (Mountpoints; formatted with fs)
```

<img width="435" alt="image" src="https://github.com/user-attachments/assets/62af8ca9-dc4b-412c-9385-2bff619cbbbf" />


![image](https://github.com/user-attachments/assets/acbf4615-0539-4e02-98cc-d8fe2619b760)

![image](https://github.com/user-attachments/assets/176b6239-2529-48ed-af54-5e96325f086f)
Using `pvs`, `vgs` and `lvs`, we can inspect them

![image](https://github.com/user-attachments/assets/05d28321-949f-45ce-9a4d-cebb82805555)

Use `lsblk` to get a broad view — “what storage is connected, what's a partition, what's an LVM logical volume?”

Use `pvs` to focus on the building blocks of LVM setups — “which disks/partitions are being managed by LVM and how?”

`lsblk`:
What is on this system? (all block devices and layout)

`pvs`:
Which devices are used as LVM physical volumes? (LVM only)

`vgs` — shows info per-volume group in LVM
`lvs` — shows info per-logical volume in LVM


`pvs` along with VG

```
[rksupport@VR-POLARIS-VW-D27E2E4 ~]$ sudo pvs -a -o +vg_name,lv_name
  PV         VG   Fmt  Attr PSize    PFree VG   LV
  /dev/sda1            ---        0     0
  /dev/sda2  root lvm2 a--   <63.00g    0  root home
  /dev/sda2  root lvm2 a--   <63.00g    0  root os
  /dev/sda2  root lvm2 a--   <63.00g    0  root reserved
  /dev/sdb   opt  lvm2 a--  <100.00g    0  opt  minio
  /dev/sdc   opt  lvm2 a--  <700.00g    0  opt  polaris
  /dev/sdc   opt  lvm2 a--  <700.00g    0  opt  monitoring
  /dev/sdc   opt  lvm2 a--  <700.00g    0  opt  staging
  /dev/sdc   opt  lvm2 a--  <700.00g    0  opt  reserved
  /dev/sdc   opt  lvm2 a--  <700.00g    0  opt  monitoring
  /dev/sdd   opt  lvm2 a--  <450.00g    0  opt  mysql
  /dev/sdd   opt  lvm2 a--  <450.00g    0  opt  mq
  /dev/sde   opt  lvm2 a--  1020.00m    0  opt  minio
  /dev/sdf   opt  lvm2 a--  1020.00m    0  opt  minio
  /dev/sdg   opt  lvm2 a--  1020.00m    0  opt  polaris
  /dev/sdh             ---        0     0
  /dev/sdi             ---        0     0
```


The `blkid` command in Linux is a tool to locate and display block device attributes, such as filesystem type, UUID, and label.

```
[rksupport@VR-POLARIS-VW-D27E2E4 ~]$ sudo blkid
/dev/sda1: UUID="c2707e9e-2e27-43a5-91d1-0d25e1f08dd2" TYPE="ext4"
/dev/sda2: UUID="ZeML8J-Nae7-pMCZ-PBis-PHel-RdXO-6iEK8a" TYPE="LVM2_member"
/dev/sde: UUID="tN9ohY-yhKb-Lpwh-oJbR-5R1E-z1YO-5fOHwC" TYPE="LVM2_member"
/dev/sdb: UUID="ZUPGiI-cqEs-0nHi-Zr7S-2bZl-V1Yk-sHI2sz" TYPE="LVM2_member"
/dev/sdc: UUID="G6BrpZ-m9C6-mECZ-RDTL-fHW9-d0rv-Ggvt0v" TYPE="LVM2_member"
/dev/mapper/root-os: UUID="36686917-4f3d-4395-8b29-2ab4738455c5" TYPE="xfs"
/dev/sdd: UUID="yX2shS-qSly-x6ic-RXel-MeDl-HncV-6Z12mz" TYPE="LVM2_member"
/dev/mapper/opt-minio: UUID="41906eed-ef27-495e-9034-e5d547eff490" TYPE="xfs"
/dev/mapper/opt-polaris: UUID="9aa1d5ca-2ee5-4295-b9a8-16b657aa975d" TYPE="xfs"
/dev/mapper/root-home: UUID="0232af07-1aef-41d9-bb3a-f0a5c843b7b7" TYPE="xfs"
/dev/mapper/opt-monitoring: UUID="6079ccd3-6c43-4b7d-976f-45a5a71fe49a" TYPE="xfs"
/dev/mapper/opt-staging: UUID="7a24813a-a10c-4081-ba10-18fb1747a39f" TYPE="xfs"
/dev/mapper/opt-mysql: UUID="02195633-663c-4f2a-b10b-f782795f7849" TYPE="xfs"
/dev/mapper/opt-mq: UUID="176fb9c8-ca43-4590-a99b-50728cd6e91c" TYPE="xfs"
/dev/sdf: UUID="owxQp8-q80i-1TTR-98M3-wd7n-MkBS-vx5lvI" TYPE="LVM2_member"
/dev/sdg: UUID="plOHfP-3YGd-6vac-DKC6-rWUS-X103-ZexUPJ" TYPE="LVM2_member"
```

<img width="1168" alt="image" src="https://github.com/user-attachments/assets/9f7df455-567c-4044-abc4-75e7f5afe32b" />

`sudo fdisk -l` lists all the partitions and disks connected to your Linux machine.


Your Linux system's filesystem table, aka `fstab`, is a configuration table designed to ease the burden of mounting and unmounting file systems to a machine. It is a set of rules used to control how different filesystems are treated each time they are introduced to a system.

The output of fstab

```
[rksupport@VR-POLARIS-VW-D27E2E4 ~]$ cat /etc/fstab

#
# /etc/fstab
# Created by anaconda on Mon Jan 20 15:59:48 2025
#
# Accessible filesystems, by reference, are maintained under '/dev/disk'
# See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info
#
/dev/mapper/root-os     /                       xfs     defaults        0 0
UUID=c2707e9e-2e27-43a5-91d1-0d25e1f08dd2 /boot                   ext4    defaults        1 2
/dev/mapper/root-home   /home                   xfs     defaults,nosuid        0 0
tmpfs /dev/shm tmpfs defaults,relatime,nodev,noexec,nosuid 0 0
/dev/mapper/opt-minio /opt/minio xfs defaults 0 2
/dev/mapper/opt-monitoring /opt/monitoring xfs defaults 0 2
/dev/mapper/opt-mq /opt/mq xfs defaults 0 2
/dev/mapper/opt-mysql /opt/mysql xfs defaults 0 2
/dev/mapper/opt-polaris /opt/polaris xfs defaults 0 2
/dev/mapper/opt-staging /opt/staging xfs defaults 0 2
```

Explanation 

<img width="844" alt="image" src="https://github.com/user-attachments/assets/69a6676c-fce7-4d1d-a716-da49ec5f5d25" />

```
# c entry details:
# /dev/mapper/root-os     /                       xfs     defaults        0 1
# * /dev/mapper/root-os <--- defines the storage device (i.e. /dev/sda2)
# * /                   <--- tells the mount command where it should mount the
#                            <file system> to.
# * xfs                 <--- defines the file system type of the device or
#                            partition to be mounted
# * defaults            <--- define particular options for filesystems. Some
#                            options relate only to the filesystem itself. Some
#                            of the more common options are:
#                            auto  - file system will mount automatically at boot
#                            exec  - allow the execution binaries that are on
#                                    that partition (default).
#                            ro    - mount the filesystem read only
#                            rw    - mount the filesystem read-write
#                            sync  - I/O should be done synchronously
#                            async - I/O should be done asynchronously
#                            nouser - only allow root to mount the filesystem
#                            defaults - default mount settings (equivalent to rw,
#                                       suid,dev,exec,auto,nouser,async).
#                            suid - allow the operation of suid, and sgid bits.
# * 0                   <--- is used by the dump utility to decide when to make
#                            a backup. When installed, dump checks the entry and
#                            uses the number to decide if a file system should
#                            be backed up. Possible entries are 0 and 1. If 0,
#                            dump will ignore the file system, if 1, dump will
#                            make a backup. Most users will not have dump
#                            installed, so they should put 0 for the <dump>
#                            entry
# * 1                       <--- fsck reads the <pass> number and determines in
#                            which order the file systems should be checked.
#                            Possible entries are 0, 1, and 2. The root file
#                            system should have the highest priority, 1, all
#                            other file systems you want to have checked should
#                            get a 2. File systems with a <pass> value 0 will
#                            not be checked by the fsck utility.
#
```

### Useful commands
<img width="775" alt="image" src="https://github.com/user-attachments/assets/a6343cab-f243-4109-ad0e-db2aff00deac" />

