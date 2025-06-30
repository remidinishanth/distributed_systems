---
layout: page
title: "Socket Programming"
category: "networking"
---

<img width="1241" alt="image" src="https://github.com/user-attachments/assets/9c02ba2d-e569-48c6-95ce-3b9e68dbf447">


Ref: https://www.cs.dartmouth.edu/~campbell/cs50/socketprogramming.html

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/08025bfc-5c3b-4b9f-9bc4-435ad4cf5ae7">

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/3b0b924f-bab3-4ef0-ba86-1cfcae998d18)


![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/0360e4e3-429c-4469-8831-7a818e0c0312)

<img width="734" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/79d5badd-020b-4bd3-a696-c306470f91c2">


<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/02b5b096-f32e-40ad-bb75-d671d54e1994">

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/6c9fba23-3241-4e18-b9bd-76f0fc4259be">

### Data flow between wire and app

<img width="471" alt="image" src="https://github.com/user-attachments/assets/191e43ab-b6aa-4e18-b65f-d81db26e28e9">


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

![image](https://github.com/user-attachments/assets/529f2e1f-58d4-4a11-9892-ac12928ad900)


<img width="1120" alt="image" src="https://github.com/user-attachments/assets/4e5cdf0f-3f06-4083-b28a-3ff0faed8f75">

<img width="652" alt="image" src="https://github.com/user-attachments/assets/c4ce2f02-1b98-4ae4-85ae-f91c0531b2a4">

The `accept()` function creates a new socket from the first connection request for the specified `socket_descriptor` and returns the file descriptor of the new socket. The file descriptor of this new socket is used in the `read()` and `write()` functions to send and receive data to and from the client node.

Ref: https://people.cs.rutgers.edu/~pxk/417/notes/index.html
