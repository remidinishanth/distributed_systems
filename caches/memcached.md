<img width="1107" height="559" alt="image" src="https://github.com/user-attachments/assets/140d7c24-51ec-4565-a36e-842e24ac51d0" />


> memcached is a high-performance, distributed memory object caching system, generic in nature, 
but originally intended for use in speeding up dynamic web applications by alleviating database load.

<img width="2390" height="1678" alt="image" src="https://github.com/user-attachments/assets/f53690c0-cd92-4543-aff6-62dc43f01740" />

Ref: https://memcached.org/about


<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/d3b833f3-cc07-4b78-ac9b-b09b6452548f" />

<img width="2222" height="672" alt="image" src="https://github.com/user-attachments/assets/dc939ab8-aef4-42ef-926d-1a8cb1df15bd" />


<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/ab51a430-2fa8-45fc-90f4-232164577dca" />

<img width="406" height="336" alt="image" src="https://github.com/user-attachments/assets/65fa0c86-4601-46cc-8c6f-fc1a9f9bcfca" />


## How Facebook served billions of requests per second Using Memcached

Memcached is a well known, simple, inmemory caching solution. Memcached was originally developed by Brad Fitzpatrick for LiveJournal in 2003. It was originally written in Perl, but is rewritten in C by Anatoly Vorobey.

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
