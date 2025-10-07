---
layout: page
title: "Socket Programming"
category: "networking"
---

<img width="1241" alt="image" src="https://github.com/user-attachments/assets/9c02ba2d-e569-48c6-95ce-3b9e68dbf447">

Ref: https://www.cs.dartmouth.edu/~campbell/cs50/socketprogramming.html

### Sockets definition

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/08025bfc-5c3b-4b9f-9bc4-435ad4cf5ae7">

Sockets are endpoints for sending and receiving data across a network. 

Sockets are a fundamental abstraction for process-to-process / inter-process communication (IPC), particularly over networks. It can also be used within the same machine.

A socket is one endpoint of a two-way communication link between two programs running on a network. 💻 ↔️ 🌐 ↔️ 💻 

Mainly two types:
* TCP Sockets (Stream Sockets) `SOCK_STREAM`
* UDP Sockets (Datagram Sockets) `SOCK_DGRAM`

Address family: `AF_INET` for IPv4, `AF_INET6` for IPv6, or `AF_UNIX` for local IPC

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/02b5b096-f32e-40ad-bb75-d671d54e1994">

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/6c9fba23-3241-4e18-b9bd-76f0fc4259be">

### How Sockets fit in Network stack

<img width="1628" height="882" alt="image" src="https://github.com/user-attachments/assets/d1bc26f3-008d-4b14-bc80-5a9fd1a6ed28" />

### How does sockets fit in OS stack

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/0360e4e3-429c-4469-8831-7a818e0c0312)

### Data flow between wire and app

<img width="471" alt="image" src="https://github.com/user-attachments/assets/191e43ab-b6aa-4e18-b65f-d81db26e28e9">


### TCP sockets

<img width="734" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/79d5badd-020b-4bd3-a696-c306470f91c2">

The system call to create a socket is `int socket (domain, type, protocol);` 
```c
    // AF_INET: IPv4 Internet protocols, AF_INET6 for Ipv6
    // SOCK_STREAM: Provides sequenced, reliable, two-way, connection-based byte streams (TCP), SOCK_DGRAM for UDP
    // 0: Specifies the default protocol for the given domain and type
    int sockfd = socket(AF_INET, SOCK_STREAM, 0);
```



Only the server needs to bind. Bind system call `int bind (int sockfd, const struct sockaddr *my_addr, socklen_t addrlen);`

```c
struct sockaddr_in {
    short sin_family; // e.g. AF_INET
    unsigned short sin_port; // e.g. htons(3490)‏
    struct in_addr sin_addr; // see struct in_addr below
    char sin_zero[8]; // zero this if you want to
};



struct sockaddr_in my_addr; 
int sockfd; 

if ((sockfd = socket (AF_INET, SOCK_STREAM, 0) < 0) {
    printf (“Error while creating the socket</span>n”);
    exit(1);
}

bzero (&my_addr, sizeof(my_addr)); // zero structure out 
my_addr.sin_family = AF_INET; // match the socket() call
my_addr.sin_port = htons(5100); // specify port to listen on
my_addr.sin_addr.s_addr = htonl(INADDR_ANY); //allow the server to accept a client connection on any interface

if((bind(sockfd, (struct sockaddr *) &my_addr, sizeof(saddr)) < 0 {
    printf(“Error in binding</span>n”);
     exit(1);
} 
```

* We specify the IP address as `INADDR_ANY`, which allows the server to accept a client connection on any
interface, in case the server host has multiple interfaces

**Convert socket to listening socket**
* By calling listen, the socket is converted into a listening socket, on which incoming connections from clients will be accepted by the kernel.
* These three steps, `socket`, `bind`, and `listen`, are the normal steps for any TCP server to prepare what we call the listening descriptor `sockfd` in our case

Only the server needs to listen `int listen (int sockfd, int backlog)`, backlog specifies the maximum number of pending connections the kernel should queue for the socket. Listen returns `0` if OK, `-1` on error.

Only the server can accept the incoming client connections `int accept (int sockfd, struct sockaddr *fromaddr, socklen_t *addrlen)`

Clients

The client need not **bind**, **listen** or **accept**. All client needs to do is to just connect to the server.
`int connect (int sockfd, struct sockaddr *toaddr, socklen_t addrlen)`

### TCP sockets, there is connect where there is 3 way handshake

![image](https://github.com/user-attachments/assets/529f2e1f-58d4-4a11-9892-ac12928ad900)

<img width="3400" height="2926" alt="image" src="https://github.com/user-attachments/assets/e65c4c59-afba-4a11-93f3-59190614179e" />

### UDP sockets

<img width="823" height="571" alt="image" src="https://github.com/user-attachments/assets/a8f8e638-a5b5-4b3f-b19c-63f31f09f0e7" />

<img width="3400" height="2548" alt="image" src="https://github.com/user-attachments/assets/389c5094-d6b5-4591-b92f-eb456511e49d" />


## Established Socket

A socket is created by an application running in a host. The application assigns a transport protocol (TCP or UDP) and source and destination addresses to the socket. It identifies sockets by assigning numbers to them.

Note the web server has two sockets opened: one for each web page it is serving. These sockets are differentiated by the destination port numbers.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/1d15f38b-9503-42b5-9231-8714d5a66d80)

> [!NOTE]  
> One host does not assign the socket number on both sides of the communication channel. The socket numbers assigned to each socket are only used by the host that assigned them. In other words, socket number 1 created on one host may be connected to socket number 5 on another host.

> [!NOTE]  
> Based on the Well-Known source port numbers assigned to each socket, we can determine sockets 1 and 2 were created by an HTTP server application and socket 3 was created by an SMTP or email server application.

### Sockets numbers may not be same on both sides

This graphic shows a virtual TCP connection between a client and server. Note the socket numbers are not the same on both sides of the channel. Hosts create, close and number their own sockets.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/47b5a71d-0d00-41c0-9f2e-8af44fb9256e)

## Sockets

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/143f9616-dc43-4773-a845-0c79c1928b50">


![image](https://github.com/user-attachments/assets/16e999b6-ed20-4892-9485-e0fa1a699ecf)


<img width="1120" alt="image" src="https://github.com/user-attachments/assets/300008cd-4421-482d-9049-4544db916131">

<img width="652" alt="image" src="https://github.com/user-attachments/assets/a28d6424-cf25-41ab-a73d-00b2d44f053d">

### TCP

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/4e5cdf0f-3f06-4083-b28a-3ff0faed8f75">

<img width="652" alt="image" src="https://github.com/user-attachments/assets/c4ce2f02-1b98-4ae4-85ae-f91c0531b2a4">

The `accept()` function creates a new socket from the first connection request for the specified `socket_descriptor` and returns the file descriptor of the new socket. The file descriptor of this new socket is used in the `read()` and `write()` functions to send and receive data to and from the client node.

Ref: https://people.cs.rutgers.edu/~pxk/417/notes/index.html
