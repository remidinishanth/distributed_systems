Needle in a haystack: efficient storage of billions of photos

Ref: 
* https://engineering.fb.com/2009/04/30/core-infra/needle-in-a-haystack-efficient-storage-of-billions-of-photos/
* https://massivetechinterview.blogspot.com/2015/06/finding-needle-in-haystack-facebooks.html

<img width="1000" height="667" alt="image" src="https://github.com/user-attachments/assets/ad71a38a-9398-4bb6-865b-c435539162b3" />



The Photos application is one of Facebook’s most popular features. They handle billions of image scale. For each uploaded photo, Facebook generates and stores four images of different sizes, which translates to a total of 60 billion images and 1.5PB of storage.

> These numbers pose a significant challenge for the Facebook photo storage infrastructure.

<img width="1175" height="579" alt="image" src="https://github.com/user-attachments/assets/50f85038-d046-4a23-8401-c52dd20819a7" />

<img width="1110" height="756" alt="image" src="https://github.com/user-attachments/assets/5fe8cf81-9098-4854-91bd-8efeb1c8b5e4" />


<img width="780" height="451" alt="image" src="https://github.com/user-attachments/assets/f7f516bc-c7b7-43ca-9a76-3b8843556108" />

<img width="1288" height="608" alt="image" src="https://github.com/user-attachments/assets/1a643d19-6e85-4978-ae3e-afc9961c2ff9" />

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

<img width="923" height="431" alt="image" src="https://github.com/user-attachments/assets/1613080f-aaab-4b01-91f5-3f2152376b09" />

Issues

<img width="1110" height="710" alt="image" src="https://github.com/user-attachments/assets/d104bc2b-50fc-4435-9592-7c89f63da798" />

Atleast 3 disk IOPS and because of nested inode, even more
<img width="1707" height="608" alt="image" src="https://github.com/user-attachments/assets/87863699-914b-4b0f-a1f4-1413de322082" />


## Requirements

<img width="1092" height="658" alt="image" src="https://github.com/user-attachments/assets/f5e6454f-a6fd-417c-85f1-dd5ef75477c2" />

<img width="1110" height="585" alt="image" src="https://github.com/user-attachments/assets/23ff0814-6f73-432f-b74d-2230c2273970" />


## Haystack

<img width="1118" height="696" alt="image" src="https://github.com/user-attachments/assets/944bba65-46f0-484c-acfa-6944be091133" />

<img width="542" height="543" alt="image" src="https://github.com/user-attachments/assets/d20e824e-bef7-4faf-8631-a3078029c2a4" />

<img width="1269" height="547" alt="image" src="https://github.com/user-attachments/assets/48de4304-95df-424b-a7b2-5b0be0eeb399" />

### Components

<img width="1269" height="590" alt="image" src="https://github.com/user-attachments/assets/26c308ca-2f82-4a60-9fc2-970d709a88ed" />

<img width="1269" height="599" alt="image" src="https://github.com/user-attachments/assets/2ecc5c37-da17-465e-b10e-c1cc7fbe45a6" />

<img width="1269" height="599" alt="image" src="https://github.com/user-attachments/assets/3a9623ad-a9d4-4d9c-84d1-294664d35e45" />

<img width="1269" height="636" alt="image" src="https://github.com/user-attachments/assets/a2473881-b563-42bc-a9cb-d19049ce27c4" />

### Reads from Store
<img width="1269" height="584" alt="image" src="https://github.com/user-attachments/assets/c00f30b9-ecf8-4ee5-ab4c-d4ca5a5d1d9a" />


### Upload

<img width="1269" height="641" alt="image" src="https://github.com/user-attachments/assets/cf4e8176-158a-4c88-9ee8-a24465bb5da4" />

<img width="1250" height="588" alt="image" src="https://github.com/user-attachments/assets/8c848f9f-96da-49ca-8655-80394675322d" />

### Index File

<img width="1250" height="588" alt="image" src="https://github.com/user-attachments/assets/0546e13d-1d62-4b70-ad7d-fcc7430f3c3f" />

Layout of Haystack Index file
<img width="1110" height="300" alt="image" src="https://github.com/user-attachments/assets/cd4d3c3d-3ad4-447c-aaa6-4ea74338a494" />

<img width="1603" height="876" alt="image" src="https://github.com/user-attachments/assets/468d1cde-05c5-4772-9085-4dc4d82f7975" />


### Deletes

<img width="1098" height="290" alt="image" src="https://github.com/user-attachments/assets/acb8b0cc-2d60-468d-afd1-323e2da185e9" />


## Infrastructure details

<img width="810" height="408" alt="image" src="https://github.com/user-attachments/assets/20491a41-0814-4ba6-a26b-2ef99ec6ec53" />

### Storage 

<img width="786" height="484" alt="image" src="https://github.com/user-attachments/assets/a4c1baf0-bdf9-4192-a341-83342fa18598" />


### HTTP server

<img width="768" height="187" alt="image" src="https://github.com/user-attachments/assets/097cc236-d416-4ee7-a53f-98119ee0a438" />
