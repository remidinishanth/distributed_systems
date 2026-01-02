## Ceph: A Scalable, High Performance Distributed File System

The primary goals of the architecture are scalability (to hundreds of petabytes and beyond), performance, and reliability.
<img width="968" height="719" alt="image" src="https://github.com/user-attachments/assets/0e3dbef2-7b56-4400-9dac-ef2640ed59fe" />

<img width="1134" height="628" alt="image" src="https://github.com/user-attachments/assets/bb98c875-7406-4e97-b825-548c6c4a6c5b" />

Ceph is designed to be extremely scalable; it is built upon the **Reliable Autonomic Distributed Object Store (RADOS)**, a self-healing, self-managing storage layer that handles the fundamental complexity of data replication, failure detection, and recovery. 


Unlike traditional architectures that rely on centralized controller nodes—which often become performance bottlenecks or single points of failure — Ceph employs a calculated placement algorithm known as **CRUSH (Controlled Replication Under Scalable Hashing)** to distribute data across a heterogeneous cluster.

CRUSH enables Ceph clients to communicate directly with OSDs, bypassing the need for a centralized server or broker. 


Ceph uniquely delivers object, block, and file storage in one unified system.


<img width="1920" height="1342" alt="image" src="https://github.com/user-attachments/assets/52f44c4f-affb-4e5a-abe3-b839a81428a1" />
Ref: https://canonical.com/blog/ceph-storage-on-ubuntu-an-overview


<img width="655" height="292" alt="image" src="https://github.com/user-attachments/assets/419b7475-7bea-499f-b8ac-38df281d0c4a" />


<img width="1118" height="786" alt="image" src="https://github.com/user-attachments/assets/330ed975-85ac-412b-8575-71002b7f613c" />


Ceph is designed to be scalable and to have no single point of failure. 

* Object Storage Daemon (OSD) - OSDs manage data, interact with logical disks
* Monitor (MON) - Manages cluster state(monitor map, the manager map, the OSD map, and the CRUSH map) https://docs.ceph.com/en/reef/architecture/#cluster-map 
* Manager (MGR) - Provides additional features like external monitoring, dashboard
* Metadata Servers (MDS) - The Ceph metadata server daemon must be running in any Ceph cluster that runs the CephFS file system

<img width="704" height="404" alt="image" src="https://github.com/user-attachments/assets/bc7e4578-88e2-4c2c-9ff1-c4fc15076ac1" />


<img width="1040" height="678" alt="image" src="https://github.com/user-attachments/assets/e14e8fcd-8765-404b-b654-98885133cf1c" />


<img width="530" height="350" alt="image" src="https://github.com/user-attachments/assets/8508d331-a67c-46cd-a728-0e8c6dfd8256" />


<img width="1105" height="613" alt="image" src="https://github.com/user-attachments/assets/ffb9acad-56e2-4913-81cb-9c10ab4c1674" />


### What are OSD

<img width="400" height="126" alt="image" src="https://github.com/user-attachments/assets/c367b338-fae1-477f-b9e2-5cdc1d35fcb6" />

* More recent distributed file systems have adopted architectures based on object-based storage, in which conventional hard disks are replaced with intelligent object
storage devices (OSDs) which combine a CPU, network interface, and local cache with an underlying disk or RAID.
* OSDs replace the traditional block-level interface with one in which clients can read
or write byte ranges to much larger (and often variably sized) named objects, distributing low-level block allocation decisions to the devices themselves.
* Clients typically interact with a metadata server (MDS) to perform metadata operations (open, rename), while communicating directly with OSDs to perform file I/O (reads and
writes), significantly improving overall scalability.

Ref: `Ceph: A Scalable, High-Performance Distributed File System`

* The Ceph Storage Cluster receives data from Ceph Clients--whether it comes through a Ceph Block Device, Ceph Object Storage, the Ceph File System, or a custom implementation that you create by using librados.
* The data received by the Ceph Storage Cluster is stored as RADOS objects.
* Each object is stored on an Object Storage Device (this is also called an “OSD”). Ceph OSDs control read, write, and replication operations on storage drives.

<img width="1105" height="778" alt="image" src="https://github.com/user-attachments/assets/75159e46-2a0e-4e7e-8d39-b12579c26e0b" />

<img width="1105" height="829" alt="image" src="https://github.com/user-attachments/assets/5305d882-bd36-4f9a-b58e-02e51765ffc1" />

<img width="1105" height="754" alt="image" src="https://github.com/user-attachments/assets/0ccb1e3f-372f-47cd-b934-54df94520ba9" />


## Architecture

<img width="1285" height="832" alt="image" src="https://github.com/user-attachments/assets/63cd3c58-21b4-4f59-95b0-f90e82060f8f" />

<img width="849" height="517" alt="image" src="https://github.com/user-attachments/assets/7a17da7a-2024-4746-944e-0412f2176f46" />

### Pools

The Ceph storage system supports the notion of ‘Pools’, which are logical partitions for storing objects.

Ceph Clients retrieve a Cluster Map from a Ceph Monitor, and write RADOS objects to pools. The way that Ceph places the data in the pools is determined by the pool’s size or number of replicas, the CRUSH rule, and the number of placement groups in the pool.

<img width="440" height="238" alt="image" src="https://github.com/user-attachments/assets/8fee4a4a-01cf-467c-b120-17fef100948e" />


<img width="743" height="384" alt="image" src="https://github.com/user-attachments/assets/25163ead-9638-4b93-b690-4d7147e5c35d" />


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

> This layer of indirection allows Ceph to rebalance dynamically when new Ceph OSD Daemons and the underlying OSD devices come online

### CRUSH

<img width="1105" height="729" alt="image" src="https://github.com/user-attachments/assets/cfe7db7d-756a-44ba-a0db-cb84eb4921fa" />

### Summary

<img width="1079" height="806" alt="image" src="https://github.com/user-attachments/assets/e7346264-79aa-4838-baa7-8e107b941d3d" />

Ref: https://ceph.io/en/news/blog/2014/how-data-is-stored-in-ceph-cluster/


<img width="1152" height="739" alt="image" src="https://github.com/user-attachments/assets/1daf35de-9975-42b5-82f5-32ebc58106c5" />

<img width="908" height="429" alt="image" src="https://github.com/user-attachments/assets/fb3bc96b-9810-46bf-92b6-b8620539ab8f" />


<img width="2048" height="1090" alt="image" src="https://github.com/user-attachments/assets/0bb2cad5-2095-4389-946e-502c1e25eb07" />

* In this cluster, the files created (A.txt and J.txt in my diagram) are converted into several objects. These objects are then distributed into placement groups (pg) which are put into pools.
* A pool has some properties configured as how many replicas of a pg will be stored in the cluster (3 by default). Those pg will finally be physically stored into an Object Storage Daemon (OSD). An OSD stores pg (and so the objects within it) and provides access to them over the network.

Ref: https://www.dbi-services.com/blog/introduction-to-rook-ceph-for-kubernetes/

<img width="1653" height="739" alt="image" src="https://github.com/user-attachments/assets/e90b8dd7-6fd9-4ce4-b8ab-b5997354b723" />

The original paper doesn't talk about Pools

<img width="596" height="561" alt="image" src="https://github.com/user-attachments/assets/a3c549fe-567d-4913-ace1-c8097fbd2f1d" />

Ref: https://devops-insider.mygraphql.com/zh-cn/latest/ceph/ceph-mapping/ceph-mapping.html

* Objects that are stored in a Ceph cluster are put into pools.
* Pools represent logical partitions of the cluster to the outside world. For each pool a set of rules can be defined, for example, how many replications of each object must exist. The standard configuration of pools is called replicated pool.

<img width="1058" height="602" alt="image" src="https://github.com/user-attachments/assets/de3cd3bf-cf43-4ee8-9328-e1d1dd00f92a" />

### Rebalancing

<img width="1068" height="708" alt="image" src="https://github.com/user-attachments/assets/8f7ed77d-9547-4839-98e2-9ce927c74392" />


## Rados Gateway

<img width="735" height="333" alt="image" src="https://github.com/user-attachments/assets/35f8608e-14ab-4542-8518-f837b6aec65f" />

<img width="735" height="390" alt="image" src="https://github.com/user-attachments/assets/5ee6d64e-a425-4c45-9ef1-3e9209e2463d" />

## CephFS

<img width="769" height="419" alt="image" src="https://github.com/user-attachments/assets/e25c67cf-68d1-4109-969f-d160b378c0d0" />

<img width="642" height="419" alt="image" src="https://github.com/user-attachments/assets/d5f94809-3a68-4023-bc0b-1c306849b644" />

<img width="753" height="419" alt="image" src="https://github.com/user-attachments/assets/0826dd55-b1a9-46a9-a482-f2fa46ef7c3c" />

### Distributed Metadata Server

<img width="1105" height="729" alt="image" src="https://github.com/user-attachments/assets/479bf446-3020-44d3-8514-0feb594e5805" />


<img width="1105" height="637" alt="image" src="https://github.com/user-attachments/assets/cf39b001-0c27-471e-9a8a-fc105e287a5f" />

## Bluestore

For a decade, the Ceph distributed file system followed the
conventional wisdom of building its storage backend on top
of local file systems. This is a preferred choice for most distributed file systems today because it allows them to benefit
from the convenience and maturity of battle-tested code.
Ceph’s experience, however, shows that this comes at a high
price. 
* First, developing a zero-overhead transaction mechanism is challenging.
* Second, metadata performance at the local level can significantly affect performance at the distributed level.
* Third, supporting emerging storage hardware is painstakingly slow.

Ceph addressed these issues with BlueStore, a new backend designed to run directly on raw storage devices

<img width="517" height="392" alt="image" src="https://github.com/user-attachments/assets/6268352d-2f80-4cb7-9301-550cc4810fac" />


<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/028cfa19-0259-407c-be91-2aba019c6d3f" />

Ref: `File Systems Unfit as Distributed Storage Backends: Lessons from 10 Years of Ceph Evolution` paper
