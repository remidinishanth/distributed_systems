# CMU 15-445/645: Introduction to Database Systems

Notes from the CMU Database Systems course by Andy Pavlo.

Course: [CMU 15-445/645 Fall 2024](https://15445.courses.cs.cmu.edu/fall2024/)  
Video Playlist: [YouTube](https://www.youtube.com/watch?v=vdPALZ-GCfI&list=PLSE8ODhjZXjbj8BMuIrRcacnQh20hmY9g)  
Also see: Advanced Database Systems (CMU 15-721)

Also read https://maxnilz.com/docs/003-database/003-database-storage-1/

## DBMS Architecture

```
                ┌─────────────────────┐
                │    SQL Application  │
                └─────────┬───────────┘
                          │
                ┌─────────▼───────────┐
                │   Query Planning    │
                ├─────────────────────┤
                │ Operator Execution  │
                ├─────────────────────┤
                │   Access Methods    │
                ├─────────────────────┤
                │ Buffer Pool Manager │
                ├─────────────────────┤
                │    Disk Manager     │
                └─────────────────────┘
```

<img width="1427" height="669" alt="image" src="https://github.com/user-attachments/assets/c5cb6f53-493e-4d76-a59c-c845b5ba2051" />

## Course Topics

| # | Topic | Status | Notes |
|---|-------|--------|-------|
| 1 | Relational Databases | - | See [databases.md](../databases.md) |
| 2 | Storage | Covered | [Part 1](storage1.md) &#124; [Part 2](storage2.md) &#124; [Part 3](storage3.md) |
| 3 | Query Execution | TODO | |
| 4 | Query Planning / Optimization | TODO | |
| 5 | Concurrency Control | TODO | See [2PL](../2PL.md), [MVCC](../mvcc/readme.md) |
| 6 | Database Recovery | TODO | |
| 7 | Distributed Databases | TODO | |

## Storage Notes

- **[Storage Part 1: Pages & Tuples](storage1.md)** — How the DBMS represents a database in files on disk. Covers the storage manager, database pages, heap files, page directories, slotted page layout, record IDs, and tuple layout.
- **[Storage Part 2: Log-Structured Storage & Tuple Layout](storage2.md)** — Tuple-oriented reads/writes and their problems, log-structured storage (LSM trees), compaction strategies (leveled vs universal), bloom filters, and word-aligned tuple storage.
- **[Storage Part 3: Index-Organized Storage & Data Representation](storage3.md)** — B-tree vs LSM comparison, index-organized storage, NULL data type representation, and index lifecycle (CREATE/DROP).
