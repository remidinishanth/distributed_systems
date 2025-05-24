A log-structured merge-tree (LSM tree) is a data structure and storage architecture used by RocksDB.

## SSTable (Sorted String Table)
An **SSTable** is essentially a **file format**. It's a simple, immutable file on disk that stores key-value pairs, sorted strictly by their keys.

Sorted Strings Table (SSTable) is a persistent file format used by ScyllaDB, Apache Cassandra, and other NoSQL databases to take the in-memory data stored in memtables, order it for fast access, and store it on disk in a persistent, ordered, immutable set of files.

![image](https://github.com/user-attachments/assets/347e96e8-8729-414b-a825-2e984cfdc018)

* When data is committed, ScyllaDB or Cassandra stores the changes in a commitlog, which is a file that only allows appending, so writes are quick.
* Simultaneously the data is written to an in-memory cache of key/column data called a memtable.
* Periodically the memtable is flushed to persistent storage in the form of SSTables on disk.
* SSTables in Cassandra or ScyllaDB serve as the building blocks of the total data stored in the database.
* SSTables are immutable, so updates to data create a new SSTable file instead of changing the existing ones.
