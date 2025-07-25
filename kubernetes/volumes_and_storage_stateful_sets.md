---
layout: page
title: "Volumes And Storage Stateful Sets"
category: "kubernetes"
---

> Container in a pod has its own isolated filesystem, because the filesystem comes from the container’s image.

### Storage Lingo

<img width="1194" alt="image" src="https://github.com/user-attachments/assets/5f14231a-957a-4e4a-9547-d423a1546aa1" />

<img width="1197" alt="image" src="https://github.com/user-attachments/assets/fd67a4d4-7621-4155-afe7-93250a405af0" />

Ref: Kubernetes Storage Lingo 101 talk

## Volumes

Kubernetes volumes provide a way for containers in a pod to access and share data via the filesystem. 

![image](https://github.com/user-attachments/assets/2aa50e52-f653-4e8b-8f9a-a02f3dc9a34c)

<img width="1197" alt="image" src="https://github.com/user-attachments/assets/68c2796e-1b56-41b8-a39b-7b35b5372390" />


<img width="750" alt="image" src="https://github.com/user-attachments/assets/8c07890b-36dc-4b17-bdb6-45b89d73992a" />

At its core, a volume is a directory, possibly with some data in it, which is accessible to the containers in a pod.

To use a volume, specify the volumes to provide for the Pod in `.spec.volumes` and declare where to mount those volumes into containers in `.spec.containers[*].volumeMounts`.

![image](https://github.com/user-attachments/assets/2e56f26b-f725-4fbd-af16-453d01c137ec)

For each container defined within a Pod, you must independently specify where to mount each volume that the container uses.

![image](https://github.com/user-attachments/assets/f33ddd49-a7d5-4f23-b5df-7a866e248277)

![image](https://github.com/user-attachments/assets/afee4f8a-387a-402b-87df-209982e1360e)

## Special Volumes

<img width="756" alt="image" src="https://github.com/user-attachments/assets/eeff785d-da1d-461b-b518-671434333310" />


### Empty dir

An emptyDir in Kubernetes is a type of volume that is created as empty when a Pod is assigned to a Node. 
It exists as long as that Pod is running on that Node.

Provides a temporary directory that all containers in a Pod can read and write to.

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


![image](https://github.com/user-attachments/assets/cf8f5462-555d-4d6a-aff0-1b5f620cbb76)


### Secrets

Secrets are API objects in Kubernetes, which store sensitive information like passwords, keys, and tokens in base64 format. 

Secrets separate sensitive data from application code, making it easier to manage and modify independently.

`-n : no trailing newline`

![image](https://github.com/user-attachments/assets/42b8a973-a854-433a-b01a-e99d8f625c25)


![image](https://github.com/user-attachments/assets/d1880dc5-8cac-4b35-b690-ac6a0f57b843)


![image](https://github.com/user-attachments/assets/dc0b940f-e37a-43f6-8332-7bf910bfbfc0)


Ref: https://yuminlee2.medium.com/kubernetes-secrets-4287b5a83606

## Why do we need PV?

If you have your Pod on GCE, then you can port it to OnPrem or to AWS
<img width="1197" alt="image" src="https://github.com/user-attachments/assets/ab1274c0-fb64-4a3e-8238-bfbd182935d0" />


![image](https://github.com/user-attachments/assets/b894b271-cb05-47e8-a0fa-74fb5198ea59)

![image](https://github.com/user-attachments/assets/17931b9f-81d1-41d5-be28-43b009ab4f37)

<img width="1038" alt="image" src="https://github.com/user-attachments/assets/ab1c7900-7494-4401-808e-588b6893b1f5" />

Think in terms of Cluster Admin Persona and User Persona. Admins provide storage, users need not worry, they can just port the application.

<img width="953" alt="image" src="https://github.com/user-attachments/assets/d4c58e3b-4930-4546-bbe8-9b8f95f5564a" />


<img width="1031" alt="image" src="https://github.com/user-attachments/assets/413bddfc-a690-4237-86ec-488e71b4becc" />

As soon as you create the claim, Kubernetes finds the appropriate PersistentVolume
and binds it to the claim. The PersistentVolume’s capacity must be large enough to
accommodate what the claim requests.

PVC doesn't have any details regarding where the storage is, etc. **Portability**.

![image](https://github.com/user-attachments/assets/384e3899-6068-477a-a98e-8538c661fe7f)

<img width="998" alt="image" src="https://github.com/user-attachments/assets/f517add5-70a4-4bd6-9798-2ce996ff20bd" />


* Volume decouples the storage from the Container. Its lifecycle is coupled to a pod. It enables safe container restarts and sharing data between containers in a pod.

* Persistent Volume decouples the storage from the Pod. Its lifecycle is independent. It enables safe pod restarts and sharing data between pods. Allows data to persist even if the Pod is deleted.

### Host Path volumes

<img width="1035" alt="image" src="https://github.com/user-attachments/assets/f55ae0d0-5ee9-4b24-b523-4c0781ade0c0" />

If you’re thinking of using a hostPath volume as the place to store a database’s
data directory, think again. Because the volume’s contents are stored on a specific
node’s filesystem, when the database pod gets rescheduled to another node, it will no
longer see the data.

When an application running in a pod needs to persist data to disk and have that
same data available even when the pod is rescheduled to another node, you can’t use
any of the volume types we’ve mentioned so far. 
Because this data needs to be accessible from any cluster node, it must be stored on some type of network-attached storage (NAS).

Ref: Kubernetes in Action

<img width="1197" alt="image" src="https://github.com/user-attachments/assets/b85d5550-72e1-41fd-9ba4-9fe2be709a55" />

#### HostPath is different from Local PV

<img width="1197" alt="image" src="https://github.com/user-attachments/assets/d3dde3c1-1a5b-46dd-bc37-bc74cb71b457" />

<img width="891" alt="image" src="https://github.com/user-attachments/assets/1579ff95-266e-4f9f-b6fd-c6131fd162e8" />

<img width="1535" alt="image" src="https://github.com/user-attachments/assets/b8bf653d-bb3c-4fe3-ad56-311014281f0a" />


## Storage

Awesome tutorial 
* https://medium.com/@dunefro/part-1-4-container-attached-storage-with-openebs-understand-volume-provisioning-in-kubernetes-e7d7497dfe7f
* https://www.digihunch.com/2021/06/kubernetes-storage-explained/

![image](https://github.com/user-attachments/assets/da5f45cb-b404-4125-a52e-349133323357)

Ref: https://seifrajhi.github.io/blog/kubernetes-storage-deep-dive/

<img width="738" alt="image" src="https://github.com/user-attachments/assets/a97aae2a-db8e-4c91-83f7-3ce3d9a62528" />

Kubernetes Storage Options — Persistent Volumes (PV), Persistent Volume Claims (PVC), Storage Classes (SC).
* Persistent Volume — low-level representation of a storage volume.
* Persistent Volume Claim — binding between a Pod and a Persistent Volume.
* Storage Class — allows for dynamic provisioning of Persistent Volumes.

![image](https://github.com/user-attachments/assets/7d3c03c8-478a-4662-a224-d3a97176fe76)

<img width="1026" alt="image" src="https://github.com/user-attachments/assets/78328ede-3760-4da7-965c-6e07f3b34d01" />

### Kubernetes docs

![image](https://github.com/user-attachments/assets/ad3f9889-b28f-442c-b69c-f229166bb5bd)


![image](https://github.com/user-attachments/assets/f81de816-4217-4476-9e79-ecf70d7ca664)

Ref: https://seifrajhi.github.io/blog/kubernetes-storage-deep-dive/

### Deep dive

<img width="609" alt="image" src="https://github.com/user-attachments/assets/ee2fe9af-dc81-4fba-a8e9-437ab24aa823" />

PV: A PV is a storage resource located in the cluster. Administrators can manually provision PVs, and Kubernetes can use storage classes to dynamically provision PVs.

PVC: A PVC is a storage request made by a user. It works similarly to a pod but consumes PV resources rather than node resources.

CSI: The Container Storage Interface (CSI) is a standard interface that allows container orchestrators to expose arbitrary block and file storage systems to containers they manage.

![image](https://github.com/user-attachments/assets/1c2e8580-dfe6-4a08-9ea8-38bec1507abf)

### Static Provisioning and Dynamic Provisioning

#### Static Provisioning
![image](https://github.com/user-attachments/assets/fb365f50-fb40-433a-9886-6186b54bfb4f)

#### Dynamic Provisioning

The implementation of dynamic volume provisioning is based on the StorageClass abstraction.

<img width="931" alt="image" src="https://github.com/user-attachments/assets/e43b8d82-a66b-40f1-91d8-5ef5741116c7" />

### Summary

![image](https://github.com/user-attachments/assets/a93e8128-8d5d-4823-b0e1-8a7ecad0f47e)


![image](https://github.com/user-attachments/assets/56138ffc-0459-49a2-8d7c-15f2a3e2fb27)

### CSI

In Kubernetes, in-tree storage drivers were storage plugins that were directly part of the core Kubernetes code. They were being phased out in favor of CSI (Container Storage Interface) drivers, which are plug-ins that are separate from the Kubernetes core.

Before CSI, Kubernetes provided a powerful volume plugin system. These volume plugins were “in-tree” meaning their code was part of the core Kubernetes code and shipped with the core Kubernetes binaries. 

However, adding support for new volume plugins to Kubernetes was challenging. Vendors that wanted to add support for their storage system to Kubernetes (or even fix a bug in an existing volume plugin) were forced to align with the Kubernetes release process. 

![image](https://github.com/user-attachments/assets/ea3134cf-0180-4b13-8549-43c71fa09f9d)

The Kubernetes Storage Special Interest Group (SIG) defines three methods to implement a volume plugin:
1. In-tree volume plugin [deprecated],
2. Out-of-tree FlexVolume driver [deprecated],
3. Out-of-tree CSI driver.

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

## Stateful Sets

Pod mounts volume
<img width="1016" alt="image" src="https://github.com/user-attachments/assets/e51fe713-1aae-4eba-8bc7-04bc0d47b806" />

PV and PVC are binded
![image](https://github.com/user-attachments/assets/276daeaf-f2a9-45e9-bdcb-505a14c806ec)

![image](https://github.com/user-attachments/assets/51cfd7df-7e08-4a4c-a324-d70f050779fa)

### Storage classes and Dynamic Provisioning

Representation of entire class of storage
![image](https://github.com/user-attachments/assets/054fd039-5af9-4a20-be8b-6adf593d2aa1)

Admin creates storage class, and now admin no longer need to pre-provision all of the PVs

Storage classes knows how to provide the PVs on appropriate disks.

![image](https://github.com/user-attachments/assets/1f027460-0210-4045-b862-660a8ee62b34)

<img width="1129" alt="image" src="https://github.com/user-attachments/assets/2b1ad945-54dd-412d-bc91-c6a8707ab9c2" />


<img width="1199" alt="image" src="https://github.com/user-attachments/assets/f8574bc8-a2bd-40a0-ae35-4b14c26968b2" />

ReadWriteOnce (RWO): The volume can be mounted as read-write by a single **node**. This is at node level not at Pod level. If the access modes are specified as ReadWriteOncePod, the volume is constrained and can be mounted on only a single Pod.


![image](https://github.com/user-attachments/assets/a696eabb-9d44-48f5-9ca0-876a98232202)


Example:
<img width="1236" alt="image" src="https://github.com/user-attachments/assets/b6ef63da-7a45-4aee-961b-47bb91f4315c" />


#### StatefulSet vs. DaemonSet vs. Deployment
* StatefulSet: Manages stateful applications requiring stable identities and persistent storage.
* DaemonSet: Ensures a copy of a pod runs on every node for node-level services like logging.
* Deployment: Manages stateless applications with flexible, declarative updates.

### Ref
* https://www.mirantis.com/blog/kubernetes-cheat-sheet/
* https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html
