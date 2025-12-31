MinIO just does one thing - Object storage for Private cloud

<img width="1105" height="588" alt="image" src="https://github.com/user-attachments/assets/b9fdf6eb-b52c-460d-89e0-fc90a2762f2f" />

<img width="689" height="241" alt="image" src="https://github.com/user-attachments/assets/a5f7c369-978a-4944-94d7-46ce8c87c160" />


<img width="1142" height="499" alt="image" src="https://github.com/user-attachments/assets/ea1e33e0-cac3-47b9-b975-987243174e59" />

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/82863b5c-7278-4c28-981b-94dd71b614db" />


Legacy object storage architecture
<img width="1591" height="863" alt="image" src="https://github.com/user-attachments/assets/86218141-aa98-478d-9b7b-02a18b6faf71" />


## Architecture

All the nodes running distributed MinIO setup are recommended to be homogeneous, i.e. same operating system, same number of drives and same network interconnects.

Start distributed MinIO instance on n nodes with m drives each mounted at `/export1` to `/exportm` (pictured below), by running this command on all the `n` nodes:

<img width="8000" height="4500" alt="image" src="https://github.com/user-attachments/assets/19b44e38-a8c8-4daa-89ad-3d9a6854ecdd" />
* No master server, no metadata server or anything

Ref: https://github.com/minio/minio/blob/master/docs/distributed/README.md


We will see Pools and Erasure coding in the following sections
<img width="961" height="440" alt="image" src="https://github.com/user-attachments/assets/d916ce30-b3c0-46d5-9ab9-f31801b8872b" />


## MinIO Server Pools

![serverpools](https://github.com/user-attachments/assets/a25b361c-e253-4c06-983a-e95b4d0ae464)

A server pool is a set of minio server nodes which pool their drives and resources, creating a unit of expansion. All nodes in a server pool share their hardware resources in an isolated namespace.  

The other important point here involves rebalance-free, non-disruptive expansion. With MinIO’s server pool approach - rebalancing is not required to expand. Ref: https://blog.min.io/no-rebalancing-object-storage/

A MinIO cluster is built on server pools, and server pools are built on erasure sets.

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
