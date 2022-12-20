# developing_as_developer

* https://thesmartcoder.dev/fantastic-books-by-developers-for-developers/

TODO: **gRPC** - Why do we use it? Why is it becoming popular? https://people.cs.rutgers.edu/~pxk/417/notes/rpc.html

TODO: **GraphQL** - Why do we neeed it? Usecases, examples, sangria. Checkout https://medium.com/@leeshapton/mental-maps-for-teaching-graphql-to-beginners-9db9b85ac957

 * An increasingly popular approach to tailor APIs to individual use cases is to use graph-based
  APIs. A graph-based API exposes a schema composed of types,
  fields, and relationships across types. The API allows a client to
  declare what data it needs and make a request.
 * It provides a schema of the data in the API and gives clients the power to ask for exactly what they need.

**Jinja2** - used for creating templated outputs. Ref: https://ttl255.com/jinja2-tutorial-part-1-introduction-and-variable-substitution/ You can also use macros.

Read about Shared Nothing Architecture. Differences with Microservices architecture.

Read about **bazel** build and make files.

#### Read about Service discovery mechanism.
* Consul

Software Engineering at Google: Lessons Learned from Programming Over Time.

Clean code tips: (Book Clean Code : Robert C Martin) https://medium.com/storyblocks-engineering/these-four-clean-code-tips-will-dramatically-improve-your-engineering-teams-productivity-b5bd121dd150

Read: The Good Parts of AWS

CQRS stands for Command Query Responsibility Segregation.  At its heart is the notion that you can use a different model to update information than the model you use to read information. Many systems do fit a CRUD mental model, and so should be done in that style. There is an inherent
replication lag between the time a change has been applied on the write path and the read path has received and applied it, which
makes the system sequentially consistent. Ref: https://martinfowler.com/bliki/CQRS.html
* This is used in URL shortener, once we shorten the URL, we don't update the same shortened-url.

#### Sharding Strategies of key-value store
* The mapping between keys and partitions, and other metadata, is
typically maintained in a strongly-consistent configuration store,
like etcd or Zookeeper.
* Ideally, if a partition is added, only 𝐾/𝑁 keys should be shuffled
around, where 𝐾 is the number of keys and 𝑁 the number of
partitions. A hashing strategy that guarantees this property is
called stable hashing. Ring hashing is an example of stable hashing. With ring hashing,
a function maps a key to a point on a circle. The circle is then split
into partitions that can be evenly or pseudo-randomly spaced, depending on the specific algorithm. When a new partition is added,
it can be shown that most keys don’t need to be shuffled around. E.g Consistent Hashing.

#### Introduction to Distributed Systems

<img width="1468" alt="image" src="https://user-images.githubusercontent.com/19663316/208266696-6a229b84-7349-4d26-b35e-ebff3e90c4c7.png">

* TODO: https://www.freecodecamp.org/news/a-thorough-introduction-to-distributed-systems-3b91562c9b3c/

![image](https://user-images.githubusercontent.com/19663316/161559354-8bb255d2-7222-4164-9db1-2c2ac7ca6115.png)

![image](https://user-images.githubusercontent.com/19663316/161559627-b0fa0f35-6e56-46a4-af2c-482c77538cb5.png)

A Range of Interesting Problems for Distributed System Designers
* Peer-to-Peer(P2P) Systems [Gnutella, Kazaa, BitTorrent]
* Cloud Infrastructures [AWS, Azure, Google Cloud]
* Cloud Storage [Key-Value stores, NoSQL, Cassandra]
* Cloud Programming [MapReduce, Storm, Pregel]
* Coordination [Paxos, Leader Election, Snapshots]
* Managing many clients and Servers Concurrently [Concurrency Control, Replication Control]

Core Concepts of Distributed Systems
* Gossip
* Membership
* Distributed Hash Tables(DHTs)

Peer-to-peer systems use DHTs, Key-Value/NoSQL stores uses DHTs, gossip, membership.

![image](https://user-images.githubusercontent.com/19663316/161790715-9d11faed-c5a0-45bf-b51d-8fb759f06f76.png)
