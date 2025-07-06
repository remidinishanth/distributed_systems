---
layout: page
title: "Networking"
category: "kubernetes"
---

Read awesome resources: 
* https://yuminlee2.medium.com/kubernetes-service-networking-2169f04ce02a
* CNI https://yuminlee2.medium.com/kubernetes-container-network-interface-cni-ee5b21514664
* Weave net plugin: https://yuminlee2.medium.com/kubernetes-weave-net-cni-plugin-810849203c73
* Ref: Understanding Kubernetes Networking in 30 Minutes - Ricardo Katz & James Strong Conference talk

## Networking

<img width="944" alt="image" src="https://github.com/user-attachments/assets/e2d03ce5-0207-4327-83fd-34dcff0714bb" />

Which components handle these different APIs

<img width="944" alt="image" src="https://github.com/user-attachments/assets/1499bab2-702d-4918-bea5-27423a7cb632" />


<img width="1031" alt="image" src="https://github.com/user-attachments/assets/d704ad09-36d5-4b0a-93fa-ab15f5ecf32a" />

<img width="1199" alt="image" src="https://github.com/user-attachments/assets/7ee09ae7-64d7-4fd8-8544-a5e96eb60028" />


The official documentation says kubernetes networking model requires:
* all containers can communicate with all other containers without NAT
* all nodes can communicate with all containers (and vice-versa) without NAT
* the IP that a container sees itself as is the same IP that others see it as

Must Read https://mvallim.github.io/kubernetes-under-the-hood/documentation/kube-flannel.html and https://www.devopsschool.com/tutorial/kubernetes/kubernetes-cni-flannel-overlay-networking.html

Todo: https://iximiuz.com/en/posts/service-discovery-in-kubernetes/

* Every Pod will have its own network namespace and containers inside that Pod share the same IP address and ports.
  - All communication between these containers would happen via the localhost as they are all part of the same namespace.
* In Kubernetes, every node has a designated CIDR range of IPs for Pods. This would ensure that every Pod gets a unique IP address that can be seen by other Pods in the cluster and also ensures that when a new Pod is created, the IP address never overlaps.  

![image](https://github.com/user-attachments/assets/189dcfac-1abb-400e-98e4-24895a0853a0)

Ref: https://medium.com/techbeatly/kubernetes-networking-fundamentals-d30baf8a28c8

Pod-to-Service Networking

* Pods are very dynamic in nature. They may need to scale up or down based on demand. They may be created again in case of an application crash or a node failure.
* All of these events cause the Pods IP address to change and this makes networking a challenge.

![image](https://github.com/user-attachments/assets/89c02c2a-0e81-4330-89d0-2f16d8f2d0bc)


![image](https://github.com/user-attachments/assets/d130a27a-345e-4434-b762-0ea0f6c9807f)

Ref: https://blog.purestorage.com/purely-technical/kubernetes-cluster-networking-components/

CNI Plugin
<img width="1108" alt="image" src="https://github.com/user-attachments/assets/56d736c7-1b0d-4e17-973b-765684f0fb02" />


![image](https://github.com/user-attachments/assets/3dd1bf94-6979-4de5-9962-a1dce0890366)

Flannel is a Container Network Interface (CNI) plugin for Kubernetes that provides a simple overlay network, essential for pod-to-pod communication across different nodes in the cluster. It assigns unique subnets to each node and encapsulates packets at the host level, allowing pods to communicate even if they reside on different physical machines. 


While flannel was originally designed for Kubernetes, it is a generic overlay network that can be used as a simple alternative to existing software defined networking solutions. More specifically, flannel gives each host an IP subnet (/24 by default) from which the Docker daemon is able to allocate IPs to the individual containers. Each address corresponds to a container, so that all containers in a system may reside on different hosts.

![image](https://github.com/user-attachments/assets/6172d743-ba0c-4460-80ac-64333218a5b7)

It works by first configuring an overlay network, with an IP range and the size of the subnet for each host. For example, one could configure the overlay to use 10.1.0.0/16 and each host to receive a /24 subnet. Host A could then receive 10.1.15.1/24 and host B could get 10.1.20.1/24. Flannel uses etcd to maintain a mapping between allocated subnets and real host IP addresses. For the data path, flannel uses UDP to encapsulate IP datagrams to transmit them to the remote host.

<img width="888" alt="image" src="https://github.com/user-attachments/assets/813c0c3c-85c1-4b8c-988a-3761b7ebe020" />


