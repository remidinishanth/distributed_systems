---
layout: page
title: "Cassandra"
category: "cassandra"
---

### Goal

>  Goal: How can we build an ultra-high performance, low-latency multi-table database that can grow incrementally to handle massive-scale data?

Cassandra is a large-scale open-source NoSQL wide-column distributed database. It was developed at Facebook and its design was influenced by Amazon Dynamo and Google Bigtable.
* Cassandra was initially designed at Facebook using a staged event-driven architecture (SEDA).
* This initial design implemented a combination of Amazon’s Dynamo distributed storage and replication techniques and Google’s Bigtable data and storage engine model.

<img width="1516" alt="image" src="https://github.com/user-attachments/assets/05658a19-7c38-4ee6-912e-1f177f4dac82">


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


* The key goals of Cassandra were **high performance**, **incremental scalability**, and **high availability**.
  * Since the CAP theorem tells us we must choose between consistency and availability if a network partition occurs, Cassandra opted for high availability at the expense of consistency, implementing an eventually consistent model.
  * This also helps with performance since there is no need to take and release locks across all nodes that store replicas of data that’s being modified.

![image](https://github.com/user-attachments/assets/1bbe43b9-ed7e-4afa-aaaf-c1cd640a6a55)



Cassandra explicitly chooses not to implement operations that require cross-partition coordination as they are typically slow and hard to provide highly available global semantics. For example, Cassandra **does not support**:
* Cross-partition transactions
* Distributed joins
* Foreign keys or referential integrity.

### Wide-columns storage

NoSQL databases are often designed as **wide-column** data stores. 

* Traditional relational databases comprise tables that have a small, fixed set of fields, or columns, that are defined when the table is created (e.g., user ID, name, address, city, state, country, phone). 
* In wide-column data stores, a row may have an arbitrary number of columns and columns may be created dynamically after the tables are created. The only columns that are expected to be present are the key (which is essential to identify a row uniquely and look up a row efficiently) and any columns that user software deems to be mandatory for the application.

The concept of wide columns changes the way columns are used in a data store. With a traditional database, the programmer is aware of the columns (fields) within a table. A wide-column table, on the other hand, may contain different columns for different rows and the programmer may need to iterate over their names.

![image](https://github.com/user-attachments/assets/4120a0a6-b59f-451d-a1a6-dbdd1d928008)

For example a user might have multiple phone numbers that, instead of being stored in a separate table with foreign keys that identify the owner of each number, are co-located with the user’s information. Each number is a new column (e.g., phone-1, phone-2, …). There’s also no limit to the number of columns in a row.


* In Bigtable, columns might be a list of URLs. In the Bigtable example, a URL of a web page serves as a row key but column names within the URL have names of URLs that contain links to that page and the value of each of those columns contains the link text that appears within the page.
* We can also consider a related example where a row key identifies a user and each column name is a URL that the user visited and its value is the time of the visit. This is a powerful departure from conventional database fields because the column name itself may be treated as a form of value: by iterating over the columns, one can obtain a list of URLs and then look up their associated value (e.g., link text or time of visit).

![image](https://github.com/user-attachments/assets/415fe5e7-6f87-409b-81ba-35d1eda0777b)


Wide-column data stores such as Bigtable and Cassandra essentially have no limit on the number of columns that can be associated with a row. In the case of Cassandra, the limit is essentially the storage capacity of the node.

![image](https://github.com/user-attachments/assets/13a62e15-c17f-4c8a-85cc-45a7cdf42f38)
