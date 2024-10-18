## NGINX

Ref: https://aosabook.org/en/v2/nginx.html

![image](https://github.com/user-attachments/assets/c2c79287-044e-4929-ba96-50c217793b47)


nginx ("engine x") is an 
* HTTP web server,
* reverse proxy,
* content cache,
* load balancer,
* TCP/UDP proxy server, and
* mail proxy server.

NGINX is open source software that powers web servers and enables reverse proxying, caching, load balancing, and media streaming. 
It was originally designed as a web server with high performance and reliability. 

Besides functioning as an HTTP server, NGINX acts as a proxy server for email (IMAP, POP3, and SMTP) and a reverse proxy and load balancer for HTTP, TCP, and UDP servers.

![image](https://github.com/user-attachments/assets/2a3481ec-f2ee-45de-8b79-0bc8c21267ae)


#### Web server

NGINX is a web server that renders the pages we have developed using HTML, CSS, and JavaScript. It is one of the topmost web servers among the server’s available in the market. NGINX is preferred for its single-threaded, event-driven, non-blocking and master-slave architecture.

#### Reverse Proxy

As a reverse proxy, NGINX sits between clients and backend servers. It receives client requests, forwards them to the backend server, and then sends the server's response back to the client.

Key Features:
* Load Balancing: Distributes client requests across multiple backend servers.
* SSL Termination: Handles SSL/TLS encryption and decryption.
* Caching: Can cache responses from backend servers to reduce load and latency.

![image](https://github.com/user-attachments/assets/a5341aeb-b1b3-4b7f-9831-5e983ac46d58)

![image](https://github.com/user-attachments/assets/2e921037-b0dc-4b3e-a8e4-1f70dba324a1)


## Alternatives to Nginx
* Apache HTTP Server (Apache)
* LiteSpeed
* HAProxy
* Traefik (doesn't support caching)
* Caddy

## Why it was created?

NGINX was created to address limitations in the then-dominant web server, **Apache HTTP Server**, particularly related to handling high concurrency (large numbers of simultaneous connections) efficiently.

Originally, Apache architecture matched the then-existing operating systems and hardware, but also the state of the Internet, where a website was typically a standalone physical server running a single instance of Apache. It was architected to spawn a copy of itself for each new connection, which was not suitable for nonlinear scalability of a website.

Aimed at solving the C10K problem of 10,000 simultaneous connections, Nginx was written with a different architecture in mind—one that is much more suitable for nonlinear scalability in both the number of simultaneous connections and requests per second. nginx is event-based, so it does not follow Apache's style of spawning new processes or threads for each web page request.

* High Concurrency - Apache Server used a multi-threaded or multi-process model where each connection consumed a thread or process, leading to high memory use and inefficiency under heavy load.
  - NGINX was designed with an event-driven, asynchronous, and non-blocking architecture. This allows it to handle thousands of simultaneous connections within a handful of worker processes. Instead of creating a new thread or process for each connection, NGINX uses a more efficient loop that can handle multiple connections as events, thus drastically reducing memory and CPU usage. 
* NGINX was designed with reverse proxying and load balancing as core features.

Ref: https://aosabook.org/en/v2/nginx.html


![image](https://github.com/user-attachments/assets/aa8bc0c4-7366-42ca-8295-ac4f142bde85)

![image](https://github.com/user-attachments/assets/b07d9b47-cb6f-425b-9e91-8f22e42388f9)


## Architecture:

NGINX leads the pack in web performance, and it’s all due to the way the software is designed. Whereas many web servers and application servers use a simple threaded or process‑based architecture, NGINX stands out with a sophisticated event‑driven architecture that enables it to scale to hundreds of thousands of concurrent connections on modern hardware.

NGINX follows the master-slave architecture by supporting **event-driven**, **asynchronous**, and **non-blocking** model.

![image](https://github.com/user-attachments/assets/7f0ae307-179f-4a87-9433-4de01e4b19a1)

### Overview

Traditional process- or thread-based models of handling concurrent connections involve handling each connection with a separate process or thread, and blocking on network or input/output operations. Depending on the application, it can be very inefficient in terms of memory and CPU consumption. Spawning a separate process or thread requires preparation of a new runtime environment, including allocation of heap and stack memory, and the creation of a new execution context. Additional CPU time is also spent creating these items, which can eventually lead to poor performance due to thread thrashing on excessive context switching. All of these complications manifest themselves in older web server architectures like Apache's. This is a tradeoff between offering a rich set of generally applicable features and optimized usage of server resources.

From the very beginning, nginx was meant to be a specialized tool to achieve more performance, density and economical use of server resources while enabling dynamic growth of a website, so it has followed a different model. It was actually inspired by the ongoing development of advanced event-based mechanisms in a variety of operating systems. What resulted is a modular, event-driven, asynchronous, single-threaded, non-blocking architecture which became the foundation of nginx code.

nginx uses multiplexing and event notifications heavily, and dedicates specific tasks to separate processes. Connections are processed in a highly efficient run-loop in a limited number of single-threaded processes called workers. Within each worker nginx can handle many thousands of concurrent connections and requests per second.

#### How Does NGINX Work?
NGINX uses a predictable process model that is tuned to the available hardware resources:

* The master process performs the privileged operations such as reading configuration and binding to ports, and then creates a small number of child processes (the next three types).
* The cache loader process runs at startup to load the disk‑based cache into memory, and then exits. It is scheduled conservatively, so its resource demands are low.
* The cache manager process runs periodically and prunes entries from the disk caches to keep them within the configured sizes.
* The worker processes do all of the work! They handle network connections, read and write content to disk, and communicate with upstream servers.

![image](https://github.com/user-attachments/assets/f3f11c0c-6bd9-4281-8424-5718ed25d351)

![image](https://github.com/user-attachments/assets/e4cc5fda-e3ce-402a-a792-8dcab0eeae4e)


#### Event Driven Architecture

When a connection is received, it is added to a queue of events that the server needs to process. Instead of blocking on the connection and waiting for a response, Nginx is free to continue processing other events in the queue.

Nginx uses an event loop, which is a programming construct that repeatedly checks for new events and handles them as they arrive. When an event is received, Nginx determines the appropriate action to take based on its configuration and the type of event. This may involve serving static content, proxying requests to another server, or handling SSL encryption.

By using an event-driven architecture and a single thread, Nginx can handle a large number of connections with minimal overhead.

![image](https://github.com/user-attachments/assets/1172a275-dff6-4d4d-9d96-26b3a908a63b)


#### Workers Model

As previously mentioned, nginx doesn't spawn a process or thread for every connection. Instead, worker processes accept new requests from a shared "listen" socket and execute a highly efficient run-loop inside each worker to process thousands of connections per worker. There's no specialized arbitration or distribution of connections to the workers in nginx; this work is done by the OS kernel mechanisms. Upon startup, an initial set of listening sockets is created. workers then continuously accept, read from and write to the sockets while processing HTTP requests and responses.

The run-loop is the most complicated part of the nginx worker code. It includes comprehensive inner calls and relies heavily on the idea of asynchronous task handling. Asynchronous operations are implemented through modularity, event notifications, extensive use of callback functions and fine-tuned timers. Overall, the key principle is to be as non-blocking as possible. The only situation where nginx can still block is when there's not enough disk storage performance for a worker process.

Because nginx does not fork a process or thread per connection, memory usage is very conservative and extremely efficient in the vast majority of cases. nginx conserves CPU cycles as well because there's no ongoing create-destroy pattern for processes or threads.

> The NGINX configuration recommended in most cases – running one worker process per CPU core – makes the most efficient use of hardware resources.

General rule (not specific to Nginx) For an IO intensive you can schedule hundreds of threads but for compute-intensive workload it should be proportional to number of cores.

#### Inside worker

Each NGINX worker process is initialized with the NGINX configuration and is provided with a set of listen sockets by the master process.

![image](https://github.com/user-attachments/assets/0dd9b5ac-6d33-47c0-a7f9-a69cb32b2f69)

The NGINX worker processes begin by waiting for events on the listen sockets (accept_mutex and kernel socket sharding). Events are initiated by new incoming connections. These connections are assigned to a
state machine – the HTTP state machine is the most commonly used, but NGINX also implements state machines for stream (raw TCP) traffic and for a number of mail protocols (SMTP, IMAP, and POP3).

![image](https://github.com/user-attachments/assets/1254ca7c-8738-4f78-861a-913a1aa4542c)

Most web servers that perform the same functions as NGINX use a similar state machine – the difference lies in the implementation.

![image](https://github.com/user-attachments/assets/840b2549-b814-4bb4-9193-5b69eb743e33)

![image](https://github.com/user-attachments/assets/504332b4-249b-4242-90a8-02cc0237d343)
