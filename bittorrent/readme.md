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

#### Basics of original BitTorrent (BT) protocol:
* Create a `.torrent` file, which contains meta-information about the file (file name, length, info about pieces that comprise the file, URL of tracker).
* Have a **tracker**. A server that knows the identity of all the peers involved in your file transfer.
* To download:
  - **Peer** contacts tracker.
  - Tracker responds with list of other peers involved in transfer.
  - Peer connects to these other peers, begins to transfer blocks (see below).
  - Some peers are **seeders**: Already have the entire file (maybe servers that host the file, or just nice peers who are sticking around).

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


### Architecture

BitTorrent in its original
form matches the **hybrid** peer-to-peer concept. 

It’s all about the torrent file, the centralized tracker
and the associated swarm of peers.

The BitTorrent architecture normally consists of the following entities:
- a static metainfo file (a “torrent file”)
- a ‘tracker’
- an original downloader (“seed”)
- the end user downloader (“leecher”)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/4074576e-e92c-4ab2-86ac-dd8890670044)

#### Torrent file

* The first step in publishing a file using BitTorrent is to
create a metainfo file from the file that you want to
publish. 
* The metainfo file is called a “torrent”. The
torrent file contains the filename, size, hashing
information and the URL of the “tracker”.
* The “torrent” is needed by anyone who wants to download the file the
torrent is created from. The torrent file can be distributed
by e-mail, IRC, http etc.
* To download or “seed” a file, you need a
BitTorrent client. The BitTorrent client is a free
application that administrates the download procedure.
* A BitTorrent download is
started by opening the torrent file in the BitTorrent client.

#### Tracker

* The tracker keeps a log of peers that are currently
downloading a file, and helps them find each other.
* The tracker is not directly involved in the transfer of data and
does not have a copy of the file.
* The tracker and the downloading users exchange information
using a simple protocol on top of HTTP.

* First, the user gives information to the tracker about
which file it’s downloading, ports it’s listening on etc. The responce from the tracker is a list
of other users which are downloading the same file and information on how to contact them.

* This group of peers that all share the same torrent represents a ‘**swarm**’. 

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/9801d540-5f0c-4904-b712-ab0d4a21c019)

#### What’s “seeding”?

Seeding is the act of uploading to a torrent data stream. As a key part of the technology, seeding is what allows for data redundancy when other seeders go offline, as well as a boost to overall throughput/data speeds when other peers want to download a file.

* An original downloader known as a “seed” has to be started.
* A “seed” is a user that has the entire file.
* A downloading user that has nothing or only parts of a file is called a “leecher”.
* The “seed” must upload at least one complete copy of the file.
* Once an entire copy is distributed amongst the other downloaders, the ‘seed’ can stop uploading and the download
will still continue for all downloaders as long as there are enough people downloading the
file, and all pieces of the file are available.

#### What’s being a “peer”?
A peer is anyone connected to a torrent file, and downloading or uploading data to the collective network.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/21e4a0aa-c64f-439c-a754-1d640a5e8919)


### Ref
* https://iq.opengenus.org/bittorrent-architecture/
* https://blogs.umass.edu/Techbytes/2015/04/28/bittorrent-an-explanation-of-the-protocol/
* https://web.cs.ucla.edu/classes/cs217/05BitTorrent.pdf
