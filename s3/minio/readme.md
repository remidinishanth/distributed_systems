MinIO just does one thing - Object storage for Private cloud

<img width="1105" height="588" alt="image" src="https://github.com/user-attachments/assets/b9fdf6eb-b52c-460d-89e0-fc90a2762f2f" />

<img width="1142" height="499" alt="image" src="https://github.com/user-attachments/assets/ea1e33e0-cac3-47b9-b975-987243174e59" />

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/82863b5c-7278-4c28-981b-94dd71b614db" />


All the nodes running distributed MinIO setup are recommended to be homogeneous, i.e. same operating system, same number of drives and same network interconnects.


Start distributed MinIO instance on n nodes with m drives each mounted at `/export1` to `/exportm` (pictured below), by running this command on all the `n` nodes:

<img width="8000" height="4500" alt="image" src="https://github.com/user-attachments/assets/19b44e38-a8c8-4daa-89ad-3d9a6854ecdd" />


Ref: https://github.com/minio/minio/blob/master/docs/distributed/README.md

## MinIO Server Pools

![serverpools](https://github.com/user-attachments/assets/a25b361c-e253-4c06-983a-e95b4d0ae464)

A server pool is a set of minio server nodes which pool their drives and resources, creating a unit of expansion. All nodes in a server pool share their hardware resources in an isolated namespace. 

## Erasure coding

Data shards contain a portion of a given object. Parity shards contain a mathematical representation of the object used for rebuilding Data shards.

Erasure Coding (EC): MinIO employs Reed-Solomon erasure coding to break objects into data and parity shards, distributing them across multiple drives to ensure fault tolerance. For example, in a 16-drive setup, data can be split into 12 data shards and 4 parity shards, allowing the system to rebuild data even if up to 4 drives fail.

<img width="1030" height="540" alt="image" src="https://github.com/user-attachments/assets/45f2609b-43c4-4988-99d7-b6a2c173d17a" />

The value K here constitutes the read quorum for the deployment. The erasure set must therefore have at least K healthy drives in the erasure set to support read operations.

<img width="960" height="540" alt="image" src="https://github.com/user-attachments/assets/83ffabbc-95dc-4ac7-8fa7-832d27d59b87" />

Ref: https://blog.min.io/erasure-coding-vs-raid/

Not only does MinIO erasure coding protect objects against data loss in the event that multiple drives and nodes fail, MinIO also protects and heals at the object level. 
* The ability to heal one object at a time is a dramatic advantage over systems such as RAID that heal at the volume level.
* A corrupt object could be restored in MinIO in seconds vs. hours in RAID.


MinIO protects against `BitRot`, or silent data corruption, which can have many different causes such as power current spikes, bugs in disk firmware and even simply aging drives. 
* MinIO uses the `HighwayHash` algorithm to compute a hash on read and verify it on write from the application, across the network and to the storage media.
* This process is highly efficient - it can achieve hashing speeds over 10 GB/sec on a single core on Intel CPUs - and has minimal impact on normal read/write operations across the erasure set. https://github.com/google/highwayhash

## Put and Get Operation

### Storing an Object (The PUT Request)
When a client sends an object to the cluster, MinIO follows a specific sequence to ensure data is stored safely and evenly distributed.

* Step 1: Hashing: The object name is processed by a deterministic hash function to create a unique hash value.

* Step 2: Drive Selection: A modulus function is applied to that hash value. The result determines the specific set of drives (erasure set) where the data will live.

* Step 3: Erasure Coding: Simultaneously, the Erasure Code Engine processes the object data. It breaks the object into:
  - Data blocks: The actual content.
  - Parity blocks: Redundancy data for recovery.

* Step 4: Writing: These blocks are written to the prescribed drives.

Note: MinIO uses `SipHash` for this process. This algorithm ensures that objects are distributed evenly across all drives, resulting in near-uniform disk utilization.

<img width="1606" height="929" alt="image" src="https://github.com/user-attachments/assets/25a0614e-a95f-41e3-adcf-e5278acca6f0" />

### Retrieving an Object (The GET Request)
To retrieve data, MinIO reverses the logic used during the write process.

* Step 1: Location Calculation: The client requests the file by name. MinIO runs the name through the same hash and modulus functions used during the PUT request to identify the correct drives immediately.

* Step 2: Retrieval: The system reads the object shards (blocks) from those specific drives.

* Step 3: Reassembly & Verification: The shards are passed back through the Erasure Code Engine. The engine reassembles the original object and verifies its integrity ("sanity check") to ensure no corruption occurred.

* Step 4: Delivery: The verified object is sent back to the client.

## Site to Site Replication

<img width="827" height="427" alt="image" src="https://github.com/user-attachments/assets/dbb719ec-3363-45a8-922e-ebcc9a63dd14" />
