---
layout: page
title: "Tcp Udp"
category: "networking"
---

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/2faf6170-a273-439c-b0a4-693d5d86e9ad)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/ac4cf316-47f4-4908-8138-33696a3498b5)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/427c1cee-dfd5-4b80-b1c6-d7c91dbae734)

Note how the TCP protocol requires more information and overhead to guarantee data delivery.
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6a2a293d-816a-4fff-9bfa-084b22018671)


#### Common Applications and their Protocols:
    * HTTP (Hypertext Transfer Protocol):
        Transport Protocol: TCP
        Description: Used for transferring hypertext documents on the World Wide Web. It is the foundation of any data exchange on the Web.
        Rationale: HTTP requires a reliable, connection-oriented protocol to ensure the accurate and ordered delivery of web pages and associated resources.

    * HTTPS (Hypertext Transfer Protocol Secure):
        Transport Protocol: TCP
        Description: The secure version of HTTP, used for secure communication over a computer network. It is widely used for secure data transmission over the Internet.
        Rationale: Similar to HTTP, HTTPS requires the reliability and security provided by TCP for secure and ordered data transfer.

    * FTP (File Transfer Protocol):
        Transport Protocol: Primarily TCP
        Description: Used for transferring files between computers on a network. FTP can use both active and passive modes, but the control connection is always established over TCP.
        Rationale: FTP relies on a reliable connection to ensure the successful transfer of files, making TCP the suitable choice.

    * SFTP (SSH File Transfer Protocol):
        Transport Protocol: Typically runs over TCP
        Description: A secure file transfer protocol that provides file access, file transfer, and file management over a secure data stream.
        Rationale: SFTP, like FTP, requires a reliable and secure connection, making TCP the preferred choice.

    * SMTP (Simple Mail Transfer Protocol):
        Transport Protocol: TCP
        Description: Used for the transmission of emails between computers. It is a text-based protocol where one server sends messages to another server.
        Rationale: Email transmission requires reliable and ordered delivery of messages, which TCP provides.

    * POP3 (Post Office Protocol 3):
        Transport Protocol: Primarily TCP
        Description: Used for retrieving emails from a server. It allows an email client to download email messages from the server.
        Rationale: Similar to SMTP, POP3 relies on TCP for reliable and ordered retrieval of email messages.

    * IMAP (Internet Message Access Protocol):
        Transport Protocol: Primarily TCP
        Description: An email retrieval and storage protocol, more advanced than POP3. It allows multiple devices to access and manipulate the same mailbox on a server.
        Rationale: IMAP, like POP3, requires a reliable and ordered connection to manage and access email messages.

    * NFS (Network File System):
        Transport Protocol: Typically uses UDP or TCP
        Description: A distributed file system protocol allowing a user on a client computer to access files over a network as if those files were local.
        Rationale: NFS can use either UDP or TCP depending on the specific requirements. UDP may be preferred for lower overhead, while TCP offers reliability.

    * DNS (Domain Name System):
        Transport Protocol: Primarily uses UDP (for queries) and TCP (for zone transfers and large responses)
        Description: Resolves domain names to IP addresses and vice versa. UDP is used for most DNS queries, and TCP is used when the response data size exceeds the DNS protocol limits.
        Rationale: DNS queries, which are typically small and frequent, use UDP for lower overhead. TCP is reserved for larger responses and zone transfers where reliability is more critical.

    * SNMP (Simple Network Management Protocol):
        Transport Protocol: Primarily uses UDP
        Description: Used for collecting and organizing information about managed devices on an IP network and for modifying that information to change device behavior.
        Rationale: SNMP typically uses UDP for its simplicity and lower overhead. While it doesn't guarantee delivery, SNMP is often used for real-time monitoring where timely updates are crucial.

In summary, the choice between TCP and UDP for these protocols is influenced by factors such as the need for reliable and ordered delivery, security considerations, and the characteristics of the data being transmitted.

## Notes

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/ba24e5e8-1029-42b8-a83e-43e7243eaf7b">

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/1147d021-a67d-4216-8f07-02b80c60131e">

<img width="1120" alt="image" src="https://github.com/user-attachments/assets/5b8c0ee3-a8c6-40f5-8f17-be2d27aeb60d">


Ref: https://people.cs.rutgers.edu/~pxk/417/notes/index.html
