## Components

<img width="1710" height="928" alt="image" src="https://github.com/user-attachments/assets/25a569e2-5d90-4334-9442-9ba721328302" />
Ref: https://gear.hermygong.com/p/seaweeds/

### Volume Server

In SeaweedFS, a volume is a single file consisting of many small files. When a master server starts, it sets the volume file maximum size to 30GB (see: -volumeSizeLimitMB). At volume server initialization, it will create 8 of these volumes (see: -max).

Each volume has its own TTL and replication.

Ref: https://github.com/seaweedfs/seaweedfs/wiki/Components

### Volume Files Structure

<img width="1104" height="1226" alt="image" src="https://github.com/user-attachments/assets/194e12d8-78fe-403d-b6af-346baec95f85" />

<img width="1944" height="774" alt="image" src="https://github.com/user-attachments/assets/3fcd6b47-609d-4463-9e28-2aa8cdd2e706" />

Ref: https://github.com/seaweedfs/seaweedfs/wiki/Volume-Files-Structure

### S3 changes

<img width="836" height="378" alt="image" src="https://github.com/user-attachments/assets/21b8445f-10de-443b-a1d3-7816c2ec7d02" />

<img width="836" height="568" alt="image" src="https://github.com/user-attachments/assets/5a972039-601c-4fe8-b407-55e7920be026" />


Ref: SeaweedFS S3 API in 2025: Enterprise‑grade security and control - Chris Lu, SeaweedFS KubeCon
