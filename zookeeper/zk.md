## ZooKeeper

> “Because coordinating distributed systems is a Zoo”

### Background
* Developed at Yahoo! Research
* Started as sub-project of Hadoop, now a top-level Apache project
* Development is driven by application needs

### ZooKeeper in the Hadoop ecosystem

<img width="1145" alt="image" src="https://user-images.githubusercontent.com/19663316/210209115-0ddf7bd3-9c2c-41c3-8c93-5afa78a5d72b.png">

![image](https://user-images.githubusercontent.com/19663316/210208983-7d1917d7-6ce4-4301-b26b-5a16825a3788.png)

![image](https://user-images.githubusercontent.com/19663316/210209213-f95f370c-9db6-48af-97b3-d668678523a9.png)

### Motivation

* In the past: a single program running on a single computer with a single CPU
* Today: applications consist of independent programs running on a changing set of computers
* Difficulty: coordination of those independent programs
* Developers have to deal with coordination logic and application logic at the same time

> ZooKeeper: designed to relieve developers from writing coordination logic code.

A centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services
* Distributed, **Consistent** Data Store
* Highly Available
* High performance
* Strictly ordered access 


### Highly Available
Tolerates the loss of a **minority** of ensemble members and still function.
* To tolerate a loss of `n` members, we need atleast `2 * n + 1` nodes because for `n` to be the minority, we need atleast `2 * n + 1` nodes.
* It's recommended to have odd number(3 or 5) of nodes because we want to have majority surviving to continue to function, You don't get any benefit by having 6 nodes instead of 5 nodes, for both 5 or 6 nodes, we can only have loss of 2 nodes.

### High Performance
* All data is stored **in memory**
* Performance measured around 50,000
operations/second
* Particularly fast for read performance, _built for read dominant workloads_

### Strictly ordered access 
* **Atomic Writes**
* In the order you sent them
* Changes always seen in the order they occurred
* **Reliable**, no writes acked will be dropped 

### Basic Cluster Interactions

<img width="967" alt="image" src="https://user-images.githubusercontent.com/19663316/210209709-8c3e372d-dc6a-4440-b014-3cc66c36f528.png">

### Uses of Zookeeper
* Naming service
  - Identifying nodes in a cluster by name (“DNS” for nodes)
* Configuration management
  - Up-to-date system config info for a joining node
* Cluster management
  - Joining / leaving of nodes, real-time node status
* Leader election
  - Electing a node as leader for coordination purpose
* Locking and synchronization service
* Highly reliable data registry
