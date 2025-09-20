Source: https://mysqlentomologist.blogspot.com/2021/01/linux-proc-filesystem-for-mysql-dbas.html

## Proc

![image](https://github.com/user-attachments/assets/43504491-a1e9-430c-9fb5-e8dbf0878562)


### Julia Evans

![image](https://github.com/user-attachments/assets/ea68d72c-9867-4dde-ac10-f4ad2e491e26)

![image](https://github.com/user-attachments/assets/b53fbeb3-bf51-46c9-bed4-ab16c0873e14)


Basically, the proc filesystem is a pseudo-filesystem which provides an interface to kernel data structures. It is commonly mounted at /proc:   


`/etc/issue` is a text file which contains a message or system identification to be printed before the login prompt.

The `cat /etc/os-release` command displays the operating system's name and version information on a Linux system
```
openxs@ao756:~$ cat /etc/issue
Ubuntu 16.04.7 LTS \n \l
```

`mount` displays a list of all currently mounted filesystems. 
The output typically includes the device, the mount point, the filesystem type, and mount options.

```
openxs@ao756:~$ mount | grep '/proc'
proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
...
```


`/proc` is a temporary filesystem which linux initializes, in order to keep track of resources and states of each process,
as well as system wide resources.

<img width="1076" height="490" alt="image" src="https://github.com/user-attachments/assets/ec52687a-21cc-46a0-9ac9-07a5da8d6629" />

<img width="1152" height="1192" alt="image" src="https://github.com/user-attachments/assets/ee25c7f9-cee6-4f03-bbe3-6523a1120cdd" />


Every command line tool like ps,top,vmstat,free derives its results from parsing `/proc` directory.

<img width="1132" height="1380" alt="image" src="https://github.com/user-attachments/assets/40298b8a-54f5-4aa2-8c15-8b3a22ae9e29" />

<img width="5819" height="4115" alt="image" src="https://github.com/user-attachments/assets/79dbd0be-9282-4b8a-9419-ebf0eed6aaaa" />

<img width="816" height="693" alt="image" src="https://github.com/user-attachments/assets/3573c4d5-28f5-448f-a2f2-163315ed0d71" />

<img width="4096" height="4094" alt="image" src="https://github.com/user-attachments/assets/a45922d7-b09e-48d0-9cac-8954c499771e" />


Most  of it is read-only, but some files allow to change kernel variables. 

```
openxs@ao756:~$ ps aux | grep mysqld
...
mysql    30580  0.7  8.1 746308 313984 ?       Sl   Jan02   9:55 /usr/sbin/mysqld --daemonize --pid-file=/var/run/mysqld/mysqld.pid
```


So there will be the `/proc/30580` directory with the following content:

<img width="1000" height="502" alt="image" src="https://github.com/user-attachments/assets/49a89a0f-33b7-4c0c-a3b9-086b8eb3684d" />

Highlighted are the files and directories are more useful. 
Note also "Permission denied" messages above that you may get while accessing some files in `/proc`, 
even related to the processes you own. You may still need root/sudo access (or belong to some dedicated group) to read them.

<img width="1620" height="1500" alt="image" src="https://github.com/user-attachments/assets/deb12ffd-11a3-4105-8bab-04f00c761b20" />



<img width="1702" height="356" alt="image" src="https://github.com/user-attachments/assets/83964633-b1c2-4871-8a8c-3f2d6f81a3c5" />

### tasks

<img width="2246" height="1736" alt="image" src="https://github.com/user-attachments/assets/90e52e44-437f-4a64-b8cc-0f19ac297a85" />

### cmdline

<img width="2170" height="700" alt="image" src="https://github.com/user-attachments/assets/9e7e0111-dbee-4d0e-aca2-26e92b133d07" />

### comm

<img width="2172" height="768" alt="image" src="https://github.com/user-attachments/assets/cf02ed5a-6668-480e-b1da-f52f7511af50" />

### coredump_filter

<img width="1892" height="1784" alt="image" src="https://github.com/user-attachments/assets/f167ee99-8d2b-4dc3-a193-d07ed9435ef4" />

### environ

<img width="1814" height="1346" alt="image" src="https://github.com/user-attachments/assets/6ce2c84b-209a-481c-91e5-03ceed13aeca" />

### fd

<img width="2302" height="1752" alt="image" src="https://github.com/user-attachments/assets/26877482-e4ab-4291-a09e-891f509bb330" />

<img width="2248" height="986" alt="image" src="https://github.com/user-attachments/assets/a555a0b6-1d83-49ec-bdf3-0f68f82eef93" />

### fidinfo

<img width="2290" height="1742" alt="image" src="https://github.com/user-attachments/assets/ab947164-088c-4bff-988a-4319272a8b8f" />

<img width="2224" height="810" alt="image" src="https://github.com/user-attachments/assets/5759e992-8b3e-47a1-ba39-ad39af57dc96" />

### io

<img width="2256" height="1482" alt="image" src="https://github.com/user-attachments/assets/635671c4-08f5-4c01-8bfd-e2b15712dd39" />

### limits

<img width="2286" height="1654" alt="image" src="https://github.com/user-attachments/assets/2689fa5d-fe9a-4b63-9c3d-9f0cd2fa8d35" />

### maps

<img width="2264" height="1722" alt="image" src="https://github.com/user-attachments/assets/09281f15-7a93-4832-a082-cc29f225585a" />

### status

<img width="2184" height="1738" alt="image" src="https://github.com/user-attachments/assets/8634aac0-cf32-4732-84c5-f16127e6f8bd" />
