## TODO
* https://homepages.uc.edu/~thomam/Intro_Unix_Text/TOC.html
* https://dsf.berkeley.edu/cs262/FFS.pdf
* https://ops.tips/blog/what-is-slash-proc/

![image](https://github.com/user-attachments/assets/488bc987-89d4-4ff2-bd7d-769acb27c96b)

![image](https://github.com/user-attachments/assets/a412aa46-bb58-4d64-b41f-0aefb2ae5516)


<img width="924" alt="image" src="https://github.com/user-attachments/assets/30981023-cd70-4df6-8d6f-8f4089137ced" />

### Components
<img width="931" alt="image" src="https://github.com/user-attachments/assets/7508763f-51ef-4e57-993d-865446446a22" />

### File system
![image](https://github.com/user-attachments/assets/93114437-78ab-4e05-b40d-93089a411205)

![image](https://github.com/user-attachments/assets/61fb391d-c984-4c63-be07-06637f4156ec)

![image](https://github.com/user-attachments/assets/a49ae909-5042-4fa1-b703-9e4e7e62dce5)

![image](https://github.com/user-attachments/assets/5f708dde-102f-4dd1-90b1-ad87fb6402b4)


![image](https://github.com/user-attachments/assets/a133ec99-96e7-47e8-a12c-e9b7335d38df)

## Proc

![image](https://github.com/user-attachments/assets/43504491-a1e9-430c-9fb5-e8dbf0878562)


### Julia Evans

![image](https://github.com/user-attachments/assets/ea68d72c-9867-4dde-ac10-f4ad2e491e26)

![image](https://github.com/user-attachments/assets/b53fbeb3-bf51-46c9-bed4-ab16c0873e14)

## fid

![image](https://github.com/user-attachments/assets/ca74a8ea-00e1-413f-9818-f888f437b8be)

![image](https://github.com/user-attachments/assets/6901d206-f5bd-4ad7-a2bb-4efae3684cc7)

![image](https://github.com/user-attachments/assets/864e48a0-46ea-4d0a-8aab-7c3815449edd)

## Commands

![image](https://github.com/user-attachments/assets/cfb4aaeb-2a9c-453b-bc51-a71386700291)

![image](https://github.com/user-attachments/assets/56b6becc-d1e1-443b-b752-5d9132218d52)

<img width="885" alt="image" src="https://github.com/user-attachments/assets/69396e2b-318b-4030-9070-3d79d1150c04" />


### Boot Process

![image](https://github.com/user-attachments/assets/942dfde5-b494-47e8-a341-51910d70c577)

* Step 1 - When we turn on the power, BIOS (Basic Input/Output System) or UEFI (Unified Extensible Firmware Interface) firmware is loaded from non-volatile memory, and executes POST (Power On Self Test).

* Step 2 - BIOS/UEFI detects the devices connected to the system, including CPU, RAM, and storage.

* Step 3 - Choose a booting device to boot the OS from. This can be the hard drive, the network server, or CD ROM.

* Step 4 - BIOS/UEFI runs the boot loader (GRUB - Grand Unified Bootloader), which provides a menu to choose the OS or the kernel functions.

* Step 5 - After the kernel is ready, we now switch to the user space. The kernel starts up systemd as the first user-space process, which manages the processes and services, probes all remaining hardware, mounts filesystems, and runs a desktop environment.

* Step 6 - systemd activates the default. target unit by default when the system boots. Other analysis units are executed as well.

* Step 7 - The system runs a set of startup scripts and configure the environment.

* Step 8 - The users are presented with a login window. The system is now ready.

### Network

![image](https://github.com/user-attachments/assets/65f54108-b5c7-4785-b62d-5d3ffcb01813)

![image](https://github.com/user-attachments/assets/a717401b-1261-43cc-8539-84244c9dff51)

### Performance Observability Tools

<img width="1063" alt="image" src="https://github.com/user-attachments/assets/591106ec-23fc-4ccc-80fb-8540f4801b1e" />

![image](https://github.com/user-attachments/assets/530c50a0-742b-40b7-81b7-4707cb90434b)


![image](https://github.com/user-attachments/assets/b8d47d33-dd4b-450f-94a7-e0135a83b236)

Ref: https://www.brendangregg.com/linuxperf.html
