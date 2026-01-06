---
layout: page
title: "Raft Distributed Consensus"
category: "consensus"
tags: ["raft", "consensus", "distributed-systems", "fault-tolerance"]
description: "Understanding the Raft consensus algorithm for distributed systems"
---

# Raft Distributed Consensus
State machine replication

> Goal: Create a fault-tolerant distributed algorithm that enables a set of processes to agree on a sequence of events.

Distributed Consensus
* Distributed = Many nodes
* Consensus = Agreement on something

## Why do we need consensus?

* Consensus, or distributed agreement, is a recurring problem in distributed systems design.
* It is useful for things such as
  - mutual exclusion, where all processes agree on who has exclusive access to a resource, and
  - leader election, where a group of processes has to decide which of them is in charge.
* Perhaps most importantly, consensus plays a pivotal role in building replicated state machines.

Ref: https://pk.org/classes/417/notes/raft.html

## Use cases:

<img width="714" height="447" alt="image" src="https://github.com/user-attachments/assets/f0796aec-05bd-4a21-aecf-d4d865532ee8" />

## Raft

<img width="1145" height="551" alt="image" src="https://github.com/user-attachments/assets/a8cd3dbe-2ca0-459b-9b54-fb1df8a36cf1" />

Consensus Goal

<img width="1165" height="660" alt="image" src="https://github.com/user-attachments/assets/9890b478-7d53-4030-a5d8-03f7af75812b" />

<img width="1058" height="809" alt="image" src="https://github.com/user-attachments/assets/2c2ea768-cb6d-4cbd-b9db-3db75b0a1d65" />

<img width="1165" height="625" alt="image" src="https://github.com/user-attachments/assets/542cf195-bdca-4879-aa8e-5d69f764c49c" />

## Implementation

States

<img width="951" height="462" alt="image" src="https://github.com/user-attachments/assets/f713a988-422f-4230-ac85-93ffa7bfeee3" />

RPCs

<img width="1095" height="495" alt="image" src="https://github.com/user-attachments/assets/e1865d90-1bcc-4a02-b55f-832bd1325dfe" />

Terms

<img width="1161" height="592" alt="image" src="https://github.com/user-attachments/assets/64ff5827-d551-4e69-b8a9-e86650bcb879" />

### Leader Election

<img width="1161" height="616" alt="image" src="https://github.com/user-attachments/assets/06e947bd-221b-46c7-bc8c-a93ef8911acd" />

Server State transitions
<img width="1034" height="700" alt="image" src="https://github.com/user-attachments/assets/616004c8-3406-442d-bb4b-18b315f3990b" />


<img width="1161" height="611" alt="image" src="https://github.com/user-attachments/assets/9cc78e4b-7c39-4fa0-b016-4395e7fa28ca" />

<img width="1161" height="535" alt="image" src="https://github.com/user-attachments/assets/faffeae8-a279-4818-93ad-023504c71678" />

Log Replication

<img width="1161" height="535" alt="image" src="https://github.com/user-attachments/assets/d54114ee-6a0a-4bce-9679-1ef61262b176" />

<img width="1161" height="580" alt="image" src="https://github.com/user-attachments/assets/ae759b98-13d1-458d-a837-3632f600ca42" />

<img width="1157" height="528" alt="image" src="https://github.com/user-attachments/assets/091eda39-0d12-4df6-a7f1-cd035bc2d1d3" />


<img width="711" height="680" alt="image" src="https://github.com/user-attachments/assets/024d4d9d-506a-4b92-9dc3-be0b5502bdc7" />

Possible logs of followers

<img width="564" height="650" alt="image" src="https://github.com/user-attachments/assets/bfc52954-7026-4d01-895f-dc3a56fc874f" />

<img width="564" height="666" alt="image" src="https://github.com/user-attachments/assets/959b8015-9aa3-4bbe-8d4e-6a365d1865ce" />


## Summary

![image](https://user-images.githubusercontent.com/19663316/215156748-3d94f01c-b752-4801-afc3-6f0cb5f5a198.png)

source: https://www.hashicorp.com/resources/raft-consul-consensus-protocol-explained
