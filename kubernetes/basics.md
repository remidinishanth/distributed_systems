Terms with which you should be familiar:

* Cluster – Group of physical or virtual servers wherein  Kubernetes is installed. A set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.
* Node (Master) – Physical or virtual server that controls the Kubernetes cluster
* Node (Worker) – Physical or virtual servers where  workloads run in a given container technology 
* Pods – A pod is the smallest unit in a Kubernetes cluster. A pod may contain one or more containers. Group of containers and volumes which share the same network namespace 
* Labels – User defined Key:Value pair associated to Pods  
* Master – Control plane components which provide access  point for admins to manage cluster workloads 
* Service – An abstraction which serves as a proxy for a group of Pods performing a “service”

![](images/k8s_i1.png)

![](images/k8s_i3.png)

Masters manage the cluster and the nodes are used to host the running applications.

![](images/k8s_i4.png)

A Deployment is responsible for creating and updating instances of your application. 
After creating application instances, a Deployment Controller continuously watches them, and replaces an instance if the 
Node hosting it goes down or it is deleted. This provides a self-healing mechanism to address machine failure and machine maintenance.

![](images/k8s_i5.png)

A Pod is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), 
IP address and information about how to run them. Pods always run on Nodes.

Containers should only be scheduled together in a single Pod if they are tightly coupled and need to share resources such as disk.

![](images/k8s_i6.png)

A node is a worker machine in Kubernetes and may be a VM or physical machine, depending on the cluster. Multiple Pods can run on one Node.

![](images/k8s_i7.png)

A Kubernetes Service is an abstraction layer which defines a **logical set of Pods** and enables external traffic exposure, load balancing and service discovery for those Pods.

REF: https://www.mirantis.com/blog/kubernetes-cheat-sheet/
