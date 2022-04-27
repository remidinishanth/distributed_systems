Most physical machines are underused, working on 15% of their capacity at best, leaving 85% unused. vSphere comes with the solution to this problem, allowing us to utilize most of the hardware. On a single physical machine, multiple instances of virtual machines can be installed, running different OS and apps, your entire infrastructure can be virtualized, even if it consists of multiple connected physical machines and storage. And with that comes increased amounts of data generated, and you probably want to secure this data, even though VMs are a more secure way of keeping your data, it is not 100% bulletproof. This is why you need to have a backup of your virtual machines in place.

Ref: https://xopero.com/blog/en/2021/09/03/vmware-vsphere-vcenter-and-esxi-definitions-and-differences/#:~:text=To%20recap%2C%20vSphere%20is%20a,on%20a%20few%20ESXi%20servers.


VMware vSphere is VMware's virtualization platform, which transforms data centers into aggregated computing infrastructures that include CPU, storage, and networking resources. 

![image](https://user-images.githubusercontent.com/19663316/165442472-6896206e-d47f-4f21-9577-0d1011fe7a99.png)

The two core components of vSphere are **ESXi** and **vCenter Server**. 
* ESXi is the virtualization platform where you create and run virtual machines and virtual appliances. 
* vCenter Server is the service through which you manage multiple hosts connected in a network and pool host resources.

More explanation:
* ESXi – a type 1 hypervisor, or in other words a bare-metal hypervisor. It is a virtualization software allowing you to create multiple instances of virtual machines on one physical host. It is installed directly on your physical machine and does not require an underlying operating system to function. ESXi runs on hosts to manage the execution of the virtual machines (VMs) and allocates resources to them as needed.
* vCenter Server – an advanced server management software that provides a centralized platform for controlling vSphere environments. vCenter Server allows you to automate and deliver a virtual infrastructure.
* vCenter Client – a management interface that enables users to remotely connect to vCenter Server.
* Virtual Machine File System (VMFS) – a cluster file system for ESXi virtual machines.

![image](https://user-images.githubusercontent.com/19663316/165449300-d443a921-8b34-4c38-8f65-80d2d0a5b6a3.png)

![image](https://user-images.githubusercontent.com/19663316/165449755-00e43334-d468-4273-a37f-0b9110885e00.png)

![image](https://user-images.githubusercontent.com/19663316/165449774-5fa784e7-cdf1-4516-9a53-58870ffa4a6c.png)

![image](https://user-images.githubusercontent.com/19663316/165446915-82524330-f3f7-4ecc-af14-a2b191a24e6e.png)

VMware vCloud Director (VMware vCD) is a platform with multi-tenant support for managing software-defined data centers (SDDC) and providing infrastructure as a service (IaaS) to customers. 
