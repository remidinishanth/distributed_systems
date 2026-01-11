<img width="502" height="516" alt="image" src="https://github.com/user-attachments/assets/794cefcd-0950-4418-bef8-30408126b405" />

## Components

<img width="1710" height="928" alt="image" src="https://github.com/user-attachments/assets/25a569e2-5d90-4334-9442-9ba721328302" />
Ref: https://gear.hermygong.com/p/seaweeds/

## Blob storage

<img width="1158" height="557" alt="image" src="https://github.com/user-attachments/assets/a7c602b1-bc3d-4729-9e8f-d976712db515" />

<img width="1184" height="603" alt="image" src="https://github.com/user-attachments/assets/82bb0567-4294-47dc-99b4-dfe0ca190723" />

<img width="1180" height="622" alt="image" src="https://github.com/user-attachments/assets/b431b55a-1d41-4a39-a116-50192cb66558" />

<img width="1136" height="476" alt="image" src="https://github.com/user-attachments/assets/424ce6e2-017b-41b9-ace1-1a953bb231db" />

Other Blobstore operations
<img width="1136" height="547" alt="image" src="https://github.com/user-attachments/assets/e900b20c-8c09-42fe-8286-64735ee21edf" />


### Write

<img width="1554" height="765" alt="image" src="https://github.com/user-attachments/assets/fba8e5cd-88d5-4b19-81eb-ef989d5dd1bd" />

<img width="1024" height="590" alt="image" src="https://github.com/user-attachments/assets/1d227518-fcdc-40fb-ac84-889e77ef86a4" />


### Read

<img width="1554" height="752" alt="image" src="https://github.com/user-attachments/assets/4b15a1cc-f7f4-4671-b42c-d2529415fccf" />

<img width="1024" height="709" alt="image" src="https://github.com/user-attachments/assets/735bfe02-ff1a-43cd-b4c1-47a91003b3aa" />


## File Storage

<img width="1157" height="577" alt="image" src="https://github.com/user-attachments/assets/96678169-fcde-4f50-a97d-aeba4632cd9f" />

Filer Architecture

<img width="1157" height="555" alt="image" src="https://github.com/user-attachments/assets/cffd90ce-2a70-4bb5-887a-92734a82bf4f" />

<img width="1090" height="508" alt="image" src="https://github.com/user-attachments/assets/3638b85c-4cba-4e32-b7fb-6fe2c1760313" />

<img width="1024" height="666" alt="image" src="https://github.com/user-attachments/assets/15a14a26-c4e9-4723-a1b8-1a9a18d5ffec" />
Ref: https://www.a-programmer.top/2021/06/19/SeaweedFS%E5%88%9D%E6%8E%A2/

<img width="1180" height="644" alt="image" src="https://github.com/user-attachments/assets/a78f14d7-8e56-4816-acab-2d00dd4740a4" />

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/f530b26e-f825-4c44-ac3e-a048feb44164" />

<img width="1113" height="498" alt="image" src="https://github.com/user-attachments/assets/ddb0eccf-0dbd-4c37-a8af-6cfa2d413a06" />

### Filer Store Data Model

<img width="2230" height="1056" alt="image" src="https://github.com/user-attachments/assets/67623e7a-9a55-4448-a9ee-7138d8691486" />

<img width="281" height="344" alt="image" src="https://github.com/user-attachments/assets/b6fd9c41-f1f0-4fc3-af2d-369f928d13d6" />


## Volume Server

<img width="1113" height="547" alt="image" src="https://github.com/user-attachments/assets/f38a01a2-9da2-4040-8395-616f163e674b" />

Volume
<img width="1136" height="450" alt="image" src="https://github.com/user-attachments/assets/b3c9bd8c-2b69-47e2-a382-2a4aab239613" />


<img width="856" height="569" alt="image" src="https://github.com/user-attachments/assets/bf432e51-7953-4b85-a1fc-6c0f4cdf7498" />

In SeaweedFS, a volume is a single file consisting of many small files. When a master server starts, it sets the volume file maximum size to 30GB (see: -volumeSizeLimitMB). At volume server initialization, it will create 8 of these volumes (see: -max).

Each volume has its own TTL and replication.

Ref: https://github.com/seaweedfs/seaweedfs/wiki/Components

### Volume Files Structure

<img width="1104" height="1226" alt="image" src="https://github.com/user-attachments/assets/194e12d8-78fe-403d-b6af-346baec95f85" />

<img width="1944" height="774" alt="image" src="https://github.com/user-attachments/assets/3fcd6b47-609d-4463-9e28-2aa8cdd2e706" />

Ref: https://github.com/seaweedfs/seaweedfs/wiki/Volume-Files-Structure

## Architecture

<img width="2020" height="1280" alt="image" src="https://github.com/user-attachments/assets/0a79d086-f742-431e-9d37-f62f5ce38bcf" />

## Design Philosophy

<img width="1296" height="433" alt="image" src="https://github.com/user-attachments/assets/5f48b0c7-4ab5-476d-a47f-814349f29d88" />

## High Availability

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/efef40ca-d4b9-4d23-ad2e-7103cb2d2275" />

In Master, How Raft is Used
* Leader Election: Multiple master servers form a Raft cluster to elect a leader. Only the leader can assign new volume IDs.
* Volume ID Assignment: When a new volume needs to be created, the leader:
  - Gets the current max volume ID
  - Increments it
  - Replicates this new max via Raft to ensure all masters agree

Also master manages, all these
* Assign file ID	- `Leader only (proxied if needed)`
* Volume creation	- `Leader only`
* NextVolumeId (Raft write)	- `Leader only + barrier`
* Volume lookup	Leader uses local topology, non-leader queries too
* Client connections	- Any master, but redirected to leader

<img width="1554" height="742" alt="image" src="https://github.com/user-attachments/assets/d51248cb-6c91-4203-bdcb-ebfb3f55006b" />


## Replication

<img width="1111" height="603" alt="image" src="https://github.com/user-attachments/assets/4bd33da7-c396-4de4-b775-18a834225894" />

<img width="1157" height="497" alt="image" src="https://github.com/user-attachments/assets/22d6410c-b393-46de-84e4-fa412c26629b" />

## Erasure Coding

* SeaweedFS uses Reed-Solomon erasure coding with a default 10+4 scheme (10 data shards + 4 parity shards = 14 total shards). 
* This allows you to lose up to 4 volume servers and still recover your data.
* Only volumes with this fullness ratio 80% or higher will be erasure coded, configurable

```
/*

Steps to apply erasure coding to .dat .idx files
0. ensure the volume is readonly
1. client call VolumeEcShardsGenerate to generate the .ecx and .ec00 ~ .ec13 files
2. client ask master for possible servers to hold the ec files
3. client call VolumeEcShardsCopy on above target servers to copy ec files from the source server
4. target servers report the new ec files to the master
5.   master stores vid -> [14]*DataNode
6. client checks master. If all 14 slices are ready, delete the original .idx, .idx files

*/
```

Ref: seaweedfs/weed/server/volume_grpc_erasure_coding.go

## S3 changes

<img width="836" height="378" alt="image" src="https://github.com/user-attachments/assets/21b8445f-10de-443b-a1d3-7816c2ec7d02" />

<img width="836" height="568" alt="image" src="https://github.com/user-attachments/assets/5a972039-601c-4fe8-b407-55e7920be026" />

Ref: SeaweedFS S3 API in 2025: Enterprise‑grade security and control - Chris Lu, SeaweedFS KubeCon

Changes related to this `S3 data path skips filer`  https://github.com/seaweedfs/seaweedfs/pull/7481
* Check this file `weed/s3api/s3api_object_handlers_put.go` previously it used Filer Proxy `proxyReq, err := http.NewRequest(http.MethodPut, uploadUrl, body)`, now S3 api directly talks to volume server

<img width="1000" height="938" alt="image" src="https://github.com/user-attachments/assets/49020b3f-1779-453c-8d6d-cf249b5a2ae0" />

Change

```mermaid

flowchart TB
    subgraph OLD["Before v4.01"]
        direction TB
        C1[S3 Client] --> S1[S3 API Server]
        S1 -->|"HTTP proxy<br/>ALL data + metadata"| F1[Filer]
        F1 -->|"Read/Write data"| V1[Volume Server]
        
        style F1 fill:#cc4444,color:#fff
        style S1 fill:#4466aa,color:#fff
    end

    subgraph NEW["NEW Architecture (v4.01+)"]
        direction TB
        C2[S3 Client] --> S2[S3 API Server]
        S2 -->|"gRPC<br/>metadata only"| F2[Filer]
        S2 -->|"HTTP direct<br/>data streaming"| V2[Volume Server]
        
        style F2 fill:#44aa66,color:#fff
        style S2 fill:#4466aa,color:#fff
        style V2 fill:#44aa66,color:#fff
    end
```

### Write Path

```mermaid
sequenceDiagram
    participant Client as S3 Client
    participant S3API as S3 API Server
    participant Filer as Filer (gRPC)
    participant Volume as Volume Server

    Note over Client,Volume: PUT Object - Direct Volume Access

    Client->>S3API: PUT /bucket/key (data)

    rect rgba(144, 238, 144, 0.3)
        Note right of S3API: Step 1: Get Volume Assignment
        S3API->>Filer: AssignVolume (gRPC)
        Filer-->>S3API: {volumeId, fileId, url, JWT}
    end

    rect rgba(173, 216, 230, 0.3)
        Note right of S3API: Step 2: Upload Data DIRECTLY
        loop For each 8MB chunk
            S3API->>Volume: POST http://volume:8080/{fid} (chunk data + JWT)
            Volume-->>S3API: {size, eTag, fid}
        end
    end

    rect rgba(255, 182, 193, 0.3)
        Note right of S3API: Step 3: Save Metadata Only
        S3API->>Filer: CreateEntry (gRPC)
        Note over Filer: Stores: chunks[], size,<br/>ETag, SSE metadata,<br/>user metadata, etc.
        Filer-->>S3API: OK
    end

    S3API-->>Client: 200 OK + ETag
```

### Read Path

```mermaid
sequenceDiagram
    participant Client as S3 Client
    participant S3API as S3 API Server
    participant Filer as Filer (gRPC)
    participant Volume as Volume Server

    Note over Client,Volume: GET Object - Direct Volume Access

    Client->>S3API: GET /bucket/key

    rect rgba(255, 182, 193, 0.3)
        Note right of S3API: Step 1: Fetch Metadata Only
        S3API->>Filer: LookupDirectoryEntry (gRPC)
        Filer-->>S3API: Entry{chunks[], size, attrs, extended}
    end

    rect rgba(144, 238, 144, 0.3)
        Note right of S3API: Step 2: Resolve Volume URLs
        Note over S3API: Uses FilerClient's<br/>cached vidMap<br/>(no gRPC per chunk!)
        S3API->>S3API: lookupFileIdFn(volumeId)
    end

    rect rgba(173, 216, 230, 0.3)
        Note right of S3API: Step 3: Stream Data DIRECTLY
        S3API->>Volume: GET http://volume:8080/{fid} + JWT
        Volume-->>S3API: chunk data (streaming)
        S3API-->>Client: data (streaming passthrough)
    end
```

### How Large file is written to S3?

```mermaid
sequenceDiagram
    autonumber
    participant C as Client
    participant S3 as S3 API Server
    participant F as Filer
    participant M as Master
    participant V1 as Volume Server 1
    participant V2 as Volume Server 2

    Note over V1,M: Periodically send heartbeats to Master
    
    C->>S3: PUT /bucket/object (4GB stream)
    
    Note over S3: Stream data, chunk in 8MB buffers<br/>(max 4 buffers = 32MB)
    
    rect rgb(240, 248, 255)
        Note over S3,V2: Repeat for each 8MB chunk (streaming)
        S3->>F: AssignVolume (gRPC)
        F->>M: Assign request
        M-->>F: Return Fid + Volume URL
        F-->>S3: Return Fid + Volume URL
        S3->>V1: POST chunk data (HTTP)
        V1->>V2: Replicate
        Note over V1,V2: Strong consistency:<br/>response after replication completes
        V2-->>V1: Ack
        V1-->>S3: Return chunk size + ETag
    end
    
    Note over S3: All chunks uploaded,<br/>collect FileChunk metadata
    
    S3->>F: CreateEntry (gRPC)<br/>(path, chunks[], attributes)
    F-->>S3: Success
    S3-->>C: 200 OK + ETag
```
