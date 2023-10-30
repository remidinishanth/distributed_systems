## File-Sharing: Getting a File from One Person (Machine) to Another
#### Can use client/server:
* Client requests file, server responds with the data.
* HTTP, FTP work this way.
* Downsides: Single point of failure, expensive, doesn’t scale.
#### Could use CDNs:
* Buy multiple servers, put them near clients to decrease latency.
* No single point of failure, scales better.
