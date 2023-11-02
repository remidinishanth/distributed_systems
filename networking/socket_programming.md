Ref: https://www.cs.dartmouth.edu/~campbell/cs50/socketprogramming.html

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/08025bfc-5c3b-4b9f-9bc4-435ad4cf5ae7">

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/3b0b924f-bab3-4ef0-ba86-1cfcae998d18)


![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/0360e4e3-429c-4469-8831-7a818e0c0312)

<img width="734" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/79d5badd-020b-4bd3-a696-c306470f91c2">


<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/02b5b096-f32e-40ad-bb75-d671d54e1994">

<img width="873" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/6c9fba23-3241-4e18-b9bd-76f0fc4259be">

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
