# Storage Part 3: Index-Organized Storage & Data Representation

> CMU 15-445/645 — Lecture 5: Storage Models & Compression (partial)  
> Ref: https://15445.courses.cs.cmu.edu/fall2024/slides/05-storage3.pdf

## Table of Contents
- [Trees in Database Storage](#trees-in-database-storage)
- [Index-Organized Storage](#index-organized-storage)
- [NULL Data Types](#null-data-types)
- [Indexes](#indexes)

---

## Trees in Database Storage

Two fundamental tree-based approaches to organizing data on disk:

<img width="1491" height="787" alt="image" src="https://github.com/user-attachments/assets/33cf915c-8712-4c5f-a49a-98bc217efb52" />

| | B-Tree | Log-Structured (LSM) |
|---|--------|---------------------|
| **Unit** | Node = Page (KBs) | SSTable file (MBs) |
| **Write pattern** | In-place update, may touch multiple pages | Sequential append to log |
| **Read** | O(log n) traversal via pointers | Check MemTable → bloom filters → SSTables level by level |
| **Pointers** | Pointers across nodes across levels | No pointers across SSTables |
| **Write cost** | Writes may update multiple pages (node splits) | Writes are fast (sequential I/O) but reads may be slow |
| **Size** | Same page size across the tree | SST size grows at lower levels |

**Key insight:** B-Tree pays maintenance costs **upfront** (on writes), whereas LSM pays for it **later** (on reads and during compaction).

### Observation

Both storage approaches discussed so far rely on **indexes** to find individual tuples — indexes are necessary because the tables are inherently unsorted.

But what if the DBMS could keep tuples sorted automatically using an index?

<img width="1491" height="649" alt="image" src="https://github.com/user-attachments/assets/a6463b1f-0b6a-4f01-a2cf-92571bfd04ea" />

## Index-Organized Storage

The DBMS stores a table's tuples as the value of an index data structure (typically a **B+Tree**).
- Still uses a page layout that looks like a slotted page
- Tuples are sorted in a page based on a key

Used by: SQLite, MySQL (InnoDB), Oracle, SQL Server.

<img width="1491" height="752" alt="image" src="https://github.com/user-attachments/assets/0fce49f2-7b5b-4393-8db8-4d115440fc9c" />

```
Index-Organized Page Layout:

  ┌──────────────────────────────────────────┐
  │               Header                      │
  ├────────┬────────┬────────┬────────────────┤
  │ key+   │ key+   │ key+   │                │
  │ offset │ offset │ offset │   (free)       │
  ├────────┴────────┴────────┤                │
  │          ...              │                │
  │                           ├────────────────┤
  │                           │   Tuple #3     │
  │         ◄─────────────────│   Tuple #2     │
  │                           │   Tuple #1     │
  └───────────────────────────┴────────────────┘
  Tuples are sorted in key-order within the page.
  Inner nodes contain keys + child page pointers.
  Leaf nodes contain keys + full tuple data.
```

## NULL Data Types

How does the DBMS represent NULL values in a tuple?

<img width="1491" height="752" alt="image" src="https://github.com/user-attachments/assets/5ecf59a9-7043-4576-9ccb-d7dbff6a8697" />

Three choices:

**Choice #1: Null Column Bitmap Header** (most common)
- Store a bitmap in a centralized header that specifies which attributes are NULL
- One bit per attribute — very space efficient
- Example: for a tuple with 8 columns, a single byte bitmap where `1` means NULL

```
Tuple Header:  [ visibility | null_bitmap: 0b00100100 ]
                                              ↑    ↑
                                            col5  col2 are NULL
```

**Choice #2: Special Values**
- Designate a specific value to represent NULL for each data type
- Example: `INT32_MIN` for integers, empty string for text
- Drawback: reduces the valid value range for that type

**Choice #3: Per Attribute Null Flag**
- Store a flag per attribute that marks whether the value is NULL
- Must use more than just a single bit because this messes up word alignment
- Wastes more space than the bitmap approach

## Indexes

An index is a data structure that enables fast lookup of tuples based on a subset of their attributes (the "key").

<img width="1491" height="752" alt="image" src="https://github.com/user-attachments/assets/9e9dae93-aa8f-4a07-8aa6-5773e3e459fd" />

### CREATE INDEX

Building an index on an existing table:
1. Scan the entire table and populate the index
2. Record changes made by transactions that modified the table while the index was being built
3. When the scan completes, **lock the table** and resolve changes that were missed after the scan started

### DROP INDEX

Removing an index:
1. Just drop the index **logically** from the catalog
2. It only becomes "invisible" when the transaction that dropped it **commits**
3. All existing transactions will still have to update the index until the drop commits

> Indexes themselves are covered in much more detail in later lectures (B+Tree indexes, hash indexes, etc.)
