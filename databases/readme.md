Todo: https://www.youtube.com/watch?v=vdPALZ-GCfI&list=PLSE8ODhjZXjbj8BMuIrRcacnQh20hmY9g

F2023 #00 - Course Overview & Logistics (CMU Intro to Database Systems)

Also there is Adv Databases course

![image](https://github.com/user-attachments/assets/9ac0d78a-9437-47df-b122-87b754b9929f)

## SQL

### Atomicity (all-or-nothing guarantee)
Atomicity: The system can only be in the state it was before the operation or after the operation, not something in between.

In the context of ACID, atomicity is not about concurrency. It does not describe what happens if several processes try to access the same data at the same time, because that is covered under the letter I, for isolation.

Atomicity simplifies this problem: if a transaction was aborted, the application can be sure that it didn’t change anything, so it can safely be retried.

### Consistency

The idea of ACID consistency is that you have certain statements about your data (invariants) that must always be true — for example, in an accounting system, credits and debits across all accounts must always be balanced.

However, this idea of consistency depends on the application’s notion of invariants, and it’s the application’s responsibility to define its transactions correctly so that they preserve consistency. 

Atomicity, isolation, and durability are properties of the database, whereas consis‐ tency (in the ACID sense) is a property of the application.

### Isolation levels

Concurrently running transactions shouldn’t interfere with each other. For example, if one transaction makes several writes, then another transaction should see either all or none of those writes, but not some subset.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/427f890e-a716-47ba-b274-b3caaeaad5af)

![image](https://github.com/user-attachments/assets/e9f83dd0-16d0-40a6-946a-d5b5e42d1587)

🔹 Serializalble: This is the highest isolation level. Concurrent transactions are guaranteed to be executed in sequence.

🔹 Repeatable Read: Data read during the transaction stays the same as the transaction starts.

🔹 Read Committed: Data modification can only be read after the transaction is committed.

🔹 Read Uncommitted: The data modification can be read by other transactions before a transaction is committed.

![image](https://github.com/user-attachments/assets/24cd023b-1b63-43eb-afa2-02acf7b944c6)


#### Dirty Read

Violating isolation: one transaction reads another transaction’s uncommitted writes (a “dirty read”)

![image](https://github.com/user-attachments/assets/096db127-df30-460f-9f72-58ec3ad69b54)
Ref: https://maxnilz.com/docs/003-database/015-concurrency-control/

No need to even write, you can just read too and you might end up reading uncommitted write

![image](https://github.com/user-attachments/assets/49cf0f86-f9f2-4e1e-bc6d-2dbb7c025bb0)

### Summary
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/24b3af22-f1b1-4bbb-ab25-0da20bea4dd7)

### SQL Databases

* Cloud Agnostic: Oracle, Microsoft SQL Server, IBM DB2, PostgreSQL, and MySQL
* AWS: Hosted PostgreSQL and MySQL in Relational Database Service (RDS)
* Microsoft Azure: Hosted SQL Server as Azure SQL Database
* Google Cloud: Hosted PostgreSQL and MySQL in Cloud SQL, and also horizontally scaling Cloud Spanner

![image](https://github.com/user-attachments/assets/e91eb4ea-e525-4b7e-8bea-d715f7b99977)


Nice read about Cockroach DB https://www.cockroachlabs.com/blog/distributed-sql-key-value-store/

<img width="1074" alt="image" src="https://github.com/user-attachments/assets/e24257e3-44df-4450-92f8-5a4e24d8a07b">
