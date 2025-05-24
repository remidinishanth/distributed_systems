A log-structured merge-tree (LSM tree) is a data structure and storage architecture used by RocksDB.

## SSTable (Sorted String Table)
An SSTable is essentially a file format. It's a simple, immutable file on disk that stores key-value pairs, sorted strictly by their keys.
