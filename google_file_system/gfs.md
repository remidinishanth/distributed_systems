## The Google File System

Sanjay Ghemawat, Howard Gobioff, and Shun-Tak Leung (SOSP 2003)

Google needs a distributed file system that matches its applications needs. Solution: GFS.

GFS: _Scalable distributed file system for large distributed data-intensive applications._

### Why are we reading this paper?
* Distributed storage is a key abstraction
* Incorporates many of the recurring themes in Distributed Systems: 
  * parallel performance 
  * fault tolerance 
  * replication and 
  * consistency.
* Successful real-world design(Academics didn't use single master). BigTable, MapReduce built on top of GFS.
* Well-written systems paper - details from the application to the network.
* Successfully applied the single master and weak consistency.


### Why is distributed storage hard?

* high performance ⇒ shard data over many servers (To achieve high performance, a common strategy is to shard data over many servers)
* many servers ⇒ constant faults (Having many servers lead to constant faults)
* fault tolerance ⇒ replication (To implement fault tolerance, one strategy is to make use of replication)
* replication ⇒ potential inconsistencies (With different replica of the same data, however, potential inconsistencies between each replica will occur)
* better consistency ⇒ low performance

In GFS, we will see that consistency is traded off for simpler design, greater performance, and high availability.


### What would we like for consistency?
* Ideal model: same behavior as a single server
* Ideal server executes client operations one at a time; even if multiple clients issued operations concurrently
  * reads reflect previous writes even if server crashes and restarts all clients see the same data
  * thus: suppose `C1` and `C2` write **concurrently**, and after the writes have completed, `C3` and `C4` read. what can they see?
    * `C1`: `Wx(1)`
    * `C2`: `Wx(2)`
    * `C3`:         `Rx?`
    * `C4`:             `Rx?`
  * answer: either `1` or `2`, but both have to see the same value.
  * This is a "strong" consistency model.


### Replication for fault-tolerance makes strong consistency tricky.
  * a simple but broken replication scheme:
    - two replica servers, S1 and S2
    - clients send writes to both, in parallel
    - clients send reads to either
  * In our example, C1's and C2's write messages could arrive in
    - different orders at the two replicas
    - if C3 reads S1, it might see x=1
    - if C4 reads S2, it might see x=2
  * or what if S1 receives a write, but 
    - the client crashes before sending the write to S2?
  * that's not strong consistency!
  * better consistency usually requires communication to
    - ensure the replicas stay in sync -- can be slow!
  * lots of tradeoffs possible between performance and consistency

## GFS Key Ideas

### Motivating Example

#### Google
* Crawl the whole web
* Store it all on “one big disk”
* Process users’ searches on “one big CPU”
* More storage, CPU required than one PC can offer
* Custom parallel supercomputer: expensive (so much so, not really available today)

#### Cluster of PCs as Supercomputer
* More than 15,000 commodity-class PC's.
* Multiple clusters distributed worldwide.
* Thousands of queries served per second.
* One query reads 100's of MB of data.
* One query consumes 10's of billions of CPU cycles.
* Google stores dozens of copies of the entire Web!

Conclusion: Need **large**, **distributed**, highly **fault-tolerant** file system. 

### Serving a Google Search query

Steps in query answering:

* DNS picks out a cluster geographically close to user.
* A Google Web server machine (GWS) at the cluster is chosen based on load-balancing.
* Query is sent to index servers. Each index holds inverted index for a random subset of documents.
* Index server returns list of docid's, with relevance scores, sorted in decreasing order.
* GWS merges lists, gets overall list sorted by query.
* Full text of actual documents are divided among file servers. GWS sends each file server the query plus the list of docid's on that server.
* File server returns document summary for that query (excerpt containing query word with query words highlighted)
* GWS assembles summaries, advertisements from ad server, spelling correction suggestions, returns to user.

![image](https://user-images.githubusercontent.com/19663316/210136235-b3c4ecbf-b7a9-4d35-808f-3aced4b620e3.png)

"On average, a single query in Google reads hundreds of megabytes of data and consumes tens of billions of CPU cycles. Supporting a peak request stream of thousands of queries per second requires an infrastructure comparable in size to that of the largest supercomputer installations. Combining more than 15,000 commodity-class PC's with fault-tolerant software creates a solution that is more cost-effective than a comparable system build out of a smaller number of high-end servers."

### Google Platform Characteristics
* 100s to 1000s of PCs in cluster
* Cheap, commodity parts in PCs
* Many modes of failure for each PC:
  - App bugs, OS bugs
  - Human error
  - Disk failure, memory failure, net failure, power supply failure
  - Connector failure
* **Monitoring**, **fault tolerance**, **auto-recovery** essential

### Design Motivations
1. GFS runs on a large number of machines. Failures occur regularly, so **fault-tolerance** and auto-recovery need to be built in.
2. File sizes are much **larger**. Standard I/O assumptions (e.g. block size) have to be reexamined.
3. Record appends are the prevalent form of writing. Need good semantics for concurrent appends to the same file by multiple clients.
4. Google applications and GFS are both designed in-house - so they can and should be co-designed. 


### What were the problems GFS was trying to solve?
Google needed a large-scale and high-performant unified storage system for many of its internal services such as MapReduce, web crawler services. 

In particular, this storage system must:
* Be global. Any client can access (read/write) any file. This allows for sharing of data among different applications.
* Support automatic sharding of large files over multiple machines. This improves performance by allowing parallel processes on each file chunk and also deals with large files that cannot fit into a single disk.
* Support automatic recovery from failures.
* Be optimized for sequential access to huge files and for read and append operations which are the most common.

In particular, GFS is optimized for high sustained bandwidth (target applications place a premium on processing data in bulk at a high rate), but not necessarily for low latency (GFS is typically used for internal services and is not client-facing).


### Architecture

```
  clients (library, RPC -- but not visible as a UNIX FS)
  
  coordinator tracks file names
  
  chunkservers store 64 MB chunks
  
  big files split into 64 MB chunks, on lots of chunkservers
    big chunks -> low book-keeping overhead
  
  each chunk replicated on 3 chunkservers
```

![image](https://user-images.githubusercontent.com/19663316/210136483-f4048736-f908-48e9-8743-c155b244eb6d.png)


* GFS consists of a single master and multiple chunkservers and is accessed by multiple clients. 
* Files are divided into fixed-sized chunks of 64MB. 
* Each chunk has an immutable and globally unique chunk handler, which is assigned by the master at the time of chunk creation. 
* By default, each file chunk is replicated on 3 different chunkservers.


### Analogy with File system

On a single-machine FS:
* An upper layer maintains the metadata.
* A lower layer (i.e. disk) stores the data in units called “blocks”.
* Upper layer store

In the GFS:
* A master process maintains the metadata.
* A lower layer (i.e. a set of chunkservers) stores the data in units called “chunks”. 

<img width="943" alt="image" src="https://user-images.githubusercontent.com/19663316/210136593-cc479533-3ce6-4162-b088-7951ed3ff6c2.png">

Very important: **data flow is decoupled from control flow**
* Clients interact with the master for metadata operations
* Clients interact directly with chunkservers for all files operations
* This means performance can be improved by scheduling expensive data flow
based on the network topology

Neither the clients nor the chunkservers cache file data
* Working sets are usually too large to be cached, chunkservers can use Linux’s buffer cache
