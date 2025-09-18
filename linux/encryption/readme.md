Nice blog https://wiki.archlinux.org/title/Data-at-rest_encryption

## Ways of Encryption

* Application side, Read more at https://www.percona.com/blog/transparent-data-encryption-tde/
  - Server side encryption(MySQL stores the tables etc encrypted, but when we read etc we get unencrypted data) - the encryption process is transparent to the user, allowing them to access and manipulate the data as usual without worrying about the encryption and decryption process.
  - Client side encryption(Client writes encrypted data to server, say we encrypt and store that encrypted data in MySQL rows)
 
* Disk level encryption that encrypts all data stored on a disk or storage device
  - Several tools are available for implementing disk-level encryption, including BitLocker for Windows, dm-crypt for Linux, and FileVault for MacOS.

<img width="1347" height="656" alt="image" src="https://github.com/user-attachments/assets/1d7f8ec2-04fb-43cf-9d03-54d1e156144a" />



> All data-at-rest encryption methods operate in such a way that even though the disk actually holds encrypted data, the operating system and applications "see" it as the corresponding normal readable data as long as the cryptographic container (i.e. the logical part of the disk that holds the encrypted data) has been "unlocked" and mounted.
Ref: https://wiki.archlinux.org/title/Data-at-rest_encryption

For this to happen, some "secret information" (usually in the form of a keyfile and/or passphrase) needs to be supplied by the user, from which the actual encryption key can be derived (and stored in the kernel keyring for the duration of the session).

<img width="1862" height="1442" alt="image" src="https://github.com/user-attachments/assets/5a4c9a2c-83f8-4deb-a9f6-6e3356975445" />


> Device Mapper - Framework provided by the Linux Kernel, used to map physical block devices to higher level virtual block devices

> DM-Crypt - A target used with device mapper that provides transparent encryption. Allows us to create a virtual block device and have all data be encrypted on the fly before being committed to disk and can decrypt in the same way for reads.

> LUKS - Linux Unified Key Setup Provides an efficient user-friendly way to store and manage keys. Without LUKS, DM-Crypt can be more cumbersome and error-prone.

## Device Mapper

The device mapper is a framework provided by the Linux kernel for mapping physical block devices onto higher-level virtual block devices. 

It forms the foundation of the logical volume manager (LVM), software RAIDs and dm-crypt disk encryption, and offers additional features such as file system snapshots.

* The Device Mapper is a kernel driver that provides a framework for volume management.
* It provides a generic way of creating mapped devices, which may be used as logical volumes. It does not specifically know about volume groups or metadata formats.

<img width="640" height="494" alt="image" src="https://github.com/user-attachments/assets/9fd64a8a-ba21-46e7-b05c-2078dfad983f" />

## dm-crypt

`dm-crypt` is the Linux kernel's device mapper(dm) crypto target.

<img width="1936" height="1080" alt="image" src="https://github.com/user-attachments/assets/21250f3b-69f2-4b58-92ee-91f903698d19" />

https://gitlab.com/cryptsetup/cryptsetup

<img width="300" height="900" alt="image" src="https://github.com/user-attachments/assets/bdd1e970-9a36-4f28-821d-f752ea84fa67" />

<img width="1024" height="1024" alt="image" src="https://github.com/user-attachments/assets/ff1b7116-4185-45f1-96d4-a145c90abe77" />

Raw DM-Crypt requires manual key management. You have to handle key derivation, storage, and potential header formats yourself, which is prone to errors like weak key generation or incompatibility across systems.

## Linux Unified Key Setup (LUKS)

LUKS builds on DM-Crypt by standardizing the on-disk format for encrypted volumes, focusing on key management. It's essentially a header format that stores metadata (like encryption parameters and key slots) at the beginning of the block device.

LUKS (Linux Unified Key Setup) is the standard on-disk format for disk encryption on Linux systems, based on the dm-crypt kernel module.

