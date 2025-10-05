Also refer to https://timilearning.com/posts/mit-6.824/lecture-16-memcache-at-facebook/ and https://blog.bytebytego.com/p/how-facebook-served-billions-of-requests

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

<img width="1247" height="717" alt="image" src="https://github.com/user-attachments/assets/f7916d43-8d8e-4537-9561-c4ed599de38a" />


<img width="2390" height="1678" alt="image" src="https://github.com/user-attachments/assets/f53690c0-cd92-4543-aff6-62dc43f01740" />

Ref: https://memcached.org/about

Ref: https://engineering.fb.com/2013/04/15/core-infra/scaling-memcache-at-facebook/

<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/d3b833f3-cc07-4b78-ac9b-b09b6452548f" />

<img width="2222" height="672" alt="image" src="https://github.com/user-attachments/assets/dc939ab8-aef4-42ef-926d-1a8cb1df15bd" />


<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/ab51a430-2fa8-45fc-90f4-232164577dca" />

<img width="406" height="336" alt="image" src="https://github.com/user-attachments/assets/65fa0c86-4601-46cc-8c6f-fc1a9f9bcfca" />


## How Facebook served billions of requests per second Using Memcached

<img width="1064" height="697" alt="image" src="https://github.com/user-attachments/assets/69e45fc0-5958-4185-aa91-3c55f598ee9c" />


Memcached is a well known, simple, inmemory caching solution. Memcached was originally developed by Brad Fitzpatrick for LiveJournal in 2003. It was originally written in Perl, but is rewritten in C by Anatoly Vorobey.

<img width="1226" height="595" alt="image" src="https://github.com/user-attachments/assets/52ded5e8-61c1-4143-83d3-eb8149abc328" />

Ref: https://www.linuxjournal.com/article/7451

<img width="1295" height="1039" alt="image" src="https://github.com/user-attachments/assets/95982ff2-4d9e-44d2-adac-fcb476909381" />

Facebook took up the open-source version of Memcached and enhanced it to build a distributed key-value store. This enhanced version was known as `Memcache`.

* The open-source version Facebook started with provides a singlemachine in-memory hash table.
* > memcached provides no server-to-server coordination; it is an in-memory hash table running on a single server
* They took this basic building block, made it more efficient, and used it to build a distributed key-value store
that can process billions of requests per second that supports the world’s largest social network. 

<img width="603" height="304" alt="image" src="https://github.com/user-attachments/assets/fbaa7888-5725-4d8a-826f-e93c002d7f51" />


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

<img width="911" height="597" alt="image" src="https://github.com/user-attachments/assets/a50cbb6b-c727-42d7-a469-edbacfe6eef1" />


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

<img width="1396" height="1600" alt="image" src="https://github.com/user-attachments/assets/e9ae3f06-cc56-4697-9839-b02ce09febe8" />


**Leases**: 
* Facebook introduced a new mechanism we call leases to address
two problems: stale sets and thundering herds.
* A stale set occurs when a web server sets a value in memcache
that does not reflect the latest value that should be
cached.
* This can occur when concurrent updates to memcache get reordered.

How lease token solves this

* Intuitively, a memcached instance gives a lease to a
client to set data back into the cache when that client experiences a cache miss. 
* The lease is a 64-bit token bound to the specific key the client originally requested.
* The client provides the lease token when setting the value in the cache.
* With the lease token, memcached can verify and determine whether the data should be stored and thus arbitrate concurrent writes.
* Verification can fail if memcached has invalidated the lease token due to receiving a delete request for that item.
* Leases prevent stale sets in a manner similar to how load-link/storeconditional operates

<img width="663" height="468" alt="image" src="https://github.com/user-attachments/assets/0bd37ae8-8618-4a02-9ab5-b8902b533065" />

<img width="1573" height="1600" alt="image" src="https://github.com/user-attachments/assets/1837fcf0-c0c4-4e93-ac16-14d2a7748d41" />



* A thundering herd happens when a specific key undergoes heavy read and write activity. 
* As the write activity repeatedly invalidates the recently set values, many reads default to the more costly
path. The lease mechanism solves both problems.

All servers see a cache miss and everyone reaches out to database, increasing the load on the database.

<img width="645" height="436" alt="image" src="https://github.com/user-attachments/assets/17d86474-5eff-42d2-a49b-b101c462bc59" />

Caches arbitrates access to the database:
* A slight modification to leases also mitigates thundering herds.
* Each memcached server regulates the rate at which it returns tokens.
* By default, we configure these servers to return a token only once every 10 seconds per
key.
* Requests for a key’s value within 10 seconds of a token being issued results in a special notification telling
the client to wait a short amount of time.
* Typically, the client with the lease will have successfully set the data
within a few milliseconds.
* Thus, when waiting clients retry the request, the data is often present in cache.

<img width="1311" height="817" alt="image" src="https://github.com/user-attachments/assets/a2f9503c-15c4-4aea-a2e6-073ea7920c29" />


### Many memcache servers in one cluster 

When you add more webservers, we would need more memcache servers

<img width="1316" height="604" alt="image" src="https://github.com/user-attachments/assets/024d5696-f509-4f03-a628-0e028ea7b0be" />

This would lead to every server communicating to every memcache server, all to all communication

<img width="917" height="692" alt="image" src="https://github.com/user-attachments/assets/752a2edc-c1bf-439b-9312-f7a26dda5d80" />

One of the Problem is

When server wants some values and it does a wide parallel fetch.
<img width="655" height="293" alt="image" src="https://github.com/user-attachments/assets/c9e38549-7c39-4347-9e6a-8a9232599118" />

When the server returns the responses, we would see packet drops on the client side because of network congestion

<img width="655" height="488" alt="image" src="https://github.com/user-attachments/assets/751a95ba-c1ba-4e57-bade-d00d1ace9ca6" />

Memcache clients implement flowcontrol mechanisms to limit incast congestion.
* When a client requests a large number of keys, the responses can overwhelm components such as rack and cluster
switches if those responses arrive all at once.
* Clients therefore use a sliding window mechanism to control the number of outstanding requests.
* When the client receives a response, the next request can be sent.
* Similar to TCP’s congestion control, the size of this sliding window grows slowly upon a successful request and shrinks
when a request goes unanswered.
* The window applies to all memcache requests independently of destination;
whereas TCP windows apply only to a single stream.

### Multiple clusters

* It is tempting to buy more web and memcached servers to scale a cluster as demand increases. 
* However, naıvely scaling the system does not eliminate all problems.
* > Highly requested items will only become more popular as more web servers are added to cope with increased
user traffic.
* Incast congestion also worsens as the number of memcached servers increases.
* We therefore split `our web and memcached servers` into multiple frontend clusters.
* These clusters, along with a storage cluster that contain the databases, define a region.
* This region architecture also allows for smaller failure domains and a tractable network configuration.
* We trade replication of data for more independent failure domains, tractable network configuration, and a reduction of incast congestion.


The all to all communication limits horizontal scalability.

<img width="655" height="428" alt="image" src="https://github.com/user-attachments/assets/c0837e3e-0784-4fa5-a5a4-6ce18f9ad999" />

Now we will need to keep the caches consistent

<img width="870" height="594" alt="image" src="https://github.com/user-attachments/assets/fde43e7a-5d75-4041-b049-90b52a0da73d" />


* SQL statements that modify authoritative state are amended to include memcache keys that need to be
invalidated once the transaction commits.
* We deploy invalidation daemons (named mcsqueal) on every
database.
* Each daemon inspects the SQL statements that its database commits, extracts any deletes, and broadcasts
these deletes to the memcache deployment in every frontend cluster in that region.

<img width="1221" height="948" alt="image" src="https://github.com/user-attachments/assets/387eed73-3b09-4dc9-a799-91c17defb263" />

Inter-cluster bandwidth is less than Intra-cluster bandwidth.
<img width="1221" height="948" alt="image" src="https://github.com/user-attachments/assets/0f821203-7072-4fa5-9716-9a4229156e77" />


### Geographically distributed clusters

<img width="949" height="628" alt="image" src="https://github.com/user-attachments/assets/b847ae6c-53fe-46bf-8c73-ec0020b3c1ae" />


Facebook is ok to write to master DB across because fb is read heavy system with 2 orders of higher magnitude

<img width="635" height="458" alt="image" src="https://github.com/user-attachments/assets/57e2646e-6fcd-49c8-8eeb-4673b6b4214f" />


<img width="1253" height="948" alt="image" src="https://github.com/user-attachments/assets/1f87287f-b832-4873-9be1-4703f3314f21" />

<img width="1253" height="927" alt="image" src="https://github.com/user-attachments/assets/28a78291-b422-4521-863a-ad79593c2bd6" />


## Memory allocation

<img width="1406" height="1258" alt="image" src="https://github.com/user-attachments/assets/a1110814-a54f-4815-923b-79311bcd3e39" />

<img width="314" height="512" alt="image" src="https://github.com/user-attachments/assets/5eb25943-4e2d-4212-95ff-9a9cdaa01715" />

<img width="1280" height="720" alt="image" src="https://github.com/user-attachments/assets/fc729e86-aad1-4d16-bf20-30fd1c1d19b9" />

<img width="776" height="840" alt="image" src="https://github.com/user-attachments/assets/10eaec9d-4408-4b57-8e06-a0f583efafbe" />


One of the improvements Facebook made to memcached was moving to a smaller exponential so there is not as much waste in storing values in chunks. Instead of 2^n for the slab allocation, the latest versions of memcached use a much smaller growth exponential, 1.25^n, so you will see slabs with sizes 1KB, 1.25KB, 1.56KB, etc… This means that instead of 25% waste on average, you should see closer to 10%. Effectively you regain 15% of your memcached memory just by installing the latest version!

