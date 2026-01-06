<img width="502" height="516" alt="image" src="https://github.com/user-attachments/assets/794cefcd-0950-4418-bef8-30408126b405" />

## Components

<img width="1710" height="928" alt="image" src="https://github.com/user-attachments/assets/25a569e2-5d90-4334-9442-9ba721328302" />
Ref: https://gear.hermygong.com/p/seaweeds/

## Blob storage

<img width="1158" height="557" alt="image" src="https://github.com/user-attachments/assets/a7c602b1-bc3d-4729-9e8f-d976712db515" />

<img width="1184" height="603" alt="image" src="https://github.com/user-attachments/assets/82bb0567-4294-47dc-99b4-dfe0ca190723" />

<img width="1180" height="622" alt="image" src="https://github.com/user-attachments/assets/b431b55a-1d41-4a39-a116-50192cb66558" />


### Write

<img width="1554" height="765" alt="image" src="https://github.com/user-attachments/assets/fba8e5cd-88d5-4b19-81eb-ef989d5dd1bd" />

<img width="1024" height="590" alt="image" src="https://github.com/user-attachments/assets/cbe79b6d-8e2f-47c0-b4e4-a827891e458d" />


### Read

<img width="1554" height="752" alt="image" src="https://github.com/user-attachments/assets/4b15a1cc-f7f4-4671-b42c-d2529415fccf" />

<img width="1024" height="709" alt="image" src="https://github.com/user-attachments/assets/61d72d8c-f829-415d-80e9-ae28b70a7b9f" />


## File Storage

<img width="1024" height="666" alt="image" src="https://github.com/user-attachments/assets/15a14a26-c4e9-4723-a1b8-1a9a18d5ffec" />
Ref: https://www.a-programmer.top/2021/06/19/SeaweedFS%E5%88%9D%E6%8E%A2/

<img width="1180" height="644" alt="image" src="https://github.com/user-attachments/assets/a78f14d7-8e56-4816-acab-2d00dd4740a4" />


<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/f530b26e-f825-4c44-ac3e-a048feb44164" />

## Design Philosophy

<img width="1296" height="433" alt="image" src="https://github.com/user-attachments/assets/5f48b0c7-4ab5-476d-a47f-814349f29d88" />


## High Availability

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/efef40ca-d4b9-4d23-ad2e-7103cb2d2275" />

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/d51248cb-6c91-4203-bdcb-ebfb3f55006b" />

### Volume Server

<img width="1113" height="547" alt="image" src="https://github.com/user-attachments/assets/f38a01a2-9da2-4040-8395-616f163e674b" />

<img width="856" height="569" alt="image" src="https://github.com/user-attachments/assets/bf432e51-7953-4b85-a1fc-6c0f4cdf7498" />

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
