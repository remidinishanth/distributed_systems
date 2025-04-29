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


### Character vs Block device

![image](https://github.com/user-attachments/assets/d2e80327-aaaa-4603-a0b7-b4ad8fa5ed01)

#### Character devices
* A character (char) device is one that can be accessed as a stream of bytes (like afile); a char driver is in charge of implementing this behavior. Such a driver usually implements at least the open, close, read, and write system calls.
* The text console (/dev/console) and the serial ports (/dev/ttyS0 and friends) are examples
of char devices, as they are well represented by the stream abstraction.
* Char devices are accessed by means of filesystem nodes, such as /dev/tty1 and /dev/lp0.
  - The only relevant difference between a char device and a regular file is that you can always move back and forth in the regular file, whereas most char devices are just data channels, which you can only access sequentially. There exist, nonetheless, char devices that look like data areas, and you can move back and
forth in them; for instance, this usually applies to frame grabbers, where the
applications can access the whole acquired image using mmap or lseek.

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
