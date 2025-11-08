> For understanding virtual memory, start with this: _every address generated
by a user program is a virtual address_. The OS is just providing an illusion
to each process, specifically that it has its own large and private memory; with
some hardware help, the OS will turn these pretend virtual addresses into real
physical addresses, and thus be able to locate the desired information.

From 3 easy pieces of Operating system
