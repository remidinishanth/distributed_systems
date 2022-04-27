Most physical machines are underused, working on 15% of their capacity at best, leaving 85% unused. vSphere comes with the solution to this problem, allowing us to utilize most of the hardware. On a single physical machine, multiple instances of virtual machines can be installed, running different OS and apps, your entire infrastructure can be virtualized, even if it consists of multiple connected physical machines and storage. And with that comes increased amounts of data generated, and you probably want to secure this data, even though VMs are a more secure way of keeping your data, it is not 100% bulletproof. This is why you need to have a backup of your virtual machines in place.

Ref: https://xopero.com/blog/en/2021/09/03/vmware-vsphere-vcenter-and-esxi-definitions-and-differences/#:~:text=To%20recap%2C%20vSphere%20is%20a,on%20a%20few%20ESXi%20servers.


* ESXi – a type 1 hypervisor, or in other words a bare-metal hypervisor. It is a virtualization software allowing you to create multiple instances of virtual machines on one physical host. It is installed directly on your physical machine and does not require an underlying operating system to function. ESXi runs on hosts to manage the execution of the virtual machines (VMs) and allocates resources to them as needed.
* vCenter Server – an advanced server management software that provides a centralized platform for controlling vSphere environments. vCenter Server allows you to automate and deliver a virtual infrastructure.
* vCenter Client – a management interface that enables users to remotely connect to vCenter Server.
* Virtual Machine File System (VMFS) – a cluster file system for ESXi virtual machines.
