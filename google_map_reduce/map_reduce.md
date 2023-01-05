### MapReduce

A framework for large-scale parallel processing

> Goal: Create a distributed computing framework to process data on a massive scale.


### MapReduce Etymology

* MapReduce was created at Google in 2004 by Jeffrey Dean and Sanjay Ghemawat. 
* The name is inspired from map and reduce functions in the LISP programming language. 
* In LISP, the map function takes as parameters a function and a set of values. That function is then applied to each of the values.
  - length function to each item `(map ‘length ‘(() (a) (ab) (abc)))` to `(0 1 2 3)` 
* The reduce function is given a binary function and a set of values as parameters. It combines all the values together using the binary function.
  - add function in reduce `(reduce #'+ '(0 1 2 3))` to `6`


### MapReduce Library

A simple programming model that applies to many large-scale
computing problems

Hide messy details in MapReduce runtime library:
* automatic parallelization
* load balancing
* network and disk transfer optimization
* handling of machine failures
* robustness
* improvements to core library benefit all users of library!

### Typical problem solved by MapReduce
* Read a lot of data
* Map: extract something you care about from each record
* Shuffle and Sort
* Reduce: aggregate, summarize, filter, or transform
* Write the results

Outline stays the same, Map and Reduce change to fit the problem
