### Why should we reading the GFS paper?
* Incorporates many of the recurring themes in Distributed Systems: parallel performance, fault tolerance, replication, consistency.
* Successful real-world design(Academics didn't use single master). BigTable, MapReduce built on top of GFS.
* Well-written systems paper - details from the application to the network.


### Why is distributed storage hard?
* high performance -> shard data over many servers
* many servers -> constant faults
* fault tolerance -> replication
* replication -> potential inconsistencies
* better consistency -> low performance
