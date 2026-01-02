---
layout: page
title: "S3 Storage"
category: "s3"
---

* Structured data (such as relational data),
* Semi-structured data (such as JSON, XML, and CSV files), and
* Unstructured data (such as photos, videos, email, web pages, sensor data, and audio files).

<img width="1700" height="674" alt="image" src="https://github.com/user-attachments/assets/865ffd0b-9479-4086-adf0-bc70f9d3ee0f" />


Blob storage is a type of cloud storage for unstructured data. 
* Unstructured data is data that doesn't adhere to a particular data model or definition, such as text or binary data.
* A "blob," which is short for **B**inary **L**arge **Ob**ject, is a mass of data in binary form that does not necessarily conform to any file format.
* Blob storage keeps these masses of data in non-hierarchical storage areas called data lakes.


## S3, Azure Blob Storage, Google Cloud Storage

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/1b2b5c2e-fe49-4384-896e-a7bc7df54387)

### Google Cloud storage
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/42b7d597-18c6-40aa-9bdd-4736ea06a70b)
* Encryption at Rest with keys you manage or Google manage. No need to worry about security.

* Based on the frequency of usage, we have storage classes

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/06093b29-1d9e-456f-a57b-f1d12c6fdd3e)


## Object Storage Definition
Object storage is a technology that manages data as objects. All data is stored in one large repository which may be distributed across multiple physical storage devices, instead of being divided into files or folders.

It is easier to understand object-based storage when you compare it to more traditional forms of storage – file and block storage.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/17eea566-1f2e-4978-9a46-0f9652a9871f)

<img width="720" height="540" alt="image" src="https://github.com/user-attachments/assets/8e6b5470-c07e-49e1-9bc1-3de0b3503192" />

### File storage
File storage stores data in folders. This method, also known as hierarchical storage, simulates how paper documents are stored. When data needs to be accessed, a computer system must look for it using its path in the folder structure.

### Block storage
Block storage splits a file into separate data blocks, and stores each of these blocks as a separate data unit. Each block has an address, so the storage system can find data without needing a path to a folder. This also allows data to be split into smaller pieces and stored in a distributed manner. Whenever a file is accessed, the storage system software assembles the file from the required blocks.

### Object storage
* In object storage systems, data blocks that make up a file or “object”, together with its metadata, are all kept together. Extra metadata is added to each object, which makes it possible to access data with no hierarchy. All objects are placed in a unified address space. In order to find an object, users provide a unique ID.

* Objects are discrete units of data that are stored in a structurally flat data environment. There are no folders, directories, or complex hierarchies as in a file-based system. Each object is a simple, self-contained repository that includes the data, metadata (descriptive information associated with an object), and a unique identifying ID number (instead of a file name and file path).

* Cloud object storage systems distribute this data across multiple physical devices but allow users to access the content efficiently from a single, virtual storage repository.

<img width="1095" height="780" alt="image" src="https://github.com/user-attachments/assets/87674af0-d4e5-4961-a2b3-42e3094b5dd7" />



<img width="1496" height="740" alt="image" src="https://github.com/user-attachments/assets/4412fec1-4354-4c23-8d9a-1228a115e9c4" />

The design philosophy of object storage is very similar to that of the UNIX file system. 
* In UNIX, when we save a file in the local file system, it does not save the filename and file data together.
* Instead, the filename is stored in a data structure called "inode" and the file data is stored in different disk locations.
* The inode contains a list of file block pointers that point to the disk locations of the file data.
* When we access a local file, we first fetch the metadata in the inode. We then read the file data by following the file block pointers to the actual disk locations.

The object storage works similarly. The inode becomes the metadata store that stores all the object metadata. The hard disk becomes the data store that stores the object data. In the UNIX file system, the inode uses the file block pointer to record the location of data on the hard disk.


<img width="559" height="471" alt="image" src="https://github.com/user-attachments/assets/d3d69c0e-a578-44d7-97bd-68d92a354ef9" />


## Multi-part upload

<img width="693" height="389" alt="image" src="https://github.com/user-attachments/assets/fdcf7090-1144-494b-b132-51eeb4152b6f" />


## Metadata in the Object store

Metadata is critical to object storage technology. 

* With object storage, objects are kept in a single bucket and are not files inside of folders.
* Instead, object storage combines the pieces of data that make up a file, adds all the user-created metadata to that file, and attaches a custom identifier.
* This creates a flat structure, called a bucket, as opposed to hierarchical or tiered storage.
* This lets you retrieve and analyze any object in the bucket, no matter the file type, based on its function and characteristics.

Object storage is the ideal storage for data lakes because it delivers an architecture for large amounts of data, with each piece of data stored as an object, and the object metadata provides a unique identifier for easier access. 

#### What is a data lake?

A data lake is a centralized repository that allows you to store all your structured and unstructured data at any scale. You can store your data as-is, without having to first structure the data, and run different types of analytics—from dashboards and visualizations to big data processing, real-time analytics, and machine learning to guide better decisions.

You can seamlessly and nondisruptively increase storage from gigabytes to petabytes of content, paying only for what you use. 


Ref: https://www.ibm.com/topics/object-storage

