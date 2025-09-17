Ways of Encryption

<img width="1347" height="656" alt="image" src="https://github.com/user-attachments/assets/1d7f8ec2-04fb-43cf-9d03-54d1e156144a" />

## Device Mapper

The device mapper is a framework provided by the Linux kernel for mapping physical block devices onto higher-level virtual block devices. 

It forms the foundation of the logical volume manager (LVM), software RAIDs and dm-crypt disk encryption, and offers additional features such as file system snapshots.

* The Device Mapper is a kernel driver that provides a framework for volume management.
* It provides a generic way of creating mapped devices, which may be used as logical volumes. It does not specifically know about volume groups or metadata formats.

<img width="640" height="494" alt="image" src="https://github.com/user-attachments/assets/9fd64a8a-ba21-46e7-b05c-2078dfad983f" />

https://gitlab.com/cryptsetup/cryptsetup

LUKS (Linux Unified Key Setup) is the standard on-disk format for disk encryption on Linux systems, based on the dm-crypt kernel module.

<img width="600" height="1800" alt="image" src="https://github.com/user-attachments/assets/bdd1e970-9a36-4f28-821d-f752ea84fa67" />

<img width="1936" height="1080" alt="image" src="https://github.com/user-attachments/assets/21250f3b-69f2-4b58-92ee-91f903698d19" />
