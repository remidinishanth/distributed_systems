## Types of datastore:
Databases come in a variety of genres, such as 
* Relational (Postgres),
* Key-Value (Riak, Redis, Cassandra, Amazon DynamoDB),
* Columnar - stores data in columns instead of rows (HBase),
* Document-oriented (MongoDB, CouchDB), and
* Graph (Neo4J).

### Relational

Relational database management systems (RDBMSs)
are set-theory-based systems implemented as two-dimensional tables with
rows and columns. The canonical means of interacting with an RDBMS is by
writing queries in Structured Query Language (SQL).

Importantly, tables can join and morph
into new, more complex tables, because of their mathematical basis in relational (set) theory.

Examples: MySQL, H2, HSQLDB, SQLite and Postgres.

### Key Value

As the name implies, a KV store pairs keys to values in much the same way that a map (or
hashtable) would in any popular programming language.

A filesystem could be considered a
key-value store, if you think of the file path as the key and the file contents
as the value. 

Because the KV moniker demands so little, databases of this
type can be incredibly performant in a number of scenarios but generally
won’t be helpful when you have complex query and aggregation needs.

Examples: Memcached (and its cousins memcachedb and membase), Voldemort, Redis
and Riak.

### Columnar

Databases are so named because the important
aspect of their design is that data from a given column (in the two-dimensional
table sense) is stored together.

In columnoriented databases, adding columns is quite inexpensive and is done on a
row-by-row basis.

With respect to structure, columnar is about midway between relational and key-value.
