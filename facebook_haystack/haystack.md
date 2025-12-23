Needle in a haystack: efficient storage of billions of photos

Ref: 
* https://engineering.fb.com/2009/04/30/core-infra/needle-in-a-haystack-efficient-storage-of-billions-of-photos/
* https://massivetechinterview.blogspot.com/2015/06/finding-needle-in-haystack-facebooks.html

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

<img width="657" height="636" alt="image" src="https://github.com/user-attachments/assets/3b7947b8-8c9a-40ce-9ae4-d935a23ec108" />

Since each image is stored in its own file, there is an enormous amount of metadata generated on the storage tier due to the namespace directories and file inodes. The amount of metadata far exceeds the caching abilities of the NFS storage tier, resulting in multiple I/O operations per photo upload or read request. The whole photo serving infrastructure is bottlenecked on the high metadata overhead of the NFS storage tier, which is one of the reasons why Facebook relies heavily on CDNs to serve photos. Two additional optimizations were deployed in order to mitigate this problem to some degree:

* `Cachr`: a caching server tier caching smaller Facebook “profile” images.
* NFS file handle cache – deployed on the photo serving tier eliminates some of the NFS storage tier metadata overhead

The major lesson we learned is that CDNs by themselves do not offer a practical solution to serving photos on a
social networking site. 

> CDNs do effectively serve the hottest photos— profile pictures and photos that have
been recently uploaded—but a social networking site like Facebook also generates a large number of requests
for less popular (often older) content, which we refer to as the long tail.

<img width="1118" height="696" alt="image" src="https://github.com/user-attachments/assets/683e247c-33a7-46e4-9509-e786e1cae96b" />


<img width="1290" height="641" alt="image" src="https://github.com/user-attachments/assets/50a057d9-f772-4128-b7f1-69f4a49e8c81" />


<img width="1288" height="608" alt="image" src="https://github.com/user-attachments/assets/1a643d19-6e85-4978-ae3e-afc9961c2ff9" />

<img width="923" height="431" alt="image" src="https://github.com/user-attachments/assets/1613080f-aaab-4b01-91f5-3f2152376b09" />

## Requirements

<img width="1092" height="658" alt="image" src="https://github.com/user-attachments/assets/f5e6454f-a6fd-417c-85f1-dd5ef75477c2" />


## Haystack

<img width="810" height="408" alt="image" src="https://github.com/user-attachments/assets/20491a41-0814-4ba6-a26b-2ef99ec6ec53" />

<img width="542" height="543" alt="image" src="https://github.com/user-attachments/assets/d20e824e-bef7-4faf-8631-a3078029c2a4" />
