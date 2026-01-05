## Components

<img width="1710" height="928" alt="image" src="https://github.com/user-attachments/assets/25a569e2-5d90-4334-9442-9ba721328302" />
Ref: https://gear.hermygong.com/p/seaweeds/

## Blob storage

<img width="1554" height="732" alt="image" src="https://github.com/user-attachments/assets/8a9df9d6-f6a9-428d-b5a2-d47fb3dc243a" />

<img width="1554" height="835" alt="image" src="https://github.com/user-attachments/assets/5d23436b-e7cd-474d-95f7-ac6ca5d481cf" />

### Write

<img width="1554" height="765" alt="image" src="https://github.com/user-attachments/assets/fba8e5cd-88d5-4b19-81eb-ef989d5dd1bd" />

### Read

<img width="1554" height="752" alt="image" src="https://github.com/user-attachments/assets/4b15a1cc-f7f4-4671-b42c-d2529415fccf" />

## File Storage

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/f530b26e-f825-4c44-ac3e-a048feb44164" />

## Design Philosophy

<img width="1296" height="433" alt="image" src="https://github.com/user-attachments/assets/5f48b0c7-4ab5-476d-a47f-814349f29d88" />


## High Availability

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/efef40ca-d4b9-4d23-ad2e-7103cb2d2275" />

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/d51248cb-6c91-4203-bdcb-ebfb3f55006b" />

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
