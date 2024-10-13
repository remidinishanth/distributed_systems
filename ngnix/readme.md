## NGNIX

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

#### Web server

NGINX is a web server that renders the pages we have developed using HTML, CSS, and JavaScript. It is one of the topmost web servers among the server’s available in the market. NGINX is preferred for its single-threaded, event-driven, non-blocking and master-slave architecture.

#### Reverse Proxy

As a reverse proxy, NGINX sits between clients and backend servers. It receives client requests, forwards them to the backend server, and then sends the server's response back to the client.

Key Features:
* Load Balancing: Distributes client requests across multiple backend servers.
* SSL Termination: Handles SSL/TLS encryption and decryption.
* Caching: Can cache responses from backend servers to reduce load and latency.

## Architecture:

NGINX leads the pack in web performance, and it’s all due to the way the software is designed. Whereas many web servers and application servers use a simple threaded or process‑based architecture, NGINX stands out with a sophisticated event‑driven architecture that enables it to scale to hundreds of thousands of concurrent connections on modern hardware.

NGINX follows the master-slave architecture by supporting event-driven, asynchronous and non-blocking model.

![image](https://github.com/user-attachments/assets/7f0ae307-179f-4a87-9433-4de01e4b19a1)

#### How Does NGINX Work?
NGINX uses a predictable process model that is tuned to the available hardware resources:

* The master process performs the privileged operations such as reading configuration and binding to ports, and then creates a small number of child processes (the next three types).
* The cache loader process runs at startup to load the disk‑based cache into memory, and then exits. It is scheduled conservatively, so its resource demands are low.
* The cache manager process runs periodically and prunes entries from the disk caches to keep them within the configured sizes.
* The worker processes do all of the work! They handle network connections, read and write content to disk, and communicate with upstream servers.
