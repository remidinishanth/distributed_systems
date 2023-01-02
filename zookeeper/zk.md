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
* As long as a majority of the servers are available, the ZooKeeper service will be available.
* To tolerate a loss of `n` members, we need atleast `2 * n + 1` nodes because for `n` to be the minority, we need atleast `2 * n + 1` nodes.
* Its good to form an Ensemble of odd number of nodes - as n (even number) nodes tends to allow same number of failure as of n-1 (odd number) nodes.
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

* ZooKeeper is replicated. Like the distributed processes it coordinates, ZooKeeper itself is intended to be replicated over a sets of hosts called an **ensemble**.
  - ZooKeeper service is an ensemble of servers that use replication (high availability)  
* During startup, When a leader doesn’t exist in the ensemble, ZooKeeper runs a leader election algorithm in the ensemble of servers. 
  - **One leader** and remaining all followers.
* Clients connect to a single ZooKeeper server. The client maintains a TCP connection through which it sends requests, gets responses, gets watch events, and sends heart beats. If the TCP connection to the server breaks, the client will connect to a different server.
  - Can read from **any** ZooKeeper server
  - Writes go through the leader & need majority consensus 

When one server goes down, clients will see a disconnect event and client will re-connect themselves to another member of the quorum.
<img width="967" alt="image" src="https://user-images.githubusercontent.com/19663316/210221305-6eb69a3d-fa1c-4f68-9165-4ea6db24891f.png">

### Zookeeper data structure

<img width="1047" alt="image" src="https://user-images.githubusercontent.com/19663316/210221770-a4c6b378-720f-46ad-9f56-5d2819a6b95a.png">

* Nodes can contain data, have children, or both
* Ephemeral nodes are associated with the session that created them. 
  - These nodes exists as long as the session that created the znode is active. When the session ends the znode is deleted.
  - They cannot have children, and disappear when that session ends
* Sequential nodes have an ever-increasing number attached to them 

#### File system analogy
* The name space provided by ZooKeeper is much like that of a standard file system. A name is a sequence of path elements separated by a slash (/). Every node in ZooKeeper's name space is identified by a path.
* Unlike is standard file systems, **each node in a ZooKeeper namespace can have data** associated with it as well as children. It is like having a file-system that allows a file to also be a directory. (ZooKeeper was designed to store coordination data: status information, configuration, location information, etc., so the data stored at each node is usually small, in the byte to kilobyte range.) 

![image](https://user-images.githubusercontent.com/19663316/210222011-ff442d5e-d0db-412f-80bf-f290823e1000.png)

* Basically, every node in a ZooKeeper tree is a ZNode.
  - **znode**: in-memory data node in ZooKeeper, organised in a hierarchical namespace (the data tree) 

* Znodes maintain a stat structure that includes version numbers for data changes, ACL changes, and timestamps, to allow cache validations and coordinated updates. 
  - Each time a znode's data changes, the version number increases. 
  - For instance, whenever a client retrieves data it also receives the version of the data.

* The data stored at each znode in a namespace is read and written atomically. Reads get all the data bytes associated with a znode and a write replaces all the data. Each node has an Access Control List (ACL) that restricts who can do what.

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
