<img width="1694" height="352" alt="image" src="https://github.com/user-attachments/assets/37baac18-1935-4a71-a628-7261a7b13cea" />Nice blog https://wiki.archlinux.org/title/Data-at-rest_encryption

## Ways of Encryption

* Application side, Read more at https://www.percona.com/blog/transparent-data-encryption-tde/
  - Server side encryption(MySQL stores the tables etc encrypted, but when we read etc we get unencrypted data) - the encryption process is transparent to the user, allowing them to access and manipulate the data as usual without worrying about the encryption and decryption process.
  - Client side encryption(Client writes encrypted data to server, say we encrypt and store that encrypted data in MySQL rows)
 
* Disk level encryption that encrypts all data stored on a disk or storage device
  - Several tools are available for implementing disk-level encryption, including BitLocker for Windows, dm-crypt for Linux, and FileVault for MacOS.

<img width="1347" height="656" alt="image" src="https://github.com/user-attachments/assets/1d7f8ec2-04fb-43cf-9d03-54d1e156144a" />

-------

> All data-at-rest encryption methods operate in such a way that even though the disk actually holds encrypted data, the operating system and applications "see" it as the corresponding normal readable data as long as the cryptographic container (i.e. the logical part of the disk that holds the encrypted data) has been "unlocked" and mounted.
Ref: https://wiki.archlinux.org/title/Data-at-rest_encryption

For this to happen, some "secret information" (usually in the form of a keyfile and/or passphrase) needs to be supplied by the user, from which the actual encryption key can be derived (and stored in the kernel keyring for the duration of the session).

<img width="1862" height="1442" alt="image" src="https://github.com/user-attachments/assets/5a4c9a2c-83f8-4deb-a9f6-6e3356975445" />

<img width="1576" height="838" alt="image" src="https://github.com/user-attachments/assets/5584d73e-e93e-4265-8823-ba75c17076d7" />


> Device Mapper - Framework provided by the Linux Kernel, used to map physical block devices to higher level virtual block devices

> DM-Crypt - A target used with device mapper that provides transparent encryption. Allows us to create a virtual block device and have all data be encrypted on the fly before being committed to disk and can decrypt in the same way for reads.

> LUKS - Linux Unified Key Setup Provides an efficient user-friendly way to store and manage keys. Without LUKS, DM-Crypt can be more cumbersome and error-prone.

<img width="1694" height="352" alt="image" src="https://github.com/user-attachments/assets/87b12d15-ebe0-4a03-8306-b71297137c58" />


## Device Mapper

The device mapper is a framework provided by the Linux kernel for mapping physical block devices onto higher-level virtual block devices. 

It forms the foundation of the logical volume manager (LVM), software RAIDs and dm-crypt disk encryption, and offers additional features such as file system snapshots.

* The Device Mapper is a kernel driver that provides a framework for volume management.
* It provides a generic way of creating mapped devices, which may be used as logical volumes. It does not specifically know about volume groups or metadata formats.

<img width="640" height="494" alt="image" src="https://github.com/user-attachments/assets/9fd64a8a-ba21-46e7-b05c-2078dfad983f" />

## dm-crypt

`dm-crypt` is the Linux kernel's device mapper(dm) crypto target. A transparent disk encryption subsystem in the Linux kernel. It is implemented as a device mapper target and may be stacked on top of other device mapper transformations. 

dm-crypt is a transparent disk encryption subsystem. That being said, it's better suited to encrypt disks and partitions. It can encrypt files, but they have to be mapped as devices for this to work. You can still encrypt files by using loop devices, cryptsetup will even automatically create those loop devices as needed. https://unix.stackexchange.com/questions/275707/how-can-i-encrypt-a-file-with-dm-crypt

Ref: https://wiki.gentoo.org/wiki/Custom_Initramfs#Encrypted_keyfile

If you want to encrypt only one file, GnuPG could be a better tool. Example: `gpg -c filename`

<img width="1936" height="1080" alt="image" src="https://github.com/user-attachments/assets/21250f3b-69f2-4b58-92ee-91f903698d19" />

https://gitlab.com/cryptsetup/cryptsetup

When you unlock an encrypted volume, cryptsetup creates a new
device mapping that applications can access like any regular storage device. The actual encryption and decryption work is
performed transparently by the kernel’s device-mapper dm-crypt driver.

<img width="300" height="900" alt="image" src="https://github.com/user-attachments/assets/bdd1e970-9a36-4f28-821d-f752ea84fa67" />

<img width="1024" height="1024" alt="image" src="https://github.com/user-attachments/assets/ff1b7116-4185-45f1-96d4-a145c90abe77" />

Raw DM-Crypt requires manual key management. You have to handle key derivation, storage, and potential header formats yourself, which is prone to errors like weak key generation or incompatibility across systems.

## Linux Unified Key Setup (LUKS)

LUKS builds on DM-Crypt by standardizing the on-disk format for encrypted volumes, focusing on key management. It's essentially a header format that stores metadata (like encryption parameters and key slots) at the beginning of the block device.

LUKS is the standard on-disk format for disk encryption on Linux systems, based on the dm-crypt kernel module.

LUKS, the Linux Unified Key Setup, is a standard for disk encryption.
* It adds a standardized header at the start of the device, a keyslot area directly behind the header and the bulk data area behind that.
* The whole set is called a 'LUKS container'. The device that a LUKS container resides on is called a 'LUKS device'.

### Demo

<img width="1024" height="551" alt="image" src="https://github.com/user-attachments/assets/706ae5a9-0dcd-46e1-a3c7-8301eef82776" />

```
cryptsetup --verbose --cipher aes-xts-plain64 --key-size 512 --iter-time=4000 --hash sha512 --key-file=pw.txt luksFormat /dev/sdc
```

* `--cipher aes-xts-plain64`: Specifies the encryption algorithm.
  - `aes`: The Advanced Encryption Standard, a very common and secure cipher.
  -  `xts`: A block cipher mode (XTS-AES) designed for disk encryption that prevents an attacker from seeing patterns in the data.
  -  `plain64`: The initialization vector (IV) mode.

* `--key-size 512`: Sets the size of the master encryption key in bits. A 512-bit key is very strong.

* `--iter-time=4000`: This is a security feature to make brute-force attacks much harder. It tells the system to spend 4000 milliseconds (4 seconds) repeatedly hashing your passphrase to derive the actual decryption key. A longer time means it's much slower for an attacker to guess passwords.

* `--hash sha512`: The hashing algorithm used in the key derivation process. SHA-512 is a secure hashing algorithm.
* `--key-file=pw.txt` : Instead of prompting you to type a passphrase, this tells cryptsetup to use the content of the file pw.txt as the passphrase. This is useful for scripting.


```
time cryptsetup open --type luks /dev/sdc cryptdemo --key-file=pw.txt
```
* `--type luks`: Explicitly tells the command that this is a LUKS device.
* `cryptdemo`: The name for the decrypted virtual device that will be created. You will see it at `/dev/mapper/cryptdemo`

<img width="1292" height="1140" alt="image" src="https://github.com/user-attachments/assets/8c8ed32d-da51-4ea7-b40f-3d9803ee4df2" />


## Trusted Platform Module (TPM)

<img width="624" height="351" alt="image" src="https://github.com/user-attachments/assets/26589a84-3da2-4117-8bd7-538a148be545" />


In Linux, a TPM is a hardware chip on your computer's motherboard that acts as a secure cryptoprocessor to store cryptographic keys, perform cryptographic operations, and ensure the integrity of the system's boot process.

Linux systems use TPMs for key management, device authentication, and to bind secrets, like disk encryption keys, to the system's secure state through Platform Configuration Registers (PCRs). 

The TPM specification is an operating system agnostic, international standard (from the Trusted Computing Group and International Standards Organization) for a secure cryptoprocessor, which is a dedicated microprocessor designed to secure hardware by integrating cryptographic keys into devices.
