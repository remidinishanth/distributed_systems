## Summary of Raft

* Raft is an algorithm for managing a replicated (basically append-only) log over a cluster of nodes.
* When you combine this with a state machine you get a stateful, distributed application. Log entries act as commands for the state machine.
* When a node in the Raft cluster crashes, it is brought up to date by sending (also called "replaying") all commands in the log through the state machine.
  - This can be made more efficient by implementing an application-specific concept of state snapshots.


## Simple example 
* https://notes.eatonphil.com/minimal-key-value-store-with-hashicorp-raft.html

Hashicorp raft implementation https://github.com/hashicorp/raft?tab=readme-ov-file talks about two examples
* Raft gRPC Example - Utilizing the Raft repository with gRPC https://github.com/Jille/raft-grpc-example 
* Raft-based KV-store Example - Uses Hashicorp Raft to build a distributed key-value store https://github.com/otoolep/hraftd
  - https://philipotoole.com/building-a-distributed-key-value-store-using-raft/ This blog talks about it


Also checkout GopherCon2023 talk "Build Your Own Distributed System Using Go"
