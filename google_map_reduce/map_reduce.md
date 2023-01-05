### MapReduce

A framework for large-scale parallel processing

> Goal: Create a distributed computing framework to process data on a massive scale.

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
