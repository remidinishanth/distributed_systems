---
layout: page
title: "Devices and Disk Layout"
category: "linux"
description: "Linux devices, disk layout, and filesystem configuration"
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


> Special files for char drivers are identified by a “c” in the first column of the output of ls -l. Block devices appear in /dev as well, but they are identified by a “b.” The focus of this chapter is on char devices, but much of the following information applies to block devices as well.

