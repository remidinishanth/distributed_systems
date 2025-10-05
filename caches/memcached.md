<img width="1107" height="559" alt="image" src="https://github.com/user-attachments/assets/140d7c24-51ec-4565-a36e-842e24ac51d0" />

Pre-memcache
<img width="1289" height="940" alt="image" src="https://github.com/user-attachments/assets/fdb27233-14d4-4351-9e97-f8f762238630" />

<img width="1311" height="911" alt="image" src="https://github.com/user-attachments/assets/46bf6b19-e8af-45b2-bcd4-3ba36c7f3d6b" />

Adding memcache

<img width="1311" height="885" alt="image" src="https://github.com/user-attachments/assets/dc13ef73-99b6-4ef9-b890-954aeca4b17d" />

<img width="1311" height="825" alt="image" src="https://github.com/user-attachments/assets/587947c6-8f6e-4737-8dea-81120b2daaf9" />

## What is memcache

> memcached is a high-performance, distributed memory object caching system, generic in nature, 
but originally intended for use in speeding up dynamic web applications by alleviating database load.

<img width="2390" height="1678" alt="image" src="https://github.com/user-attachments/assets/f53690c0-cd92-4543-aff6-62dc43f01740" />

Ref: https://memcached.org/about

Ref: https://engineering.fb.com/2013/04/15/core-infra/scaling-memcache-at-facebook/

<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/d3b833f3-cc07-4b78-ac9b-b09b6452548f" />

<img width="2222" height="672" alt="image" src="https://github.com/user-attachments/assets/dc939ab8-aef4-42ef-926d-1a8cb1df15bd" />


<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/ab51a430-2fa8-45fc-90f4-232164577dca" />

<img width="406" height="336" alt="image" src="https://github.com/user-attachments/assets/65fa0c86-4601-46cc-8c6f-fc1a9f9bcfca" />


## How Facebook served billions of requests per second Using Memcached

Memcached is a well known, simple, inmemory caching solution. Memcached was originally developed by Brad Fitzpatrick for LiveJournal in 2003. It was originally written in Perl, but is rewritten in C by Anatoly Vorobey.

<img width="1226" height="595" alt="image" src="https://github.com/user-attachments/assets/52ded5e8-61c1-4143-83d3-eb8149abc328" />

Ref: https://www.linuxjournal.com/article/7451

<img width="1295" height="1039" alt="image" src="https://github.com/user-attachments/assets/95982ff2-4d9e-44d2-adac-fcb476909381" />


Facebook took up the open-source version of Memcached and enhanced it to build a distributed key-value store. This enhanced version was known as `Memcache`.

* The open-source version Facebook started with provides a singlemachine in-memory hash table.
* > memcached provides no server-to-server coordination; it is an in-memory hash table running on a single server
* They took this basic building block, made it more efficient, and used it to build a distributed key-value store
that can process billions of requests per second that supports the world’s largest social network. 


The following properties greatly influence their design.
* First, users consume an order of magnitude more content than they create. This behavior results in a workload
dominated by fetching data and suggests that caching
can have significant advantages.
* Second, our read operations fetch data from a variety of sources such as
MySQL databases, HDFS installations, and backend services. This heterogeneity requires a flexible caching
strategy able to store data from disparate sources.

Major design goals:
* Any change must impact a userfacing or operational issue. Optimizations that have limited scope are rarely considered.
* They treat the probability of reading transient stale data as a parameter to
be tuned, similar to responsiveness. Willing to expose slightly stale data in exchange for insulating a
backend storage service from excessive load.

### How requests are served

<img width="1456" height="873" alt="image" src="https://github.com/user-attachments/assets/f3ddd7c9-9741-4cbd-be61-74e632dad787" />

> They choose to delete cached data instead of updating it because deletes are idempotent.

### Consistent Hashing

Items are distributed across the memcached servers through consistent hashing.

Consistent Hashing is a technique that allows the distribution of a set of keys across multiple nodes in a way that minimizes the impact of node failures or additions. 

For example, Each key is assigned to the node that falls closest to it in a clockwise direction.

<img width="1600" height="1461" alt="image" src="https://github.com/user-attachments/assets/e643f11a-14af-4e50-8e59-1557206a1d79" />

Clients maintain a map of all available servers, which is updated through an auxiliary configuration system.


#### **Client-server communication**: 
* Memcached servers do not communicate with each other.
* When appropriate, we embed the complexity of the system into a stateless
client rather than in the memcached servers.
  - This greatly
simplifies memcached and allows us to focus on making
it highly performant for a more limited use case.
* Keeping the clients stateless enables rapid iteration in the
software and simplifies our deployment process.
* Client logic is provided as two components. a library that can be embedded into applications or as a standalone proxy
named mcrouter. This proxy presents a memcached
server interface and routes the requests/replies to/from
other servers.

<img width="1628" height="1908" alt="image" src="https://github.com/user-attachments/assets/f4887672-a38a-423e-9751-4ab5c83096c1" />


### Reducing latency

At Facebook's scale, a single web request can trigger hundreds of fetch requests to retrieve data from Memcached servers. Consider a scenario where a user loads a popular page containing numerous posts and comments. 

<img width="1600" height="1123" alt="image" src="https://github.com/user-attachments/assets/b73c3b03-1147-495b-b09a-7080466e4419" />


#### Parallel requests and batching**: 
* They structure our webapplication code to minimize the number of network
round trips necessary to respond to page requests.
* They construct a directed acyclic graph (DAG) representing
the dependencies between data. By analyzing the DAG, the web server can determine the optimal order and grouping of data fetches.
* A web server uses this DAG to maximize the number of items that can be fetched concurrently. On average these batches consist
of 24 keys per request.


#### Using UDP 
Facebook employed a clever strategy to optimize network communication between the web servers and the Memcache server.

* For fetch requests, Facebook configured the clients to use UDP instead of TCP. 
* As you may know, UDP is a connectionless protocol and much faster than TCP. By using UDP, the clients can send fetch requests to the Memcache servers with less network overhead, resulting in faster request processing and reduced latency.
* However, UDP has a drawback: it doesn’t guarantee the delivery of packets. If a packet is lost during transmission, UDP doesn’t have a built-in mechanism to retransmit it. 
* To handle such cases, they treated UDP packet loss as a cache miss on the client side. If a response isn’t received within a specific timeframe, the client assumes that the data is not available in the cache and proceeds to fetch it from the primary data source.


### Problems with Caching

**Leases**: 
* Facebook introduced a new mechanism we call leases to address
two problems: stale sets and thundering herds.
* A stale set occurs when a web server sets a value in memcache
that does not reflect the latest value that should be
cached.
* This can occur when concurrent updates to memcache get reordered.

<img width="1311" height="936" alt="image" src="https://github.com/user-attachments/assets/123c1bbf-48f5-4482-9a40-ee5b4bfc87f3" />

* A thundering herd happens when a specific key undergoes heavy read and write activity. 
* As the write activity repeatedly invalidates the recently set values, many reads default to the more costly
path. The lease mechanism solves both problems.

<img width="1311" height="817" alt="image" src="https://github.com/user-attachments/assets/a2f9503c-15c4-4aea-a2e6-073ea7920c29" />

<img width="1221" height="948" alt="image" src="https://github.com/user-attachments/assets/387eed73-3b09-4dc9-a799-91c17defb263" />

<img width="1221" height="948" alt="image" src="https://github.com/user-attachments/assets/0f821203-7072-4fa5-9716-9a4229156e77" />

<img width="1253" height="948" alt="image" src="https://github.com/user-attachments/assets/1f87287f-b832-4873-9be1-4703f3314f21" />

<img width="1253" height="927" alt="image" src="https://github.com/user-attachments/assets/28a78291-b422-4521-863a-ad79593c2bd6" />
