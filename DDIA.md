Some notes at https://notes.shichao.io/books/ and https://notes.shichao.io/dda/

Hierarchical model 
* https://mariadb.com/kb/en/understanding-the-hierarchical-database-model/
* https://www.db-book.com/db6/appendices-dir/e.pdf
* https://www.ibm.com/docs/en/zos-basic-skills?topic=product-ims-database-manager

The name of an IMS segment becomes the table name in an SQL query, and the name of a field becomes the column name in the SQL query.

A fundamental difference between segments in a hierarchical database and tables in a relational database is that, in a hierarchical database, segments are implicitly joined with each other. In a relational database, you must explicitly join two tables. A segment instance in a hierarchical database is already joined with its parent segment and its child segments, which are all along the same hierarchical path. In a relational database, this relationship between tables is captured by foreign keys and primary keys.
