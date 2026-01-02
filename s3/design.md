Ref: https://www.codesmith.io/blog/amazon-s3-storage-diagramming-system-design

<img width="1168" height="610" alt="image" src="https://github.com/user-attachments/assets/b8dc2118-d57d-4169-8657-ee3b391a6b11" />

**Bucket:** A logical container for objects. The bucket name is globally unique. To upload data to S3, we must first create a bucket. 

**Object:** An object is an individual piece of data we store in a bucket. It contains object data (also called payload) and metadata. Object data can be any sequence of bytes we want to store. The metadata is a set of name-value pairs that describe the object.




<img width="1496" height="788" alt="image" src="https://github.com/user-attachments/assets/4efdb68d-d240-4045-87e8-e3e315cee38e" />

<img width="1496" height="740" alt="image" src="https://github.com/user-attachments/assets/58af9939-8535-4cdb-bd58-030dcf701410" />


<img width="969" height="698" alt="image" src="https://github.com/user-attachments/assets/3f5317e9-0deb-46b9-8ce8-ce6d249af493" />

## Timeline of features

<img width="759" height="671" alt="image" src="https://github.com/user-attachments/assets/37d16048-97a7-4089-b172-6fff0dd4d0d7" />

Ref: https://highscalability.com/behind-aws-s3s-massive-scale/

<img width="904" height="450" alt="image" src="https://github.com/user-attachments/assets/3ba82a6a-9a6b-41cf-b6f5-aa0aa95b8418" />

<img width="901" height="414" alt="image" src="https://github.com/user-attachments/assets/09209329-5c2c-4c94-a959-d95dd652ed34" />

<img width="1095" height="780" alt="image" src="https://github.com/user-attachments/assets/ee76d7a6-7116-43a2-9872-6eeb30432f69" />


<img width="1036" height="407" alt="image" src="https://github.com/user-attachments/assets/afcdf63d-38e7-4ae5-9729-6e5f5300e4d6" />

<img width="1406" height="1226" alt="image" src="https://github.com/user-attachments/assets/d64329e5-4e73-4fba-adb2-a63dbd0b81cc" />

## Architecture

S3 is said to be composed of more than 300 microservices.

It tries to follow the core design principle of simplicity.

You can distinct its architecture by four high-level services:
* a front-end fleet with a REST API
* a namespace service
* a storage fleet full of hard disks
* a storage management fleet that does background operations, like replication and tiering.

<img width="879" height="650" alt="image" src="https://github.com/user-attachments/assets/6c9fe2fa-17af-4178-8818-2125973d9069" />

<img width="1600" height="1088" alt="image" src="https://github.com/user-attachments/assets/5b163daa-97a1-4225-b74b-f3418f8362f6" />


## Upload

<img width="1470" height="814" alt="image" src="https://github.com/user-attachments/assets/29fc2991-04b4-429a-8068-0b549703d859" />


<img width="2196" height="2646" alt="image" src="https://github.com/user-attachments/assets/05ded656-3837-4ed4-85eb-09784618d7d2" />


### Storage Fleet

<img width="802" height="617" alt="image" src="https://github.com/user-attachments/assets/b8051fa1-f417-4229-a1e4-02c90739fcff" />

### Hard Drives

<img width="792" height="581" alt="image" src="https://github.com/user-attachments/assets/5cfdedcf-6a00-4f24-9fec-52963171f21c" />

<img width="1600" height="841" alt="image" src="https://github.com/user-attachments/assets/a1f1330a-0dd8-43cb-b989-916b7431b4c9" />

### Replication

<img width="746" height="615" alt="image" src="https://github.com/user-attachments/assets/c623f769-b8c6-42f4-a567-534596f318de" />

### Heat Management at Scale

<img width="1492" height="1502" alt="image" src="https://github.com/user-attachments/assets/0ebaa341-22d9-4d3b-9188-f82ceddaa4be" />

<img width="1600" height="889" alt="image" src="https://github.com/user-attachments/assets/8d42b67b-20d9-4683-8bc9-23e8e2d0df82" />

But as the system aggregates millions of workloads, the underlying traffic to the storage flattens out remarkably. The aggregate demand results in a smoothened out, more predictable throughput.

When you aggregate on a large enough scale, a single workload cannot influence the aggregate peak.

The problem then becomes much easier to solve - you simply need to balance out a smooth demand rate across many disks.
