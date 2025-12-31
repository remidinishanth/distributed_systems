All the nodes running distributed MinIO setup are recommended to be homogeneous, i.e. same operating system, same number of drives and same network interconnects.


Start distributed MinIO instance on n nodes with m drives each mounted at `/export1` to `/exportm` (pictured below), by running this command on all the `n` nodes:

<img width="8000" height="4500" alt="image" src="https://github.com/user-attachments/assets/19b44e38-a8c8-4daa-89ad-3d9a6854ecdd" />


Ref: https://github.com/minio/minio/blob/master/docs/distributed/README.md
