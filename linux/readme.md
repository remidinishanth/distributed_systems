![image](https://github.com/user-attachments/assets/488bc987-89d4-4ff2-bd7d-769acb27c96b)

![image](https://github.com/user-attachments/assets/a412aa46-bb58-4d64-b41f-0aefb2ae5516)


![image](https://github.com/user-attachments/assets/5240466a-baa2-4fba-bbae-f35e91dbbba3)

![image](https://github.com/user-attachments/assets/93114437-78ab-4e05-b40d-93089a411205)

![image](https://github.com/user-attachments/assets/61fb391d-c984-4c63-be07-06637f4156ec)

![image](https://github.com/user-attachments/assets/a49ae909-5042-4fa1-b703-9e4e7e62dce5)

![image](https://github.com/user-attachments/assets/5f708dde-102f-4dd1-90b1-ad87fb6402b4)


![image](https://github.com/user-attachments/assets/a133ec99-96e7-47e8-a12c-e9b7335d38df)

### Commands

![image](https://github.com/user-attachments/assets/cfb4aaeb-2a9c-453b-bc51-a71386700291)

### Boot Process

![image](https://github.com/user-attachments/assets/942dfde5-b494-47e8-a341-51910d70c577)

* Step 1 - When we turn on the power, BIOS (Basic Input/Output System) or UEFI (Unified Extensible Firmware Interface) firmware is loaded from non-volatile memory, and executes POST (Power On Self Test).

* Step 2 - BIOS/UEFI detects the devices connected to the system, including CPU, RAM, and storage.

* Step 3 - Choose a booting device to boot the OS from. This can be the hard drive, the network server, or CD ROM.

* Step 4 - BIOS/UEFI runs the boot loader (GRUB), which provides a menu to choose the OS or the kernel functions.

* Step 5 - After the kernel is ready, we now switch to the user space. The kernel starts up systemd as the first user-space process, which manages the processes and services, probes all remaining hardware, mounts filesystems, and runs a desktop environment.

* Step 6 - systemd activates the default. target unit by default when the system boots. Other analysis units are executed as well.

* Step 7 - The system runs a set of startup scripts and configure the environment.

* Step 8 - The users are presented with a login window. The system is now ready.
