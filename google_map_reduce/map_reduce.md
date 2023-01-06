### MapReduce

A framework for large-scale parallel processing.

> Goal: Create a distributed computing framework to process data on a massive scale.

MapReduce is a software framework for processing (large) data sets in a distributed fashion over a several machines.

MapReduce framework/library allows programmers without any experience with parallel and distributed systems to easily utilize the resources of a large distributed system.

### Motivation

* Inverted index is storing a mapping from content, such as words or numbers, to its documents where the word is present on the web. Indexing is the process by which search engines organize information before a search to enable super-fast responses to queries.
* To index `20+` billion web pages, assuming each is of `20KB` size, we need to process `400+` terabytes of data.
  - `20+` billion web pages `x 20KB` = `400+` terabytes
  - One computer can read `30-35 MB/sec` from disk, so it takes four months to read the web
  - `~1,000` hard drives just to store the web
  - Good news: same problem with `1000` machines, `< 3` hours
* Most such computations are conceptually straightforward. However, the input data is usually large and the computations have to be distributed across hundreds or thousands of machines in order to finish in a reasonable amount of time.
* But to solve the problem on `1000` machines, we will need to write programs to handle
  - communication and coordination: parallelize the computation and distribute the data
  - handle failures and recovering from machine failure (all the time!)
  - status reporting, debugging, optimization and locality
* Similar difficulty repeats for every problem Google wants to solve
* As a reaction to this complexity, Google designed an abstraction that allows us to express the simple computations we were trying to perform but hides the messy details in MapReduce runtime library:
  - automatic parallelization
  - load balancing
  - data distribution: network and disk transfer optimization
  - fault tolerance: handling of machine failures and robustness
  - improvements to core library benefit all users of library!

### MapReduce Etymology

* MapReduce was created at Google in 2004 by Jeffrey Dean and Sanjay Ghemawat. 
* The name is inspired from map and reduce functions in the LISP programming language. 
* In LISP, the map function takes as parameters a function and a set of values. That function is then applied to each of the values.
  - length function to each item `(map ‘length ‘(() (a) (ab) (abc)))` to `(0 1 2 3)` 
* The reduce function is given a binary function and a set of values as parameters. It combines all the values together using the binary function.
  - add function in reduce `(reduce #'+ '(0 1 2 3))` to `6`


### Programming Model

* The computation takes a set of **input** key/value pairs, and produces a set of **output** key/value pairs. 
* The user of the MapReduce library expresses the computation as two functions: **Map** and **Reduce**.
* Map, written by the user, takes an input pair and produces a set of **intermediate** key/value pairs.
* The MapReduce library **groups together all intermediate values associated with the same intermediate key** I and passes them
to the Reduce function.
* The Reduce function, also written by the user, accepts an **intermediate key I** and a **set of values for that key**. It
**merges together these values** to form a possibly smaller set of values. 


### Word-count example

<img width="1154" alt="image" src="https://user-images.githubusercontent.com/19663316/210795017-6205fe34-f237-4151-904c-31dec4b9684f.png">

### Execution Overview

<img width="890" alt="image" src="https://user-images.githubusercontent.com/19663316/210792948-4460abf7-4fc5-4db4-ade5-0f96100ab517.png">

### Typical problem solved by MapReduce
* Read a lot of data
* **Map**: extract something you care about from each record
* Shuffle and Sort
* **Reduce**: aggregate, summarize, filter, or transform
* Write the results

Outline stays the same, **Map** and **Reduce** functions change to fit the problem
