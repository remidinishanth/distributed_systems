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


### Overview

<img width="863" height="666" alt="image" src="https://github.com/user-attachments/assets/52416eed-7a48-4486-b3d5-5849dde96d7d" />

Ref: https://jaceklaskowski.gitbooks.io/apache-kafka/content/kafka-overview.html



https://stackoverflow.com/questions/41744506/difference-between-stream-processing-and-message-processing

## Event broker vs Message queue
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a4bdffbf-bfff-4bf7-94e0-fa0b14ae5d42)


#### MESSAGE QUEUE
Messages are put onto a queue and a consumer consumes the message and processes them. Messages are acknowledged as consumed and deleted afterwards. Messages are split between consumers which makes it hard to communicate system with events.

Example of this would be Amazon SQS. Publish messages to the queue and then listen to them, process them and they are removed from the queue.

#### EVENT BROKER
Event brokers are a push system, they push these events downstream to consumers. Example of this would be Amazon EventBridge.

Ref: https://serverlessland.com/event-driven-architecture/visuals/message-queue-vs-event-broker


<img width="1379" height="809" alt="image" src="https://github.com/user-attachments/assets/91af6d2a-090f-4a99-9332-9e7fe60ac025" />

<img width="1456" height="809" alt="image" src="https://github.com/user-attachments/assets/1d88cb63-72bb-483d-bcc9-930a4d246785" />

### How Google PubSub achieves similar functionality

By separating Topic and Subscription

<img width="2538" height="1305" alt="image" src="https://github.com/user-attachments/assets/5d0c08ca-7ba2-4482-bbcc-e8382cbd1b9a" />

AMQP Protocol

<img width="1361" height="644" alt="image" src="https://github.com/user-attachments/assets/5b797e2d-573e-40cb-8e49-0959e59a3bda" />

RabbitMQ supports different types of Exchanges, to achieve Pubsub-like functionality, fanout with multiple queues as subscription is good
<img width="1180" height="720" alt="image" src="https://github.com/user-attachments/assets/6e1ed711-9b12-4e6d-baf8-0dcb8ba955d1" />

Similarly, in SNS + SQS for achieving similar functionality

<img width="723" height="448" alt="image" src="https://github.com/user-attachments/assets/1fcbb4dc-5661-464c-b035-c63597cba63c" />


### Kafka

<img width="1466" height="773" alt="image" src="https://github.com/user-attachments/assets/2b6f75b9-b44a-49a4-b835-9bec6fe36a11" />


Read at https://www.oreilly.com/library/view/kafka-the-definitive/9781491936153/ch04.html


![1693476406513](https://github.com/remidinishanth/distributed_systems/assets/19663316/f1978c9a-b706-4806-ab38-04581e4c22b8)

### Topics and Paritions

<img width="589" height="666" alt="image" src="https://github.com/user-attachments/assets/faee0cab-dd82-4a0a-b2c0-e0c3167775a5" />

Ref: Kafka white paper

<img width="589" height="307" alt="image" src="https://github.com/user-attachments/assets/91efd441-83ba-4c59-9618-dae607f42a06" />


<img width="673" height="369" alt="image" src="https://github.com/user-attachments/assets/68c43db3-0a3c-42c1-b4cd-66b75aa26f65" />

<img width="850" height="444" alt="image" src="https://github.com/user-attachments/assets/d5754de4-fbad-4bda-ac46-ad9aaebe92d8" />


<img width="1495" height="712" alt="image" src="https://github.com/user-attachments/assets/df44e52d-6a48-4883-b7c9-90953b548451" />

#### Consumers

<img width="1462" height="705" alt="image" src="https://github.com/user-attachments/assets/ff0cc06a-809b-4771-a7a3-2822b897c37e" />


<img width="1272" height="755" alt="image" src="https://github.com/user-attachments/assets/e048a411-edc7-40d5-b933-3dbb899217b5" />

<img width="1456" height="914" alt="image" src="https://github.com/user-attachments/assets/8a1491f5-e0d2-4d75-aa18-1a3d1f368289" />

Ref: https://stackoverflow.com/questions/36203764/how-can-i-scale-kafka-consumers to read about scaling of consuming

Write scalability
<img width="1668" height="1104" alt="image" src="https://github.com/user-attachments/assets/7d5ec3f6-2e0e-426f-baa8-c3c4c0ae7cc9" />

Read scalability
<img width="1638" height="1148" alt="image" src="https://github.com/user-attachments/assets/8797b247-fd7a-41b0-bf7d-f09d64880ab4" />

Ref: https://www.instaclustr.com/blog/the-power-of-kafka-partitions-how-to-get-the-most-out-of-your-kafka-cluster/

### Zookeeper

<img width="1490" height="814" alt="image" src="https://github.com/user-attachments/assets/c70f474a-f0a4-4aad-a45d-8ff1452602c1" />


## TODO:
Zuul architecture, https://www.youtube.com/watch?v=6w6E_B55p0E
