---
layout: page
title: "Readme"
category: "docker"
---

<img width="2020" height="774" alt="image" src="https://github.com/user-attachments/assets/911412e5-b386-4ece-9511-43c5b501c91c" />

<img width="2004" height="1398" alt="image" src="https://github.com/user-attachments/assets/362f1996-ab98-4b8d-b0ea-39612ce73ffa" />

<img width="2076" height="1392" alt="image" src="https://github.com/user-attachments/assets/2e605634-7013-423b-b1d0-5c433d66e269" />


## Two core concepts:
* Namespaces: Keeps your processes separated in isolated groups
* Cgroups: Controls the resources allocated

![image](https://github.com/user-attachments/assets/111bff81-55e7-471c-b60a-2c1f34d0a8a9)

![image](https://github.com/user-attachments/assets/b949a89b-d710-48f9-9baf-a675dc3a3e58)

### Namespaces
![image](https://github.com/user-attachments/assets/cc3688c7-cfd1-4ff6-9851-7642c672aee4)


```
docker run traefik
```

```
pstree -spa 66560

systemd,1 --system --deserialize 18
  └─containerd-shim,66535 -namespace moby -id 0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 -address /run/containerd/containerd.sock
      └─traefik,66560 traefik
```

<img width="1570" height="1264" alt="image" src="https://github.com/user-attachments/assets/075e7630-9cb2-4561-b6fe-8864cf4ee844" />

> nsenter - run program in different namespaces

#### Filesystem (Mount Namespace) comparision
<img width="1694" height="1852" alt="image" src="https://github.com/user-attachments/assets/ae55839b-758b-471c-866a-82255c4a91df" />

Similary, we can check
* `sudo nsenter -t 66560 -u -- hostname` and `hostname` for Hostname (UTS Namespace)
* `sudo nsenter -t 66560 -u -- ip addr` and `ip addr` for Network (Net Namespace)

#### Implementation

<img width="1442" height="1382" alt="image" src="https://github.com/user-attachments/assets/cfd0357f-de6c-420e-a27d-5dfe4f7b48af" />


```c

#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>

/* Define a stack for clone, stack size 1M */
#define STACK_SIZE (1024 * 1024)

static char container_stack [ STACK_SIZE ] ;

char * const container_args [] = {
    "/bin/bash" ,
    NULL
} ;

int container_main(void* arg)
{
    /* Looking at the PID of the child process,
    we can see that the pid of the output child process is 1 */
    printf("Container [%5d] - inside the container!\n", getpid());
    sethostname("container",10);
    execv(container_args[0], container_args);
    printf("Something's wrong!\n");
    return 1;
}

int main()
{
    printf("Parent [%5d] - start a container!\n", getpid());
    /* PID namespace - CLONE_NEWPID */
    int container_pid = clone(container_main, container_stack+STACK_SIZE, 
            CLONE_NEWUTS | CLONE_NEWPID | SIGCHLD, NULL); 
    waitpid(container_pid, NULL, 0);
    printf("Parent - container stopped!\n");
    return 0;
}
```

Output

```
hchen@ubuntu:~$ sudo ./pid
Parent [ 3474] - start a container!
Container [ 1] - inside the container!
root@container:~# echo $$
1
```

Ref: https://coolshell.cn/articles/17010.html

### Cgroups
![image](https://github.com/user-attachments/assets/0061da2d-84b4-4f73-8d62-38e4fd806895)

<img width="1758" height="864" alt="image" src="https://github.com/user-attachments/assets/e10e9df1-19ec-4ba2-a5ed-ff9ae1ecec8e" />

```
➜  ~ head -n 1 /proc/66560/cgroup

12:pids:/docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154
```


On most Linux systems, this very large number(2^63 - 1) is used to represent an "unlimited" or "no-limit" setting within cgroups.
```
➜  ~ cat /sys/fs/cgroup/memory/docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154/memory.limit_in_bytes

9223372036854771712
```

### Deep Dive into Docker Internals - Union Filesystem

https://martinheinz.dev/blog/44


### Overlay filesystem

What is overlay fs https://wiki.archlinux.org/title/Overlay_filesystem 

![image](https://github.com/user-attachments/assets/c52fb749-e7a6-4f5d-ae3a-d1486be7029f)

#### Ubuntu example
<img width="1100" height="667" alt="image" src="https://github.com/user-attachments/assets/d5ba6ecc-a397-44b1-90a3-4b99eaf68721" />

#### How containers use this
<img width="985" height="588" alt="image" src="https://github.com/user-attachments/assets/41c4b677-0943-4b00-aea6-e0494e242808" />

If container writes any files, it doesn't modify anything in lower layers

<img width="1706" height="690" alt="image" src="https://github.com/user-attachments/assets/f513fa8a-837d-4e31-8559-b444a594d6a7" />


```
docker run -d traefik
c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6
```

Running `mount -t overlay` inside the docker
```
docker exec -it c702369a8429 sh
/ # mount -t overlay
overlay on / type overlay (rw,relatime,lowerdir=/var/lib/docker/overlay2/l/CK3RK6RKLXTDLCVT7J6XUNJFYI:/var/lib/docker/overlay2/l/MJZW5RC5EQX5QV64ZQFI5YRA6V:/var/lib/docker/overlay2/l/XG3WJGGNM4CP67RWANTABIWBOL:/var/lib/docker/overlay2/l/X32XXQFB6ADFFO2FLDCVIV6J2K:/var/lib/docker/overlay2/l/T72XWGVHJ6FWJXBYGSBLRK6FPE,upperdir=/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/diff,workdir=/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/work)
```

```
docker inspect c702369a8429 | grep  GraphDriver -A 8
        "GraphDriver": {
            "Data": {
                "ID": "c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6",
                "LowerDir": "/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413-init/diff:/var/lib/docker/overlay2/0107d134713b05fc02091a41f1da372a9c9a0b7442f0c6a9ec130ace13940fe8/diff:/var/lib/docker/overlay2/8e8803ebddca09cd58274141eed8e426ddb4d3b96273cdda29c61f17ca20513b/diff:/var/lib/docker/overlay2/6b075fb9786d41cae6451f6ccc4e7708133646b57f45460394508e63a0da822b/diff:/var/lib/docker/overlay2/8beff5c84e30b1915a9017f659232bacde302c7386b5a9b7e4196b3932492780/diff",
                "MergedDir": "/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/merged",
                "UpperDir": "/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/diff",
                "WorkDir": "/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/work"
            },
            "Name": "overlay2"
```            

Also on the host by searching merged dir `/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/merged`

```
➜  ~ mount | grep 79ded441a3bd88ad3721bf119dc6266904
overlay on /var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/merged type overlay (rw,relatime,lowerdir=/var/lib/docker/overlay2/l/CK3RK6RKLXTDLCVT7J6XUNJFYI:/var/lib/docker/overlay2/l/MJZW5RC5EQX5QV64ZQFI5YRA6V:/var/lib/docker/overlay2/l/XG3WJGGNM4CP67RWANTABIWBOL:/var/lib/docker/overlay2/l/X32XXQFB6ADFFO2FLDCVIV6J2K:/var/lib/docker/overlay2/l/T72XWGVHJ6FWJXBYGSBLRK6FPE,upperdir=/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/diff,workdir=/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/work)
```


```
➜  ~ sudo findmnt --target /var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/merged
TARGET                                                                                           SOURCE  FSTYPE  OPTIONS
/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/merged overlay overlay rw,relatime,lowerdir=/var/lib/docker/overlay2/l/CK3RK6RKLXTDLCVT7J6XUNJFYI:/var/lib/docker/overlay2/l
➜  ~
```

In the following, we can see detailed map of the filesystem environment for docker container process 14188:
* The Root Filesystem (OverlayFS) `1121 ... / ... overlay overlay rw,lowerdir=...,upperdir=...,workdir=...`
* Virtual Filesystems (/proc, /dev, /sys) `1122 proc 1123 tmpfs 1125 sysfs ...`
* Cgroups Mounts `/sys/fs/cgroup/*` `1127 ...`
* Container-Specific Configuration Files `/resolv.conf /hostname /hosts`

```
➜  ~ sudo cat /proc/14188/mountinfo
1121 994 0:80 / / rw,relatime - overlay overlay rw,lowerdir=/var/lib/docker/overlay2/l/CK3RK6RKLXTDLCVT7J6XUNJFYI:/var/lib/docker/overlay2/l/MJZW5RC5EQX5QV64ZQFI5YRA6V:/var/lib/docker/overlay2/l/XG3WJGGNM4CP67RWANTABIWBOL:/var/lib/docker/overlay2/l/X32XXQFB6ADFFO2FLDCVIV6J2K:/var/lib/docker/overlay2/l/T72XWGVHJ6FWJXBYGSBLRK6FPE,upperdir=/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/diff,workdir=/var/lib/docker/overlay2/79ded441a3bd88ad3721bf119dc626690444ce58c9ed378f5a1b923667abe413/work
1122 1121 0:87 / /proc rw,nosuid,nodev,noexec,relatime - proc proc rw
1123 1121 0:88 / /dev rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1124 1123 0:89 / /dev/pts rw,nosuid,noexec,relatime - devpts devpts rw,gid=5,mode=620,ptmxmode=666
1125 1121 0:90 / /sys ro,nosuid,nodev,noexec,relatime - sysfs sysfs ro
1126 1125 0:91 / /sys/fs/cgroup rw,nosuid,nodev,noexec,relatime - tmpfs tmpfs rw,mode=755
1127 1126 0:29 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/systemd ro,nosuid,nodev,noexec,relatime master:11 - cgroup cgroup rw,xattr,name=systemd
1128 1126 0:31 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/perf_event ro,nosuid,nodev,noexec,relatime master:14 - cgroup cgroup rw,perf_event
1129 1126 0:32 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/cpu,cpuacct ro,nosuid,nodev,noexec,relatime master:15 - cgroup cgroup rw,cpu,cpuacct
1130 1126 0:33 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/cpuset ro,nosuid,nodev,noexec,relatime master:16 - cgroup cgroup rw,cpuset
1131 1126 0:34 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/blkio ro,nosuid,nodev,noexec,relatime master:17 - cgroup cgroup rw,blkio
1132 1126 0:35 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/rdma ro,nosuid,nodev,noexec,relatime master:18 - cgroup cgroup rw,rdma
1133 1126 0:36 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/net_cls,net_prio ro,nosuid,nodev,noexec,relatime master:19 - cgroup cgroup rw,net_cls,net_prio
1134 1126 0:37 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/devices ro,nosuid,nodev,noexec,relatime master:20 - cgroup cgroup rw,devices
1135 1126 0:38 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/freezer ro,nosuid,nodev,noexec,relatime master:21 - cgroup cgroup rw,freezer
1136 1126 0:39 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/hugetlb ro,nosuid,nodev,noexec,relatime master:22 - cgroup cgroup rw,hugetlb
1137 1126 0:40 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/memory ro,nosuid,nodev,noexec,relatime master:23 - cgroup cgroup rw,memory
1138 1126 0:41 /docker/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6 /sys/fs/cgroup/pids ro,nosuid,nodev,noexec,relatime master:24 - cgroup cgroup rw,pids
1139 1123 0:86 / /dev/mqueue rw,nosuid,nodev,noexec,relatime - mqueue mqueue rw
1140 1123 0:92 / /dev/shm rw,nosuid,nodev,noexec,relatime - tmpfs shm rw,size=65536k
1141 1121 8:1 /var/lib/docker/containers/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6/resolv.conf /etc/resolv.conf rw,relatime - ext4 /dev/sda1 rw,errors=remount-ro,data=ordered
1142 1121 8:1 /var/lib/docker/containers/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6/hostname /etc/hostname rw,relatime - ext4 /dev/sda1 rw,errors=remount-ro,data=ordered
1143 1121 8:1 /var/lib/docker/containers/c702369a8429445312f561631ef8871ed9b8c055551151e549190398fef936e6/hosts /etc/hosts rw,relatime - ext4 /dev/sda1 rw,errors=remount-ro,data=ordered
995 1122 0:87 /bus /proc/bus ro,nosuid,nodev,noexec,relatime - proc proc rw
996 1122 0:87 /fs /proc/fs ro,nosuid,nodev,noexec,relatime - proc proc rw
997 1122 0:87 /irq /proc/irq ro,nosuid,nodev,noexec,relatime - proc proc rw
998 1122 0:87 /sys /proc/sys ro,nosuid,nodev,noexec,relatime - proc proc rw
999 1122 0:87 /sysrq-trigger /proc/sysrq-trigger ro,nosuid,nodev,noexec,relatime - proc proc rw
1012 1122 0:93 / /proc/acpi ro,relatime - tmpfs tmpfs ro
1013 1122 0:88 /null /proc/interrupts rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1014 1122 0:88 /null /proc/kcore rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1015 1122 0:88 /null /proc/keys rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1016 1122 0:88 /null /proc/timer_list rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1017 1122 0:88 /null /proc/sched_debug rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1018 1122 0:94 / /proc/scsi ro,relatime - tmpfs tmpfs ro
1019 1125 0:95 / /sys/firmware ro,relatime - tmpfs tmpfs ro
```

Ignore typos my AI generated image
<img width="1024" height="1024" alt="image" src="https://github.com/user-attachments/assets/290338f7-f00e-4140-8368-aff3e271d9a7" />


![image](https://github.com/user-attachments/assets/65039f74-d7f1-471d-9c7d-e12dab06f2a4)


## Concepts
![image](https://github.com/user-attachments/assets/7ccf890e-bc91-433e-bcc4-2a7a412374c4)


![image](https://github.com/user-attachments/assets/d0c42348-1143-4ab5-8c3b-3c31c05fd5da)


![image](https://github.com/user-attachments/assets/5ccf578c-1d62-4e70-af84-a09f6d06632b)

![image](https://github.com/user-attachments/assets/bcf96052-df30-44ae-800f-43c5455af245)

<img width="554" height="418" alt="image" src="https://github.com/user-attachments/assets/800bbe98-a47f-4532-9730-545a83af3954" />

<img width="958" height="510" alt="image" src="https://github.com/user-attachments/assets/11eabf61-ccd6-48e0-b5fd-52d12ce02899" />

<img width="1782" height="1080" alt="image" src="https://github.com/user-attachments/assets/6eb0aea4-cfb3-4b7b-a8bf-8f27e974d8c9" />

<img width="2859" height="4096" alt="image" src="https://github.com/user-attachments/assets/22810829-3bda-404e-b7c8-24643112fa2f" />


## What happens under the hood when we create a new container on Linux?

* When the command is fired from CLI by the user, it makes an API call to the docker daemon, which then calls containerD via GRPC, which further calls shim process and runC.
*  ContainerD handles execution/lifecycle operations like start, stop, pause and unpause. OCI (Open Container Initiative) layer does the interface with the kernel.
* RunC spins up the container and exits, however shim remains connected to the container. This is also the case when multiple containers are spun up.

<img width="1198" height="988" alt="image" src="https://github.com/user-attachments/assets/02f82160-aab8-4ad7-91d1-7de0e753056d" />

<img width="2437" height="1530" alt="image" src="https://github.com/user-attachments/assets/807bc203-ae39-479b-a160-80417f4cfe65" />

Ref: https://stackoverflow.com/questions/46649592/dockerd-vs-docker-containerd-vs-docker-runc-vs-docker-containerd-ctr-vs-docker-c

<img width="492" height="452" alt="image" src="https://github.com/user-attachments/assets/2c2516af-4900-474f-aab4-021a4efbebf2" />


## TODO

https://blog.quarkslab.com/digging-into-runtimes-runc.html

Ref: https://terenceli.github.io/%E6%8A%80%E6%9C%AF/2021/12/22/runc-internals-1

Ref: https://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/

## History

A Brief History of Containers (by Jeff Victor & Kir Kolyshkin)

<img width="1938" height="1170" alt="image" src="https://github.com/user-attachments/assets/350bd7ad-b9d6-4ec5-bd64-6ec22ffadfc2" />

* 2005: Open VZ (Open Virtuzzo) is an operating system-level virtualization technology for Linux which **uses a patched Linux kernel** for virtualization, isolation, resource management and checkpointing. The code was not released as part of the official Linux kernel.

* Process Containers (launched by Google in 2006) was designed for limiting, accounting and isolating resource usage (CPU, memory, disk I/O, network) of a collection of processes. It was renamed “Control Groups (cgroups)” a year later and eventually merged to Linux kernel 2.6.24.

* LXC (LinuX Containers) was the first, most complete implementation of Linux container manager. It was implemented in 2008 using cgroups and Linux namespaces, and it works on a single Linux kernel **without requiring any patches**.

* Docker also used LXC in its initial stages and later replaced that container manager with its own library, libcontainer.

Ref: https://www.aquasec.com/blog/a-brief-history-of-containers-from-1970s-chroot-to-docker-2016/
