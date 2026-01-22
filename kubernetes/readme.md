---
layout: page
title: "Kubernetes"
category: "container-orchestration"
tags: ["kubernetes", "containers", "orchestration", "microservices"]
description: "Container orchestration with Kubernetes - concepts, architecture, and practical examples"
---

Also checkout: https://docs.google.com/document/d/10w2ynbbVisw6XHhq6ohB5pVVlXh8hdCWTHahc4ieQYQ/edit

TODO: https://cybozu.github.io/introduction-to-kubernetes/introduction-to-kubernetes.html

and https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html


### Why we need kubernetes

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/734f817d-5e89-47f9-a782-fc72daeeca41)

Ref: https://newsletter.francofernando.com/p/kubernetes

<img width="1071" alt="image" src="https://github.com/user-attachments/assets/dac3573c-599b-4eaa-8e75-71c690b58786" />

<img width="672" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/549f0cec-ac99-438d-ba7b-f668fe85c762">

Ref: http://bit.ly/9stepsawesome

## Architecture

<img width="1065" alt="image" src="https://github.com/user-attachments/assets/7fd5b490-d95e-459d-b6e5-bc95d4ab0960" />

### K8S components in a nutshell
* YAML file & kubectl cmdline. 
* K8s node: CRI, Kubelet, Kubeproxy  
* K8s master: ETCd, APIserver, Controller & Scheduler

![image](https://github.com/user-attachments/assets/35dc72ea-62bd-4bbd-a9fa-29af82204cc0)

![image](https://github.com/user-attachments/assets/21192418-ccde-4508-96aa-395216e52312)

![image](https://github.com/user-attachments/assets/321cc113-bb7f-4fae-be96-a7676c7ec853)

<img width="1455" alt="image" src="https://github.com/user-attachments/assets/0107343f-c0c2-4721-a29f-4af5792055b3" />

### K8s Controller

> How to get from the current state to the declared target state?

![image](https://github.com/user-attachments/assets/a5de85a6-a61a-47aa-8608-4c142097c010)

<img width="1409" alt="image" src="https://github.com/user-attachments/assets/2c9dabb1-dc95-422c-b700-9673f4e1efc5" />

<img width="1167" alt="image" src="https://github.com/user-attachments/assets/4b46c862-18f3-4b11-bdb1-5758a032ef8a" />

![image](https://github.com/user-attachments/assets/c7d85057-4782-4712-8604-08ca1886d539)

![image](https://github.com/user-attachments/assets/6819d3e8-4cfa-459e-8b08-31400391a8eb)


<img width="1437" alt="image" src="https://github.com/user-attachments/assets/7dffcf71-1cf2-41f7-9f79-7be5b2472d35" />

### Concepts

<img width="1202" alt="image" src="https://github.com/user-attachments/assets/b1e1d210-cc98-48a4-8674-cf06448d729a" />

<img width="1345" alt="image" src="https://github.com/user-attachments/assets/fc51c6be-30a7-4579-ae2f-03fbc8bc12a9" />

<img width="960" alt="image" src="https://github.com/user-attachments/assets/46e0fa27-977b-40e4-9af7-9b1606c224cf" />

### Containers, Pods, Deployment

<img width="1192" alt="image" src="https://github.com/user-attachments/assets/b6c6e11e-c6c8-42d7-a901-855ae297a32c" />


![image](https://github.com/user-attachments/assets/53eb8a8b-0319-448e-9de7-acd797ba0fde)

<img width="1279" alt="image" src="https://github.com/user-attachments/assets/a24363aa-d6a8-4392-be48-9feec2ff7e92" />

<img width="1406" alt="image" src="https://github.com/user-attachments/assets/cc5fa89b-d819-4257-be58-47006f7ac1ae" />

### Operator

Operators are Kubernetes extensions that use custom resources to manage applications and their components.

<img width="1189" alt="image" src="https://github.com/user-attachments/assets/77e14f6c-5829-4ac5-88bc-5123d1772f28" />

<img width="1358" alt="image" src="https://github.com/user-attachments/assets/c01c9756-62c3-42c7-b1c6-abe93b43f04b" />

### Stateful service

<img width="1249" alt="image" src="https://github.com/user-attachments/assets/b7538d96-24d6-4f58-aae7-254f46c6abb5" />

