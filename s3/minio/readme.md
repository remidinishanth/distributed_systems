All the nodes running distributed MinIO setup are recommended to be homogeneous, i.e. same operating system, same number of drives and same network interconnects.


Start distributed MinIO instance on n nodes with m drives each mounted at `/export1` to `/exportm` (pictured below), by running this command on all the `n` nodes:

<img width="8000" height="4500" alt="image" src="https://github.com/user-attachments/assets/19b44e38-a8c8-4daa-89ad-3d9a6854ecdd" />


Ref: https://github.com/minio/minio/blob/master/docs/distributed/README.md


## Erasure coding

Data shards contain a portion of a given object. Parity shards contain a mathematical representation of the object used for rebuilding Data shards.

<img width="1030" height="540" alt="image" src="https://github.com/user-attachments/assets/45f2609b-43c4-4988-99d7-b6a2c173d17a" />

The value K here constitutes the read quorum for the deployment. The erasure set must therefore have at least K healthy drives in the erasure set to support read operations.



