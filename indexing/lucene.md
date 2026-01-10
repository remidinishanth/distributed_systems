* Ref: https://www.baeldung.com/lucene
* https://j.blaszyk.me/tech-blog/exploring-apache-lucene-index/


### Elastic Search

Elastic Search is an open-source search and analytics engine based on the Apache Lucene library.

An Elasticsearch index is a collection of documents that are related to each other. Elasticsearch stores data as JSON documents. Each document correlates a set of keys (names of fields or properties) with their corresponding values (strings, numbers, Booleans, dates, arrays of values, geolocations, or other types of data).

Elasticsearch uses a data structure called an inverted index, which is designed to allow very fast full-text searches. An inverted index lists every unique word that appears in any document and identifies all of the documents each word occurs in.

During the indexing process, Elasticsearch stores documents and builds an inverted index to make the document data searchable in near real-time. Indexing is initiated with the index API, through which you can add or update a JSON document in a specific index.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/81160f50-1185-4085-a349-83bc5ea984b8)

### ELK stack

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/db1a816f-d426-4692-85e1-2948ae0c4a50)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a8a5e424-9185-4de8-a3d6-fdfc09332b26)


![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/bd2a4323-1067-4f08-b41d-3fe723dc5e4d)

Ref: https://www.elastic.co/what-is/elasticsearch"

### Indexing
Simply put, Lucene uses an “inverted indexing” of data – instead of mapping pages to keywords, it maps keywords to pages just like a glossary at the end of any book.

This allows for faster search responses, as it searches through an index, instead of searching through text directly.

Lucene stores documents as an “inverted index”, in that it keeps track of the list of unique words per document (so words to documents, as opposed to documents to words).

# TODO

* Ref: https://www.baeldung.com/lucene
* https://j.blaszyk.me/tech-blog/exploring-apache-lucene-index/


### Elastic Search

Elastic Search is an open-source search and analytics engine based on the Apache Lucene library.

An Elasticsearch index is a collection of documents that are related to each other. Elasticsearch stores data as JSON documents. Each document correlates a set of keys (names of fields or properties) with their corresponding values (strings, numbers, Booleans, dates, arrays of values, geolocations, or other types of data).

Elasticsearch uses a data structure called an inverted index, which is designed to allow very fast full-text searches. An inverted index lists every unique word that appears in any document and identifies all of the documents each word occurs in.

During the indexing process, Elasticsearch stores documents and builds an inverted index to make the document data searchable in near real-time. Indexing is initiated with the index API, through which you can add or update a JSON document in a specific index.

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/81160f50-1185-4085-a349-83bc5ea984b8)

### ELK stack

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/db1a816f-d426-4692-85e1-2948ae0c4a50)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a8a5e424-9185-4de8-a3d6-fdfc09332b26)


![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/bd2a4323-1067-4f08-b41d-3fe723dc5e4d)

Ref: https://www.elastic.co/what-is/elasticsearch
