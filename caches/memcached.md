<img width="1107" height="559" alt="image" src="https://github.com/user-attachments/assets/140d7c24-51ec-4565-a36e-842e24ac51d0" />


> memcached is a high-performance, distributed memory object caching system, generic in nature, 
but originally intended for use in speeding up dynamic web applications by alleviating database load.

<img width="2390" height="1678" alt="image" src="https://github.com/user-attachments/assets/f53690c0-cd92-4543-aff6-62dc43f01740" />

Ref: https://memcached.org/about


<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/d3b833f3-cc07-4b78-ac9b-b09b6452548f" />


<img width="941" height="559" alt="image" src="https://github.com/user-attachments/assets/ab51a430-2fa8-45fc-90f4-232164577dca" />

<img width="406" height="336" alt="image" src="https://github.com/user-attachments/assets/65fa0c86-4601-46cc-8c6f-fc1a9f9bcfca" />


## How Facebook served billions of requests per second Using Memcached

Facebook had to deal with these issues early on because of its popularity. At any point in time, millions of people were accessing Facebook from all over the world.

In terms of software design, this meant a few important requirements:

* Facebook had to support real-time communication.
* They had to build capabilities for on-the-fly content aggregation.
* Scale to handle billions of user requests.
* Store trillions of items across multiple geographic locations.

To achieve these goals, Facebook took up the open-source version of Memcached and enhanced it to build a distributed key-value store. This enhanced version was known as `Memcache`.

