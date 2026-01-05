## Components

<img width="1710" height="928" alt="image" src="https://github.com/user-attachments/assets/25a569e2-5d90-4334-9442-9ba721328302" />
Ref: https://gear.hermygong.com/p/seaweeds/

### Volume Server

In SeaweedFS, a volume is a single file consisting of many small files. When a master server starts, it sets the volume file maximum size to 30GB (see: -volumeSizeLimitMB). At volume server initialization, it will create 8 of these volumes (see: -max).

Each volume has its own TTL and replication.

Ref: https://github.com/seaweedfs/seaweedfs/wiki/Components
