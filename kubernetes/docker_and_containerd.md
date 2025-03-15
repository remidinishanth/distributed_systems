## Background

Kubernetes project created the container runtime interface (CRI) to make
the runtime layer pluggable. This means you can pick and choose the best runtimes
for your needs. For example, some runtimes provide better isolation, whereas others
provide better performance.

Kubernetes 1.24 finally removed support for Docker as a runtime as it was bloated
and overkill for what Kubernetes needed. Since then, most new Kubernetes clusters
ship with containerd as the default runtime. 

Fortunately, containerd is a stripped-down version of Docker optimized for Kubernetes, and it
fully supports applications containerized by Docker. In fact, Docker, containerd, and
Kubernetes all work with images and containers that implement the Open Container
Initiative (OCI) standards.

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

![image](https://github.com/user-attachments/assets/cfe2c01d-c2fe-4a5a-91f8-dbcc2539d10a)


![image](https://github.com/user-attachments/assets/a40ed94c-6320-43c0-9e4c-a855675ef3cf)

The containerd architecture consists of the following components:
* containerd-shim: A shim is a process that runs inside the container and provides a communication channel between the container and containerd.
* containerd-daemon: The daemon is the main process that manages containers. It is responsible for creating, starting, stopping, and deleting containers. It also manages the container’s lifecycle, including the start and stop of its processes.
* containerd-cri: The CRI plugin provides a gRPC interface to containerd. It is used by container orchestration platforms like Kubernetes to manage containers.


![image](https://github.com/user-attachments/assets/7201cfc1-ea2f-45ca-a7f8-248bf90d1de8)

Ref: https://collabnix.com/containerd-vs-docker-whats-the-difference/


In Linux environments, container management tools like Docker are built on a more granular set of container tools: runc and containerd.
![image](https://github.com/user-attachments/assets/a07661ed-9cae-474a-b539-72ae2a5fc2c7)

* runc is a Linux command-line tool for creating and running containers according to the OCI container runtime specification.
* containerd is a daemon that manages container life cycle from downloading and unpacking the container image to container execution and supervision.
