Source: https://mysqlentomologist.blogspot.com/2021/01/linux-proc-filesystem-for-mysql-dbas.html

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

Every command line tool like ps,top,vmstat,free derives its results from parsing `/proc` directory.

<img width="1132" height="1380" alt="image" src="https://github.com/user-attachments/assets/40298b8a-54f5-4aa2-8c15-8b3a22ae9e29" />


Most  of it is read-only, but some files allow to change kernel variables. 

```
openxs@ao756:~$ ps aux | grep mysqld
...
mysql    30580  0.7  8.1 746308 313984 ?       Sl   Jan02   9:55 /usr/sbin/mysqld --daemonize --pid-file=/var/run/mysqld/mysqld.pid
```

So there will be the `/proc/30580` directory with the following content:

<img width="1000" height="502" alt="image" src="https://github.com/user-attachments/assets/49a89a0f-33b7-4c0c-a3b9-086b8eb3684d" />

I highlighted the files and directories I consider most useful. 
Note also "Permission denied" messages above that you may get while accessing some files in `/proc`, 
even related to the processes you own. You may still need root/sudo access (or belong to some dedicated group) to read them.

<img width="1702" height="356" alt="image" src="https://github.com/user-attachments/assets/83964633-b1c2-4871-8a8c-3f2d6f81a3c5" />

