Ref: https://www.codesmith.io/blog/amazon-s3-storage-diagramming-system-design

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
