
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/8449fc09-9d06-4c22-a118-1b6e7e70550e)

# Five Layer Software Model Overview
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b3e5b738-0c80-4a15-b3d2-89b56cf34b85)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/7d0515ff-14a2-4669-adef-2fbcb2ab4295)

### Sending message
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/8b494d60-5d12-4631-bbf8-afc16a1f14a9)

### Receiving message
<img width="1706" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/264fa4de-dfe2-4c5b-a4d3-b21fb947cadd">

## Summary
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/cb75ef43-d7a6-4f0d-a172-8c404e02e174)

Ref: https://microchipdeveloper.com/tcpip:tcp-ip-five-layer-model

## Application Layer (Layer 5)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b710bde9-eba5-4f2e-beb1-d9d63921bf63)

Read more at https://microchipdeveloper.com/tcpip:common-tcp-ip-applications (Awesome resource)

## Transport Layer (Layer 4)

The header added to messages by the Transport layer includes more than just the source and destination port numbers. Here we are showing all the information included in TCP and UDP headers.

Note how the TCP protocol requires more information and overhead to guarantee data delivery.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/b525663c-b138-4bcf-80c2-216b63e8cf02)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a022b8c9-07ad-40d5-88c2-63b3e4523915)

## Network Layer (Layer 3)

* When transmitting data, this layer adds a header containing the source and destination IP addresses to the to the data received from the Transport layer. The packet it creates will then be forwarded to the MAC or Data Link layer.

* When receiving data, this layer is used to determine if the packet received by the host contains the host’s IP address. If it does, the data is forwarded up to the Transport layer.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/9542eb24-3b60-4812-b627-b428ffaee683)

## Data Link Layer (Layer 2)

All hosts that have an IP address also have a MAC (Media Access Controller) address. Unlike IP addresses which are virtual, MAC addresses are fixed hardware based addresses that never change.

This layer uses a Media Access Controller (MAC) to generate the frames that will be transmitted. As the name suggests, the MAC controls the physical transmission media.

### Ethernet and WiFi Frame Format

As you probably guessed, the Data Link layer adds more than just the source and destination MAC addresses to the packet. 

Note that the MAC for Ethernet and WiFi are different and generate different frames.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6124a9d0-2aed-4244-b593-044899377c01)

## Physical Layer (Layer 1)

It sends and receives signals on the physical wire or antenna to transmit the bits found in frames.

There is a PHY found at the end of every network interface (e.g. end of wire or antenna).

## Summary

<img width="1563" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/36edf3f2-986c-4940-af99-4f80a47c21a5">

