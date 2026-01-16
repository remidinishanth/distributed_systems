<img width="1139" height="726" alt="image" src="https://github.com/user-attachments/assets/b09e88b4-9b1f-4112-a418-a43187b8f838" />
Ref: https://github.com/minio/directpv?tab=readme-ov-file

DirectPV is designed to be lightweight and scalable to tens of thousands of drives. It is made up of three components - Controller, Node Driver, UI.

<img width="1240" height="935" alt="image" src="https://github.com/user-attachments/assets/90844b41-063b-4a67-b9d0-61721cfa6a7f" />


Ref: https://blog.min.io/introducing-directpv/ and https://www.min.io/directpv

> DirectPV is similar to LocalPV but dynamically provisioned.

Ref: https://blog.min.io/the-story-of-directpv/

## How does it choose drives

<img width="1632" height="1788" alt="image" src="https://github.com/user-attachments/assets/09c406a6-0762-44b0-80ab-2281ecb1096e" />


<img width="751" height="799" alt="image" src="https://github.com/user-attachments/assets/fd3809a5-7b97-42f3-842c-b3f417dc2828" />

Ref: https://docs.min.io/enterprise/minio-directpv/resource-management/scheduling/



## DirectPV CSI driver architecture

DirectPV is implemented as per the [CSI specification](https://github.com/container-storage-interface/spec/blob/master/spec.md). It comes with the below components run as Pods in Kubernetes.
* `Controller`
* `Node server`

When DirectPV contains legacy volumes from `DirectCSI` (old name of Direct PV), the below additional components also run as Pods.
* `Legacy controller `
* `Legacy node server`

## Controller
The Controller runs as `Deployment` Pods named `controller`, which are three replicas located in any Kubernetes nodes. In the three replicas, one instance is elected to serve requests. Each pod contains below running containers:
* `CSI provisioner` - Bridges volume creation and deletion requests from `Persistent Volume Claim` to CSI controller.
* `Controller` - Controller server which honors CSI requests to create, delete and expand volumes.
* `CSI resizer` - Bridges volume expansion requests from `Persistent Volume Claim` to CSI controller.

### Controller server
Controller server runs as container `controller` in a `controller` `Deployment` Pod. It handles below requests:
* `Create volume` - Controller server creates new `DirectPVVolume` CRD after reversing requested storage space on suitable `DirectPVDrive` CRD. For more information, refer to the [Volume scheduling guide](./volume-scheduling.md)
* `Delete volume` - Controller server deletes `DirectPVVolume` CRD for unbound volumes after releasing previously reserved space in `DirectPVDrive` CRD.
* `Expand volume` - Controller server expands `DirectPVVolume` CRD after reversing requested storage space in `DirectPVDrive` CRD.

Below is a workflow diagram
```
┌────────────┐                                               ┌────────────┐
│            │ Create Event ┌─────────────┐ CreateVolume API │            │   ┌────────────────────┐
│            │------------->│     CSI     │----------------->│            │-->│  DirectPVDrive CRD │
│ Persistent │ Delete Event │ Provisioner │ DeleteVolume API │            │   └────────────────────┘
│   Volume   │------------->│             │----------------->│ Controller │
│   Claim    │              └─────────────┘                  │   Server   │
│   (PVC)    │                                               │            │   ┌────────────────────┐
│            │ Update Event ┌─────────────┐ ExpandVolume API │            │-->│ DirectPVVolume CRD │
│            │------------->│ CSI Resizer │----------------->│            │   └────────────────────┘
│            │              └─────────────┘                  └────────────┘
└────────────┘
```

## Node server
Node server runs as `DaemonSet` Pods named `node-server` in all or selected Kubernetes nodes. Each node server Pod runs on a node independently. Each pod contains below running containers:
* `Node driver registrar` - Registers node server to kubelet to get CSI RPC calls.
* `Node server` - Honors stage, unstage, publish, unpublish and expand volume RPC requests.
* `Node controller` - Honors CRD events from `DirectPVDrive`, `DirectPVVolume`, `DirectPVNode` and `DirectPVInitRequest`.
* `Liveness probe` - Exposes `/healthz` endpoint to check node server liveness by Kubernetes.

Below is a workflow diagram
```
┌─────────┐                    ┌────────┐                 ┌──────────────────────────────────┐    ┌────────────────────┐
│         │  StageVolume RPC   │        │   StageVolume   │ * Create data directory          │    │                    │
│         │------------------->│        │---------------->│ * Set xfs quota                  │<-->│                    │
│         │                    │        │                 │ * Bind mount staging target path │    │                    │
│         │                    │        │                 └──────────────────────────────────┘    │                    │
│         │ PublishVolume RPC  │        │  PublishVolume  ┌──────────────────────────────────┐    │                    │
│         │------------------->│        │---------------->│ * Bind mount target path         │<-->│                    │
│ Kubelet │                    │  Node  │                 └──────────────────────────────────┘    │ DirectPVDrive CRD  │
│         │UnpublishVolume RPC │ Server │ UnpublishVolume ┌──────────────────────────────────┐    │ DirectPVVolume CRD │
│         │------------------->│        │---------------->│ * Unmount target path            │<-->│                    │
│         │                    │        │                 └──────────────────────────────────┘    │                    │
│         │ UnstageVolume RPC  │        │  UnstageVolume  ┌──────────────────────────────────┐    │                    │
│         │------------------->│        │---------------->│ * Unmount staging target path    │<-->│                    │
│         │                    │        │                 └──────────────────────────────────┘    │                    │
│         │  ExpandVolume RPC  │        │   ExpandVolume  ┌──────────────────────────────────┐    │                    │
│         │------------------->│        │---------------->│ * Set xfs quota                  │<-->│                    │
└─────────┘                    └────────┘                 └──────────────────────────────────┘    └────────────────────┘
```



Ref: https://github.com/minio/directpv/blob/master/docs/architecture.md
