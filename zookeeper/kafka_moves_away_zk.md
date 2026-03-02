Consensus is the most complex problem in Distributed Systems.

Here's a simple explanation of how it works in Apache Kafka 👇

<img width="2798" height="1842" alt="image" src="https://github.com/user-attachments/assets/6882a16a-3ad6-459c-87d5-452ad48f3f62" />

https://www.automq.com/blog/what-is-kafka-zookeeper-in-kafka

You can think of Apache ZooKeeper as a “cluster manager” for Apache Kafka. While Kafka handles the heavy lifting of moving data (messages), ZooKeeper manages the metadata and coordination required to keep the distributed system stable and consistent. ZooKeeper handles important tasks such as:

* Choosing a leader broker (called controller election)
* Keeping track of topics, partitions, and replicas
* Detecting when a broker joins or leaves the cluster

<img width="1000" height="660" alt="image" src="https://github.com/user-attachments/assets/03d4f0e0-fb8c-43a6-b5cd-4038b8f440be" />

Ref: https://www.confluent.io/learn/zookeeper-kafka/#apache-kafka-without-zookeeper-introduction-to-kraft

### Apache Kafka until 4.0
📖 Kafka used ZooKeeper to solve distributed consensus up until Apache Kafka 4.0 (March 2025).

ZooKeeper solves consensus through its own consensus algorithm called ZAB (ZooKeeper Atomic Broadcast).

Kafka used ZK to elect one single Controller broker amongst the whole cluster (all possible brokers).

All cluster-wide metadata decisions were taken by this single Controller node - e.g electing regular partition leaders.
Those actions were then consistently persisted in ZK (through Zab).

This is called a centralized coordination model. There's one broker that calls the shots.

### From 4.0

<img width="2126" height="994" alt="image" src="https://github.com/user-attachments/assets/7da227b5-98c4-4ac3-a1a1-58295d26b1d5" />

<img width="1314" height="860" alt="image" src="https://github.com/user-attachments/assets/63c01a15-fde6-4ae8-8271-046650491954" />

<img width="1200" height="675" alt="image" src="https://github.com/user-attachments/assets/7ee6688b-9e3e-4312-ae81-f4ffdf21e684" />



📖 In Apache Kafka 4.0 (and after), Kafka moved on to use its own consensus algorithm to elect a controller.

Kafka now uses N controller brokers (usually 3) that run a Raft-like consensus algorithm called KRaft.

The controllers choose a leader amongst them via the Raft-based algorithm. That leader assumes the Active Controller role and starts taking cluster-wide decisions.

The decisions are persisted in a single metadata log where all the cluster metadata is stored. All controllers replicate this log. 👌

Regular brokers read this log too - but they only read committed updates.

An update in the metadata log is considered committed only when a majority (a quorum) of controller nodes persists it.

In this way, all leader election throughout Kafka is done INDIRECTLY through the single quorum: 👇

The regular elections are made by the active controller, committed through the quorum and propagated through asynchronous replication.

In other words:
• controller election is done through KRaft
• leader election is done through the Controller 💡
• brokers assume leadership as they learn about it through the log ⏳

### Alternative

An alternative, more decentralized, Raft-native design is RedPanda's. 🐼

Two key differences:

1. each partition is its own separate Raft group 🧠

All leader election is done through its own Raft quorum - i.e with Raft's RequestVote RPC

2. their metadata log has a Raft quorum consisting of ALL nodes in the cluster

Unlike just the controllers in Kafka, any RedPanda broker is eligible to become the active controller there. This is done through a shared Raft quorum consisting of the whole cluster of nodes.

<img width="1600" height="880" alt="image" src="https://github.com/user-attachments/assets/565996f9-c80e-4460-978e-985333ed98d0" />


Ref: https://x.com/BdKozlovski/status/1968314511037018510
