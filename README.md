# developing_as_developer

* https://thesmartcoder.dev/fantastic-books-by-developers-for-developers/

TODO: gRPC - Why do we use it? Why is it becoming popular?

TODO: GraphQL - Why do we neeed it? Usecases, examples, sangria. Checkout https://medium.com/@leeshapton/mental-maps-for-teaching-graphql-to-beginners-9db9b85ac957

 * An increasingly popular approach to tailor APIs to individual use cases is to use graph-based
  APIs. A graph-based API exposes a schema composed of types,
  fields, and relationships across types. The API allows a client to
  declare what data it needs and make a request.

Jinja2 - used for creating templated outputs. Ref: https://ttl255.com/jinja2-tutorial-part-1-introduction-and-variable-substitution/ You can also use macros.

Read about Shared Nothing Architecture. Differences with Microservices architecture.

Read about bazel build and make files.

Read about Service discovery mechanism.

Software Engineering at Google: Lessons Learned from Programming Over Time.

Clean code tips: (Book Clean Code : Robert C Martin) https://medium.com/storyblocks-engineering/these-four-clean-code-tips-will-dramatically-improve-your-engineering-teams-productivity-b5bd121dd150

Read: The Good Parts of AWS

CQRS stands for Command Query Responsibility Segregation.  At its heart is the notion that you can use a different model to update information than the model you use to read information. Many systems do fit a CRUD mental model, and so should be done in that style. There is an inherent
replication lag between the time a change has been applied on the write path and the read path has received and applied it, which
makes the system sequentially consistent. Ref: https://martinfowler.com/bliki/CQRS.html
* This is used in URL shortener, once we shorten the URL, we don't update the same shortened-url.

Sharding Strategies of key-value store
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
