# Storage Part 1: Pages & Tuples

> CMU 15-445/645 — Lecture 3: Database Storage  
> Ref: https://15445.courses.cs.cmu.edu/fall2024/slides/03-storage1.pdf

## Table of Contents
- [Disk-Oriented DBMS](#disk-oriented-dbms)
- [File Storage](#file-storage)
- [Heap File](#heap-file)
  - [Page Directory](#page-directory)
- [Page Layout](#page-layout)
  - [Approach 1: Tuple-Oriented](#approach-1)
  - [Slotted Pages](#slotted-pages)
  - [Postgres Heap Page Example](#more-concrete-example)
- [Record IDs](#record-ids)
- [Tuple Layout](#tuple-layout)
  - [Denormalized Tuple Data](#denormalized-tuple-data)
- [Conclusion](#conclusion)

---

## Disk-Oriented DBMS

The DBMS stores a database as files on disk and moves data between disk and memory as needed.

Two key problems:
- **Problem #1**: How the DBMS represents the database in files on disk
- **Problem #2**: How the DBMS manages its memory and moves data back-and-forth from disk

<img width="1491" height="507" alt="image" src="https://github.com/user-attachments/assets/a711e202-b98a-4087-88cf-bfdd694b3ece" />


<img width="2854" height="1410" alt="image" src="https://github.com/user-attachments/assets/72fad3be-0a1d-4737-ba91-614cb789de51" />

<img width="1477" height="809" alt="image" src="https://github.com/user-attachments/assets/bd801437-aa11-46bd-9043-cc8b08befb37" />


## File Storage

The **storage manager** is responsible for maintaining a database's files.
- Organizes files as a collection of **pages**
- Tracks data read/written to pages
- Tracks available space
- Some do their own I/O scheduling to improve spatial and temporal locality

A DBMS typically does **not** maintain multiple copies of a page on disk.

<img width="1279" height="470" alt="image" src="https://github.com/user-attachments/assets/3d6dbbd5-3b92-4e94-a04c-46c74b06766e" />


<img width="1123" height="809" alt="image" src="https://github.com/user-attachments/assets/4a64eb8b-dd82-430c-9c8f-95433f228006" />

### Database Pages

A **page** is a fixed-size block of data. It can contain tuples, meta-data, indexes, or log records.
- Most systems do not mix page types
- Each page has a unique **page ID**
- The DBMS uses an indirection layer to map page IDs to physical file locations

There are three different notions of "pages":
- **Hardware Page** (usually 4KB) — the largest block the storage device can guarantee an atomic (failsafe) write
- **OS Page** (usually 4KB, x64 supports 2MB/1GB huge pages)
- **Database Page** (512B–32KB) — this is what the DBMS manages

<img width="1123" height="718" alt="image" src="https://github.com/user-attachments/assets/0e0bc003-9fe9-41ac-9ff7-05297ffbc3e9" />


<img width="1528" height="857" alt="image" src="https://github.com/user-attachments/assets/728a3480-4502-466b-92f5-4870040b48c6" />

| Page Size | DBMS |
|-----------|------|
| 4 KB | SQLite, Oracle, IBM DB2, RocksDB |
| 8 KB | SQL Server, PostgreSQL |
| 16 KB | MySQL |

DBMSs that specialize in read-only workloads tend to use larger page sizes.

### Page Storage Architecture

Different DBMSs manage pages in files on disk in different ways:
- **Heap File Organization** (focus of this lecture)
- Tree File Organization
- Sequential / Sorted File Organization (ISAM)
- Hashing File Organization

At this point in the hierarchy, we do **not** need to know anything about what is inside the pages.

## Heap File

A **heap file** is an unordered collection of pages with tuples stored in random order.
- Supports: Create / Get / Write / Delete Page
- Must also support iterating over all pages
- Need additional meta-data to track location of files and free space availability

To find a specific page: `Offset = Page# × PageSize`

<img width="1141" height="663" alt="image" src="https://github.com/user-attachments/assets/a96f4733-8626-4107-bc5f-e7f4cdb21f0e" />


<img width="1553" height="823" alt="image" src="https://github.com/user-attachments/assets/09afa63d-fb21-4cf6-8266-2e1a5a7a1d19" />
<img width="1473" height="336" alt="image" src="https://github.com/user-attachments/assets/32ef94dc-40bf-4fd1-a782-bab783a60b78" />

Two ways to find pages in a heap file:
1. **Linked List** — header page holds pointers to free page list and data page list. Requires sequential scan to find a page with free space.
2. **Page Directory** — DBMS maintains special pages that track the location of data pages and free space.

### Page directory

The DBMS maintains special pages that track the location of data pages in the database files.
- One entry per database object (table, index)
- Must keep directory pages in sync with data pages
- Also keeps meta-data about each page's contents:
  - Amount of free space per page
  - List of free / empty pages
  - Page type (data vs. meta-data)

<img width="1521" height="788" alt="image" src="https://github.com/user-attachments/assets/15a98e36-3982-4844-aed8-9321ee352549" />


## Page Layout

Every page contains a **header** of meta-data about the page's contents:
- Page Size
- Checksum
- DBMS Version
- Transaction Visibility
- Compression / Encoding Meta-data
- Schema Information

Some systems require pages to be **self-contained** (e.g., Oracle).

For any page storage architecture, we need to decide how to organize the data inside the page. Assuming row-oriented storage model:
- **Approach #1**: Tuple-oriented Storage (this lecture)
- **Approach #2**: Log-structured Storage (next lecture)
- **Approach #3**: Index-organized Storage (next lecture)

<img width="1435" height="799" alt="image" src="https://github.com/user-attachments/assets/5b8e1042-35cf-4598-9b01-82044cabf80d" />

<img width="1362" height="670" alt="image" src="https://github.com/user-attachments/assets/98ea027f-3ecc-4e72-971f-45ac99349082" />

### Approach 1

Track the number of tuples in a page, then append new tuples to the end.

**Problems with this approach:**
- Deleting a tuple leaves a gap — how to reclaim space?
- Variable-length tuples make it hard to find a specific tuple
- Fragmentation wastes space

This is why most DBMSs use **slotted pages** instead.

<img width="1024" height="559" alt="image" src="https://github.com/user-attachments/assets/17b4f871-4d84-4887-83f8-0a83bddc4443" />

### Slotted Pages

The most common page layout scheme. The slot array maps "slots" to the tuples' starting position offsets.

The header keeps track of:
- The number of used slots
- The offset of the starting location of the last slot used

**How it works:**
- The **slot array** grows from the top of the page downward
- The **tuple data** grows from the bottom of the page upward
- The page is considered full when the slot array and tuple data meet in the middle

<img width="1452" height="823" alt="image" src="https://github.com/user-attachments/assets/a8257f74-af0c-42af-82c3-4d4a28b04ad9" />

<img width="2786" height="2551" alt="image" src="https://github.com/user-attachments/assets/48824a24-8de4-42bd-bace-0536803b436d" />


<img width="1238" height="380" alt="image" src="https://github.com/user-attachments/assets/42c7c080-8109-412f-81c2-efa41b2c4c86" />

**Operations on a slotted page:**
- **Insert**: Append slot to array, write tuple at end of free space, update header
- **Delete**: Mark slot as empty. The DBMS can then compact the page to reclaim space, updating slot offsets without changing slot numbers (so external references via record IDs remain valid)
- **Compaction**: Slide tuples to remove gaps. Only slot-to-offset mappings change — external record IDs (page + slot#) stay the same

Ref: https://www.cs.swarthmore.edu/~soni/cs44/f18/Labs/lab2.html

### More concrete example

<img width="2626" height="960" alt="image" src="https://github.com/user-attachments/assets/4a0cb782-1e01-4056-983b-b58ba9438c4b" />

Ref: Internal Layout of a Heap Table File of Postgres SQL https://www.interdb.jp/pg/pgsql01/03.html

#### Writing of a Tuple

Suppose a table composed of one page that contains just one heap tuple. The pd_lower of this page points to the first line pointer, and both the line pointer and the pd_upper point to the first heap tuple. 

<img width="2578" height="542" alt="image" src="https://github.com/user-attachments/assets/863d177b-7c29-401f-87c1-1eb7b6d5254b" />

When pd_lower and pd_upper meet, the page is full — no more tuples can be inserted.

When the second tuple is inserted, it is placed after the first one. The second line pointer is appended to the first one, and it points to the second tuple. The pd_lower changes to point to the second line pointer, and the pd_upper to the second heap tuple. 

#### Reading Heap tuples

Two typical access methods, sequential scan and B-tree index scan, are outlined here:
* (a) **Sequential scan** – It reads all tuples in all pages sequentially by scanning all line pointers in each page.
* (b) **B-tree index scan** – It reads an index file that contains index tuples, each of which is composed of an index key and a TID that points to the target heap tuple.
If the index tuple with the key that you are looking for has been found1, PostgreSQL reads the desired heap tuple using the obtained TID value.

<img width="2412" height="1608" alt="image" src="https://github.com/user-attachments/assets/5e3d5b86-e9d2-43c5-8589-04304504680c" />


### Record IDs

The DBMS assigns each logical tuple a unique **record identifier** that represents its physical location in the database.
- Composed of: File ID + Page ID + Slot #
- Most DBMSs do not store record IDs in the tuple itself
- SQLite uses **ROWID** as the true primary key and stores it as a hidden attribute

Applications should **never** rely on these IDs to mean anything.

| DBMS | Record ID | Size |
|------|-----------|------|
| PostgreSQL | `ctid` (page, slot) | 6 bytes |
| SQLite | `ROWID` | 8 bytes |
| SQL Server | `%%physloc%%` | 8 bytes |
| Oracle | `ROWID` (file, block, row, object) | 10 bytes |

<img width="1492" height="838" alt="image" src="https://github.com/user-attachments/assets/a7353355-8cc1-4b47-80cc-c506ae7fa5a3" />

## Tuple Layout

A tuple is essentially a sequence of bytes. These bytes do not have to be contiguous.

It is the job of the DBMS to interpret those bytes into attribute types and values.

<img width="2382" height="990" alt="image" src="https://github.com/user-attachments/assets/51400ece-f87c-4b69-9639-c74f32372950" />

### Tuple Header

Each tuple is prefixed with a **header** that contains meta-data:
- **Visibility info** — used for concurrency control (MVCC). Which transactions can see this tuple?
- **Bit map for NULL values** — which attributes are NULL?

We do **not** need to store meta-data about the schema in the tuple — that lives in the system catalog.

<img width="1500" height="570" alt="image" src="https://github.com/user-attachments/assets/2e04e24d-ac5f-4d0f-93d6-acd99cb0a8a1" />

### Tuple Data

Attributes are typically stored in the order specified in `CREATE TABLE`. This is done for software engineering simplicity, but it might be more efficient to lay them out differently (e.g., for alignment — covered in next lecture).

<img width="1500" height="754" alt="image" src="https://github.com/user-attachments/assets/c2fa8aec-b84b-4691-8c55-52e41c7a52e1" />

### DENORMALIZED TUPLE DATA

The DBMS can physically **denormalize** (e.g., "pre-join") related tuples and store them together in the same page.
- Potentially reduces the amount of I/O for common workload patterns
- Can make updates more expensive (must update multiple copies)

Not a new idea — IBM System R did this in the 1970s. Several NoSQL DBMSs do this today (RethinkDB, CouchDB, MongoDB, RavenDB, MarkLogic) without calling it physical denormalization.

<img width="1024" height="518" alt="image" src="https://github.com/user-attachments/assets/7a21bfb3-061d-4b27-9e5e-e66b4dc35c1e" />

<img width="1500" height="801" alt="image" src="https://github.com/user-attachments/assets/99ec8946-e550-40f5-8f87-ce9215967e49" />

## CONCLUSION
* Database is organized in pages.
* Different ways to track pages.
* Different ways to store pages.
* Different ways to store tuples.
