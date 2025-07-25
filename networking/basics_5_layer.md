---
layout: page
title: "Five Layer Software Model Overview"
category: "networking"
description: "Understanding the five-layer networking model and protocols"
---

# Five Layer Software Model Overview

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b3e5b738-0c80-4a15-b3d2-89b56cf34b85)

![image](https://github.com/user-attachments/assets/1259ee00-3132-4205-9d57-428f324096ed)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/7d0515ff-14a2-4669-adef-2fbcb2ab4295)

---

### Sending message
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/8b494d60-5d12-4631-bbf8-afc16a1f14a9)

### Receiving message
<img width="1706" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/264fa4de-dfe2-4c5b-a4d3-b21fb947cadd">

---

## Summary
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/cb75ef43-d7a6-4f0d-a172-8c404e02e174)

Ref: https://microchipdeveloper.com/tcpip:tcp-ip-five-layer-model

## Application Layer (Layer 5)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b710bde9-eba5-4f2e-beb1-d9d63921bf63)

Read more at https://microchipdeveloper.com/tcpip:common-tcp-ip-applications (Awesome resource)

## Transport Layer (Layer 4)

The first duty of a transport-layer protocol is to provide **process-to-process** communication. 
A process is an application-layer entity (running program) that uses the services
of the transport layer.

The network layer is responsible for communication at the computer level (host-
to-host communication). A network-layer protocol can deliver the message only to the
destination computer. However, this is an incomplete delivery. The message still needs
to be handed to the correct process. This is where a transport-layer protocol takes over.
A transport-layer protocol is responsible for delivery of the message to the appropriate
process.

<img width="673" alt="image" src="https://github.com/user-attachments/assets/5211c4a2-48ca-4885-916d-4e7669a61ce6">

<details>
  <summary> More details </summary>

* The local host and the remote host are defined using IP addresses.
* To define the processes, we need second identifiers, called **port numbers**. In
the TCP/IP protocol suite, the port numbers are integers between 0 and 65,535 (16 bits).
* The client program defines itself with a port number, called the ephemeral port
number. The word ephemeral means “short-lived” and is used because the life of a
client is normally short. An ephemeral port number is recommended to be greater than
1023 for some client/server programs to work properly.
* The server process must also define itself with a port number. This port number,
however, cannot be chosen randomly. If the computer at the server site runs a server
process and assigns a random number as the port number, the process at the client site
that wants to access that server and use its services will not know the port number. TCP/
IP has decided to use universal port numbers for servers; these are called **well-known port numbers**.
Every client process knows the well-known port number of the corresponding server process.  

Communication using Port numbers
<img width="641" alt="image" src="https://github.com/user-attachments/assets/0c20bed8-1b21-4a16-a830-befeca38e025">


<img width="641" alt="image" src="https://github.com/user-attachments/assets/7f067ec0-0b90-45e9-8433-c60744f2a120">

A transport-layer protocol in the TCP suite needs both the IP address and the port number,
at each end, to make a connection. The combination of an IP address and a port number is
called a **socket address**.


Whenever an entity accepts items from more than one source, this is referred to as
multiplexing (many to one); whenever an entity delivers items to more than one source,
this is referred to as demultiplexing (one to many). The transport layer at the source
performs multiplexing; the transport layer at the destination performs demultiplexing.

Three client processes are running at the client site: P1, P2, and P3. The processes P1 and P3 need to
send requests to the corresponding server process running in a server. 

<img width="641" alt="image" src="https://github.com/user-attachments/assets/e2ed19b6-95a2-4346-9fa5-cb85f71aeae3">

---
  
</details>  


The header added to messages by the Transport layer includes more than just the source and destination port numbers. Here we are showing all the information included in TCP and UDP headers.

Note how the TCP protocol requires more information and overhead to guarantee data delivery.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b525663c-b138-4bcf-80c2-216b63e8cf02)

### TCP

* TCP (Transmission Control Protocol) provides reliable byte stream (connection-oriented) service. 
* This layer of software ensures that packets arrive at the application in order and lost or corrupt packets are retransmitted.
* The transport layer keeps track of the destination so the application can have the illusion of a connected data stream.

TCP tries to give a datagram some of the characteristics of a virtual circuit network. The TCP layer will send sequence numbers along with each packet of data, buffer received data in memory so they can be presented to the application in order, acknowledge received packets, and request a retransmission of missing or corrupt packets.

### UDP
* The User Datagram Protocol (UDP) is a connectionless, unreliable transport protocol. It
does not add anything to the services of IP except for providing process-to-process
communication instead of host-to-host communication.
* UDP is a very simple protocol using a minimum of overhead. Sending a small message using
UDP takes much less interaction between the sender and receiver than using TCP.
* UDP packets, called user datagrams, have a fixed-size header of 8 bytes made up
of four fields, each of 2 bytes (16 bits). 

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a022b8c9-07ad-40d5-88c2-63b3e4523915)

## Network Layer (Layer 3)

* When transmitting data, this layer adds a header containing the source and destination IP addresses to the to the data received from the Transport layer. The packet it creates will then be forwarded to the MAC or Data Link layer.

* When receiving data, this layer is used to determine if the packet received by the host contains the host’s IP address. If it does, the data is forwarded up to the Transport layer.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/9542eb24-3b60-4812-b627-b428ffaee683)

<img width="610" alt="image" src="https://github.com/user-attachments/assets/3ede1bd1-7635-43c3-8419-ad87a8da971b">

#### Network layer and Data Link Layer
* To send an IP packet out, the system needs to identify the link layer destination address (MAC, or Media Access Control address) on the local area network that corresponds to the desired IP destination (it may be the address of a router if the packet is going to a remote network). 
* The **Address Resolution Protocol**, or ARP, accomplishes this. It works by broadcasting a request containing an IP address (the message asks, do you know the corresponding MAC address for this IP address?) and then waiting for a response from the computer with the corresponding IP address. 
* To avoid doing this for every outgoing packet, ARP maintains a cache of most recently used addresses.

## Data Link Layer (Layer 2)

All hosts that have an IP address also have a MAC (Media Access Controller) address. Unlike IP addresses which are virtual, MAC addresses are fixed hardware based addresses that never change.

This layer uses a Media Access Controller (MAC) to generate the frames that will be transmitted. As the name suggests, the MAC controls the physical transmission media.

#### Summary
We can consider the data-link layer as two sublayers. 
* The upper sublayer is responsible for data-link control, and
* The lower sublayer is responsible for resolving access to the shared media.

Data-link control (DLC) deals with the design and procedures for communication between two adjacent nodes: node-to-node communication. 
This sublayer is responsible for framing and error control. 
Error control deals with data corruption during transmission.

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/d8d16298-e031-4b46-87d6-373e8771dd18">


### Ethernet and WiFi Frame Format

As you probably guessed, the Data Link layer adds more than just the source and destination MAC addresses to the packet. 

Note that the MAC for Ethernet and WiFi are different and generate different frames.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6124a9d0-2aed-4244-b593-044899377c01)

## Physical Layer (Layer 1)

It sends and receives signals on the physical wire or antenna to transmit the bits found in frames.

There is a PHY found at the end of every network interface (e.g. end of wire or antenna).

## Summary

<img width="1563" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/36edf3f2-986c-4940-af99-4f80a47c21a5">

<img width="1021" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/c31c8968-0be9-4618-802d-2b2ed803154d">

<img width="1054" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/c135e5ff-0ffa-421a-ab48-6e172f07e823">

<img width="777" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/9c526693-d1e2-4aa1-bf9f-6b00b75588f7">

### Data Communications and Networking by Forouzan

The computer with logical address `A` and physical address `10` needs to send a packet to the computer with logical address `P` and physical address `95`.

<img width="849" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/42ad6ea1-aadb-485f-9d00-4f9751c52f47">

* The network layer needs to find the physical address of the next hop before the packet can be delivered.
* The network layer consults its routing table and finds the logical address of the next hop (router I) to be F.
* The ARP finds the physical address of router 1 that corresponds to the logical address of 20.
* Now the network layer passes this address to the data link layer, which in tum, encapsulates the packet with physical destination address 20 and physical source address 10.

* The frame is received by every device on LAN 1, but is discarded by all except router 1, which finds that the destination physical address in the frame matches with its own physical address.
* The router decapsulates the packet from the frame to read the logical destination address P. Since the logical destination address does not match the
router's logical address, the router knows that the packet needs to be forwarded.

* The router consults its routing table and ARP to find the physical destination address of the next hop (router 2), creates a new frame, encapsulates the packet, and sends it to router 2.
* Note the physical addresses in the frame. The source physical address changes from 10 to 99.
* The destination physical address changes from 20 (router 1 physical address) to 33 (router 2 physical address).
* The logical source and destination addresses must remain the same; otherwise the packet will be lost.

> The physical addresses(MAC address) will change from hop to hop, but the logical addresses(IP address) usually(there are some exceptions to
this rule) remain the same.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/d51f3aab-151e-4e59-b0d9-9bcf1e50f99b)

<img width="1562" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/814b299a-fb25-4b8e-b838-651a4b92e360">

Ref: https://people.cs.rutgers.edu/~pxk/417/notes/pdf/01c-networking-slides.pdf"
category: "networking"
---

## 7 Layer vs 5 layer

![image](https://github.com/user-attachments/assets/6d937997-2529-4853-a3f4-c892b7ac797e)

### Goal

> Goal: Enable computers to communicate with each other; create the machine-to-machine and process-to-process communication channels.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/8449fc09-9d06-4c22-a118-1b6e7e70550e)

---

# Five Layer Software Model Overview
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b3e5b738-0c80-4a15-b3d2-89b56cf34b85)

---

![image](https://github.com/user-attachments/assets/1259ee00-3132-4205-9d57-428f324096ed)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/7d0515ff-14a2-4669-adef-2fbcb2ab4295)

---

### Sending message
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/8b494d60-5d12-4631-bbf8-afc16a1f14a9)

### Receiving message
<img width="1706" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/264fa4de-dfe2-4c5b-a4d3-b21fb947cadd">

---

## Summary
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/cb75ef43-d7a6-4f0d-a172-8c404e02e174)

Ref: https://microchipdeveloper.com/tcpip:tcp-ip-five-layer-model

## Application Layer (Layer 5)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b710bde9-eba5-4f2e-beb1-d9d63921bf63)

Read more at https://microchipdeveloper.com/tcpip:common-tcp-ip-applications (Awesome resource)

## Transport Layer (Layer 4)

The first duty of a transport-layer protocol is to provide **process-to-process** communication. 
A process is an application-layer entity (running program) that uses the services
of the transport layer.

The network layer is responsible for communication at the computer level (host-
to-host communication). A network-layer protocol can deliver the message only to the
destination computer. However, this is an incomplete delivery. The message still needs
to be handed to the correct process. This is where a transport-layer protocol takes over.
A transport-layer protocol is responsible for delivery of the message to the appropriate
process.

<img width="673" alt="image" src="https://github.com/user-attachments/assets/5211c4a2-48ca-4885-916d-4e7669a61ce6">

<details>
  <summary> More details </summary>

* The local host and the remote host are defined using IP addresses.
* To define the processes, we need second identifiers, called **port numbers**. In
the TCP/IP protocol suite, the port numbers are integers between 0 and 65,535 (16 bits).
* The client program defines itself with a port number, called the ephemeral port
number. The word ephemeral means “short-lived” and is used because the life of a
client is normally short. An ephemeral port number is recommended to be greater than
1023 for some client/server programs to work properly.
* The server process must also define itself with a port number. This port number,
however, cannot be chosen randomly. If the computer at the server site runs a server
process and assigns a random number as the port number, the process at the client site
that wants to access that server and use its services will not know the port number. TCP/
IP has decided to use universal port numbers for servers; these are called **well-known port numbers**.
Every client process knows the well-known port number of the corresponding server process.  

Communication using Port numbers
<img width="641" alt="image" src="https://github.com/user-attachments/assets/0c20bed8-1b21-4a16-a830-befeca38e025">


<img width="641" alt="image" src="https://github.com/user-attachments/assets/7f067ec0-0b90-45e9-8433-c60744f2a120">

A transport-layer protocol in the TCP suite needs both the IP address and the port number,
at each end, to make a connection. The combination of an IP address and a port number is
called a **socket address**.


Whenever an entity accepts items from more than one source, this is referred to as
multiplexing (many to one); whenever an entity delivers items to more than one source,
this is referred to as demultiplexing (one to many). The transport layer at the source
performs multiplexing; the transport layer at the destination performs demultiplexing.

Three client processes are running at the client site: P1, P2, and P3. The processes P1 and P3 need to
send requests to the corresponding server process running in a server. 

<img width="641" alt="image" src="https://github.com/user-attachments/assets/e2ed19b6-95a2-4346-9fa5-cb85f71aeae3">

---
  
</details>  


The header added to messages by the Transport layer includes more than just the source and destination port numbers. Here we are showing all the information included in TCP and UDP headers.

Note how the TCP protocol requires more information and overhead to guarantee data delivery.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b525663c-b138-4bcf-80c2-216b63e8cf02)

### TCP

* TCP (Transmission Control Protocol) provides reliable byte stream (connection-oriented) service. 
* This layer of software ensures that packets arrive at the application in order and lost or corrupt packets are retransmitted.
* The transport layer keeps track of the destination so the application can have the illusion of a connected data stream.

TCP tries to give a datagram some of the characteristics of a virtual circuit network. The TCP layer will send sequence numbers along with each packet of data, buffer received data in memory so they can be presented to the application in order, acknowledge received packets, and request a retransmission of missing or corrupt packets.

### UDP
* The User Datagram Protocol (UDP) is a connectionless, unreliable transport protocol. It
does not add anything to the services of IP except for providing process-to-process
communication instead of host-to-host communication.
* UDP is a very simple protocol using a minimum of overhead. Sending a small message using
UDP takes much less interaction between the sender and receiver than using TCP.
* UDP packets, called user datagrams, have a fixed-size header of 8 bytes made up
of four fields, each of 2 bytes (16 bits). 

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a022b8c9-07ad-40d5-88c2-63b3e4523915)

## Network Layer (Layer 3)

* When transmitting data, this layer adds a header containing the source and destination IP addresses to the to the data received from the Transport layer. The packet it creates will then be forwarded to the MAC or Data Link layer.

* When receiving data, this layer is used to determine if the packet received by the host contains the host’s IP address. If it does, the data is forwarded up to the Transport layer.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/9542eb24-3b60-4812-b627-b428ffaee683)

<img width="610" alt="image" src="https://github.com/user-attachments/assets/3ede1bd1-7635-43c3-8419-ad87a8da971b">

#### Network layer and Data Link Layer
* To send an IP packet out, the system needs to identify the link layer destination address (MAC, or Media Access Control address) on the local area network that corresponds to the desired IP destination (it may be the address of a router if the packet is going to a remote network). 
* The **Address Resolution Protocol**, or ARP, accomplishes this. It works by broadcasting a request containing an IP address (the message asks, do you know the corresponding MAC address for this IP address?) and then waiting for a response from the computer with the corresponding IP address. 
* To avoid doing this for every outgoing packet, ARP maintains a cache of most recently used addresses.

## Data Link Layer (Layer 2)

All hosts that have an IP address also have a MAC (Media Access Controller) address. Unlike IP addresses which are virtual, MAC addresses are fixed hardware based addresses that never change.

This layer uses a Media Access Controller (MAC) to generate the frames that will be transmitted. As the name suggests, the MAC controls the physical transmission media.

#### Summary
We can consider the data-link layer as two sublayers. 
* The upper sublayer is responsible for data-link control, and
* The lower sublayer is responsible for resolving access to the shared media.

Data-link control (DLC) deals with the design and procedures for communication between two adjacent nodes: node-to-node communication. 
This sublayer is responsible for framing and error control. 
Error control deals with data corruption during transmission.

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/d8d16298-e031-4b46-87d6-373e8771dd18">


### Ethernet and WiFi Frame Format

As you probably guessed, the Data Link layer adds more than just the source and destination MAC addresses to the packet. 

Note that the MAC for Ethernet and WiFi are different and generate different frames.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6124a9d0-2aed-4244-b593-044899377c01)

## Physical Layer (Layer 1)

It sends and receives signals on the physical wire or antenna to transmit the bits found in frames.

There is a PHY found at the end of every network interface (e.g. end of wire or antenna).

## Summary

<img width="1563" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/36edf3f2-986c-4940-af99-4f80a47c21a5">

<img width="1021" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/c31c8968-0be9-4618-802d-2b2ed803154d">

<img width="1054" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/c135e5ff-0ffa-421a-ab48-6e172f07e823">

<img width="777" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/9c526693-d1e2-4aa1-bf9f-6b00b75588f7">

### Data Communications and Networking by Forouzan

The computer with logical address `A` and physical address `10` needs to send a packet to the computer with logical address `P` and physical address `95`.

<img width="849" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/42ad6ea1-aadb-485f-9d00-4f9751c52f47">

* The network layer needs to find the physical address of the next hop before the packet can be delivered.
* The network layer consults its routing table and finds the logical address of the next hop (router I) to be F.
* The ARP finds the physical address of router 1 that corresponds to the logical address of 20.
* Now the network layer passes this address to the data link layer, which in tum, encapsulates the packet with physical destination address 20 and physical source address 10.

* The frame is received by every device on LAN 1, but is discarded by all except router 1, which finds that the destination physical address in the frame matches with its own physical address.
* The router decapsulates the packet from the frame to read the logical destination address P. Since the logical destination address does not match the
router's logical address, the router knows that the packet needs to be forwarded.

* The router consults its routing table and ARP to find the physical destination address of the next hop (router 2), creates a new frame, encapsulates the packet, and sends it to router 2.
* Note the physical addresses in the frame. The source physical address changes from 10 to 99.
* The destination physical address changes from 20 (router 1 physical address) to 33 (router 2 physical address).
* The logical source and destination addresses must remain the same; otherwise the packet will be lost.

> The physical addresses(MAC address) will change from hop to hop, but the logical addresses(IP address) usually(there are some exceptions to
this rule) remain the same.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/d51f3aab-151e-4e59-b0d9-9bcf1e50f99b)

<img width="1562" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/814b299a-fb25-4b8e-b838-651a4b92e360">

Ref: https://people.cs.rutgers.edu/~pxk/417/notes/pdf/01c-networking-slides.pdf 
