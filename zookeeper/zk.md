## ZooKeeper

A centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services
* Distributed, **Consistent** Data Store
* Highly Available
* High performance
* Strictly ordered access 

### Background
* Developed at Yahoo! Research
* Started as sub-project of Hadoop, now a top-level Apache project
* Development is driven by application needs

### ZooKeeper in the Hadoop ecosystem

<img width="1145" alt="image" src="https://user-images.githubusercontent.com/19663316/210209115-0ddf7bd3-9c2c-41c3-8c93-5afa78a5d72b.png">

![image](https://user-images.githubusercontent.com/19663316/210208983-7d1917d7-6ce4-4301-b26b-5a16825a3788.png)

![image](https://user-images.githubusercontent.com/19663316/210209213-f95f370c-9db6-48af-97b3-d668678523a9.png)


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
