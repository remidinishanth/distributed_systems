---
layout: page
title: "Readme"
category: "docker"
---

Two core concepts:
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

![image](https://github.com/user-attachments/assets/65039f74-d7f1-471d-9c7d-e12dab06f2a4)


## Concepts
![image](https://github.com/user-attachments/assets/7ccf890e-bc91-433e-bcc4-2a7a412374c4)


![image](https://github.com/user-attachments/assets/d0c42348-1143-4ab5-8c3b-3c31c05fd5da)


![image](https://github.com/user-attachments/assets/5ccf578c-1d62-4e70-af84-a09f6d06632b)

![image](https://github.com/user-attachments/assets/bcf96052-df30-44ae-800f-43c5455af245)


## History

A Brief History of Containers (by Jeff Victor & Kir Kolyshkin)

<img width="1938" height="1170" alt="image" src="https://github.com/user-attachments/assets/350bd7ad-b9d6-4ec5-bd64-6ec22ffadfc2" />

* 2005: Open VZ (Open Virtuzzo) is an operating system-level virtualization technology for Linux which **uses a patched Linux kernel** for virtualization, isolation, resource management and checkpointing. The code was not released as part of the official Linux kernel.

* Process Containers (launched by Google in 2006) was designed for limiting, accounting and isolating resource usage (CPU, memory, disk I/O, network) of a collection of processes. It was renamed “Control Groups (cgroups)” a year later and eventually merged to Linux kernel 2.6.24.

* LXC (LinuX Containers) was the first, most complete implementation of Linux container manager. It was implemented in 2008 using cgroups and Linux namespaces, and it works on a single Linux kernel **without requiring any patches**.

* Docker also used LXC in its initial stages and later replaced that container manager with its own library, libcontainer.

Ref: https://www.aquasec.com/blog/a-brief-history-of-containers-from-1970s-chroot-to-docker-2016/
