File descriptors are an index into a file descriptor table stored by the kernel. The kernel creates a file descriptor in response to an open call and associates the file descriptor with some abstraction of an underlying file-like object, be that an actual hardware device, or a file system or something else entirely. Consequently a process's read or write calls that reference that file descriptor are routed to the correct place by the kernel to ultimately do something useful.

<img width="1081" alt="image" src="https://github.com/user-attachments/assets/1b450486-8fb7-4894-a75d-d1ee063c8290" />

### File Descriptor
* is an integer that refers to an open file description
* the mapping of decriptor to description is local to a specific process
* the open file description is a kernel data structure, not directly accessible by the process

![image](https://github.com/user-attachments/assets/e15baf3d-1266-467b-9525-59052e8a81cf)

* other file descriptors are assigned as open() is called
* dup() and dup2() can be used to manipulate file descriptors

Ref: https://www.cs.fsu.edu/~baker/opsys/notes/unixfiles.html

![image](https://github.com/user-attachments/assets/0762cb57-cb75-4b3a-aec3-88ef63e15c39)

## Virtual File System
<img width="785" alt="image" src="https://github.com/user-attachments/assets/c20d2465-292e-4793-ad46-fb0df67a8df3" />


<img width="597" alt="image" src="https://github.com/user-attachments/assets/de5cc6f9-45c0-489b-9b61-b37d1a9c7235" />

