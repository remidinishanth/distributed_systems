Needle in a haystack: efficient storage of billions of photos

Ref: https://engineering.fb.com/2009/04/30/core-infra/needle-in-a-haystack-efficient-storage-of-billions-of-photos/

<img width="1000" height="667" alt="image" src="https://github.com/user-attachments/assets/ad71a38a-9398-4bb6-865b-c435539162b3" />



The Photos application is one of Facebook’s most popular features. They handle billions of image scale. For each uploaded photo, Facebook generates and stores four images of different sizes, which translates to a total of 60 billion images and 1.5PB of storage.

> These numbers pose a significant challenge for the Facebook photo storage infrastructure.

<img width="1175" height="579" alt="image" src="https://github.com/user-attachments/assets/50f85038-d046-4a23-8401-c52dd20819a7" />


<img width="780" height="451" alt="image" src="https://github.com/user-attachments/assets/f7f516bc-c7b7-43ca-9a76-3b8843556108" />



## NFS photo infrastructure
The old photo infrastructure consisted of several tiers:

* Upload tier receives users’ photo uploads, scales the original images and saves them on the NFS storage tier.
* Photo serving tier receives HTTP requests for photo images and serves them from the NFS storage tier.
* NFS storage tier built on top of commercial storage appliances.


Since each image is stored in its own file, there is an enormous amount of metadata generated on the storage tier due to the namespace directories and file inodes. The amount of metadata far exceeds the caching abilities of the NFS storage tier, resulting in multiple I/O operations per photo upload or read request. The whole photo serving infrastructure is bottlenecked on the high metadata overhead of the NFS storage tier, which is one of the reasons why Facebook relies heavily on CDNs to serve photos. Two additional optimizations were deployed in order to mitigate this problem to some degree:

* Cachr: a caching server tier caching smaller Facebook “profile” images.
* NFS file handle cache – deployed on the photo serving tier eliminates some of the NFS storage tier metadata overhead


<img width="1288" height="608" alt="image" src="https://github.com/user-attachments/assets/1a643d19-6e85-4978-ae3e-afc9961c2ff9" />
