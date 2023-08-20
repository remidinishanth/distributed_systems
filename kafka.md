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
