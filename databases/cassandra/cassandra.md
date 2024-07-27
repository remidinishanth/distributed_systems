### Goal

>  Goal: How can we build an ultra-high performance, low-latency multi-table database that can grow incrementally to handle massive-scale data?

Cassandra is a large-scale open-source NoSQL wide-column distributed database. It was developed at Facebook and its design was influenced by Amazon Dynamo and Google Bigtable.

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
* At Facebook this meant the system was required to handle a very high write throughput, billions
of writes per day, and also scale with the number of users.
