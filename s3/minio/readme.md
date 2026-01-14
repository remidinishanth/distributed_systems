MinIO just does one thing - Object storage for Private cloud

<img width="1105" height="588" alt="image" src="https://github.com/user-attachments/assets/b9fdf6eb-b52c-460d-89e0-fc90a2762f2f" />

<img width="1326" height="744" alt="image" src="https://github.com/user-attachments/assets/226c6dad-f468-4140-998d-86641a9115eb" />


<img width="689" height="241" alt="image" src="https://github.com/user-attachments/assets/a5f7c369-978a-4944-94d7-46ce8c87c160" />


<img width="1142" height="499" alt="image" src="https://github.com/user-attachments/assets/ea1e33e0-cac3-47b9-b975-987243174e59" />

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/82863b5c-7278-4c28-981b-94dd71b614db" />

* Minio adopts a metadata-free database design for high performance, avoiding the metabase becoming a performance bottleneck for the entire system, and limiting failures to a single cluster, so that no other clusters are involved.
* Minio is also fully compatible with the S3 interface, so it can also be used as a gateway to provide S3 access to the outside world.
* Use both Minio Erasure code and checksum to prevent hardware failures. Even if you lose more than half of your hard drive, you can still recover from it. (N/2)-1 node failure is also allowed in the distribution

<img width="1181" height="619" alt="image" src="https://github.com/user-attachments/assets/82ac9b76-b73d-4bf0-a084-bcdc3ceaac49" />

<img width="1209" height="672" alt="image" src="https://github.com/user-attachments/assets/3bf05817-6bb7-4d08-86df-db43fe8b0f71" />

<img width="1104" height="676" alt="image" src="https://github.com/user-attachments/assets/e82fa817-cbfa-4e66-aac9-1c06b528884f" />


### Legacy object storage architecture
<img width="1591" height="863" alt="image" src="https://github.com/user-attachments/assets/86218141-aa98-478d-9b7b-02a18b6faf71" />


## Architecture

All the nodes running distributed MinIO setup are recommended to be homogeneous, i.e. same operating system, same number of drives and same network interconnects.

Start distributed MinIO instance on n nodes with m drives each mounted at `/export1` to `/exportm` (pictured below), by running this command on all the `n` nodes:

<img width="8000" height="4500" alt="image" src="https://github.com/user-attachments/assets/19b44e38-a8c8-4daa-89ad-3d9a6854ecdd" />

* No master server, no metadata server or anything

Ref: https://github.com/minio/minio/blob/master/docs/distributed/README.md


We will see Pools and Erasure coding in the following sections
<img width="961" height="440" alt="image" src="https://github.com/user-attachments/assets/d916ce30-b3c0-46d5-9ab9-f31801b8872b" />



### Decentralized architecture

Minio adopts a decentralized shared-nothing architecture, where object data is scattered and stored on multiple hard disks on different nodes, providing unified namespace access and load balancing between servers through load balancing or DNS rounding

## MinIO Server Pools

![serverpools](https://github.com/user-attachments/assets/a25b361c-e253-4c06-983a-e95b4d0ae464)

A server pool is a set of minio server nodes which pool their drives and resources, creating a unit of expansion. All nodes in a server pool share their hardware resources in an isolated namespace.  

The other important point here involves rebalance-free, non-disruptive expansion. With MinIO’s server pool approach - rebalancing is not required to expand. Ref: https://blog.min.io/no-rebalancing-object-storage/

A MinIO cluster is built on server pools, and server pools are built on erasure sets.


## Code Structure

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

Interfaces in MinIO

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
    
    class xlStorageDiskIDCheck {
        -storage StorageAPI
        -diskID string
        -healthCheck bool
        +WriteMetadata() error
        +ReadVersion() FileInfo
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
    StorageAPI <|.. xlStorageDiskIDCheck : implements
    
    xlStorageDiskIDCheck o-- StorageAPI : wraps
    
    style ObjectLayer fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style StorageAPI fill:#e1f5ff,stroke:#01579b,stroke-width:3px,color:#000
    style erasureServerPools fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style erasureSets fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style erasureObjects fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#000
    style xlStorage fill:#f1f8e9,stroke:#33691e,stroke-width:2px,color:#000
    style Erasure fill:#e8f5e9,stroke:#2e7d32,stroke-width:2px,color:#000
```

## Erasure coding

* Erasure sets, built per server pool, are sets of nodes and drives to which MinIO applies erasure coding to protect data from loss and corruption. 
* Erasure coding breaks objects into data and parity blocks and can use these blocks to reconstruct missing or corrupted blocks if necessary. With MinIO’s highest level of protection (8 parity or EC:8), you may lose up to half of the total drives and still recover data.
* For example, in a 16-drive setup, data can be split into 12 data shards and 4 parity shards, allowing the system to rebuild data even if up to 4 drives fail.

Data shards contain a portion of a given object. Parity shards contain a mathematical representation of the object used for rebuilding Data shards.

<img width="1030" height="540" alt="image" src="https://github.com/user-attachments/assets/45f2609b-43c4-4988-99d7-b6a2c173d17a" />

The value K here constitutes the read quorum for the deployment. The erasure set must therefore have at least K healthy drives in the erasure set to support read operations.


Here say we have small object which only has 1 part - `part.1`, now in this case, we have 2 data blocks and 2 parity blocks for this part.
<img width="960" height="540" alt="image" src="https://github.com/user-attachments/assets/83ffabbc-95dc-4ac7-8fa7-832d27d59b87" />

Ref: https://blog.min.io/erasure-coding-vs-raid/

Not only does MinIO erasure coding protect objects against data loss in the event that multiple drives and nodes fail, MinIO also protects and heals at the object level. 
* The ability to heal one object at a time is a dramatic advantage over systems such as RAID that heal at the volume level.
* A corrupt object could be restored in MinIO in seconds vs. hours in RAID.


MinIO protects against `BitRot`, or silent data corruption, which can have many different causes such as power current spikes, bugs in disk firmware and even simply aging drives. 
* MinIO uses the `HighwayHash` algorithm to compute a hash on read and verify it on write from the application, across the network and to the storage media.
* This process is highly efficient - it can achieve hashing speeds over 10 GB/sec on a single core on Intel CPUs - and has minimal impact on normal read/write operations across the erasure set. https://github.com/google/highwayhash

### Read request

<img width="758" height="799" alt="image" src="https://github.com/user-attachments/assets/7367da38-6071-49c5-9934-7c4ae10027b6" />

### Write request

Two cases:
* Case 1: Parity < 50% of drives
  - Write Quorum = Parity
* Case 2: Parity = 50% of drives
  - Write Quorum = Parity + 1

> If parity equals 1/2 (half) the number of erasure set drives, write quorum equals parity + 1 (one) to avoid data inconsistency due to 'split brain' scenarios.

<img width="758" height="900" alt="image" src="https://github.com/user-attachments/assets/7d221209-56e3-4924-96ba-2c0d74e9248f" />


## Put and Get Operation

### Storing an Object (The PUT Request)

<img width="1137" height="911" alt="image" src="https://github.com/user-attachments/assets/7c0955af-93ee-418d-9115-9c560a92708d" />

- Choosing an erasure set for the object is decided during `PutObject()`, object names are used to find the right erasure set using the following pseudo code.

```go
// hashes the key returning an integer.
func sipHashMod(key string, cardinality int, id [16]byte) int {
        if cardinality <= 0 {
                return -1
        }
        sip := siphash.New(id[:])
        sip.Write([]byte(key))
        return int(sip.Sum64() % uint64(cardinality))
}
```

Input for the key is the object name specified in `PutObject()`, returns a unique index. This index is one of the erasure sets where the object will reside. This function is a consistent hash for a given object name i.e for a given object name the index returned is always the same.



When a client sends an object to the cluster, MinIO follows a specific sequence to ensure data is stored safely and evenly distributed.

* Step 1: Hashing: The object name is processed by a deterministic hash function to create a unique hash value.

* Step 2: Drive Selection: A modulus function is applied to that hash value. The result determines the specific set of drives (erasure set) where the data will live.

* Step 3: Erasure Coding: Simultaneously, the Erasure Code Engine processes the object data. It breaks the object into:
  - Data blocks: The actual content.
  - Parity blocks: Redundancy data for recovery.

* Step 4: Writing: These blocks are written to the prescribed drives.

Note: MinIO uses `SipHash` for this process. This algorithm ensures that objects are distributed evenly across all drives, resulting in near-uniform disk utilization.

<img width="1606" height="929" alt="image" src="https://github.com/user-attachments/assets/25a0614e-a95f-41e3-adcf-e5278acca6f0" />

<img width="3192" height="1766" alt="image" src="https://github.com/user-attachments/assets/9f0b687c-5923-49c5-8eaf-8f9131dbefaf" />

For example, with 5 data blocks and 3 parity blocks

<img width="1452" height="895" alt="image" src="https://github.com/user-attachments/assets/ba80d04b-8806-41ff-bd96-574dcf06a89d" />

```mermaid
graph TB
    Start[Client: PUT Object] --> CheckExisting{Object Already<br/>Exists?}
    
    CheckExisting -->|Yes| UseExistingPool[Use Same Pool<br/>as Existing Object]
    CheckExisting -->|No| SelectPool[Select Pool Based on<br/>Available Space]
    
    SelectPool --> CalcSpace[Calculate Available Space<br/>for Each Pool]
    CalcSpace --> FilterPools[Filter Pools:<br/>- Skip Suspended<br/>- Skip Rebalancing<br/>- Check Disk Space]
    FilterPools --> WeightedRandom[Weighted Random Selection<br/>Based on Available Space]
    
    WeightedRandom --> PoolSelected[Pool Selected]
    UseExistingPool --> PoolSelected
    
    PoolSelected --> HashObject[Hash Object Name<br/>Using SipHash/CRC]
    HashObject --> SelectSet[Select Erasure Set<br/>setIndex = hash mod numSets]
    
    SelectSet --> CreateMetadata[Create FileInfo Metadata<br/>with Distribution Order]
    CreateMetadata --> CalcDistribution[Calculate Distribution:<br/>hashOrder based on object name]
    
    CalcDistribution --> ErasureEncode[Erasure Encode Object<br/>Split into Data + Parity Shards]
    
    ErasureEncode --> ShardCalc[For N drives in set:<br/>Data Shards = K<br/>Parity Shards = M<br/>N = K + M]
    
    ShardCalc --> ShuffleDisks[Shuffle Disks According<br/>to Distribution Order]
    ShuffleDisks --> WriteShard[Write Each Shard to<br/>Corresponding Drive]
    
    WriteShard --> WriteMetadata[Write xl.meta with:<br/>- Erasure Info<br/>- Distribution Array<br/>- Shard Index]
    
    WriteMetadata --> Complete[Write Complete]
    
    style Start fill:#e1f5ff
    style Complete fill:#d4edda
    style SelectPool fill:#fff3cd
    style HashObject fill:#fff3cd
    style ErasureEncode fill:#f8d7da
    style WriteShard fill:#d4edda
```

We are searching all the server pools in parallel to see if we find the object using the deterministic erasure set.


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
    
    subgraph "Layer 9: Storage API"
        S1[xlStorage<br/>Disk 1]
        S2[xlStorage<br/>Disk 2]
        S3[xlStorage<br/>Disk 16]
    end
    
    subgraph "Layer 10: Physical Disks"
        D1["/disk1/<br/>part.1<br/>~875 KB"]
        D2["/disk2/<br/>part.1<br/>~875 KB"]
        D3["/disk16/<br/>part.1<br/>~875 KB"]
    end
    
    subgraph "Metadata Write"
        M1[Create xl.meta<br/>with FileInfo]
        M2[Write to all<br/>16 disks]
        M3[Check quorum<br/>12/16]
        M4{Success?}
    end
    
    subgraph "Final"
        F1[✅ Return<br/>ObjectInfo]
        F2[❌ Revert<br/>& Error]
    end
    
    C --> H
    H --> A
    A --> SP
    SP --> ES
    ES --> EO
    EO --> L1
    
    L1 --> L2
    L2 --> RS
    RS --> L3
    
    L3 --> W1
    L3 --> W2
    L3 --> W3
    L3 --> W4
    L3 --> W5
    
    W1 --> S1
    W2 --> S2
    W5 --> S3
    
    S1 --> D1
    S2 --> D2
    S3 --> D3
    
    W1 --> L4
    W2 --> L4
    W3 --> L4
    W4 --> L4
    W5 --> L4
    
    L4 --> L5
    L5 -->|Yes| L1
    L5 -->|No| M1
    
    M1 --> M2
    M2 --> M3
    M3 --> M4
    
    M4 -->|Yes| F1
    M4 -->|No| F2
    
    style C fill:#e1f5ff
    style RS fill:#fff3cd
    style L4 fill:#fff3cd
    style M3 fill:#fff3cd
    style M4 fill:#fff3cd
    style F1 fill:#d4edda
    style F2 fill:#f8d7da
    style D1 fill:#cfe2ff
    style D2 fill:#cfe2ff
    style D3 fill:#ffc9c9
```


### Retrieving an Object (The GET Request)
To retrieve data, MinIO reverses the logic used during the write process.

* Step 1: Location Calculation: The client requests the file by name. MinIO runs the name through the same hash and modulus functions used during the PUT request to identify the correct drives immediately.

* Step 2: Retrieval: The system reads the object shards (blocks) from those specific drives.

* Step 3: Reassembly & Verification: The shards are passed back through the Erasure Code Engine. The engine reassembles the original object and verifies its integrity ("sanity check") to ensure no corruption occurred.

* Step 4: Delivery: The verified object is sent back to the client.

## Healing

<img width="868" height="930" alt="image" src="https://github.com/user-attachments/assets/e90b355a-6810-4648-823b-5baddacf7d64" />

<img width="688" height="945" alt="image" src="https://github.com/user-attachments/assets/0f3218db-1d01-4b99-a652-20172683fb60" />

Ref: https://minio-docs.tf.fo/operations/concepts/healing

<img width="1136" height="946" alt="image" src="https://github.com/user-attachments/assets/f3711710-afca-4123-8ba6-0c07469876b7" />

Ref: https://minio-docs.tf.fo/operations/data-recovery

## Site to Site Replication

<img width="827" height="427" alt="image" src="https://github.com/user-attachments/assets/dbb719ec-3363-45a8-922e-ebcc9a63dd14" />


## Use of distributed locking

<img width="2540" height="1216" alt="image" src="https://github.com/user-attachments/assets/231e0ec3-20e3-45c9-8c2e-b703ed8e4e9b" />

<img width="1375" height="872" alt="image" src="https://github.com/user-attachments/assets/6e6a6e02-44b8-4e32-b552-22a134de5f40" />

Ref: https://blog.min.io/minio-dsync-a-distributed-locking-and-syncing-package-for-go/

### Distributed lock management

Similar to distributed databases, Minio suffers from data consistency issues: while one client reads an object, another client may be modifying or deleting the object. To avoid inconsistencies. Minio specifically designed and implemented the dsync distributed lock manager to control data consistency.

* A lock request from any one node is broadcast to all online nodes in the cluster
* If consent is received from N/2+1 nodes, the acquisition is successful
* There is no master node, each node is peered to each other, and the stale lock detection mechanism is used between nodes to determine the status of nodes and the lock status
* Due to the simple design, it is relatively rough. It has certain defects, and supports up to 32 nodes. Scenarios where lock loss cannot be avoided. However, the available needs are basically met.

Lock throughput decreases as cluster grows.

Ref: https://e-whisper.com/posts/9462/

## MinIO Object Storage Gateway [Deprecated]

In addition to being a storage system service, Minio can also be used as a gateway, and the backend can be used with distributed file systems such as NAS systems and HDFS systems, or third-party storage systems such as S3 and OSS. With the Minio gateway, S3-compatible APIs can be added to these back-end systems for easy management and portability, because S3 APIs are already a de facto label in the object storage world.

<img width="984" height="760" alt="image" src="https://github.com/user-attachments/assets/402f972a-ea67-4c56-966f-938ad454f3a9" />

<img width="1024" height="800" alt="image" src="https://github.com/user-attachments/assets/f28dbef2-d8fa-41e2-9952-8476ef0ade6d" />

* MinIO introduced the gateway feature early on to help make the S3 API ubiquitous. From legacy POSIX-based SAN/NAS systems to modern cloud storage services, the different MinIO gateway modules brought S3 API compatibility where it did not exist previously.
* The primary objective was to provide sufficient time to port the applications over a modern cloud-native architecture.
* In the gateway mode, MinIO ran as a stateless proxy service, performing inline translation of the object storage functions from the S3 API to their corresponding equivalent backend functions.
* At any given time, the MinIO gateway service could be turned off and the only loss was S3 compatibility. The objects were always written to the backend in their native format, be it NFS or Azure Blob, or HDFS. 

<img width="697" height="354" alt="image" src="https://github.com/user-attachments/assets/5b0d98dd-b760-4cf6-98ca-17349637d92f" />


* The Gateway was initially developed to allow customers to use the S3 API to work with backends, such as NFS, Azure Blob and HDFS, that would not otherwise support it.
* The S3 API is ubiquitous (thanks in part to MinIO Gateway), but if we were to continue developing the MinIO Gateway, we would simply be perpetuating older technologies that are neither high-performance nor cloud-native. Also, addressing the ongoing technical challenges required to maintain MinIO Gateway for each backend are time and resource intensive so it makes much more sense to deprecate it entirely.

Reason for deprecation:
* The S3 API has evolved considerably since we started, and what began as inline translation morphed into something much more.
* Critical S3 capabilities like versioning, bucket replication, immutability/object locking, s3-select, encryption, and compression couldn’t be supported in the gateway mode without introducing a proprietary backend format.
* It would defeat the purpose of the gateway mode because the backend could no longer be read directly without the help of the gateway service.
* The backends would merely act as storage media for the gateway and you might as well run MinIO in server mode. Thus it became a compromise that MinIO no longer wanted to engage in. This meant it was time for us to let go. 

Ref: https://blog.min.io/minio-gateway-migration/ and https://blog.min.io/deprecation-of-the-minio-gateway/ 

## Beyond the Basics: What Else Can MinIO Do?

While we have only scratched the surface, MinIO is packed with advanced features:

* Versioning: Keep multiple versions of an object. Accidentally overwrote a file? No problem, just revert to an earlier version!

* Object Locking (WORM): Enforce "Write Once, Read Many" protection, making data immutable for compliance or security. Once written, it cannot be changed or deleted for a set period.

* Lifecycle Management: Automatically move or delete objects after a certain time, saving storage costs.

* Identity and Access Management (IAM): Control who can access what, just like in a big cloud environment. You can create users, groups, and define fine grained policies.

* Replication: Copy data automatically across different MinIO instances for disaster recovery and high availability.


## Encryption

<img width="1245" height="626" alt="image" src="https://github.com/user-attachments/assets/0571b920-39c8-494e-972d-0d48d7b06592" />

<img width="1245" height="626" alt="image" src="https://github.com/user-attachments/assets/b73e3512-bc87-4115-8689-d9f5872562fe" />

<img width="1245" height="678" alt="image" src="https://github.com/user-attachments/assets/eb419eca-2356-4086-ab6c-f453c25094d4" />


## TODO
* https://blog.min.io/minio-versioning-metadata-deep-dive/


## MinIO Backend Storage Metadata on Nodes

https://blog.min.io/minio-versioning-metadata-deep-dive/

Inside `.minio.sys/format.json` on a node

`minio1/.minio.sys$ cat format.json  | jq`

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

Explanation: Cluster has:

* 1 erasure set with 12 disks (UUIDs in the sets array)
* Deployment ID: f9a7a6ba-39d9-4483-bb47-fe86518bdc67 (the id field)
* Distribution Algorithm: SIPMOD+PARITY (the distributionAlgo field)
* This disk: 9ae64de8-1c75-46df-b09d-ad8b97f95313 (position 2 in the set)
