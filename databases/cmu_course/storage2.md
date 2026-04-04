# Storage Part 2: Log-Structured Storage & Tuple Storage

> CMU 15-445/645 — Lecture 4: Database Storage (Part 2)  
> Ref: https://15445.courses.cs.cmu.edu/fall2024/slides/04-storage2.pdf

## Table of Contents
- [Tuple-Oriented Storage: Reads](#tuple-oriented-storage-reads)
- [Tuple-Oriented Storage: Writes](#tuple-oriented-storage-writes)
- [Log-Structured Storage](#log-structure-storage)
  - [Compaction](#compaction)
  - [LSM Trees](#lsm-trees)
- [Storage of Tuples](#storage-of-tuples)
  - [Word Alignment](#word-alignment)
  - [Large Values](#large-values)
  - [System Catalogs](#system-catalogs)

---

## Tuple-Oriented Storage: Reads

<img width="1436" height="836" alt="image" src="https://github.com/user-attachments/assets/3dbfcb17-aaff-45c7-a5c9-79650fd385d2" />

Get an existing tuple using its **record ID**:
1. Check page directory to find location of page
2. Retrieve the page from disk (if not in memory)
3. Find offset in page using slot array

The DBMS relies on **indexes** to find individual tuples because the tables are inherently unsorted. But what if the DBMS could keep tuples sorted automatically using an index?


Reads and the need for indexing

<img width="1157" height="773" alt="image" src="https://github.com/user-attachments/assets/6b2a5cfe-0aa0-4182-8311-3483b1b96626" />

### Observation

Both tuple-oriented and log-structured storage approaches rely on **indexes** to find individual tuples. Indexes are necessary because the tables are inherently unsorted.

What if the DBMS could keep tuples sorted automatically using an index?

<img width="1190" height="578" alt="image" src="https://github.com/user-attachments/assets/e45eba14-19a6-4e30-9a42-de95392a6b29" />

### Index-Organized Storage

The DBMS stores a table's tuples as the value of an index data structure (typically a **B+Tree**).
- Still uses a page layout that looks like a slotted page
- Tuples are sorted in a page based on a key
- Used by: SQLite, MySQL (InnoDB), Oracle, SQL Server

> "B+Tree pays maintenance costs upfront, whereas LSMs pay for it later."

<img width="1551" height="852" alt="image" src="https://github.com/user-attachments/assets/83395fe8-bbdc-4b6a-879c-0e4e481a79b6" />

## Tuple-Oriented Storage: Writes

<img width="1211" height="786" alt="image" src="https://github.com/user-attachments/assets/de0d4fbe-8ff4-471c-b5a1-7dfadcd89b74" />

**Insert** a new tuple:
1. Check page directory to find a page with a free slot
2. Retrieve the page from disk (if not in memory)
3. Check slot array to find empty space in page that will fit

**Update** an existing tuple using its record ID:
1. Check page directory to find location of page
2. Retrieve the page from disk (if not in memory)
3. Find offset in page using slot array
4. If new data fits, overwrite existing data
5. Otherwise, mark existing tuple as deleted and insert new version in a different page



<img width="1182" height="797" alt="image" src="https://github.com/user-attachments/assets/5eb62483-61e8-4ab7-9ba7-4de10e21ac8d" />

### Problems with Tuple-Oriented Storage

1. **Fragmentation** — Pages are not fully utilized (unusable space, empty slots after deletions)
2. **Useless Disk I/O** — DBMS must fetch entire page to update one tuple
3. **Random Disk I/O** — Worst case: each tuple to update is on a separate page

What if the DBMS **cannot** overwrite data in pages and could only create new pages?  
Examples: some object stores, HDFS, Google Colossus.

## Log Structure Storage

Instead of storing tuples in pages and updating them in-place, the DBMS maintains a **log** that records changes to tuples.
- Each log entry represents a tuple **PUT** or **DELETE** operation
- Originally proposed as log-structured merge trees (LSM Trees) in 1996

The DBMS applies changes to an in-memory data structure (**MemTable**) and then writes out the changes sequentially to disk (**SSTable**).

<img width="1182" height="746" alt="image" src="https://github.com/user-attachments/assets/60557ba3-7f3b-4d1e-aae0-bb6f69a30c6d" />

![lsm](./animation.gif)

<img width="1501" height="868" alt="image" src="https://github.com/user-attachments/assets/bee6bd6c-b44c-4ac3-819c-cbd852122182" />

**Read path**: To find a key, the DBMS checks:
1. **MemTable** (in memory) — newest data
2. **SummaryTable** (in memory) — stores min/max key per SSTable + bloom filter per level
3. **Level 0 → Level 1 → ... → Level k** SSTables on disk

**SSTable** (Sorted String Table): A file of key-value pairs, sorted by key (low → high).
- Each log record must contain the tuple's unique identifier
- PUT records contain the tuple contents
- DELETE marks the tuple as deleted
- The DBMS appends log records to the end of the file without checking previous records

<img width="1435" height="817" alt="image" src="https://github.com/user-attachments/assets/22d0f72b-ea94-46da-82cf-c2c47c9100ed" />

<img width="1351" height="868" alt="image" src="https://github.com/user-attachments/assets/28cd99c0-287c-44fb-a3c9-4caae6f1f339" />

Periodically compact SSTables to reduce wasted space and speed up reads.
- Use a **sort-merge** algorithm: only keep the "latest" value for each key
- Newer records take precedence over older ones
- DELETEs remove the key entirely from the merged output

### Compaction

Two main compaction strategies:

<img width="1444" height="826" alt="image" src="https://github.com/user-attachments/assets/54f9dfae-7981-4848-8e4d-8ab13af8aff5" />

**Leveled Compaction** (used by RocksDB):
- Data is organized into levels with SSTable **size limits** per level
- SSTables in a level are **non-overlapping** on key ranges (except Level 0)
- Level 0 contains SSTables recently flushed from memory — these **may** have overlapping ranges
- Compactions merge a file from one level into the next lower level, maintaining sorted, non-overlapping key ranges
- Better for **read-heavy** workloads (fewer SSTables to check per read)

<img width="1498" height="653" alt="image" src="https://github.com/user-attachments/assets/2a3c953e-f26b-4d18-a27d-d793595b0022" />

<img width="1447" height="653" alt="image" src="https://github.com/user-attachments/assets/ad823b81-4206-433f-95ec-4978fff5b674" />

**Universal (Tiered) Compaction** (used by Cassandra):
- SSTables reside in a single "universal" level (no multi-level hierarchy)
- DBMS triggers compaction when too many SSTables overlap in key ranges or exceed size thresholds
- Better for **write-heavy** workloads and time-oriented queries

| | Leveled | Universal (Tiered) |
|---|---------|-------------------|
| Write amplification | Higher | Lower |
| Read performance | Better (fewer SSTables) | Worse (more SSTables to scan) |
| Space amplification | Lower | Higher |
| Best for | Read-heavy workloads | Write-heavy workloads |

### LSM trees

The **LSM Tree** (Log-Structured Merge Tree) is the overall architecture that ties together MemTables, WAL, SSTables, and compaction.

<img width="880" height="557" alt="image" src="https://github.com/user-attachments/assets/2651081d-8bd3-428c-a140-7b08ebc8deaf" />

**Write-Ahead Log (WAL)**: Before writing to the MemTable, the DBMS first writes the operation to a sequential log on disk. This ensures that in-memory writes aren't lost if the system crashes before flushing to an SST file.

We use WAL so that in-memory writes aren't lost before we create an immutable SST file 
<img width="3334" height="2284" alt="image" src="https://github.com/user-attachments/assets/a20329d2-344b-4511-99f9-ea954f79f38c" />

<img width="1929" height="1285" alt="image" src="https://github.com/user-attachments/assets/2317949e-b045-495d-b4c1-97dcf5982e52" />

**Write path:**
1. Write operation goes to **WAL** (sequential write on disk) for durability
2. Simultaneously written to the active **MemTable** (in memory, typically a skip list or red-black tree)
3. When MemTable is full → becomes **immutable** (read-only)
4. Immutable MemTable is **flushed** to disk as a new Level 0 SST file
5. Background **compaction** merges SST files into lower levels

**Read path:**
1. Check active MemTable
2. Check immutable MemTable (if exists)
3. Check SST files level by level (L0 → L1 → ... → Lk)
4. Use bloom filters to skip levels that definitely don't contain the key

<img width="1491" height="787" alt="image" src="https://github.com/user-attachments/assets/7e9bb547-9a85-4f89-adca-fe1da55023be" />


LSM Tree, Optimization for reads, use a bloom filter

<img width="2076" height="1260" alt="image" src="https://github.com/user-attachments/assets/3ddcf072-23fc-4124-bbb4-22ae5313ba88" />


<img width="2048" height="1152" alt="image" src="https://github.com/user-attachments/assets/d468a72f-220c-41e2-a9ba-34d090a130e5" />

**Bloom Filters** are probabilistic data structures that answer: "Is this key *possibly* in this SSTable?"
- **False positives** possible (says "maybe" when key is absent) — triggers an unnecessary disk read
- **False negatives** impossible (if it says "no", the key is definitely not there) — safely skip the SSTable
- Dramatically reduces the number of SSTables read during a lookup
- Stored in the SummaryTable in memory for fast access

<img width="1190" height="648" alt="image" src="https://github.com/user-attachments/assets/3bb84819-4d71-4b04-a581-83f2a342051f" />

**History of LSM Trees:**

| Year | System | Significance |
|------|--------|-------------|
| 1992 | LSF (Log-Structured Filesystem) | Design and implementation of a log-structured file system |
| 1996 | LSM Tree | Original paper by O'Neil et al. |
| 2006 | Bigtable | Google's distributed storage for structured data |
| 2007 | HBase | Open-source Bigtable clone on Hadoop |
| 2010 | Cassandra | Facebook's decentralized structured storage |
| 2011 | LevelDB | Google's fast persistent key-value store |
| 2013 | RocksDB | Facebook's fork of LevelDB, optimized for SSDs |
| 2015 | TSM Tree | InfluxDB's Time Structured Merge Tree |

<img width="1211" height="571" alt="image" src="https://github.com/user-attachments/assets/1ca445f6-a968-419d-98b3-6d6b8a9f1814" />

Log-structured storage is more common today, partly due to the proliferation of **RocksDB** (used by CockroachDB, TiDB, YugabyteDB, and many others).

**Downsides of log-structured storage:**
- **Write Amplification** — data is rewritten multiple times during compaction (once to L0, again when compacted to L1, etc.)
- **Compaction is Expensive** — CPU and I/O intensive; can interfere with foreground operations

### Conclusion (Log-Structured Storage)

Log-structured storage is an alternative approach to tuple-oriented architecture:
- Ideal for **write-heavy** workloads because it maximizes sequential disk I/O
- The storage manager is not entirely independent from the rest of the DBMS

## Storage of Tuples

A tuple is essentially a sequence of bytes prefixed with a **header** that contains meta-data about it.

The DBMS's **catalogs** contain the schema information about tables that the system uses to figure out the tuple's layout.

### Data Layout

Tuples are stored as `unsigned char[]` byte arrays. The DBMS interprets these bytes into attribute types using `reinterpret_cast<type*>(address)`.

<img width="1138" height="758" alt="image" src="https://github.com/user-attachments/assets/7a75fb53-ad3b-4b36-92f0-df706b672f1f" />

<img width="1559" height="622" alt="image" src="https://github.com/user-attachments/assets/7f87999c-e529-427e-a0e9-b74606da6685" />

### Word Alignment

Word alignment

<img width="1559" height="772" alt="image" src="https://github.com/user-attachments/assets/f3a03737-29aa-4f5a-8043-f35c36e6e26d" />

All attributes in a tuple must be **word aligned** to enable the CPU to access them without unexpected behavior or additional work. On a 64-bit system, each word is 8 bytes (64 bits).

Two approaches to ensure word alignment:

Padding
<img width="1559" height="772" alt="image" src="https://github.com/user-attachments/assets/d0809056-fa45-4015-9269-c371a3b51c40" />

**Approach 1: Padding** — Add empty bits after attributes to round up the storage size to the next largest word size. Simple but wastes space.

Reordering
<img width="1559" height="772" alt="image" src="https://github.com/user-attachments/assets/7fd60a1c-b874-4500-8dc8-70f68df1eac1" />

**Approach 2: Reordering** — Switch the order of attributes in the tuple's physical layout so they are naturally aligned. May still need some padding, but generally more space-efficient than pure padding.

Example: `CREATE TABLE foo (id INT, cdate TIMESTAMP, color CHAR(2), zipcode INT)`

Sizes: id=32-bit, cdate=64-bit, color=16-bit, zipcode=32-bit

```
Naive layout (no alignment):
┌──────┬──────────┬───────┬─────────┬─────────┐
│  id  │  cdate   │ color │ zipcode │ (waste) │
│ 32b  │   64b    │  16b  │   32b   │         │
└──────┴──────────┴───────┴─────────┴─────────┘
  Problem: cdate straddles a 64-bit word boundary

Padding:
┌──────┬──PAD─┬──────────┬───────┬──PAD─┬─────────┬──PAD─────┐
│  id  │ (32) │  cdate   │ color │ (16) │ zipcode │  (32)    │
│ 32b  │      │   64b    │  16b  │      │   32b   │          │
└──────┴──────┴──────────┴───────┴──────┴─────────┴──────────┘
  Every field aligned, but lots of wasted padding bytes

Reordering (most efficient):
┌──────────┬──────┬─────────┬───────┬──PAD─┐
│  cdate   │  id  │ zipcode │ color │ (16) │
│   64b    │ 32b  │   32b   │  16b  │      │
└──────────┴──────┴─────────┴───────┴──────┘
  Reorder by size (largest first) → minimal padding needed
```

<img width="1415" height="837" alt="image" src="https://github.com/user-attachments/assets/dd6e1d2e-4015-4799-8f43-8d595d58737a" />

<img width="1462" height="837" alt="image" src="https://github.com/user-attachments/assets/e4698f79-6980-46b7-a5fa-c1cf32766b4d" />

### Large Values

Most DBMSs do **not** allow a tuple to exceed the size of a single page. To store values larger than a page, the DBMS uses separate **overflow** storage pages.

| DBMS | Overflow Threshold | Mechanism |
|------|-------------------|-----------|
| PostgreSQL | > 2 KB | TOAST (The Oversized-Attribute Storage Technique) |
| MySQL | > 1/2 page size | Overflow pages |
| SQL Server | > page size | Overflow pages |

The tuple stores a pointer (size + location) to the overflow page instead of the full value. Potential optimizations: overflow compression, "German Strings" (inline prefix + pointer).

<img width="1138" height="685" alt="image" src="https://github.com/user-attachments/assets/def8f765-6d92-4335-8675-4f9ce76e3004" />

### System Catalogs

You can query the DBMS's internal **`INFORMATION_SCHEMA`** catalog to get info about the database.
- ANSI standard set of read-only views
- Provides info about all tables, views, columns, and procedures

DBMSs also have non-standard shortcuts:

<img width="1138" height="758" alt="image" src="https://github.com/user-attachments/assets/6da83404-63ac-4f57-8823-c97a75f719a4" />

| Method | DBMS |
|--------|------|
| `SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE table_catalog = '<db>'` | SQL-92 standard |
| `\d` | PostgreSQL |
| `SHOW TABLES;` | MySQL |
| `.tables` | SQLite |
