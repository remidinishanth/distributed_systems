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

![image](https://github.com/user-attachments/assets/c52fb749-e7a6-4f5d-ae3a-d1486be7029f)

```
➜  ~ docker ps
CONTAINER ID   IMAGE                                                    COMMAND                  CREATED             STATUS             PORTS                                                  NAMES
0ac949292b65   traefik                                                  "/entrypoint.sh trae…"   About an hour ago   Up About an hour   80/tcp                                                 adoring_bhabha
```

```
➜  ~ docker inspect 0ac949292b65 | grep  GraphDriver -A 8
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926-init/diff:/var/lib/docker/overlay2/54d68871e061924a04e9f5929876c1f62799ef505e2850bfb453e44d5c785982/diff:/var/lib/docker/overlay2/44eac91bcc11009b4bc33aa25b9d23d9b175f31b548d458b1db0f9acd61d829c/diff:/var/lib/docker/overlay2/da89c4703e3774e3a5ee4f0821232d19a7d4a006c60cf253c9d656c9c1e0910d/diff:/var/lib/docker/overlay2/4f9ef368c98bdfe533018873f57eb8a2aaa1815ed967ebf31683e0b3d5224975/diff",
                "MergedDir": "/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/merged",
                "UpperDir": "/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/diff",
                "WorkDir": "/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/work"
            },
            "Name": "overlay2"
        },
```

```
➜  ~ docker inspect adoring_bhabha | grep MergedDir

                "MergedDir": "/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/merged",
```

```
➜  ~ mount | grep 45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926
overlay on /var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/merged type overlay (rw,relatime,lowerdir=/var/lib/docker/overlay2/l/MEJZY4W3TTOIOOSQFSI76TU5IR:/var/lib/docker/overlay2/l/TXPWWDBIAR4MKWCSSEQXD2G5GX:/var/lib/docker/overlay2/l/OR3RPUZHUFDK3NKMHT24VXK254:/var/lib/docker/overlay2/l/EN5ZHUWKYHDCGGK57EMO27PGDJ:/var/lib/docker/overlay2/l/ME4EEDCZPBSQLDBRU4CANROOQH,upperdir=/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/diff,workdir=/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/work)
```


```
➜  ~ sudo findmnt --target /var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/merged

TARGET                                                                                           SOURCE  FSTYPE  OPTIONS
/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/merged overlay overlay rw,relatime,lowerdir=/var/lib/docker/overlay2/l/MEJZY4W3TTOIOOSQFSI76TU5IR:/var/lib/docker/overlay2/l
```

In the following, we can see:
* The Root Filesystem (OverlayFS) `1427 ... / ... overlay overlay rw,lowerdir=...,upperdir=...,workdir=...`
* Virtual Filesystems (/proc, /dev, /sys) `1428 1429 ...`
* Cgroups Mounts `1433 ...`
* Container-Specific Config Files `1447 1448 1449`

```
➜  ~ sudo cat /proc/66560/mountinfo
1427 1267 0:159 / / rw,relatime master:58 - overlay overlay rw,lowerdir=/var/lib/docker/overlay2/l/MEJZY4W3TTOIOOSQFSI76TU5IR:/var/lib/docker/overlay2/l/TXPWWDBIAR4MKWCSSEQXD2G5GX:/var/lib/docker/overlay2/l/OR3RPUZHUFDK3NKMHT24VXK254:/var/lib/docker/overlay2/l/EN5ZHUWKYHDCGGK57EMO27PGDJ:/var/lib/docker/overlay2/l/ME4EEDCZPBSQLDBRU4CANROOQH,upperdir=/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/diff,workdir=/var/lib/docker/overlay2/45629581e8696009c105d94ba07e3eb69ce4cc05cc963132cdfb5103d07fb926/work
1428 1427 0:166 / /proc rw,nosuid,nodev,noexec,relatime - proc proc rw
1429 1427 0:167 / /dev rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1430 1429 0:168 / /dev/pts rw,nosuid,noexec,relatime - devpts devpts rw,gid=5,mode=620,ptmxmode=666
1431 1427 0:169 / /sys ro,nosuid,nodev,noexec,relatime - sysfs sysfs ro
1432 1431 0:170 / /sys/fs/cgroup rw,nosuid,nodev,noexec,relatime - tmpfs tmpfs rw,mode=755
1433 1432 0:29 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/systemd ro,nosuid,nodev,noexec,relatime master:11 - cgroup cgroup rw,xattr,name=systemd
1434 1432 0:31 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/perf_event ro,nosuid,nodev,noexec,relatime master:14 - cgroup cgroup rw,perf_event
1435 1432 0:32 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/cpu,cpuacct ro,nosuid,nodev,noexec,relatime master:15 - cgroup cgroup rw,cpu,cpuacct
1436 1432 0:33 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/cpuset ro,nosuid,nodev,noexec,relatime master:16 - cgroup cgroup rw,cpuset
1437 1432 0:34 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/blkio ro,nosuid,nodev,noexec,relatime master:17 - cgroup cgroup rw,blkio
1438 1432 0:35 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/rdma ro,nosuid,nodev,noexec,relatime master:18 - cgroup cgroup rw,rdma
1439 1432 0:36 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/net_cls,net_prio ro,nosuid,nodev,noexec,relatime master:19 - cgroup cgroup rw,net_cls,net_prio
1440 1432 0:37 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/devices ro,nosuid,nodev,noexec,relatime master:20 - cgroup cgroup rw,devices
1441 1432 0:38 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/freezer ro,nosuid,nodev,noexec,relatime master:21 - cgroup cgroup rw,freezer
1442 1432 0:39 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/hugetlb ro,nosuid,nodev,noexec,relatime master:22 - cgroup cgroup rw,hugetlb
1443 1432 0:40 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/memory ro,nosuid,nodev,noexec,relatime master:23 - cgroup cgroup rw,memory
1444 1432 0:41 /docker/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154 /sys/fs/cgroup/pids ro,nosuid,nodev,noexec,relatime master:24 - cgroup cgroup rw,pids
1445 1429 0:165 / /dev/mqueue rw,nosuid,nodev,noexec,relatime - mqueue mqueue rw
1446 1429 0:171 / /dev/shm rw,nosuid,nodev,noexec,relatime - tmpfs shm rw,size=65536k
1447 1427 8:1 /var/lib/docker/containers/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154/resolv.conf /etc/resolv.conf rw,relatime - ext4 /dev/sda1 rw,errors=remount-ro,data=ordered
1448 1427 8:1 /var/lib/docker/containers/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154/hostname /etc/hostname rw,relatime - ext4 /dev/sda1 rw,errors=remount-ro,data=ordered
1449 1427 8:1 /var/lib/docker/containers/0ac949292b659a21e0037c91c7149f6fea12235ae4c5840d8448714081973154/hosts /etc/hosts rw,relatime - ext4 /dev/sda1 rw,errors=remount-ro,data=ordered
1268 1428 0:166 /bus /proc/bus ro,nosuid,nodev,noexec,relatime - proc proc rw
1269 1428 0:166 /fs /proc/fs ro,nosuid,nodev,noexec,relatime - proc proc rw
1270 1428 0:166 /irq /proc/irq ro,nosuid,nodev,noexec,relatime - proc proc rw
1271 1428 0:166 /sys /proc/sys ro,nosuid,nodev,noexec,relatime - proc proc rw
1272 1428 0:166 /sysrq-trigger /proc/sysrq-trigger ro,nosuid,nodev,noexec,relatime - proc proc rw
1273 1428 0:172 / /proc/acpi ro,relatime - tmpfs tmpfs ro
1274 1428 0:167 /null /proc/kcore rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1275 1428 0:167 /null /proc/keys rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1276 1428 0:167 /null /proc/timer_list rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1277 1428 0:167 /null /proc/sched_debug rw,nosuid - tmpfs tmpfs rw,size=65536k,mode=755
1278 1428 0:173 / /proc/scsi ro,relatime - tmpfs tmpfs ro
1279 1431 0:174 / /sys/firmware ro,relatime - tmpfs tmpfs ro
```

![image](https://github.com/user-attachments/assets/65039f74-d7f1-471d-9c7d-e12dab06f2a4)


## Concepts
![image](https://github.com/user-attachments/assets/7ccf890e-bc91-433e-bcc4-2a7a412374c4)


![image](https://github.com/user-attachments/assets/d0c42348-1143-4ab5-8c3b-3c31c05fd5da)


![image](https://github.com/user-attachments/assets/5ccf578c-1d62-4e70-af84-a09f6d06632b)

![image](https://github.com/user-attachments/assets/bcf96052-df30-44ae-800f-43c5455af245)

<img width="554" height="418" alt="image" src="https://github.com/user-attachments/assets/800bbe98-a47f-4532-9730-545a83af3954" />


## What happens under the hood when we create a new container on Linux?

* When the command is fired from CLI by the user, it makes an API call to the docker daemon, which then calls containerD via GRPC, which further calls shim process and runC.
*  ContainerD handles execution/lifecycle operations like start, stop, pause and unpause. OCI (Open Container Initiative) layer does the interface with the kernel.
* RunC spins up the container and exits, however shim remains connected to the container. This is also the case when multiple containers are spun up.

<img width="1198" height="988" alt="image" src="https://github.com/user-attachments/assets/02f82160-aab8-4ad7-91d1-7de0e753056d" />

<img width="958" height="510" alt="image" src="https://github.com/user-attachments/assets/11eabf61-ccd6-48e0-b5fd-52d12ce02899" />

<img width="2437" height="1530" alt="image" src="https://github.com/user-attachments/assets/807bc203-ae39-479b-a160-80417f4cfe65" />

Ref: https://stackoverflow.com/questions/46649592/dockerd-vs-docker-containerd-vs-docker-runc-vs-docker-containerd-ctr-vs-docker-c


## History

A Brief History of Containers (by Jeff Victor & Kir Kolyshkin)

<img width="1938" height="1170" alt="image" src="https://github.com/user-attachments/assets/350bd7ad-b9d6-4ec5-bd64-6ec22ffadfc2" />

* 2005: Open VZ (Open Virtuzzo) is an operating system-level virtualization technology for Linux which **uses a patched Linux kernel** for virtualization, isolation, resource management and checkpointing. The code was not released as part of the official Linux kernel.

* Process Containers (launched by Google in 2006) was designed for limiting, accounting and isolating resource usage (CPU, memory, disk I/O, network) of a collection of processes. It was renamed “Control Groups (cgroups)” a year later and eventually merged to Linux kernel 2.6.24.

* LXC (LinuX Containers) was the first, most complete implementation of Linux container manager. It was implemented in 2008 using cgroups and Linux namespaces, and it works on a single Linux kernel **without requiring any patches**.

* Docker also used LXC in its initial stages and later replaced that container manager with its own library, libcontainer.

Ref: https://www.aquasec.com/blog/a-brief-history-of-containers-from-1970s-chroot-to-docker-2016/
