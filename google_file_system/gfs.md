## The Google File System
Sanjay Ghemawat, Howard Gobioff, and Shun-Tak Leung

SOSP 2003

### Why are we reading this paper?
* distributed storage is a key abstraction
* Incorporates many of the recurring themes in Distributed Systems: 
  * parallel performance 
  * fault tolerance 
  * replication and 
  * consistency.
* Successful real-world design(Academics didn't use single master). BigTable, MapReduce built on top of GFS.
* Well-written systems paper - details from the application to the network.


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
  * Suppose the consistency is not there, for example: if the requests `C1` and `C3` go to server 1 and `C2` and `C3` go to server 2, then we might see different responses if servers doesn't have sync the write data consistently.
