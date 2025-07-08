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

Overview

![image](https://github.com/user-attachments/assets/d82374fa-f335-4889-80bb-ee977238df61)


<img width="944" alt="image" src="https://github.com/user-attachments/assets/e2d03ce5-0207-4327-83fd-34dcff0714bb" />

Which components handle these different APIs

<img width="944" alt="image" src="https://github.com/user-attachments/assets/1499bab2-702d-4918-bea5-27423a7cb632" />


<img width="1031" alt="image" src="https://github.com/user-attachments/assets/d704ad09-36d5-4b0a-93fa-ab15f5ecf32a" />

<img width="1199" alt="image" src="https://github.com/user-attachments/assets/7ee09ae7-64d7-4fd8-8544-a5e96eb60028" />


The official documentation says the Kubernetes Network Model requires that:
* Pods are all able to communicate with one another without the need to make use of Network Address Translation (NAT).
* Nodes - the machines that run the Kubernetes cluster. These can be either virtual or physical machines, or indeed anything else that is able to run Kubernetes - are also able to communicate with all the Pods, without the need for NAT.
* Each Pod will see itself with the same IP that other Pods see it as having.

<img width="1562" alt="image" src="https://github.com/user-attachments/assets/c9548c4c-b2d9-435b-ba28-beeddad87d55" />

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


## CNI Plugin

How it works

![image](https://github.com/user-attachments/assets/4050fc69-563e-4c5b-9411-35c8b769bc58)


![image](https://github.com/user-attachments/assets/af9acab9-1e5f-4215-971b-0e35655cf01c)

* In Overlay mode, a container is independent of the host’s IP address range. During cross-host communication, tunnels are established between hosts and all packets in the container CIDR block are encapsulated as packets exchanged between hosts in the underlying physical network. This mode removes the dependency on the underlying network.
* In Routing mode, hosts and containers belong to different CIDR blocks. Cross-host communication is implemented through routing. No tunnels are established between different hosts for packet encapsulation. However, route interconnection partially depends on the underlying network. For example, a reachable route must exist from the underlying network to Layer 2.
* In Underlay mode, containers and hosts are located at the same network layer and share the same position. Network interconnection between containers depends on the underlying network. Therefore, this mode is highly dependent on the underlying capabilities.

Ref: https://alibaba-cloud.medium.com/getting-started-with-kubernetes-kubernetes-cnis-and-cni-plug-ins-10c33e44ac88

### Overlay networking

Flannel is created by CoreOS for Kubernetes networking, it also can be used as a general software defined network solution for other purpose.

To achieve kubernetes' network requirements, flannel’s idea is simple: create another flat network which runs above the host network, this is the so-called overlay network. All containers(Pod) will be assigned one ip address in this overlay network, they communicate with each other by calling each other’s ip address directly.

Ref: https://www.devopsschool.com/tutorial/kubernetes/kubernetes-cni-flannel-overlay-networking.html
