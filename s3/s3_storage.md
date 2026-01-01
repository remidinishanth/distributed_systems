---
layout: page
title: "S3 Storage"
category: "s3"
---

* Structured data (such as relational data),
* Semi-structured data (such as JSON, XML, and CSV files), and
* Unstructured data (such as photos, videos, email, web pages, sensor data, and audio files).


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

### File storage
File storage stores data in folders. This method, also known as hierarchical storage, simulates how paper documents are stored. When data needs to be accessed, a computer system must look for it using its path in the folder structure.

### Block storage
Block storage splits a file into separate data blocks, and stores each of these blocks as a separate data unit. Each block has an address, so the storage system can find data without needing a path to a folder. This also allows data to be split into smaller pieces and stored in a distributed manner. Whenever a file is accessed, the storage system software assembles the file from the required blocks.

### Object storage
* In object storage systems, data blocks that make up a file or “object”, together with its metadata, are all kept together. Extra metadata is added to each object, which makes it possible to access data with no hierarchy. All objects are placed in a unified address space. In order to find an object, users provide a unique ID.

* Objects are discrete units of data that are stored in a structurally flat data environment. There are no folders, directories, or complex hierarchies as in a file-based system. Each object is a simple, self-contained repository that includes the data, metadata (descriptive information associated with an object), and a unique identifying ID number (instead of a file name and file path).

* Cloud object storage systems distribute this data across multiple physical devices but allow users to access the content efficiently from a single, virtual storage repository.

<img width="1095" height="780" alt="image" src="https://github.com/user-attachments/assets/87674af0-d4e5-4961-a2b3-42e3094b5dd7" />


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


## Timeline of features

<img width="759" height="671" alt="image" src="https://github.com/user-attachments/assets/37d16048-97a7-4089-b172-6fff0dd4d0d7" />

Ref: https://highscalability.com/behind-aws-s3s-massive-scale/
