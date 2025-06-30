---
layout: page
title: "Readme"
category: "etcd"
---

* https://www.youtube.com/watch?v=OmphHSaO1sE Awesome explanation

* https://www.ibm.com/cloud/learn/etcd
  * the **fault-tolerant open source key-value database** that serves as the **primary data backbone for Kubernetes** and other distributed platforms.

* etcd is a distributed reliable key-value store for the most critical data of a distributed system, with a focus on being:
  * Simple: well-defined, user-facing API (gRPC)
  * Secure: automatic TLS with optional client cert authentication
  * Fast: benchmarked 10,000 writes/sec
  * Reliable: properly distributed using Raft
  
![image](https://user-images.githubusercontent.com/19663316/195422255-8c417e55-8238-4181-98f7-63d8f14e4b16.png)
  
  
![image](https://user-images.githubusercontent.com/19663316/195421754-2c3920bf-5b92-4312-89f1-c3ee7001e80a.png)

* Multiple nodes cooperate with each other through the raft consensus algorithm. 
* The algorithm elects a master node as the leader, which is responsible for data synchronization and distribution. 
* Quorum is a key concept in etcd. It is defined as (n+1)/2, indicating that more than half of the nodes in the cluster constitute a quorum. 

Also read https://learnk8s.io/etcd-kubernetes
