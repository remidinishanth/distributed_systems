---
layout: page
title: "Kafka"
category: "general"
---

Also read https://pk.org/417/notes/kafka.html 

> Goal: Create a distributed messaging system to handle large-scale streams of messages.

How can a cluster of computers handle the influx of never-ending streams of data, coming from multiple sources? This data may come from industrial sensors, IoT devices scattered around the world, or log files from tens of thousands of systems in a data center.

It’s easy enough to say that we can divide the work among multiple computers but how would we exactly do that?

<img width="575" height="439" alt="image" src="https://github.com/user-attachments/assets/206e47ab-68ca-42dc-ab8b-7fd0f3ddc9c3" />


https://stackoverflow.com/questions/41744506/difference-between-stream-processing-and-message-processing

## Event broker vs Message queue
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a4bdffbf-bfff-4bf7-94e0-fa0b14ae5d42)


#### MESSAGE QUEUE
Messages are put onto a queue and a consumer consumes the message and processes them. Messages are acknowledged as consumed and deleted afterwards. Messages are split between consumers which makes it hard to communicate system with events.

Example of this would be Amazon SQS. Publish messages to the queue and then listen to them, process them and they are removed from the queue.

#### EVENT BROKER
Event brokers are a push system, they push these events downstream to consumers. Example of this would be Amazon EventBridge.

Ref: https://serverlessland.com/event-driven-architecture/visuals/message-queue-vs-event-broker


![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/e2456728-b302-40f1-907b-7d4880ac4090)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/86839c99-ca34-4e8e-8de8-06a3ed3ea4ac)

Ref: https://stackoverflow.com/questions/36203764/how-can-i-scale-kafka-consumers to read about scaling of consuming

Read at https://www.oreilly.com/library/view/kafka-the-definitive/9781491936153/ch04.html

Zuul architecture, https://www.youtube.com/watch?v=6w6E_B55p0E


![1693476406513](https://github.com/remidinishanth/distributed_systems/assets/19663316/f1978c9a-b706-4806-ab38-04581e4c22b8)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/f757eede-a6c7-4fe0-b858-3bc8c4aef877)

