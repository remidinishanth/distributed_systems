## File-Sharing: Getting a File from One Person (Machine) to Another
#### Can use client/server:
* Client requests file, server responds with the data.
* HTTP, FTP work this way.
* Downsides: Single point of failure, expensive, doesn’t scale.
#### Could use CDNs:
* Buy multiple servers, put them near clients to decrease latency.
* No single point of failure, scales better.

#### More details about HTTP and FTP
- FTP is a protocol specifically designed for file transfers, while HTTP is primarily used for serving web pages and web content.
- FTP uses two TCP ports: port 20 for sending data and port 21 for sending control commands.


## Peer-To-Peer (P2P) Networks for File-Sharing

In this model, files are distributed among the peers themselves, without relying on a central sever. 
Each peer acts as both a consumer and a contributor. They can download and upload files to other peers within the network. 
E.g: Kazaa, Napster, Gnutella and Direct Connect

Compared to the more common server-client solution, a peer-to-peer approach has
several advantages including increased robustness and resource providing, such as bandwidth,
storage space and computing power, by peers.

* Distribute the architecture to the extreme.
* Once a client downloads (part of) the file from the server, that client can upload (part of) the file to others. Put clients to work!
* In theory: Infinitely scalable.
* P2P networks create overlays on top of the underlying Internet (so do CDNs).
* Problem: What if users aren’t willing to upload?

### Problem with Traditional P2P networks

* The problem with many “traditional” peer-to-peer file sharing protocols, is that most users have different speeds downlink and uplink.
* This means that even though a user has plenty of bandwidth downlink, the speed of a file transfer to him will be
restricted by a much smaller bandwidth uplink from the user he downloads from.
#### Solution:
* Bram Cohen this problem by splitting files into smaller pieces.
* When requesting a file, the user’s computer sniffs around on the internet for people having one or more pieces of the wanted file.
* He then downloads different parts of the file from different users at the same time and utilizes his downlink
bandwidth better.

## Bit Torrent

* BitTorrent is a technology/protocol which makes the distribution of files, especially large
files, easier and less bandwidth consuming for the publisher. This is accomplished by utilizing the upload capacity of the peers that are downloading a file. 
* BitTorrent base its operation around the concept of a torrent file, a centralized tracker and an associated
swarm of peers. 
* The centralized tracker provides the different entities with an address list
over available peers.

### How to Incentivize Peers to Upload
#### Basics of original BitTorrent (BT) protocol:
* Create a `.torrent` file, which contains meta-information about the file (file name, length, info about pieces that comprise the file, URL of tracker).
* Have a tracker. A server that knows the identity of all the peers involved in your file transfer.
* To download:
 - Peer contacts tracker.
 - Tracker responds with list of other peers involved in transfer.
 - Peer connects to these other peers, begins to transfer blocks (see below).
 - Some peers are seeders: Already have the entire file (maybe servers that host the file, or just nice peers who are sticking around).

#### In the actual download, peers request blocks: pieces of pieces.
* Details/terminology doesn’t matter. Just know that blocks are small (~16KB) chunks of the file.
* Request blocks in a random order (more or less).

#### What incentivizes users to upload (UL) rather than just download(DL)ing?
* High-level: Users aren’t allowed to DL from a user unless they’re also ULing to that user.
 - So peers want mutual interest: A has to have blocks that B needs, and vice versa.
* Protocol is divided into rounds. In round n, some number of peers upload blocks to Peer X. In round n+1, Peer X will send blocks to the peers that uploaded the most in round n. (Typically, to the top four peers.)
* How do peers get started?  Each peer reserves some (small) amount of bandwidth to give away freely.

#### This method of incentivizing peers is part of what allowed P2P file-sharing to take off.
#### Lingering problem: tracker is central point of failure.
#### Most BT clients today are “trackerless”, and use Distributed Hash Tables (DHTs) instead.

## Architecture

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/9801d540-5f0c-4904-b712-ab0d4a21c019)


Ref: https://iq.opengenus.org/bittorrent-architecture/

