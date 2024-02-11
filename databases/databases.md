## Database and DBMS

A database is an organized collection of inter-related data that models some aspect of the real-world
(e.g., modeling the students in a class or a digital music store). 

People often confuse “databases” with
“database management systems” (e.g., MySQL, Oracle, MongoDB, Snowflake). A database management
system (DBMS) is the software that manages a database.

### Data Models

A data model is a collection of concepts for describing the data in database.
Examples: relational (most common), NoSQL (key/value, document, graph), array/matrix/vectors (for machine learning)

A schema is a description of a particular collection of data based on a data model.
* A schema is a blueprint of the database which specifies what fields will be present and what would be their types.
* For example an employee table will have an employee_ID column represented by a string of 10 digits and an employee_Name column with a string of 45 characters.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/efd2765d-37a3-4380-a67d-483ccebd3d6d)

## Types of datastore:
Databases come in a variety of genres, such as 
* Relational (Postgres),
* Key-Value (Riak, Redis, Amazon DynamoDB),
* Columnar - stores data in columns instead of rows (HBase),
* Document-oriented (MongoDB, CouchDB), and
* Graph (Neo4J).

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/3c4908e2-6e77-4f9f-bb86-0f78ae12be34)

### Relational

Relational database management systems (RDBMSs)
are set-theory-based systems implemented as two-dimensional tables with
rows and columns. The canonical means of interacting with an RDBMS is by
writing queries in Structured Query Language (SQL).

Importantly, tables can join and morph
into new, more complex tables, because of their mathematical basis in relational (set) theory.

> RDBMSs are built atop a set theory branch called relational algebra — a combination of selections (WHERE ...), projections (SELECT ...), Cartesian
products (JOIN ...), 

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

In column-oriented databases, adding columns is quite inexpensive and is done on a
row-by-row basis.

Each row can have a different set of columns, or none at
all, allowing tables to remain sparse without incurring a storage cost for null
values. With respect to structure, columnar is about midway between relational and key-value.

With respect to structure, columnar is about midway between relational and key-value.

Example: HBase, Cassandra, and Hypertable.

Using Google’s BigTablepaper as a blueprint, HBase is built on 
Hadoop (a mapreduce engine) and
designed for scaling horizontally on clusters of commodity hardware.

Ref: https://www.youtube.com/watch?v=IuJldwJLyFM&t=824s

### Document
Document-oriented databases store, well, documents. In short, a document
is like a hash, with a unique ID field and values that may be any of a variety
of types, including more hashes.

Examples: MongoDB, CouchDB

MongoDB is designed to be huge (the name mongo is extracted from the word humongous). 

Like Mongo, CouchDB’s native query language is JavaScript.

### Graph

Graph databases excel at dealing with highly interconnected data.

* A graph database consists of nodes and relationships between nodes. 
* Both nodes and relationships can have properties—key-value pairs—that store data.
* The real strength of graph databases is traversing through the nodes by following relationships.


### New SQL
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/6b5bb117-5efe-4f37-a183-480038f035fb)


Here are some NewSQL databases to get you started:
* 🔹 Google cloud spanner
* 🔹 CockroachDB
* 🔹 VoltDB
* 🔹 SingleStore
