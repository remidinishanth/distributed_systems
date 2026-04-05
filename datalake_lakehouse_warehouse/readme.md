# Data Lakes, Data Warehouses & Data Lakehouses

## Table of Contents
- [What is a Data Lake?](#what-is-a-data-lake)
- [What is a Data Warehouse?](#what-is-a-data-warehouse)
- [Data Lake vs Data Warehouse](#data-lake-vs-data-warehouse)
- [The Data Swamp Problem](#the-data-swamp-problem)
- [What is a Data Lakehouse?](#what-is-a-data-lakehouse)
- [Open Table Formats](#open-table-formats)
- [OLTP vs OLAP](#oltp-vs-olap)
- [Further Reading](#further-reading)

---

## What is a Data Lake?

A **data lake** is a centralized storage repository that holds vast amounts of raw data in its **native format** — structured, semi-structured, or unstructured — until it's needed for analysis.

Instead of transforming data before loading (like a data warehouse), you dump everything as-is and figure out the structure later.

```
Sources              Data Lake              Consumers
───────              ─────────              ─────────
App DBs     ──┐
Logs        ──┤      ┌────────────────┐     Spark jobs
IoT sensors ──┼────▶ │  S3 / HDFS     │ ──▶ Presto/Trino
CSVs        ──┤      │  (raw files)   │     ML pipelines
APIs        ──┘      └────────────────┘     Dashboards
```

### What's Actually In There

Just **files on cheap storage**:
- **Parquet, ORC, Avro** (structured/columnar)
- **JSON, XML, CSV** (semi-structured)
- **Images, video, logs** (unstructured)

### Key Characteristics

- **Schema-on-read** — structure is defined at query time, not when data is loaded
- **ELT** (Extract, Load, Transform) — load raw data first, transform later
- **Cheap** — built on object storage (S3, GCS, ADLS) which costs pennies per GB/month
- **Scales to petabytes** — no need to provision servers
- **Flexible** — same data feeds SQL queries, ML training, and stream processing

### Why Data Lakes Exist

| Concern | Data Lake Advantage |
|---|---|
| **Cost** | Object storage is orders of magnitude cheaper than database storage |
| **Scale** | Petabytes of data without provisioning servers |
| **Flexibility** | Store now, decide schema later |
| **Diverse workloads** | Same data feeds SQL, ML, and streaming |

---

## What is a Data Warehouse?

A **data warehouse** is a system designed for **fast analytical queries** over structured, pre-modeled data. Data is cleaned, transformed, and loaded using a defined schema before it can be queried.

```
              ETL Pipeline                         Fast Queries
              ────────────                         ────────────
Sources ──▶ Extract ──▶ Transform ──▶ Load ──▶  ┌─────────────────┐
                                                │    Warehouse    │
                                                │   (Redshift,    │ ──▶ BI Dashboards
                                                │    Snowflake,   │ ──▶ SQL Reports
                                                │    BigQuery)    │ ──▶ Business Analytics
                                                └─────────────────┘
```

### Key Characteristics

- **Schema-on-write** — data must conform to a predefined schema before loading
- **ETL** (Extract, Transform, Load) — data is cleaned and structured before it enters the warehouse
- **Optimized for reads** — indexes, materialized views, columnar storage for fast queries
- **Structured data only** — tables, rows, columns
- **Expensive** — compute and storage are often coupled

### Common Data Warehouses

| Product | Provider |
|---|---|
| **Redshift** | AWS |
| **BigQuery** | Google Cloud |
| **Snowflake** | Multi-cloud |
| **Synapse** | Azure |
| **Apache Hive** | Open source (on Hadoop) |

---

## Data Lake vs Data Warehouse

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/e24f1ec4-467c-49f3-aaff-66064c728a5c" />


### One-Line Difference

- **Data Warehouse** — Store **processed, structured** data for fast business queries.
- **Data Lake** — Store **raw, any-format** data cheaply at massive scale.

### Side-by-Side Comparison

| | Data Warehouse | Data Lake |
|---|---|---|
| **Data format** | Structured (tables, rows, columns) | Any (Parquet, JSON, logs, images, video) |
| **Schema** | Schema-on-**write** (define before loading) | Schema-on-**read** (define at query time) |
| **Processing** | **ETL** — transform then load | **ELT** — load then transform |
| **Storage cost** | Expensive | Cheap (object storage) |
| **Query performance** | Fast (indexed, optimized, materialized views) | Slower (full scans, no indexes) |
| **Users** | Business analysts, BI tools | Data engineers, data scientists |
| **Scale** | TBs (costly at PB scale) | PBs easily |
| **ACID transactions** | Yes | No (without Hudi/Iceberg/Delta) |
| **Updates/Deletes** | Native | Difficult (rewrite entire files) |
| **Examples** | Snowflake, Redshift, BigQuery | S3 + Spark, HDFS, ADLS |

### When to Use Which

**Use a Data Warehouse when:**
- You need fast, repeatable SQL queries for dashboards and reports
- Data is structured and schema is well-understood
- Business users need self-service analytics
- Data volume is moderate (TBs, not PBs)

**Use a Data Lake when:**
- You have diverse data types (logs, events, images, IoT)
- You need to store PBs of data cheaply
- Schema isn't known upfront or changes frequently
- Data scientists need raw data for ML/experimentation

---

## The Data Swamp Problem

Without governance, data lakes degrade into **data swamps**:

```
 Healthy Data Lake           Data Swamp
 ────────────────           ──────────
 ┌────────────────────┐     ┌────────────────────┐
 │     Cataloged      │     │     No catalog     │
 │  Quality-checked   │     │   Duplicate data   │
 │ Access-controlled  │ ─▶  │ No access control  │
 │      Governed      │     │   Corrupt files    │
 └────────────────────┘     │  Unknown schemas   │
                            └────────────────────┘
```

Key issues:
- **No catalog** — nobody knows what data exists or what it means
- **No quality checks** — corrupt and duplicate data accumulates silently
- **No access control** — sensitive data is exposed
- **No ACID transactions** — concurrent writes corrupt files
- **No updates/deletes** — can't comply with GDPR-style deletion requests

This is exactly the gap that **lakehouse table formats** fill.

---

## What is a Data Lakehouse?

A **data lakehouse** combines the best of both worlds: the **scalability and cost** of a data lake with the **transactional guarantees and query performance** of a data warehouse.

```
     Traditional: Two Separate Systems
     ──────────────────────────────────
     Sources ──▶ Data Lake ──▶ ETL ──▶ Data Warehouse ──▶ Dashboards
                   │
                   └──▶ ML pipelines


     Lakehouse: Unified Architecture
     ────────────────────────────────
     Sources ──▶ Lakehouse (S3 + Hudi/Iceberg/Delta)
                   │
                   ├──▶ SQL queries   (warehouse-speed)
                   ├──▶ ML pipelines  (raw access)
                   └──▶ Dashboards    (BI tools)
```

### How It Works

A lakehouse uses **open table formats** (metadata layers) on top of cheap object storage to provide database-like features:

```
 ┌────────────────────────────────────────────────┐
 │                 Query Engines                  │
 │      Spark  .  Flink  .  Presto  .  Trino      │
 ├────────────────────────────────────────────────┤
 │            Open Table Format Layer             │
 │     Hudi  .  Apache Iceberg  .  Delta Lake     │
 │                                                │
 │    ACID transactions  .  Schema enforcement    │
 │        Time travel  .  Upserts/Deletes         │
 │       Incremental queries  .  Compaction       │
 ├────────────────────────────────────────────────┤
 │                  File Formats                  │
 │            Parquet  .  ORC  .  Avro            │
 ├────────────────────────────────────────────────┤
 │                 Object Storage                 │
 │          S3  .  GCS  .  ADLS  .  HDFS          │
 └────────────────────────────────────────────────┘
```

### What the Lakehouse Adds Over a Raw Lake

| Capability | Data Lake | Lakehouse |
|---|---|---|
| **Upserts & Deletes** | Rewrite entire files | Record-level operations |
| **ACID Transactions** | No | Yes |
| **Schema Enforcement** | No (schema-on-read) | Yes (schema evolution supported) |
| **Time Travel** | No | Query any prior snapshot |
| **Incremental Reads** | Re-read everything | Pull only changed records |
| **Concurrent Writers** | Risk of corruption | Safe via concurrency control |
| **Compaction** | Manual | Automatic background optimization |

---

## Open Table Formats

The three major open table formats that enable the lakehouse pattern:

### Apache Hudi

**H**adoop **U**pserts **D**eletes and **I**ncrementals — originally built at Uber for managing large-scale ride data.

- **Two table types:**
  - **Copy-on-Write (CoW)** — updates rewrite entire Parquet files at write time. Reads are fast (pure columnar), writes are expensive.
  - **Merge-on-Read (MoR)** — updates go to delta log files, merged at read time. Writes are fast, reads pay merge cost until compaction.
- **Strengths:** Incremental processing, record-level upserts, timeline-based metadata
- **See:** [Apache Hudi notes](../apache_hudi/readme.md)

### Apache Iceberg

An open table format specification originally created at Netflix, designed to solve correctness and performance issues with Hive tables at petabyte scale.

- **Key features:** Hidden partitioning, partition evolution (change partitioning without rewriting data), snapshot isolation, schema evolution
- **Strengths:** Spec-driven (engine-agnostic), strong partition pruning, good multi-engine support
- **See:** [Apache Iceberg notes](../apache_iceberg/readme.md)

### Delta Lake

Open-source storage layer from Databricks that brings ACID transactions to Apache Spark and data lakes.

- **Key features:** Transaction log (JSON-based), schema enforcement, Z-ordering for data skipping, change data feed
- **Strengths:** Deep Spark integration, Databricks ecosystem, Unity Catalog for governance

### Quick Comparison

| | Hudi | Iceberg | Delta Lake |
|---|---|---|---|
| **Origin** | Uber | Netflix | Databricks |
| **Write model** | CoW + MoR | CoW + MoR | CoW (MoR in development) |
| **Metadata** | Timeline (on storage) | Manifest files + manifest lists | JSON transaction log |
| **Upserts** | First-class (key-based) | Merge-into | Merge-into |
| **Incremental reads** | Native (timeline) | Snapshot diffing | Change data feed |
| **Ecosystem** | Broad (Spark, Flink, Presto) | Broadest (engine-agnostic spec) | Strongest with Spark/Databricks |

---

## OLTP vs OLAP

Understanding data warehouses and lakehouses requires understanding the two fundamental database workload types:

### OLTP — Online Transaction Processing

- **Purpose:** Serve the application — fast reads and writes of individual records
- **Pattern:** Short queries touching a few rows (e.g., "get user #123", "insert order")
- **Storage model:** Row-oriented (**N-Ary Storage Model / NSM**) — stores all attributes of a tuple contiguously
- **Examples:** MySQL, PostgreSQL, Oracle, SQL Server

```
 Row-Oriented Storage (NSM):

 Page
 ┌────────────────────────────────────────────────────┐
 │ (id=1, name="Alice", age=30, addr="NYC")           │
 │ (id=2, name="Bob",   age=25, addr="SF")            │
 │ (id=3, name="Carol", age=28, addr="LA")            │
 └────────────────────────────────────────────────────┘
 Good for: SELECT * FROM users WHERE id = 1
 Bad for:  SELECT AVG(age) FROM users   (reads all columns)
```

### OLAP — Online Analytical Processing

- **Purpose:** Analyze the data — complex aggregations over large datasets
- **Pattern:** Long queries scanning millions of rows but few columns (e.g., "average revenue per region last quarter")
- **Storage model:** Column-oriented (**Decomposition Storage Model / DSM**) — stores all values of a single attribute together
- **Examples:** Redshift, BigQuery, ClickHouse, DuckDB

```
 Column-Oriented Storage (DSM):

 id column:   [1, 2, 3, ...]
 name column: ["Alice", "Bob", "Carol", ...]
 age column:  [30, 25, 28, ...]
 addr column: ["NYC", "SF", "LA", ...]

 Good for: SELECT AVG(age) FROM users  (reads only the age column)
 Bad for:  SELECT * FROM users WHERE id = 1  (must reconstruct row)
```

### Comparison

| | OLTP | OLAP |
|---|---|---|
| **Operations** | Read/write individual records | Aggregate large datasets |
| **Query complexity** | Simple (point lookups, inserts) | Complex (joins, GROUP BY, window functions) |
| **Data scope per query** | Few rows, all columns | Many rows, few columns |
| **Storage** | Row-oriented (NSM) | Column-oriented (DSM) |
| **Concurrency** | High (many users) | Low (few analysts) |
| **Data freshness** | Real-time | Periodic loads (batch) |
| **Typical systems** | MySQL, PostgreSQL | Redshift, BigQuery, Snowflake |

> For more on storage models (NSM, DSM, PAX) see [CMU 15-445 Storage Part 3](../databases/cmu_course/storage3.md)

---

## Further Reading

- [Apache Hudi — notes and talk screenshots](../apache_hudi/readme.md)
- [Apache Iceberg — specification and resources](../apache_iceberg/readme.md)
- [S3 / Object Storage basics — data lake storage layer](../s3/s3_storage_basics.md)
- [CMU 15-445 Storage Models — NSM, DSM, PAX, OLTP vs OLAP](../databases/cmu_course/storage3.md)
- [Hadoop ecosystem — HDFS, Hive, MapReduce](../hadoop/hadoop.md)
