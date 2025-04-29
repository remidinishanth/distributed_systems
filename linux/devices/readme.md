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

https://dev.to/shankarsurya035/how-to-create-lvm-partition-in-linux-dgo

LVM creates a logical volume that aggregates multiple disks to stripe reads and writes. 
It also serves as a scalable volume for maintenance.

* Logical Volume Manager (LVM) plays an important role in the Linux operating system by improving the availability, disk I/O, performance and capability of disk management.

* LVM is a widely used technique that is extremely flexible for disk management.

* This adds an extra layer between the physical disks and the file system, allowing you to create a logical volume instead of a physical disk.

* LVM allows you to easily resize, extend and decrease the logical volume when you need it.


![image](https://github.com/user-attachments/assets/8bc1e9ec-824b-4466-a098-de9888f56736)


<img width="594" alt="image" src="https://github.com/user-attachments/assets/42721ba9-0ca8-4267-a769-882bcc628f3a" />


![image](https://github.com/user-attachments/assets/2b17c6b2-287f-4106-b4ce-25aa65743f33)

### Useful commands
<img width="775" alt="image" src="https://github.com/user-attachments/assets/a6343cab-f243-4109-ad0e-db2aff00deac" />


<img width="1168" alt="image" src="https://github.com/user-attachments/assets/9f7df455-567c-4044-abc4-75e7f5afe32b" />

`sudo fdisk -l` lists all the partitions and disks connected to your Linux machine. Let's break it down step by step:
