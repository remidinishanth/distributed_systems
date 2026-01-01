<img width="1134" height="628" alt="image" src="https://github.com/user-attachments/assets/bb98c875-7406-4e97-b825-548c6c4a6c5b" />

Ceph is designed to be extremely scalable; it is built upon the **Reliable Autonomic Distributed Object Store (RADOS)**, a self-healing, self-managing storage layer that handles the fundamental complexity of data replication, failure detection, and recovery. 

Unlike traditional architectures that rely on centralized controller nodes—which often become performance bottlenecks or single points of failure — Ceph employs a calculated placement algorithm known as **CRUSH (Controlled Replication Under Scalable Hashing)** to distribute data across a heterogeneous cluster.

CRUSH enables Ceph clients to communicate directly with OSDs, bypassing the need for a centralized server or broker. 


Ceph uniquely delivers object, block, and file storage in one unified system.


<img width="1920" height="1342" alt="image" src="https://github.com/user-attachments/assets/52f44c4f-affb-4e5a-abe3-b839a81428a1" />
Ref: https://canonical.com/blog/ceph-storage-on-ubuntu-an-overview

<img width="1118" height="786" alt="image" src="https://github.com/user-attachments/assets/330ed975-85ac-412b-8575-71002b7f613c" />

<img width="849" height="517" alt="image" src="https://github.com/user-attachments/assets/7a17da7a-2024-4746-944e-0412f2176f46" />



Ceph is designed to be scalable and to have no single point of failure. 

* Object Storage Daemon (OSD) - OSDs manage data, interact with logical disks
* Monitor (MON) - Manages cluster state(monitor map, the manager map, the OSD map, and the CRUSH map) https://docs.ceph.com/en/reef/architecture/#cluster-map 
* Manager (MGR) - Provides additional features
* Metadata Servers (MDS) - The Ceph metadata server daemon must be running in any Ceph cluster that runs the CephFS file system

<img width="530" height="350" alt="image" src="https://github.com/user-attachments/assets/8508d331-a67c-46cd-a728-0e8c6dfd8256" />


### What are OSD

* More recent distributed file systems have adopted architectures based on object-based storage, in which conventional hard disks are replaced with intelligent object
storage devices (OSDs) which combine a CPU, network interface, and local cache with an underlying disk or RAID.
* OSDs replace the traditional block-level interface with one in which clients can read
or write byte ranges to much larger (and often variably sized) named objects, distributing low-level block allocation decisions to the devices themselves.
* Clients typically interact with a metadata server (MDS) to perform metadata operations (open, rename), while communicating directly with OSDs to perform file I/O (reads and
writes), significantly improving overall scalability.

Ref: `Ceph: A Scalable, High-Performance Distributed File System paper`


## Architecture

<img width="1285" height="832" alt="image" src="https://github.com/user-attachments/assets/63cd3c58-21b4-4f59-95b0-f90e82060f8f" />


## HOW :: Data is Storage Inside Ceph Cluster

<img width="1079" height="806" alt="image" src="https://github.com/user-attachments/assets/e7346264-79aa-4838-baa7-8e107b941d3d" />

Ref: https://ceph.io/en/news/blog/2014/how-data-is-stored-in-ceph-cluster/
