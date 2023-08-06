* Structured data (such as relational data),
* Semi-structured data (such as JSON, XML, and CSV files), and
* Unstructured data (such as images or media files).


## Object Storage Definition
Object storage is a technology that manages data as objects. All data is stored in one large repository which may be distributed across multiple physical storage devices, instead of being divided into files or folders.

It is easier to understand object-based storage when you compare it to more traditional forms of storage – file and block storage.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/17eea566-1f2e-4978-9a46-0f9652a9871f)

### File storage
File storage stores data in folders. This method, also known as hierarchical storage, simulates how paper documents are stored. When data needs to be accessed, a computer system must look for it using its path in the folder structure.

### Block storage
Block storage splits a file into separate data blocks, and stores each of these blocks as a separate data unit. Each block has an address, and so the storage system can find data without needing a path to a folder. This also allows data to be split into smaller pieces and stored in a distributed manner. Whenever a file is accessed, the storage system software assembles the file from the required blocks.

### Object storage
In object storage systems, data blocks that make up a file or “object”, together with its metadata, are all kept together. Extra metadata is added to each object, which makes it possible to access data with no hierarchy. All objects are placed in a unified address space. In order to find an object, users provide a unique ID.

Metadata is an important part of object storage technology. Metadata is determined by the user, and allows flexible analysis and retrieval of the data in a storage pool, based on its function and characteristics.


Objects are discrete units of data that are stored in a structurally flat data environment. There are no folders, directories, or complex hierarchies as in a file-based system. Each object is a simple, self-contained repository that includes the data, metadata (descriptive information associated with an object), and a unique identifying ID number (instead of a file name and file path).
