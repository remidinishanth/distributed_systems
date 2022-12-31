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
