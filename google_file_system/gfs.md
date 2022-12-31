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
