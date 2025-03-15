## Containerd

containerd is a high-level container runtime, aka container manager. 
To put it simply, it's a daemon that manages the complete container lifecycle on a single host: 
creates, starts, stops containers, pulls and stores images, configures mounts, networking, etc.

containerd is designed to be easily embeddable into larger systems. 

Docker uses containerd under the hood to run containers. Kubernetes can use containerd via CRI to manage containers on a single node.

![image](https://github.com/user-attachments/assets/2279da2b-78bb-4c3c-b43d-dbe2ab363c65)

![image](https://github.com/user-attachments/assets/0932ad79-dc88-43fa-85e6-71bbf87b388d)


## More details
![image](https://github.com/user-attachments/assets/476823da-c08c-4102-9322-1f82b53303b6)

Ref: https://stackoverflow.com/questions/46649592/dockerd-vs-docker-containerd-vs-docker-runc-vs-docker-containerd-ctr-vs-docker-c 
