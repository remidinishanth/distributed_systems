## ZooKeeper

A centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services
* Distributed, **Consistent** Data Store
* Highly Available
* High performance
* Strictly ordered access 

#### Background
* Developed at Yahoo! Research
* Started as sub-project of Hadoop, now a top-level Apache project
* Development is driven by application needs

### Fault tolerance
Tolerates the loss of a minority of ensemble members and still function.
* To tolerate a loss of `n` members, we need atleast `2 * n + 1` nodes because for `n` to be the minority, we need atleast `2 * n + 1` nodes.
* It's recommended to have odd number(3 or 5) of nodes because we want to have majority surviving to continue to function, You don't get any benefit by having 6 nodes instead of 5 nodes, for both 5 or 6 nodes, we can only have loss of 2 nodes.
