## Systemd

`systemd`, also known as system daemon, is a kind of init software under GNU/Linux operating system.

Systemd is a system and service manager for Linux operating systems, primarily functioning as the init system that starts and manages user space processes.

While Systemd has become the standard, some distributions still use `SysV` or `LSB` init scripts. 

Systemd acts as PID 1, meaning it's the first process started by the kernel, and it's responsible for bringing up and maintaining other services. 

Purpose of development:
* to provide a better framework for representing dependencies between services
* implements parallel startup of services at system initialization
* reduces shell overhead and replaces SysV style init

![image](https://github.com/user-attachments/assets/3afab749-b935-480d-9663-7b1631093a34)

Ref: https://docs.rockylinux.org/books/admin_guide/16-about-sytemd/

> systemd provides aggressive parallelization capabilities, uses socket and D-Bus activation for starting services, offers on-demand starting of daemons, keeps track of processes using Linux cgroups, supports snapshotting and restoring of the system state, maintains mount and automount points and implements a powerful transactional dependency-based service control logic. It can work as a drop-in replacement for sysvinit.


![image](https://github.com/user-attachments/assets/a4c59c3a-3b1b-42be-a3a4-09e8bc51af50)
