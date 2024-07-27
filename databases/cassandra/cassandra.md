### Goal

>  Goal: How can we build an ultra-high performance, low-latency multi-table database that can grow incrementally to handle massive-scale data?

Cassandra is a large-scale open-source NoSQL wide-column distributed database. It was developed at Facebook and its design was influenced by Amazon Dynamo and Google Bigtable.
* Cassandra was initially designed at Facebook using a staged event-driven architecture (SEDA).
* This initial design implemented a combination of Amazon’s Dynamo distributed storage and replication techniques and Google’s Bigtable data and storage engine model. 

### Abstract

* Cassandra is a distributed storage system for managing very
large amounts of structured data spread out across many
commodity servers, while providing highly available service
with no single point of failure.

* Cassandra system was designed to run on
cheap commodity hardware and handle high write through-
put while not sacrificing read efficiency

### Introduction

* Cassandra was designed to fulfill the storage needs of the Inbox Search problem.
  * Enables users to search through their Facebook Inbox.
  * This involves storing reverse indices of Facebook messages that users send and receive.
  * A reverse index maps message contents to messages. That is, the system can look up a word and find all messages that contain text with that word in it 
  
* At Facebook this meant the system was required to handle a very high write throughput, billions
of writes per day, and also scale with the number of users.


* The key goals of Cassandra were high performance, incremental scalability, and high availability.
 * Since the CAP theorem tells us we must choose between consistency and availability if a network partition occurs, Cassandra opted for high availability at the expense of consistency, implementing an eventually consistent model.
* This also helps with performance since there is no need to take and release locks across all nodes that store replicas of data that’s being modified.

![image](https://github.com/user-attachments/assets/1bbe43b9-ed7e-4afa-aaaf-c1cd640a6a55)



Cassandra explicitly chooses not to implement operations that require cross-partition coordination as they are typically slow and hard to provide highly available global semantics. For example, Cassandra **does not support**:
* Cross-partition transactions
* Distributed joins
* Foreign keys or referential integrity.
