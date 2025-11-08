> For understanding virtual memory, start with this: _every address generated
by a user program is a virtual address_. The OS is just providing an illusion
to each process, specifically that it has its own large and private memory; with
some hardware help, the OS will turn these pretend virtual addresses into real
physical addresses, and thus be able to locate the desired information.

From 3 easy pieces of Operating system

**every address from a user program is virtual**

Why do we need that?

Mostly ease of use: the OS will give each program the view that it
has a large contiguous address space to put its code and data into; thus, as a
programmer, you never have to worry about things like “where should I store this
variable?” because the virtual address space of the program is large and has lots
of room for that sort of thing.

<img width="1000" height="327" alt="image" src="https://github.com/user-attachments/assets/d591d3e9-9a16-47d8-8b42-d36251047a63" />


Without virtual memory, a program couldn't run if it was larger than the available physical RAM. 
Virtual memory solves this by loading only the necessary parts of the program into RAM, 
keeping the rest on the hard drive. 
When a different part of the program is needed, the operating system swaps it into RAM, moving "older," unused parts out to the disk.

Figure shows the general layout of virtual memory, which can be much larger than physical memory:
<img width="733" height="585" alt="image" src="https://github.com/user-attachments/assets/9f712fc6-ede3-4058-8143-8f79e6651c62" />

