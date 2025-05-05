### File Descriptor
* is an integer that refers to an open file description
* the mapping of decriptor to description is local to a specific process
* the open file description is a kernel data structure, not directly accessible by the process
* 0 is the file descriptor for stdin
* 1 is the file descriptor for stdout
* 2 is the file descriptor for stderr
* other file descriptors are assigned as open() is called
* dup() and dup2() can be used to manipulate file descriptors

Ref: https://www.cs.fsu.edu/~baker/opsys/notes/unixfiles.html

![image](https://github.com/user-attachments/assets/0762cb57-cb75-4b3a-aec3-88ef63e15c39)

