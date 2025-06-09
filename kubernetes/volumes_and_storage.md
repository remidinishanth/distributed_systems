## Volumes

Kubernetes volumes provide a way for containers in a pod to access and share data via the filesystem. 

![image](https://github.com/user-attachments/assets/2aa50e52-f653-4e8b-8f9a-a02f3dc9a34c)

![image](https://github.com/user-attachments/assets/f33ddd49-a7d5-4f23-b5df-7a866e248277)

<img width="750" alt="image" src="https://github.com/user-attachments/assets/8c07890b-36dc-4b17-bdb6-45b89d73992a" />

At its core, a volume is a directory, possibly with some data in it, which is accessible to the containers in a pod.

To use a volume, specify the volumes to provide for the Pod in `.spec.volumes` and declare where to mount those volumes into containers in `.spec.containers[*].volumeMounts`.

For each container defined within a Pod, you must independently specify where to mount each volume that the container uses.

### Empty dir

```
apiVersion: v1
kind: Pod
metadata:
  name: myvolumes-pod
spec:
  containers:
  - image: alpine
    imagePullPolicy: IfNotPresent
    name: myvolumes-container
    
    command: [    'sh', '-c', 'echo The Bench Container 1 is Running ; sleep 3600']
    
    volumeMounts:
    - mountPath: /demo
      name: demo-volume
  volumes:
  - name: demo-volume
    emptyDir: {}
```

### Config Maps

A ConfigMap provides a way to inject configuration data into pods.

A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.

<img width="953" alt="image" src="https://github.com/user-attachments/assets/ac7e3803-9a0d-46c9-bb16-4b588ffdc1b0" />

## Storage

<img width="609" alt="image" src="https://github.com/user-attachments/assets/ee2fe9af-dc81-4fba-a8e9-437ab24aa823" />

Kubernetes Storage Options — Persistent Volumes (PV), Persistent Volume Claims (PVC), Storage Classes (SC).
* Persistent Volume — low level representation of a storage volume.
* Persistent Volume Claim — binding between a Pod and Persistent Volume.
* Storage Class — allows for dynamic provisioning of Persistent Volumes.

<img width="931" alt="image" src="https://github.com/user-attachments/assets/e43b8d82-a66b-40f1-91d8-5ef5741116c7" />

![image](https://github.com/user-attachments/assets/7d3c03c8-478a-4662-a224-d3a97176fe76)

PV: A PV is a storage resource located in the cluster. Administrators can manually provision PVs, and Kubernetes can use storage classes to dynamically provisioned PVs.

PVC: A PVC is a storage request made by a user. It works similarly to a pod but consumes PV resources rather than node resources.

CSI: The Container Storage Interface (CSI) is a standard interface that allows container orchestrators to expose exposing arbitrary block and file storage systems to containers they manage.

### Static Provisioning and Dynamic Provisioning

![image](https://github.com/user-attachments/assets/56138ffc-0459-49a2-8d7c-15f2a3e2fb27)

### Stateful Sets
<img width="932" alt="image" src="https://github.com/user-attachments/assets/60a34942-4a9e-4b3d-81fd-cbb2ba60a05f" />

If you have a StatefulSet called tkb-sts with five replicas and the tkb-sts-3 replica fails, the controller starts a new Pod with the same name and attaches it to the surviving volumes.

* We’ve already said that StatefulSets are for applications that need Pods to be predictable
and long-lived.
* This might involve applications connecting to specific Pods rather than
letting the Service perform round-robin load balancing across all Pods.
* To make this possible, StatefulSets use a headless Service to create reliable and predictable DNS names
for every Pod. Other apps can then query DNS (the service registry) for the full list of
Pods and make direct connections.

<img width="686" alt="image" src="https://github.com/user-attachments/assets/5b625b65-be3b-4652-a64a-eee6ee328459" />

A headless Service is a regular Kubernetes Service object without a ClusterIP address
(spec.clusterIP set to None). It becomes a StatefulSet’s governing Service when you list
it in the StatefulSet config under spec.serviceName.

#### StatefulSet vs. DaemonSet vs. Deployment
* StatefulSet: Manages stateful applications requiring stable identities and persistent storage.
* DaemonSet: Ensures a copy of a pod runs on every node for node-level services like logging.
* Deployment: Manages stateless applications with flexible, declarative updates.

### Ref
* https://www.mirantis.com/blog/kubernetes-cheat-sheet/
* https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html
