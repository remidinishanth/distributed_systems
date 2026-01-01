<img width="1134" height="628" alt="image" src="https://github.com/user-attachments/assets/bb98c875-7406-4e97-b825-548c6c4a6c5b" />

Ceph is designed to be extremely scalable; it is built upon the **Reliable Autonomic Distributed Object Store (RADOS)**, a self-healing, self-managing storage layer that handles the fundamental complexity of data replication, failure detection, and recovery. 

Unlike traditional architectures that rely on centralized controller nodes—which often become performance bottlenecks or single points of failure — Ceph employs a calculated placement algorithm known as **CRUSH (Controlled Replication Under Scalable Hashing)** to distribute data across a heterogeneous cluster.

CRUSH enables Ceph clients to communicate directly with OSDs, bypassing the need for a centralized server or broker. 


Ceph uniquely delivers object, block, and file storage in one unified system.

<img width="1118" height="786" alt="image" src="https://github.com/user-attachments/assets/330ed975-85ac-412b-8575-71002b7f613c" />



## HOW :: Data is Storage Inside Ceph Cluster

<img width="1079" height="806" alt="image" src="https://github.com/user-attachments/assets/e7346264-79aa-4838-baa7-8e107b941d3d" />

Ref: https://ceph.io/en/news/blog/2014/how-data-is-stored-in-ceph-cluster/
