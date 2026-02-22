# MinIO Architecture & Object Storage Deep Dive

MinIO just does one thing - **Object storage for Private cloud**

<img width="1105" height="588" alt="image" src="https://github.com/user-attachments/assets/b9fdf6eb-b52c-460d-89e0-fc90a2762f2f" />

<img width="1326" height="744" alt="image" src="https://github.com/user-attachments/assets/226c6dad-f468-4140-998d-86641a9115eb" />

<img width="689" height="241" alt="image" src="https://github.com/user-attachments/assets/a5f7c369-978a-4944-94d7-46ce8c87c160" />

<img width="1142" height="499" alt="image" src="https://github.com/user-attachments/assets/ea1e33e0-cac3-47b9-b975-987243174e59" />

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/82863b5c-7278-4c28-981b-94dd71b614db" />

This document provides a comprehensive overview of MinIO's architecture, how it stores objects, distributes data across servers, and retrieves objects.

## Table of Contents

1. [Core Principles](#core-principles)
2. [High-Level Architecture](#high-level-architecture)
3. [System Components](#system-components)
4. [Object Storage Internals](#object-storage-internals)
5. [Erasure Coding](#erasure-coding)
6. [PUT Operation (Storing Objects)](#put-operation)
7. [GET Operation (Retrieving Objects)](#get-operation)
8. [Distributed Architecture](#distributed-architecture)
9. [Server Pools](#server-pools)
10. [Code Architecture](#code-architecture)
11. [Healing & Self-Recovery](#healing)
12. [Advanced Features](#advanced-features)
13. [Deprecated Gateway](#deprecated-gateway)
14. [Real-World Metadata Examples](#real-world-examples)

---

## Core Principles

MinIO is built on several fundamental design principles:

- **Metadata-Free Design**: No centralized metadata database. Object metadata is stored locally on each disk alongside the data (in `xl.meta` files). This eliminates the metadata bottleneck and prevents cluster-wide failures.
- **Shared-Nothing Architecture**: Each node operates independently. Data is distributed and scattered across multiple nodes and disks, with no single point of failure.
- **S3 Compatibility**: Full S3 API compatibility allows seamless migration from AWS S3 or other S3-compatible systems.
- **Erasure Coding + Bitrot Protection**: Multi-level data protection using Reed-Solomon erasure coding and HighwayHash checksums. Even if you lose more than half of your hard drives, you can still recover data. (N/2)-1 node failure is allowed in distributed mode.
- **Rebalance-Free Expansion**: Add new pools without rebalancing existing data.
- **No Master Node**: All nodes are peers in a decentralized architecture using distributed locking (dsync).

### How MinIO Compares to Legacy Storage
* Minio adopts a metadata-free database design for high performance, avoiding the metabase becoming a performance bottleneck for the entire system, and limiting failures to a single cluster, so that no other clusters are involved.
* Minio is also fully compatible with the S3 interface, so it can also be used as a gateway to provide S3 access to the outside world.
* Use both Minio Erasure code and checksum to prevent hardware failures. Even if you lose more than half of your hard drive, you can still recover from it. (N/2)-1 node failure is also allowed in the distribution

<img width="1181" height="619" alt="image" src="https://github.com/user-attachments/assets/82ac9b76-b73d-4bf0-a084-bcdc3ceaac49" />

<img width="1209" height="672" alt="image" src="https://github.com/user-attachments/assets/3bf05817-6bb7-4d08-86df-db43fe8b0f71" />

<img width="1104" height="676" alt="image" src="https://github.com/user-attachments/assets/e82fa817-cbfa-4e66-aac9-1c06b528884f" />

### Legacy Object Storage Architecture

<img width="1591" height="863" alt="image" src="https://github.com/user-attachments/assets/86218141-aa98-478d-9b7b-02a18b6faf71" />

---

## High-Level Architecture


```
┌─────────────────────────────────────────────────────────────────┐
│                      Client (S3 API)                            │
│              PUT/GET/DELETE/LIST Requests                       │
└────────────────────────┬────────────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────────────┐
│                  HTTP Server & Router                           │
│         (Authentication, Throttling, Compression)              │
└────────────────────────┬────────────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────────────┐
│              API Handlers (GET/PUT/DELETE/LIST)                │
└────────────────────────┬────────────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────────────┐
│            ObjectLayer Interface (Abstraction)                  │
└────────────────────────┬────────────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────────────┐
│         erasureServerPools (Multi-Pool Manager)                │
│  - Weighted pool selection based on available space            │
│  - Pool expansion & decommissioning                            │
└────────────────────────┬────────────────────────────────────────┘
                         │
        ┌────────────────┼────────────────┐
        │                │                │
┌───────▼────┐   ┌───────▼────┐   ┌──────▼──────┐
│  Pool 1    │   │  Pool 2    │   │  Pool N     │
│ erasureSets│   │ erasureSets│   │ erasureSets │
└───────┬────┘   └───────┬────┘   └──────┬──────┘
        │                │                │
     ┌──┴──┐          ┌──┴──┐          ┌──┴──┐
     │Set1 │          │Set1 │          │Set1 │
     └──┬──┘          └──┬──┘          └──┬──┘
        │
┌───────▼──────────────────────────────────────────────────────┐
│  erasureObjects (Single Set - Erasure Coding Logic)         │
│  - Reed-Solomon EC:M+N encoding/decoding                    │
│  - Quorum-based read/write                                  │
│  - Object-level healing                                     │
└───────┬──────────────────────────────────────────────────────┘
        │
┌───────▼──────────────────────────────────────────────────────┐
│  StorageAPI Interface (Local & Remote Disk I/O)             │
│  - xlStorage (local), storageRESTClient (remote)            │
└───────┬──────────────────────────────────────────────────────┘
        │
   ┌────┴─────────────────────┬──────────────────┐
   │                          │                  │
┌──▼──┐ ┌──────┐ ┌──────────┐│ ┌──────────────┐
│Disk1│ │Disk2 │ │  Disk..  ││ │  Disk16      │
│     │ │      │ │          ││ │              │
│     │ │      │ │          ││ │              │
└─────┘ └──────┘ └──────────┘│ └──────────────┘
                             └─ (Parallel I/O)
```

---

## System Components

### 1. HTTP Server Layer
- Handles incoming S3 API requests
- Middleware chain: Auth (Signature V4), Tracing, Throttling, GZIP compression
- Routes requests to appropriate handlers

### 2. API Handler Layer
- `GetObjectHandler`: Retrieves objects
- `PutObjectHandler`: Stores objects
- `DeleteObjectHandler`: Deletes objects
- `ListObjectsHandler`: Lists bucket contents
- Multipart upload handlers

### 3. ObjectLayer Interface
Abstract interface implemented by `erasureServerPools`, `erasureSets`, and `erasureObjects`. Provides unified S3 operations.

### 4. Erasure Server Pools
- Manages multiple independent pools
- Selects pool based on available space (weighted random)
- Enables non-disruptive expansion
- Each pool has its own erasure sets

### 5. Erasure Sets
- Routes objects to correct set using consistent hash (SipHash)
- One set contains all disks in the erasure set
- Uses deterministic placement: same object name → same set always

### 6. Erasure Objects
- Core logic for encoding/decoding
- Manages read/write quorum
- Handles object-level healing
- Coordinates with StorageAPI

### 7. Storage Layer
- **xlStorage**: Local disk I/O, metadata, file operations
- **storageRESTClient**: Remote disk via REST API
- **xlStorageDiskIDCheck**: Health wrapper for disk monitoring

---

## Object Storage Internals

### On-Disk Layout

Each disk in a MinIO cluster stores data in the following structure:

```
disk1/
├── .minio.sys/
│   ├── format.json              # Cluster configuration
│   ├── config/                  # Server configuration
│   ├── buckets/                 # Bucket metadata
│   └── tmp/                     # Temporary files during writes
│
├── bucket1/
│   ├── object1/
│   │   ├── xl.meta              # Metadata (MessagePack serialized)
│   │   └── a1b2c3d4-e5f6.../   # DataDir UUID (contains data shard)
│   │       └── part.1           # Actual shard data
│   └── object2/
│       └── ...
└── bucket2/
    └── ...
```

### xl.meta File Format

The `xl.meta` file contains critical metadata in MessagePack binary format:

```
Header:
- Magic: "XL2 "
- Version: 1.3

Versions[] (Version History):
├── Type: ObjectType, DeleteType, or LegacyType
├── ObjectV2 (if ObjectType):
│   ├── VersionID: UUID (unique version identifier)
│   ├── DataDir: UUID (data directory on disk)
│   ├── ErasureAlgorithm: ReedSolomon
│   ├── ErasureM: Number of data blocks (e.g., 12)
│   ├── ErasureN: Number of parity blocks (e.g., 4)
│   ├── ErasureBlockSize: Block size for encoding (1MB default)
│   ├── ErasureIndex: This disk's shard index (0-15)
│   ├── ErasureDist: Distribution array [disk_index_0, disk_index_1, ...]
│   ├── BitrotChecksumAlgo: HighwayHash (for integrity)
│   ├── PartNumbers: Part IDs (multipart uploads)
│   ├── PartSizes: Size of each part
│   ├── Size: Total object size
│   ├── ModTime: Modification timestamp (Unix nanoseconds)
│   ├── MetaSys: System metadata (inline data flag, etc.)
│   └── MetaUsr: User metadata (Content-Type, custom headers)
```

### format.json - Cluster Configuration

Located at `.minio.sys/format.json` on each disk:

```json
{
  "version": "1",
  "format": "xl",
  "id": "deployment-uuid",
  "xl": {
    "version": "3",
    "this": "disk-uuid",
    "sets": [
      ["disk-0", "disk-1", "disk-2", ..., "disk-15"],
      ["disk-16", "disk-17", "disk-18", ..., "disk-31"]
    ],
    "distributionAlgo": "SIPMOD"
  }
}
```

**Key Fields**:
- `id`: Cluster deployment ID (shared by all disks)
- `this`: UUID of current disk
- `sets`: Array of erasure sets, each containing disk UUIDs
- `distributionAlgo`: Algorithm used for object placement (SipHash with parity consideration)

---

## Erasure Coding

### Reed-Solomon Encoding

MinIO uses Reed-Solomon erasure coding (via `klauspost/reedsolomon` library):

**Example Configuration (16 disks)**:
- EC:12+4 (12 data shards + 4 parity shards)
- Block size: 1MB default
- Can tolerate: Up to 4 disk failures per erasure set

### How It Works

**Encoding (Write)**:
```
Original File (10MB)
        │
        ▼
┌──────────────────────────────────┐
│ Split into 1MB blocks (10 blocks)│
└──────────────────────────────────┘
        │
        ▼
┌──────────────────────────────────────────────────────┐
│ For each 1MB block:                                  │
│ ├─ Split into 12 data shards (~85KB each)           │
│ └─ Generate 4 parity shards (~85KB each)            │
│                                                      │
│ Result: 16 shards per block (12 data + 4 parity)   │
└──────────────────────────────────────────────────────┘
        │
        ▼
┌──────────────────────────────────────────────────────┐
│ Write to 16 disks in parallel:                       │
│ ├─ Disk 0: shard_0 (all blocks)                     │
│ ├─ Disk 1: shard_1 (all blocks)                     │
│ ├─ Disk 12-15: parity shards                        │
│ └─ All disks: xl.meta (metadata)                    │
└──────────────────────────────────────────────────────┘
```

**Decoding (Read)**:
```
Read Request for 10MB object
        │
        ▼
┌──────────────────────────────────┐
│ Scenario 1: All 16 disks healthy │
│ ├─ Read 12 data shards           │
│ ├─ Ignore 4 parity shards        │
│ └─ Reconstruct original data     │
└──────────────────────────────────┘

┌──────────────────────────────────┐
│ Scenario 2: 2 disks dead         │
│ ├─ Read 14 available shards      │
│ ├─ Use Reed-Solomon to recover   │
│ │  missing 2 shards              │
│ └─ Reconstruct original data     │
└──────────────────────────────────┘

┌──────────────────────────────────┐
│ Scenario 3: 5+ disks dead        │
│ └─ READ FAILS (quorum lost)      │
└──────────────────────────────────┘
```

### Read/Write Quorum

**Read Quorum**: M (data blocks)
- Need M shards available to reconstruct object
- Example: EC:12+4 → need 12 out of 16 disks

**Write Quorum**:
- If parity < 50% of drives: Write Quorum = M (data blocks)
- If parity = 50% of drives: Write Quorum = M + 1 (avoid split-brain)
- Example: EC:12+4 (16 disks) → Write Quorum = 12

### Erasure Coding Visual

<img width="1030" height="540" alt="image" src="https://github.com/user-attachments/assets/45f2609b-43c4-4988-99d7-b6a2c173d17a" />

The value K here constitutes the read quorum for the deployment. The erasure set must therefore have at least K healthy drives to support read operations.

For a small object with only 1 part (`part.1`), here we have 2 data blocks and 2 parity blocks:

<img width="960" height="540" alt="image" src="https://github.com/user-attachments/assets/83ffabbc-95dc-4ac7-8fa7-832d27d59b87" />

Ref: https://blog.min.io/erasure-coding-vs-raid/

Not only does MinIO erasure coding protect against drive and node failures, MinIO also heals at the **object level**:
- Heal one object at a time vs RAID which heals at volume level
- A corrupt object restored in seconds vs. hours in RAID

### Read Request Flow

<img width="758" height="799" alt="image" src="https://github.com/user-attachments/assets/7367da38-6071-49c5-9934-7c4ae10027b6" />

### Write Request Flow

Two cases for write quorum:
* **Case 1**: Parity < 50% of drives → Write Quorum = Parity
* **Case 2**: Parity = 50% of drives → Write Quorum = Parity + 1

> If parity equals 1/2 (half) the number of erasure set drives, write quorum equals parity + 1 to avoid data inconsistency due to 'split brain' scenarios.

<img width="758" height="900" alt="image" src="https://github.com/user-attachments/assets/7d221209-56e3-4924-96ba-2c0d74e9248f" />

### Bitrot Protection

MinIO protects against **silent data corruption** (BitRot):

- **HighwayHash Algorithm**: Computes 256-bit hash per block
- **Verification**: Hash checked on every read
- **Performance**: >10 GB/sec hashing on single Intel CPU core
- **Storage Format**: `[hash | data | hash | data | ...]`
- **Detection**: Hash mismatch → disk marked bad → reconstruction

---

## PUT Operation

### Step-by-Step Flow (Storing a 10MB Object)

```
1. CLIENT REQUEST
   │
   ├─ PUT /bucket/photos/vacation.jpg (10MB)
   │
   ▼

2. HTTP HANDLER
   │
   ├─ Parse request, extract bucket/object name
   ├─ Verify authentication (Signature V4)
   ├─ Create hash verifier (for bitrot)
   │
   ▼

3. POOL SELECTION
   │
   ├─ If object already exists:
   │  └─ Use same pool as existing version
   │
   └─ If new object:
      ├─ Calculate available space for each pool
      ├─ Filter: skip suspended/rebalancing pools
      ├─ Weighted random selection (prefer pool with most space)
      └─ Select Pool 0
   │
   ▼

4. SET SELECTION (Consistent Hashing)
   │
   ├─ Hash object name using SipHash:
   │  setIndex = sipHashMod("photos/vacation.jpg", numSets, deploymentID) % numSets
   │
   ├─ Result: Always same set for same object name (deterministic)
   └─ Select Erasure Set 3
   │
   ▼

5. CREATE METADATA
   │
   ├─ Generate UUIDs:
   │  ├─ VersionID: Unique identifier for this version
   │  └─ DataDir: Directory to store data shards
   │
   ├─ Calculate distribution order:
   │  └─ hashOrder(objectName, diskCount) = [3, 1, 4, 2, 5, ...]
   │
   ├─ Set erasure parameters:
   │  ├─ ErasureM = 12 (data blocks)
   │  ├─ ErasureN = 4 (parity blocks)
   │  └─ BlockSize = 1MB
   │
   ▼

6. ERASURE ENCODING
   │
   ├─ Read data in 1MB blocks (10 blocks total)
   │
   ├─ For each block:
   │  ├─ Split into 12 data shards (~85KB each)
   │  ├─ Compute 4 parity shards using Reed-Solomon
   │  ├─ Add HighwayHash checksum to each shard
   │  └─ Result: 16 shards per block
   │
   ▼

7. DISK ORDERING (Distribution)
   │
   ├─ Take 16 disks from Set 3
   ├─ Shuffle according to distribution order
   └─ Map shards: shard_i → disk_i
   │
   ▼

8. PARALLEL WRITES
   │
   ├─ For each of 16 disks (in parallel):
   │  │
   │  ├─ Write to temporary location:
   │  │  .minio.sys/tmp/{VersionID}/{DataDir}/part.1
   │  │
   │  ├─ Format: [block1_hash|block1_data|block2_hash|block2_data|...]
   │  │
   │  └─ Verify write success
   │
   ├─ Check write quorum: Need ≥12 successful writes
   │  └─ If <12 fail: WRITE FAILS, cleanup
   │
   ▼

9. ATOMIC RENAME
   │
   ├─ Once quorum reached:
   │  └─ Rename all temp files to final location:
   │     bucket/object/{DataDir}/part.1
   │
   ▼

10. METADATA PERSISTENCE
    │
    ├─ Create xl.meta with all object metadata
    │ ├─ Version history
    │ ├─ Erasure config
    │ ├─ Distribution array
    │ └─ Part sizes
    │
    ├─ Write xl.meta to all 16 disks (in parallel)
    │
    ├─ Verify metadata quorum (≥12 successful)
    │
    ▼

11. SUCCESS
    │
    └─ Return 200 OK + ETag to client
```

### PUT Request Overview

<img width="1137" height="911" alt="image" src="https://github.com/user-attachments/assets/7c0955af-93ee-418d-9115-9c560a92708d" />

<img width="1606" height="929" alt="image" src="https://github.com/user-attachments/assets/25a0614e-a95f-41e3-adcf-e5278acca6f0" />

<img width="3192" height="1766" alt="image" src="https://github.com/user-attachments/assets/9f0b687c-5923-49c5-8eaf-8f9131dbefaf" />

For example, with 5 data blocks and 3 parity blocks:

<img width="1452" height="895" alt="image" src="https://github.com/user-attachments/assets/ba80d04b-8806-41ff-bd96-574dcf06a89d" />

### PUT Request Sequence

```mermaid
sequenceDiagram
    participant Client
    participant MinIO as MinIO Server
    participant PoolMgr as Pool Manager
    participant SetMgr as Erasure Set Manager
    participant EC as Erasure Coder
    participant Disk1 as Drive 1
    participant Disk2 as Drive 2
    participant DiskN as Drive N

    Client->>MinIO: PUT /bucket/object
    MinIO->>PoolMgr: getPoolIdx(bucket, object, size)

    alt Object Already Exists
        PoolMgr->>PoolMgr: Query all pools in parallel
        PoolMgr->>PoolMgr: GetObjectInfo() on each pool
        PoolMgr->>PoolMgr: Object found in Pool 2
        PoolMgr-->>MinIO: Use Pool 2 (existing)
    else New Object
        PoolMgr->>PoolMgr: getServerPoolsAvailableSpace()
        PoolMgr->>PoolMgr: Filter: skip suspended/rebalancing
        PoolMgr->>PoolMgr: Weighted random selection
        PoolMgr-->>MinIO: Use Pool 1 (most space)
    end

    MinIO->>SetMgr: Hash object name
    SetMgr->>SetMgr: sipHashMod(objectName, numSets)
    SetMgr-->>MinIO: Erasure Set 3

    MinIO->>EC: Create FileInfo metadata
    EC->>EC: hashOrder(objectName, drives)
    EC->>EC: Generate distribution array
    Note over EC: Distribution: [3,1,4,2,5,...]

    MinIO->>EC: Encode object (K data + M parity)
    EC->>EC: Split into data blocks
    EC->>EC: Calculate parity using Reed-Solomon
    EC-->>MinIO: Data + Parity shards

    MinIO->>SetMgr: Shuffle disks by distribution
    SetMgr-->>MinIO: Ordered disk list

    par Write to all drives in parallel
        MinIO->>Disk1: Write shard 1 + xl.meta
        MinIO->>Disk2: Write shard 2 + xl.meta
        MinIO->>DiskN: Write shard N + xl.meta
    end

    Disk1-->>MinIO: Success
    Disk2-->>MinIO: Success
    DiskN-->>MinIO: Success

    MinIO->>MinIO: Check write quorum (K drives)
    MinIO-->>Client: 200 OK
```

### PUT Layer-by-Layer Graph

```mermaid
graph TB
    subgraph "Client"
        C[PUT /mybucket/image.jpg<br/>10 MB]
    end

    subgraph "Layer 1: HTTP Server"
        H[HTTP Router<br/>Match route]
    end

    subgraph "Layer 2: API Handler"
        A[PutObjectHandler<br/>Parse request<br/>Create PutObjReader]
    end

    subgraph "Layer 3: Server Pools"
        SP[erasureServerPools<br/>Select Pool 0<br/>based on space]
    end

    subgraph "Layer 4: Erasure Sets"
        ES[erasureSets<br/>Hash object name<br/>Select Set 3]
    end

    subgraph "Layer 5: Erasure Objects"
        EO[erasureObjects<br/>Setup EC:12+4<br/>WriteQuorum: 12]
    end

    subgraph "Layer 6: Encoding Loop"
        L1[Read Block 1<br/>1 MB]
        L2[Encode to<br/>16 shards]
        L3[Write to<br/>16 disks]
        L4[Check<br/>quorum]
        L5{More<br/>blocks?}
    end

    subgraph "Layer 7: Reed-Solomon"
        RS[Split 1 MB into<br/>12 data shards<br/>Generate 4 parity]
    end

    subgraph "Layer 8: Parallel Writes"
        W1[Disk 1<br/>Write D1]
        W2[Disk 2<br/>Write D2]
        W3[Disk 12<br/>Write D12]
        W4[Disk 13<br/>Write P1]
        W5[Disk 16<br/>Write P4]
    end

    subgraph "Metadata Write"
        M1[Create xl.meta<br/>with FileInfo]
        M2[Write to all<br/>16 disks]
        M3[Check quorum<br/>12/16]
        M4{Success?}
    end

    subgraph "Final"
        F1[✅ Return ObjectInfo]
        F2[❌ Revert & Error]
    end

    C --> H --> A --> SP --> ES --> EO --> L1
    L1 --> L2 --> RS --> L3
    L3 --> W1 & W2 & W3 & W4 & W5
    W1 & W2 & W3 & W4 & W5 --> L4
    L4 --> L5
    L5 -->|Yes| L1
    L5 -->|No| M1
    M1 --> M2 --> M3 --> M4
    M4 -->|Yes| F1
    M4 -->|No| F2

    style C fill:#e1f5ff
    style RS fill:#fff3cd
    style F1 fill:#d4edda
    style F2 fill:#f8d7da
```

### Key Decisions

**Pool Selection Algorithm** (Weighted Random):
```go
totalFreeSpace = sum of free space in all pools
choose = random(0, totalFreeSpace)
for each pool:
    if pool.freeSpace >= choose:
        select this pool
        break
    choose -= pool.freeSpace
```

**Set Selection Algorithm** (Consistent Hash):
```go
func sipHashMod(key string, cardinality int, id [16]byte) int {
    k0, k1 := binary.LittleEndian.Uint64(id[0:8]),
             binary.LittleEndian.Uint64(id[8:16])
    sum64 := siphash.Hash(k0, k1, []byte(key))
    return int(sum64 % uint64(cardinality))
}
```

---

## GET Operation

### Step-by-Step Flow (Retrieving a 10MB Object)

```
1. CLIENT REQUEST
   │
   ├─ GET /bucket/photos/vacation.jpg
   │ Optional: Range header (e.g., bytes=2097152-10485760)
   │
   ▼

2. HTTP HANDLER
   │
   ├─ Parse request
   ├─ Verify authentication
   ├─ Check preconditions (If-Match, If-Modified-Since)
   │
   ▼

3. SET LOOKUP (Same Hash as Write)
   │
   ├─ Hash object name using same SipHash
   │  → Deterministically routes to same set as original write
   │
   └─ Select Erasure Set 3
   │
   ▼

4. METADATA READING
   │
   ├─ Read xl.meta from ALL 16 disks (in parallel)
   │
   ├─ Verify quorum: Need ≥12 successful reads
   │  └─ If <12: READ FAILS
   │
   ├─ Select latest version (by ModTime)
   │
   ├─ Extract:
   │  ├─ Erasure config (M, N, block size)
   │  ├─ Part sizes and ETags
   │  ├─ Distribution order
   │  └─ Shard indices
   │
   ▼

5. PARALLEL SHARD READING
   │
   ├─ Create readers for all 16 disks
   │
   ├─ Read in parallel:
   │  ├─ Each disk returns its shard blocks
   │  ├─ Verify HighwayHash per block
   │  │  └─ Hash mismatch → mark disk as bad
   │  └─ Stop reading once we have ≥12 good shards
   │
   ▼

6. RECONSTRUCTION (If Needed)
   │
   ├─ If all 16 disks healthy:
   │  └─ Use 12 data shards directly
   │
   └─ If some disks failed/corrupted:
      ├─ Use Reed-Solomon decoder
      ├─ Reconstruct missing shards from available ones
      └─ Need at least M (12) shards to reconstruct
   │
   ▼

7. RANGE EXTRACTION (If Range Header Present)
   │
   ├─ If range requested (e.g., bytes 2-10MB):
   │  ├─ Extract only requested byte range
   │  └─ Efficient: Don't read entire object
   │
   ├─ Apply decompression (if S2 compression used)
   │
   ├─ Apply decryption (if AES-256-GCM encryption used)
   │
   ▼

8. STREAM TO CLIENT
   │
   ├─ Set Content-Length header
   ├─ Set Content-Range (if range request)
   ├─ Stream data directly to HTTP response body
   │
   ▼

9. SUCCESS
    │
    └─ Return 200 OK (or 206 Partial Content for range)
```

### Failure Scenarios

| Scenario | Disks Available | Status | Action |
|----------|-----------------|--------|--------|
| All healthy | 16/16 | ✅ Success | Read 12 data shards |
| 1 disk dead | 15/16 | ✅ Success | Read 12 data shards from remaining |
| 2 disks dead | 14/16 | ✅ Success | Read 12+ shards, reconstruct if needed |
| 4 disks dead | 12/16 | ✅ Success | Read 12 available shards (at limit) |
| 5+ disks dead | <12/16 | ❌ FAIL | Cannot read (quorum lost) |

---

## Distributed Architecture

All the nodes running a distributed MinIO setup are recommended to be homogeneous — same operating system, same number of drives, and same network interconnects.

<img width="8000" height="4500" alt="image" src="https://github.com/user-attachments/assets/19b44e38-a8c8-4daa-89ad-3d9a6854ecdd" />

* No master server, no metadata server

Ref: https://github.com/minio/minio/blob/master/docs/distributed/README.md

MinIO adopts a decentralized shared-nothing architecture, where object data is scattered and stored on multiple hard disks on different nodes, providing unified namespace access and load balancing between servers through load balancing or DNS round-robin.

### Erasure Set Organization (4 Servers × 4 Disks Each = 16 Disks Total)

```
Server 1: [D1] [D2] [D3] [D4]
Server 2: [D5] [D6] [D7] [D8]
Server 3: [D9] [D10][D11][D12]
Server 4: [D13][D14][D15][D16]

            ↓ Round-Robin Assignment ↓

Erasure Set 0: [D1, D5, D9, D13, D2, D6, D10, D14, ...]
                S1   S2   S3   S4   S1   S2   S3    S4

Key: Each set has disks from ALL servers
```

### Fault Tolerance

**If Server 3 Dies**:
```
Set contains: [D1(S1), D5(S2), D9(S3), D13(S4), ...]

After S3 failure:
├─ Available: D1(S1) ✓, D5(S2) ✓, D13(S4) ✓
├─ Dead: D9(S3) ✗
├─ Tolerance: EC:12+4 can lose up to 4 disks
│
└─ Result: SAFE - Can still read and recover
```

**If Any 4 Disks Die**:
```
Available shards: 12 (exactly at read quorum)
Parity tolerance: 4
Result: Still readable but no fault tolerance left
```

**If 5+ Disks Die**:
```
Available shards: <12 (below read quorum)
Result: UNRECOVERABLE - READ FAILS
```

---

## Server Pools

![serverpools](https://github.com/user-attachments/assets/a25b361c-e253-4c06-983a-e95b4d0ae464)

A server pool is a set of MinIO server nodes which pool their drives and resources, creating a unit of expansion. All nodes in a server pool share their hardware resources in an isolated namespace.

The other important point here involves rebalance-free, non-disruptive expansion. With MinIO's server pool approach - rebalancing is not required to expand. Ref: https://blog.min.io/no-rebalancing-object-storage/

A MinIO cluster is built on server pools, and server pools are built on erasure sets.

<img width="961" height="440" alt="image" src="https://github.com/user-attachments/assets/d916ce30-b3c0-46d5-9ab9-f31801b8872b" />

### Multi-Pool Architecture

MinIO can have multiple independent pools for expansion:

```
Cluster
├── Pool 1 (16 disks, 4 nodes)
│   ├─ Erasure Set 0
│   └─ Erasure Set 0 (shared)
│
├── Pool 2 (32 disks, 4 nodes)
│   ├─ Erasure Set 1
│   └─ Erasure Set 2
│
└── Pool 3 (48 disks, 4 nodes)
    ├─ Erasure Set 3
    ├─ Erasure Set 4
    └─ Erasure Set 5
```

### Pool Expansion (Rebalance-Free)

1. **Add new pool**: MinIO detects new endpoints at startup
2. **Update format files**: Cluster configuration updated
3. **New objects**: Distributed across all pools by available space
4. **Existing objects**: Stay in original pool (no rebalancing)
5. **Decommission**: Background migration copies objects to other pools

### Weighted Random Selection

When adding new object to new pool:
- Calculate available space: Pool1=500GB, Pool2=200GB, Pool3=300GB (total=1TB)
- Generate random number: 0-1000GB
- If 0-500: Pool1, if 500-700: Pool2, if 700-1000: Pool3
- Result: Pools filled proportionally to their capacity

---

## Code Architecture

### Layer Hierarchy

```mermaid
graph TB
    subgraph "1. Entry Point"
        A[Application Entry<br/>& Initialization]
    end

    subgraph "2. HTTP Server Layer"
        B[HTTP Server<br/>Router & Middleware]
    end

    subgraph "3. API Handler Layer"
        C[S3 API Handlers<br/>Request Processing]
    end

    subgraph "4. Object Layer Interface"
        D[ObjectLayer Interface<br/>Abstraction Layer]
    end

    subgraph "5. Erasure Server Pools"
        E[Pool Manager<br/>Multi-Pool Coordination]
    end

    subgraph "6. Erasure Sets"
        F[Set Manager<br/>Consistent Hashing]
    end

    subgraph "7. Erasure Objects"
        G[Erasure Logic<br/>Quorum & Operations]
    end

    subgraph "8. Storage Layer"
        H[Disk I/O<br/>File Operations]
    end

    subgraph "9. Metadata Layer"
        I[xl.meta Management<br/>Version Control]
    end

    subgraph "10. Healing Layer"
        J[Self-Repair<br/>Background Scanner]
    end

    subgraph "11. Erasure Coding"
        K[Reed-Solomon<br/>Data Protection]
    end

    subgraph "12. Bitrot Protection"
        L[Hash Verification<br/>Integrity Checks]
    end

    A --> B
    B --> C
    C --> D
    D --> E
    E --> F
    F --> G
    G --> H
    H --> I

    G -.Healing.-> J
    G -.Encoding.-> K
    H -.Verification.-> L

    style A fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style B fill:#f3e5f5,stroke:#4a148c,stroke-width:3px,color:#000
    style C fill:#e8f5e9,stroke:#1b5e20,stroke-width:3px,color:#000
    style D fill:#fff3e0,stroke:#e65100,stroke-width:3px,color:#000
    style E fill:#fce4ec,stroke:#880e4f,stroke-width:3px,color:#000
    style F fill:#e0f2f1,stroke:#004d40,stroke-width:3px,color:#000
    style G fill:#f1f8e9,stroke:#33691e,stroke-width:3px,color:#000
    style H fill:#e3f2fd,stroke:#0d47a1,stroke-width:3px,color:#000
    style I fill:#fef5e7,stroke:#f39c12,stroke-width:3px,color:#000
    style J fill:#fadbd8,stroke:#c0392b,stroke-width:3px,color:#000
    style K fill:#d5f4e6,stroke:#117a65,stroke-width:3px,color:#000
    style L fill:#ebdef0,stroke:#6c3483,stroke-width:3px,color:#000
```

### Detailed Layer Flow

```mermaid
graph TB
    subgraph "1. HTTP Server Layer"
        HTTP[HTTP Server<br/>xhttp.NewServer]
        Router[Mux Router<br/>mux.NewRouter]
    end

    subgraph "2. Middleware Layer"
        Auth[Authentication<br/>Signature V4]
        Trace[HTTP Tracing]
        Throttle[Request Throttling<br/>maxClients]
        GZIP[GZIP Compression]
    end

    subgraph "3. API Handler Layer"
        APIHandlers[objectAPIHandlers]
        GetObj[GetObjectHandler]
        PutObj[PutObjectHandler]
        DelObj[DeleteObjectHandler]
        ListObj[ListObjectsHandler]
    end

    subgraph "4. ObjectLayer Interface"
        ObjInterface["<b>ObjectLayer Interface</b><br/>• GetObjectNInfo<br/>• PutObject<br/>• DeleteObject<br/>• ListObjects<br/>• GetObjectInfo<br/>• Multipart Operations"]
    end

    subgraph "5. Erasure Server Pools"
        ESP[erasureServerPools<br/>implements ObjectLayer]
        Pool1[Pool 1<br/>erasureSets]
        Pool2[Pool 2<br/>erasureSets]
        PoolN[Pool N<br/>erasureSets]
    end

    subgraph "6. Erasure Sets"
        Set1[Set 1<br/>erasureObjects]
        Set2[Set 2<br/>erasureObjects]
        SetN[Set N<br/>erasureObjects]
    end

    subgraph "7. Erasure Objects Layer"
        ErasureObj[erasureObjects<br/>implements ObjectLayer]
        ECLogic[Erasure Coding Logic<br/>Reed-Solomon]
        Quorum[Read/Write Quorum]
        Healing[Self-Healing]
    end

    subgraph "8. StorageAPI Interface"
        StorageInterface["<b>StorageAPI Interface</b><br/>• ReadVersion<br/>• WriteMetadata<br/>• DeleteVersion<br/>• ReadFile/WriteAll<br/>• Volume Operations"]
    end

    subgraph "9. Storage Implementation"
        XLStorage[xlStorage<br/>implements StorageAPI]
        Remote[storageRESTClient<br/>Remote Disks]
        DiskCheck[xlStorageDiskIDCheck<br/>Health Wrapper]
    end

    subgraph "10. Disk Layer"
        LocalDisk[Local Disk I/O<br/>xl.meta files]
        RemoteDisk[Remote Disk via REST]
        Metadata[xl.meta<br/>Object Metadata]
    end

    HTTP --> Router
    Router --> Auth
    Auth --> Trace
    Trace --> Throttle
    Throttle --> GZIP
    GZIP --> APIHandlers
    APIHandlers --> GetObj
    APIHandlers --> PutObj
    APIHandlers --> DelObj
    APIHandlers --> ListObj

    GetObj --> ObjInterface
    PutObj --> ObjInterface
    DelObj --> ObjInterface
    ListObj --> ObjInterface

    ObjInterface --> ESP
    ESP --> Pool1
    ESP --> Pool2
    ESP --> PoolN

    Pool1 --> Set1
    Pool1 --> Set2
    Pool1 --> SetN

    Set1 --> ErasureObj
    ErasureObj --> ECLogic
    ErasureObj --> Quorum
    ErasureObj --> Healing

    ErasureObj --> StorageInterface

    StorageInterface --> XLStorage
    StorageInterface --> Remote
    StorageInterface --> DiskCheck

    XLStorage --> LocalDisk
    Remote --> RemoteDisk
    LocalDisk --> Metadata
    RemoteDisk --> Metadata

    style ObjInterface fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style StorageInterface fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style ESP fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style ErasureObj fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style XLStorage fill:#f1f8e9,stroke:#33691e,stroke-width:2px,color:#000
```

### Interface Class Diagram

```mermaid
classDiagram
    class ObjectLayer {
        <<interface>>
        +GetObjectNInfo() GetObjectReader
        +PutObject() ObjectInfo
        +DeleteObject() ObjectInfo
        +GetObjectInfo() ObjectInfo
        +ListObjects() ListObjectsInfo
        +MakeBucket() error
        +NewMultipartUpload() NewMultipartUploadResult
        +GetDisks() []StorageAPI
    }

    class erasureServerPools {
        -poolMeta poolMeta
        -serverPools []*erasureSets
        -deploymentID [16]byte
        +PutObject() ObjectInfo
        +GetObjectNInfo() GetObjectReader
        +getPoolIdx() int
    }

    class erasureSets {
        -sets []*erasureObjects
        -format *formatErasureV3
        -erasureDisks [][]StorageAPI
        -setCount int
        -setDriveCount int
        +PutObject() ObjectInfo
        +getHashedSet() int
    }

    class erasureObjects {
        -setDriveCount int
        -defaultParityCount int
        -getDisks func()[]StorageAPI
        -nsMutex *nsLockMap
        +PutObject() ObjectInfo
        +putObject() ObjectInfo
        +defaultWQuorum() int
        +defaultRQuorum() int
    }

    class StorageAPI {
        <<interface>>
        +ReadVersion() FileInfo
        +WriteMetadata() error
        +DeleteVersion() error
        +CreateFile() error
        +ReadFile() int64
        +MakeVol() error
        +ListVols() []VolInfo
        +GetDiskID() string
        +IsOnline() bool
    }

    class xlStorage {
        -diskPath string
        -endpoint Endpoint
        -diskID string
        -formatFile string
        +WriteMetadata() error
        +ReadVersion() FileInfo
        +CreateFile() error
        +ReadFile() int64
    }

    class storageRESTClient {
        -endpoint Endpoint
        -restClient *rest.Client
        -diskID string
        +WriteMetadata() error
        +ReadVersion() FileInfo
        +CreateFile() error
    }

    class Erasure {
        -encoder func()Encoder
        -dataBlocks int
        -parityBlocks int
        -blockSize int64
        +EncodeData() [][]byte
        +DecodeDataBlocks() error
        +ShardSize() int64
    }

    ObjectLayer <|.. erasureServerPools : implements
    ObjectLayer <|.. erasureSets : implements
    ObjectLayer <|.. erasureObjects : implements

    erasureServerPools *-- erasureSets : contains
    erasureSets *-- erasureObjects : contains
    erasureObjects --> StorageAPI : uses
    erasureObjects --> Erasure : uses

    StorageAPI <|.. xlStorage : implements
    StorageAPI <|.. storageRESTClient : implements

    style ObjectLayer fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style StorageAPI fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style erasureServerPools fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style erasureSets fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style erasureObjects fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style xlStorage fill:#f1f8e9,stroke:#33691e,stroke-width:2px,color:#000
    style Erasure fill:#e8f5e9,stroke:#2e7d32,stroke-width:2px,color:#000
```

### Key Interfaces

**ObjectLayer Interface**:
```go
type ObjectLayer interface {
    // Bucket operations
    MakeBucket(ctx, bucket, opts) error
    GetBucketInfo(ctx, bucket, opts) (BucketInfo, error)
    ListBuckets(ctx, opts) ([]BucketInfo, error)
    DeleteBucket(ctx, bucket, opts) error
    ListObjects(...) (ListObjectsInfo, error)
    ListObjectVersions(...) (ListObjectVersionsInfo, error)

    // Object operations
    GetObjectNInfo(ctx, bucket, object, rangeSpec, headers, opts) (*GetObjectReader, error)
    GetObjectInfo(ctx, bucket, object, opts) (ObjectInfo, error)
    PutObject(ctx, bucket, object, data, opts) (ObjectInfo, error)
    CopyObject(ctx, srcBucket, srcObject, dstBucket, dstObject, ...) (ObjectInfo, error)
    DeleteObject(ctx, bucket, object, opts) (ObjectInfo, error)
    DeleteObjects(ctx, bucket, objects, opts) ([]DeletedObject, []error)

    // Multipart operations
    NewMultipartUpload(ctx, bucket, object, opts) (*NewMultipartUploadResult, error)
    PutObjectPart(ctx, bucket, object, uploadID, partID, data, opts) (PartInfo, error)
    CompleteMultipartUpload(ctx, bucket, object, uploadID, parts, opts) (ObjectInfo, error)
    AbortMultipartUpload(ctx, bucket, object, uploadID, opts) error

    // Healing & Info
    HealFormat(ctx, dryRun) (HealResultItem, error)
    HealBucket(ctx, bucket, opts) (HealResultItem, error)
    HealObject(ctx, bucket, object, versionID, opts) (HealResultItem, error)
    StorageInfo(ctx, metrics bool) StorageInfo
}
```

**StorageAPI Interface**:
```go
type StorageAPI interface {
    // Metadata
    ReadVersion(ctx, origvolume, volume, path, versionID, opts) (FileInfo, error)
    WriteMetadata(ctx, origvolume, volume, path, fi) error
    DeleteVersion(ctx, volume, path, fi, ...) error

    // File operations
    ReadFile(ctx, volume, path, offset, buf, verifier) (n, error)
    CreateFile(ctx, origvolume, volume, path, size, reader) error
    ReadFileStream(ctx, volume, path, offset, length) (io.ReadCloser, error)
    AppendFile(ctx, volume, path, buf) error
    Delete(ctx, volume, path, opts) error

    // Volume operations
    MakeVol(ctx, volume) error
    ListVols(ctx) ([]VolInfo, error)
    StatVol(ctx, volume) (VolInfo, error)
    DeleteVol(ctx, volume, forceDelete bool) error

    // Disk info
    IsOnline() bool
    GetDiskID() (string, error)
    DiskInfo(ctx, opts) (DiskInfo, error)
}
```

### Implementations

| Component | Location | Role |
|-----------|----------|------|
| `erasureServerPools` | `cmd/erasure-server-pool.go` | Pool orchestration, weighted selection |
| `erasureSets` | `cmd/erasure-sets.go` | Set routing, consistent hashing |
| `erasureObjects` | `cmd/erasure-object.go` | Core put/get/delete with EC |
| `xlStorage` | `cmd/xl-storage.go` | Local disk I/O |
| `storageRESTClient` | `cmd/storage-rest-client.go` | Remote disk via REST |
| `Erasure` | `cmd/erasure-coding.go` | Reed-Solomon encode/decode |

---

## Healing

MinIO performs automatic background healing to detect and repair corrupted objects:

### Healing Mechanisms

1. **Bitrot Detection**: HighwayHash checksum verification on every read
2. **Bad Disk Detection**: Continuous health monitoring of all disks
3. **Object-Level Healing**: Corrupted objects repaired in seconds (vs RAID hours)
4. **Background Scanner**: Periodic scan of all objects to detect bitrot proactively

### Healing Flow

```
Bad block detected (hash mismatch)
        │
        ▼
Mark disk as bad
        │
        ▼
Read remaining 15 shards (12+ available)
        │
        ▼
Use Reed-Solomon to reconstruct missing shard
        │
        ▼
Repair disk by writing reconstructed shard
        │
        ▼
Verify repair with new hash
        │
        ▼
Continue serving object (healed)
```

---

## Gateway Mode (Deprecated)

MinIO introduced gateway mode early on to provide S3 API compatibility to legacy systems:

<img width="984" height="760" alt="image" src="https://github.com/user-attachments/assets/402f972a-ea67-4c56-966f-938ad454f3a9" />

<img width="1024" height="800" alt="image" src="https://github.com/user-attachments/assets/f28dbef2-d8fa-41e2-9952-8476ef0ade6d" />

**Why Deprecated**:
- Critical S3 features (versioning, replication, locking, encryption) couldn't work in gateway mode without proprietary formats
- Would defeat the purpose of direct backend access
- Better to run MinIO in server mode than as a stateless proxy
- S3 API now ubiquitous (partly due to MinIO Gateway work)

**Lessons Learned**:
- S3 API evolved significantly since gateway inception
- Inline translation is insufficient for modern S3 capabilities
- Backends become mere storage media, which is essentially running MinIO anyway

Reference: [Gateway Migration](https://blog.min.io/minio-gateway-migration/) and [Deprecation Details](https://blog.min.io/deprecation-of-the-minio-gateway/)

---

## Advanced Features

### Versioning
- Keep multiple versions of an object
- Each version has separate `xl.meta` entry
- Access previous versions without data loss

### Object Locking (WORM)
- Write Once, Read Many protection
- Objects immutable for set retention period
- Compliance and audit requirements

### Lifecycle Management
- Automatic object deletion/transition after time period
- Move to different storage classes
- Cost optimization

### Replication
- Automatic cross-cluster replication
- Disaster recovery and high availability
- Real-time synchronization

### IAM & Access Control
- User authentication (basic, LDAP, OAuth)
- Bucket policies (similar to AWS S3)
- Access key/secret pairs

### Encryption

Server-side encryption (SSE-S3, SSE-KMS) and client-side encryption support with master key rotation.

<img width="1245" height="626" alt="image" src="https://github.com/user-attachments/assets/0571b920-39c8-494e-972d-0d48d7b06592" />

<img width="1245" height="626" alt="image" src="https://github.com/user-attachments/assets/b73e3512-bc87-4115-8689-d9f5872562fe" />

<img width="1245" height="678" alt="image" src="https://github.com/user-attachments/assets/eb419eca-2356-4086-ab6c-f453c25094d4" />

---

## Distributed Locking (dsync)

MinIO avoids consistency issues using distributed locking:

### How dsync Works

1. **Lock Request**: Any node broadcasts lock request to all nodes
2. **Quorum**: If N/2+1 nodes approve → lock acquired
3. **No Master**: Every node is peer; no single authority
4. **Stale Detection**: Between-node heartbeats detect offline nodes

### Limitations
- Supports up to 32 nodes (theoretical)
- Lock throughput decreases as cluster grows
- Can lose locks in certain scenarios (acceptable for MinIO's use case)

---

## Quick Reference: Key Concepts

| Concept | Definition |
|---------|-----------|
| **Erasure Set** | Group of disks where objects are erasure coded |
| **Server Pool** | Collection of erasure sets, independent expansion unit |
| **Consistent Hash** | SipHash used to deterministically place objects |
| **Read Quorum** | Minimum shards needed to reconstruct object (M data shards) |
| **Write Quorum** | Minimum disks that must acknowledge write (M or M+1) |
| **BitRot** | Silent data corruption, detected via HighwayHash |
| **xl.meta** | Metadata file containing object info, stored on all disks |
| **DataDir** | UUID-named directory storing object's data shard |
| **Reed-Solomon** | Erasure coding algorithm enabling data reconstruction |

---

## Performance Characteristics

- **Throughput**: GB/sec performance (limited by network)
- **Latency**: Milliseconds for PUT/GET operations
- **Scalability**: Supports petabyte-scale deployments
- **Fault Tolerance**: Up to N parity disks can fail per set
- **Healing**: Object-level healing in seconds
- **Bitrot Hashing**: >10 GB/sec on single CPU core

---

## Real-World Metadata Examples

### format.json from Actual Cluster

A 12-disk single erasure set cluster:

```json
{
  "version": "1",
  "format": "xl",
  "id": "f9a7a6ba-39d9-4483-bb47-fe86518bdc67",
  "xl": {
    "version": "3",
    "this": "9ae64de8-1c75-46df-b09d-ad8b97f95313",
    "sets": [
      [
        "4199dbce-78ba-4176-846d-7423ab6cfcd9",
        "22b83b76-f883-49c8-abc8-a3cf84eb92f4",
        "9ae64de8-1c75-46df-b09d-ad8b97f95313",
        "fc1a7dde-1da7-44cc-9380-3ae3063c415c",
        "48d7881f-6e93-42ab-9d89-f27bf0648b0d",
        "b8cfec44-f88b-4193-9575-368d92eefb16",
        "ef66b6f7-3c15-45fa-aca8-52286f4750f4",
        "02b3aa13-ff62-4e46-a196-f40b6f531c23",
        "f5dd8d65-56d7-40f2-9035-b4b37e3018a5",
        "ae4e30fd-db65-4c0e-a9c1-44f50191ba20",
        "d4cf829c-b96f-4687-845c-8884a43a6397",
        "2efc58b9-253a-4ac6-ba92-a316811f896c"
      ]
    ],
    "distributionAlgo": "SIPMOD+PARITY"
  }
}
```

**Explanation**:
- 1 erasure set with 12 disks (UUIDs in the sets array)
- Deployment ID: f9a7a6ba-39d9-4483-bb47-fe86518bdc67 (shared by all disks)
- Distribution Algorithm: SIPMOD+PARITY
- This disk position: 2 in the set (9ae64de8-1c75-46df-b09d-ad8b97f95313)

### Real xl.meta Example from Small File (65 bytes)

File content: `This is test data for xl.meta debugging with erasure coding EC:4`

```json
{
  "Versions": [
    {
      "Header": {
        "EcM": 8,
        "EcN": 4,
        "Flags": 6,
        "ModTime": "2026-01-14T16:53:55.923264863+05:30",
        "Signature": "b9f71a0b",
        "Type": 1,
        "VersionID": "00000000000000000000000000000000"
      },
      "Idx": 0,
      "Metadata": {
        "Type": 1,
        "V2Obj": {
          "CSumAlgo": 1,
          "DDir": "NhHND1OVRfWzQYC/GFqfGA==",
          "EcAlgo": 1,
          "EcBSize": 1048576,
          "EcDist": [1,2,3,4,5,6,7,8,9,10,11,12],
          "EcIndex": 3,
          "EcM": 8,
          "EcN": 4,
          "ID": "AAAAAAAAAAAAAAAAAAAAAA==",
          "MTime": 1768389835923264863,
          "MetaSys": {
            "x-minio-internal-inline-data": "dHJ1ZQ=="
          },
          "MetaUsr": {
            "content-type": "text/plain",
            "etag": "eeb5a84d38f5dac272eb0d3f772c8a59"
          },
          "PartASizes": [ 65 ],
          "PartETags": null,
          "PartNums": [  1],
          "PartSizes": [ 65 ],
          "Size": 65
        },
        "v": 1740736516
      }
    }
  ]
}
```

### xl.meta from a Different Disk (EcIndex=7)

Decoding `xl.meta` from another disk in the same erasure set shows the same object but a different shard:

```json
{
  "Versions": [
    {
      "Header": { "EcM": 8, "EcN": 4, "Type": 1, "VersionID": "00000000000000000000000000000000" },
      "Metadata": {
        "V2Obj": {
          "EcDist": [1,2,3,4,5,6,7,8,9,10,11,12],
          "EcIndex": 7,
          "EcM": 8,
          "EcN": 4,
          "MetaSys": { "x-minio-internal-inline-data": "dHJ1ZQ==" },
          "MetaUsr": { "content-type": "text/plain", "etag": "eeb5a84d38f5dac272eb0d3f772c8a59" },
          "Size": 65
        }
      }
    }
  ]
}
--- INLINE DATA ---
{
  "null": {
    "bitrot_valid": true,
    "bytes": 41,
    "data_base64": "b2RpbmcgRUM6",
    "data_string": "oding EC:"
  }
}
```

Notice `EcIndex: 7` (vs `EcIndex: 3` on the other disk) — each disk holds a different shard of the same object. The `data_string` differs (`"oding EC:"` vs `"for xl.me"`) confirming each disk stores its own slice.

<img width="1024" height="559" alt="image" src="https://github.com/user-attachments/assets/c1a8bf76-f524-48fd-aaa2-9a9f6aec28a7" />

### Data Distribution Visualization

For the 65-byte file above split into EC:8+4:

```
┌────────────────────────────────────────────────────────────────┐
│                    Original File (~65 bytes)                   │
│                   "...for xl.me...oding EC:..."                │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  Erasure Split into 8 Data Shards + 4 Parity Shards:           │
│                                                                │
│  EcDist: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]               │
│           ├─────────────────────┤ ├──────────────┤             │
│           8 DATA shards          4 PARITY shards               │
│                                                                │
│  Disk EcIndex=3: Contains data shard 3 → "for xl.me"           │
│  Disk EcIndex=7: Contains data shard 7 → "oding EC:"           │
│                                                                │
│  Disks 9-12 (EcIndex 9,10,11,12): Parity shards (for recovery) │
└────────────────────────────────────────────────────────────────┘
```

### Storing a Test File

```bash
# Create test bucket and upload file
kubectl exec minio-0 -- sh -c '
mc alias set local http://localhost:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD --insecure 2>/dev/null
mc mb local/debug-bucket --insecure 2>/dev/null || true
echo "This is test data for xl.meta debugging with erasure coding EC:4" > /tmp/test-file.txt
mc cp /tmp/test-file.txt local/debug-bucket/test-file.txt --insecure
mc stat local/debug-bucket/test-file.txt --insecure
'
```

**Expected Output**:
```
Added `local` successfully.
Bucket created successfully `local/debug-bucket`.
`/tmp/test-file.txt` -> `local/debug-bucket/test-file.txt`
┌───────┬─────────────┬──────────┬────────────┐
│ Total │ Transferred │ Duration │ Speed      │
│ 65 B  │ 65 B        │ 00m00s   │ 1.13 KiB/s │
└───────┴─────────────┴──────────┴────────────┘
Name      : test-file.txt
Date      : 2026-01-14 11:23:55 UTC
Size      : 65 B
ETag      : eeb5a84d38f5dac272eb0d3f772c8a59
Type      : file
Metadata  :
  Content-Type: text/plain
```

### On-Disk Structure

For small files (≤128KB), data is inlined in xl.meta:

```
debug-bucket/
└── test-file.txt/
    └── xl.meta              # Contains metadata + inline data
```

For larger files (>128KB):

```
/data1/testbucket/test-large-file.txt/
├── xl.meta                          # Metadata (on all 12 disks)
└── <DDir-UUID>/                     # Data directory
    └── part.1                       # Actual data shard for this disk
```

### Healing Example

<img width="868" height="930" alt="image" src="https://github.com/user-attachments/assets/e90b355a-6810-4648-823b-5baddacf7d64" />

<img width="688" height="945" alt="image" src="https://github.com/user-attachments/assets/0f3218db-1d01-4b99-a652-20172683fb60" />

Ref: https://minio-docs.tf.fo/operations/concepts/healing

<img width="1136" height="946" alt="image" src="https://github.com/user-attachments/assets/f3711710-afca-4123-8ba6-0c07469876b7" />

Ref: https://minio-docs.tf.fo/operations/data-recovery

### Replication & Site-to-Site

<img width="827" height="427" alt="image" src="https://github.com/user-attachments/assets/dbb719ec-3363-45a8-922e-ebcc9a63dd14" />

### Distributed Locking (dsync) in Action

<img width="2540" height="1216" alt="image" src="https://github.com/user-attachments/assets/231e0ec3-20e3-45c9-8c2e-b703ed8e4e9b" />

<img width="1375" height="872" alt="image" src="https://github.com/user-attachments/assets/6e6a6e02-44b8-4e32-b552-22a134de5f40" />

Ref: https://blog.min.io/minio-dsync-a-distributed-locking-and-syncing-package-for-go/

---

## References

- [MinIO Official Documentation](https://min.io/docs/)
- [MinIO Blog - Erasure Coding](https://blog.min.io/erasure-coding-vs-raid/)
- [MinIO GitHub Repository](https://github.com/minio/minio)
- [Distributed Locking in MinIO](https://blog.min.io/minio-dsync-a-distributed-locking-and-syncing-package-for-go/)
- [MinIO Versioning Deep Dive](https://blog.min.io/minio-versioning-metadata-deep-dive/)
- [No Rebalancing Object Storage](https://blog.min.io/no-rebalancing-object-storage/)
- [Distributed README](https://github.com/minio/minio/blob/master/docs/distributed/README.md)
