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


Overview

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


### Kafka

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

<img width="1272" height="755" alt="image" src="https://github.com/user-attachments/assets/e048a411-edc7-40d5-b933-3dbb899217b5" />

<img width="1456" height="914" alt="image" src="https://github.com/user-attachments/assets/8a1491f5-e0d2-4d75-aa18-1a3d1f368289" />

Ref: https://stackoverflow.com/questions/36203764/how-can-i-scale-kafka-consumers to read about scaling of consuming


## TODO:
Zuul architecture, https://www.youtube.com/watch?v=6w6E_B55p0E
