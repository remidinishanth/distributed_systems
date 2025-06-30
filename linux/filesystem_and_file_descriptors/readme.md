---
layout: page
title: "Readme"
category: "filesystem_and_file_descriptors"
---

File descriptors are an index into a file descriptor table stored by the kernel. The kernel creates a file descriptor in response to an open call and associates the file descriptor with some abstraction of an underlying file-like object, be that an actual hardware device, or a file system or something else entirely. Consequently a process's read or write calls that reference that file descriptor are routed to the correct place by the kernel to ultimately do something useful.

<img width="1081" alt="image" src="https://github.com/user-attachments/assets/1b450486-8fb7-4894-a75d-d1ee063c8290" />

![image](https://github.com/user-attachments/assets/9f01897d-b63e-4da2-bfc3-ad57a125008e)
Ref: The UNIX Time- Sharing System by DM Ritchie · 1974 

### File Descriptor
* is an integer that refers to an open file description
* the mapping of decriptor to description is local to a specific process
* the open file description is a kernel data structure, not directly accessible by the process

![image](https://github.com/user-attachments/assets/e15baf3d-1266-467b-9525-59052e8a81cf)

* other file descriptors are assigned as open() is called
* dup() and dup2() can be used to manipulate file descriptors

Ref: https://www.cs.fsu.edu/~baker/opsys/notes/unixfiles.html

## File Descriptors
* How do we interact with Linux Filesystem via 𝐟𝐢𝐥𝐞 𝐝𝐞𝐬𝐜𝐫𝐢𝐩𝐭𝐨𝐫𝐬?

A file descriptor represents an open file. It is a unique number assigned by the operating system to each file. It is an 𝐚𝐛𝐬𝐭𝐫𝐚𝐜𝐭𝐢𝐨𝐧 for working with files. We need to use file descriptors to read from or write to files in our program.

Each process maintains its own file descriptor table.

🔹 In User Space
When we open a file called “fileA.txt” in Process 1234, we get file descriptor fd1, which is equal to 3. We can then pass the file descriptor to other functions, to write data to the file.

🔹 In Kernel Space
In Linux kernel, there is a 𝐩𝐫𝐨𝐜𝐞𝐬𝐬 𝐭𝐚𝐛𝐥𝐞 to maintain the data for the processes. Each process has an entry in the table. Each process maintains a file descriptor table, with file descriptors as its indices.

The file pointer points to an entry in the 𝐨𝐩𝐞𝐧 𝐟𝐢𝐥𝐞 𝐭𝐚𝐛𝐥𝐞, which has information about open files across all processes. Multiple file descriptors can point to the same file table entry. For example, file descriptor 0,1 and 2 point to the same open file table entry.

Since different open file table entries can represent the same file, it is a waste of resources to store the file static information so many times. We need another abstraction layer called ‘vnode table’ to store the static data.

![image](https://github.com/user-attachments/assets/f62807ba-e39b-4c0c-9424-f0d7fed26c1b)

![image](https://github.com/user-attachments/assets/0762cb57-cb75-4b3a-aec3-88ef63e15c39)

![image](https://github.com/user-attachments/assets/4b615ad2-6e51-48cb-84ca-3991c5dd2b13)


## Virtual File System
<img width="785" alt="image" src="https://github.com/user-attachments/assets/c20d2465-292e-4793-ad46-fb0df67a8df3" />


<img width="597" alt="image" src="https://github.com/user-attachments/assets/de5cc6f9-45c0-489b-9b61-b37d1a9c7235" />

## Memory Types
![image](https://github.com/user-attachments/assets/4dbe62ca-5ee7-4d57-b9dc-624143bb89f0)

## File permissions

![image](https://github.com/user-attachments/assets/86a85bbd-103c-43eb-baed-d2bbe8620dad)

![image](https://github.com/user-attachments/assets/28a8140f-b032-47fa-93a2-c91d5e172833)

![image](https://github.com/user-attachments/assets/6e8c7466-c6e3-48e8-8cc3-9fa3befb51c5)

<img width="587" alt="image" src="https://github.com/user-attachments/assets/7d738d8a-77e5-4bc2-b356-6071fd84c030" />
