---
layout: page
title: "Bloom Filter"
category: "general"
---

> Probabilistic Data Structure To Check Membership

![image](https://github.com/user-attachments/assets/1e8361b7-303b-4f2f-b9d9-6fff343830b3)

False positives are possible. it might say, element might exist, even thought it isn't.
* But when it doesn't exist, it would never say it exists.

Ref to https://systemdesign.one/bloom-filters-explained/ 

## Requirements
Design a data structure with the following characteristics:
* constant time complexity to test membership
* a small amount of memory to test membership
* insert and query operations should be parallelizable
* test membership can yield approximate results

## Insertion

* the item is hashed through k hash functions
* the modulo n (length of bit array) operation is executed on the output of the hash functions to identify the k array positions (buckets)
* the bits at all identified buckets are set to one

There is a probability that some bits on the array are set to one multiple times due to hash collisions.

![image](https://github.com/user-attachments/assets/220efe39-5c59-41d4-acef-14d79f358613)

## Membership test
* Verify if the hash bits at all identified buckets are set to one

![image](https://github.com/user-attachments/assets/2dc29d8f-574a-41d9-b2db-c5e7804c1221)

### False positive case

![image](https://github.com/user-attachments/assets/3b31929d-0ee9-4cef-96cf-fbefa416b8ec)

## Bloom filter calculator

The hur.st bloom filter calculator can be used to choose an optimal size for the bloom filter and explore how the different parameters interact. The accuracy of the bloom filter depends on the following:

* number of hash functions (k)
* quality of hash functions
* length of the bit array (n)
* number of items stored in the bloom filter

The properties of an optimal hash function for the bloom filter are the following:
* fast
* independent
* uniformly distributed

## Summary
![image](https://github.com/user-attachments/assets/72efdd27-e1fb-4ed4-95ca-e876f89636d5)

Ref: https://ricardoanderegg.com/posts/bloom-filters-poster/

## Applications

![image](https://github.com/user-attachments/assets/f2fefdbe-d596-4d2a-938d-3a42acfb3378)

For example, one use case of Bloom filters is the following: you have a huge list of malicious URLs. In your browser, before a user navigates to a new URL, you want to check if it’s inside the list of dangerous URLs. You can use a bloom filter to do that! It will take less space than saving the full list of URLs, and if the answer from the Bloom filter is “no” (the URL is not a malicious one), you can safely let the user visit it

![image](https://github.com/user-attachments/assets/8d0cb71b-8ae2-44af-a53d-e0adc8f5c8fa)

Reducing disk lookups for the non-existing keys.
* Log-structured merge tree (LSM) storage engine used in databases such as Cassandra uses a bloom filter to check if the key exists in the SSTable

## Bloom filter disadvantages

The limitations of the bloom filter are the following:
* bloom filter doesn’t support the delete operation
  - Use counting bloom filter, instead of bit set 
* false positives rate can’t be reduced to zero
* bloom filter on disk requires random access due to random indices generated by hash functions

Removing an item from the bloom filter is not supported because there is no possibility to identify the bits that should be cleared. There might be other items that map onto the same bits and clearing any of the bits would introduce the possibility of false negatives.

## Extensions of bloom filter

The counting bloom filter includes an array of counters for each bit in the bloom filter array. The counting bloom filter supports the delete operation.

![image](https://github.com/user-attachments/assets/5345760c-a3ce-4591-81de-969172b69b67)

Counting Bloom filters are much bigger than classical Bloom filters because additional memory has to be allocated for the counters even if most of them are zeros.
