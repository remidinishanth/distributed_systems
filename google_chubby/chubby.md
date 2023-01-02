### Coarse-Grained Lock
> Locks a set of related objects with a single lock.

A Coarse-Grained Lock is a single lock that covers many objects. It not only simplifies the locking action itself but also frees you from having to load all the members of a group in order to lock them.

### Google Chubby - Distributed lock manager

Ref: https://people.cs.rutgers.edu/~pxk/417/notes/chubby.html


> Goal: Create a highly-available centralized lease manager and file system for small files that can 
serve as a name server and configuration repository.

Chubby is a highly available and persistent **distributed lock service** and file system created by Google and used in various projects and frameworks, including Bigtable and MapReduce. Its purpose is to manage locks for resources and store configuration information for various distributed services throughout the Google cluster environment. The key design goal of Chubby was high availability and high reliability.

By default, the service runs on five computers as five active replicas. This grouping is called a **Chubby cell**. One of these replicas is elected as the **master** to serve client requests. Only the master will serve client requests; the other machines in the cell are for fault tolerance in case the master dies. The only request they will answer from clients is to tell them who the master is.

A majority of the replicas in the cell must be running for the service to work. **Paxos** is used as a leader election algorithm and is used to keep the replicas consistent (i.e., make sure that all updates take place in the same order on all replicas). Typically, there will be one Chubby cell per datacenter.

In addition to providing locks, Chubby is designed for managing relatively small amounts of data: items such as system configuration and state information for various services. Chubby provides clients with a namespace of files & directories. Every file or directory can be treated as a lock file and every file may, of course, be used to store data. The name of a lock is the name of its file: its hierarchical pathname. The interface to Chubby is not that of a native file system. There is no kernel module; and client software communicates with Chubby via an API that sends remote procedure calls to the Chubby master.

File operations are somewhat different from those offered by conventional file systems. Files can be read and written only in their entirety: there are no seek or byte-range read and write operations. When a file is opened by the client, it is downloaded and a lease for that file is established. In this way, Chubby keeps track of which clients have cached copies of a file. All writes from a client must be sent to the Chubby master (we have a write-through cache). Chubby then sends invalidations to all clients that have cached copies.

Locks are advisory and can be either exclusive (one writer) or not (multiple readers). A client can send a request to acquire a lock for a file, release it, and also assign and check sequence numbers for a lock. Clients can also subscribe to receive **events** for an open file. These events include notification of file modification, the creation of directories or files under an open directory, and lock acquisition.

Chubby is designed primarily for managing coarse-grained locks. **Fine-grained locks** are locks that are generally used for a small object, such as a row in a table of a database. They are generally held held for a short duration, seconds or less. **Coarse-grained locks** typically control larger structures, such as an entire table or an entire database. They might be held for hours or days. They are acquired rarely compared to fine-grained locks. Hence, a server that is designed for coarse-grained locks can handle more clients.

Even though Chubby uses the term “coarse-grained”, it doesn’t exactly fit a “pure” coarse/fine grained lock model, but that model does not necessarily work well in real systems. In theory, the concept of a coarse-grained lock is to grant a process a lock for a large pool of objects of a particular class and then have that process be the lock manager for those objects. This is essentially what Chubby does but Chubby doesn’t keep track of all the objects. Instead, the top (coarse) level is a set of services, such as Bigtable tables, a GFS file system, or Pregel frameworks. Chubby allows them to ensure there is a single master and to lock any critical resources that might be shared among these applications. Those applications then handle the fine-grained aspect of giving out locks for data blocks, table cells, or synchronizing communication barriers.

Because Chubby does not hold huge amounts of data but may serve thousands of clients, all Chubby data is cached in memory. For fault tolerance, all data is written through to the disk when modified, and is propagated to replicas via the Paxos consensus algorithm to ensure consistent ordering. The entire database is backed up to the Google File System (GFS) every few hours to ensure that critical configuration information is saved even if all replicas die.

**Chubby’s fault-tolerant locking makes it a good service for leader election**. If a group of processes wants to elect a leader, they can each request a lock on the same Chubby file. Whoever gets the lock is the leader (for the lease period of the lock).
