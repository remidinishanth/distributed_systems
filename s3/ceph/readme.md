<img width="1134" height="628" alt="image" src="https://github.com/user-attachments/assets/bb98c875-7406-4e97-b825-548c6c4a6c5b" />

Ceph is designed to be extremely scalable; it is built upon the **Reliable Autonomic Distributed Object Store (RADOS)**, a self-healing, self-managing storage layer that handles the fundamental complexity of data replication, failure detection, and recovery. 

Unlike traditional architectures that rely on centralized controller nodes—which often become performance bottlenecks or single points of failure — Ceph employs a calculated placement algorithm known as **CRUSH (Controlled Replication Under Scalable Hashing)** to distribute data across a heterogeneous cluster.

CRUSH enables Ceph clients to communicate directly with OSDs, bypassing the need for a centralized server or broker. 


Ceph uniquely delivers object, block, and file storage in one unified system.


<img width="1920" height="1342" alt="image" src="https://github.com/user-attachments/assets/52f44c4f-affb-4e5a-abe3-b839a81428a1" />
Ref: https://canonical.com/blog/ceph-storage-on-ubuntu-an-overview

<img width="1118" height="786" alt="image" src="https://github.com/user-attachments/assets/330ed975-85ac-412b-8575-71002b7f613c" />


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

* The Ceph Storage Cluster receives data from Ceph Clients--whether it comes through a Ceph Block Device, Ceph Object Storage, the Ceph File System, or a custom implementation that you create by using librados.
* The data received by the Ceph Storage Cluster is stored as RADOS objects.
* Each object is stored on an Object Storage Device (this is also called an “OSD”). Ceph OSDs control read, write, and replication operations on storage drives.


## Architecture

<img width="1285" height="832" alt="image" src="https://github.com/user-attachments/assets/63cd3c58-21b4-4f59-95b0-f90e82060f8f" />

<img width="849" height="517" alt="image" src="https://github.com/user-attachments/assets/7a17da7a-2024-4746-944e-0412f2176f46" />

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/028cfa19-0259-407c-be91-2aba019c6d3f" />

### Pools

The Ceph storage system supports the notion of ‘Pools’, which are logical partitions for storing objects.

Ceph Clients retrieve a Cluster Map from a Ceph Monitor, and write RADOS objects to pools. The way that Ceph places the data in the pools is determined by the pool’s size or number of replicas, the CRUSH rule, and the number of placement groups in the pool.

<img width="440" height="238" alt="image" src="https://github.com/user-attachments/assets/8fee4a4a-01cf-467c-b120-17fef100948e" />

### Placement Groups

* Tracking object placement on a per-object basis within a pool is computationally expensive at scale.
*  To facilitate high performance at scale, Ceph subdivides a pool into placement groups, assigns each individual object to a placement group, and assigns the placement group to a primary OSD.
*  If an OSD fails or the cluster re-balances, Ceph can move or replicate an entire placement group—​i.e., all of the objects in the placement groups—​without having to address each object individually. This allows a Ceph cluster to re-balance or recover efficiently.

<img width="780" height="588" alt="image" src="https://github.com/user-attachments/assets/9a5de55b-f869-4706-a559-b8babf9df5fa" />

Ref: https://docs.redhat.com/en/documentation/red_hat_ceph_storage/1.2.3/html/storage_strategies/about-placement-groups

When CRUSH assigns a placement group to an OSD, it calculates a series of OSDs—​the first being the primary. 
* Each pool has a number of placement groups (PGs) within it. CRUSH dynamically maps PGs to OSDs. When a Ceph Client stores objects, CRUSH maps each RADOS object to a PG.

<img width="770" height="224" alt="image" src="https://github.com/user-attachments/assets/6855a121-72ca-4aae-92c4-2b84047afe20" />

PGs do not own OSDs. CRUSH assigns many placement groups to each OSD pseudo-randomly to ensure that data gets distributed evenly across the cluster. 

<img width="400" height="126" alt="image" src="https://github.com/user-attachments/assets/c367b338-fae1-477f-b9e2-5cdc1d35fcb6" />

<img width="1079" height="806" alt="image" src="https://github.com/user-attachments/assets/e7346264-79aa-4838-baa7-8e107b941d3d" />

Ref: https://ceph.io/en/news/blog/2014/how-data-is-stored-in-ceph-cluster/
