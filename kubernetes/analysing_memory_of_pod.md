https://community.ibm.com/community/user/blogs/leo-varghese/2024/06/04/kubernetes-memory-metrics

When you check the running container and navigate to the folder path `/sys/fs/cgroup/memory`, 
you obtain all the memory details of the container. In this directory, 
you can find memory metrics such as usage, limits, cache, and so on.

You can exec into the Running pod and then check the following

```
/ $ cd /sys/fs/cgroup/memory/
/sys/fs/cgroup/memory $ cat memory.stat
cache 49627136
rss 59940864
rss_huge 2097152
mapped_file 1196032
swap 0
pgpgin 49075
pgpgout 23347
pgfault 51653
pgmajfault 54
inactive_anon 0
active_anon 59895808
inactive_file 35459072
active_file 14168064
unevictable 0
hierarchical_memory_limit 134217728
hierarchical_memsw_limit 134217728
total_cache 49627136
total_rss 59940864
total_rss_huge 2097152
total_mapped_file 1196032
total_swap 0
total_pgpgin 0
total_pgpgout 0
total_pgfault 0
total_pgmajfault 0
total_inactive_anon 0
total_active_anon 59895808
total_inactive_file 35459072
total_active_file 14168064
total_unevictable 0
/sys/fs/cgroup/memory $
```

Inside the pod, we can also check

```
/ $ ps -o pid,comm,vsz,rss
PID   COMMAND          VSZ  RSS
    1 main.bin         2.2g 119m
   51 sh               1752  724
   57 ps               1668  304
```

RSS(Resident set size) is the physical memory in the main memory that doesn’t correspond to anything on disk. RSS includes stacks, heaps, and anonymous memory maps.

and

```
/ $ grep -E 'Name|VmPeak|VmSize|VmRSS|Threads' /proc/1/status
Name:	main.bin
VmPeak:	 2307988 kB
VmSize:	 2307988 kB
VmRSS:	  121500 kB
Threads:	11
```

You can get to the same directory by checking the process id and then same as well

* Get the running container docker id, you can also use `docker ps | grep <pod-name>` to get the container id
* Then do docker inspect and get the Pid

```
docker inspect $(kubectl get pod <your-pod-name> -n <your-namespace> -o jsonpath='{.status.containerStatuses[0].containerID}' | cut -d'/' -f3) | grep '"Pid":'
```

And then for the whole container 

```
VR-POLARIS-VW-1697084:/proc/27038$ ps -p 27038 -o pid,user,vsz,rss,stat,start,time,command
  PID USER        VSZ   RSS STAT  STARTED     TIME COMMAND
27038 planet   2308308 122796 Ssl   Nov 14 00:03:33 /main.bin
```

And then by checking the `/proc/<pid>/cgroup`

```
VR-POLARIS-VW-1697084:/proc/27038$ cat /proc/27038/cgroup
11:freezer:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
10:pids:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
9:cpuset:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
8:devices:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
7:memory:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
6:perf_event:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
5:cpuacct,cpu:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
4:hugetlb:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
3:blkio:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
2:net_prio,net_cls:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
1:name=systemd:/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
```

The cgroup path for your pod (named pod8657690b-b81c-4ffd-84fe-585ef6f08eb7) is nested under `/kubepods/burstable/` and the full cgroup path you provided: `/kubepods/burstable/pod<Pod-UID>/<Container-ID>`

```
cat /sys/fs/cgroup/memory/kubepods/burstable/pod8657690b-b81c-4ffd-84fe-585ef6f08eb7/memory.limit_in_bytes
134217728
```

We can also get the pod uid `pod8657690b-b81c-4ffd-84fe-585ef6f08eb7` using `kubectl get pod <your-pod-name> -n <your-namespace> -o jsonpath='{.metadata.uid}'`

and container id, we can get it from describing the pod, using `kubectl describe pod <your-pod-name> -n <your-namespace>`

```
Containers:
  my-app-container:
    Container ID:   docker://55906a09ba7e4abe794f650a365ab0cc6bc280623d571e229fae8b5ca7fd4272
    Image:          my-image:latest
    ...
```

Here, `burstable` is the part of QoS. QoS (Quality of Service) is a classification system that determines how "important" your pod is

We can verify that using `kubectl get pod <your-pod-name> -n <your-namespace> -o jsonpath='{.status.qosClass}'`

<img width="3865" height="1320" alt="image" src="https://github.com/user-attachments/assets/4e62d884-f6c3-41ec-9591-d7c1a6e0bb1c" />

Each pod falls into one of three classes, based on how it defines its CPU and memory requests and limits:
* A Guaranteed pod has CPU and memory requests that are exactly the same as their limits. These get the best performance and priority.
* A Burstable pod has a lower request than limit—so it’s guaranteed a baseline, but it can use more if there’s room.
* A Best-Effort pod doesn’t define any requests or limits. It uses whatever’s left and is the first to get throttled or evicted when things get tight.

<img width="1001" height="377" alt="image" src="https://github.com/user-attachments/assets/d1ed2e94-60f9-4250-befd-22831f52dc94" />

<img width="1864" height="1541" alt="image" src="https://github.com/user-attachments/assets/fd92fada-9bfe-4379-9717-9b549f5b0d3c" />
