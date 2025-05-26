> Probabilistic Data Structure To Check Membership

<img width="1551" alt="image" src="https://github.com/user-attachments/assets/9fc707e3-0eec-48c5-ab9a-53218714ceff" />

False positives are possible. it might say, element might exist, even thought it isn't.

But when it doesn't exist, it would never say it exists.

Ref to https://systemdesign.one/bloom-filters-explained/ 

## Requirements
Design a data structure with the following characteristics:
* constant time complexity to test membership
* a small amount of memory to test membership
* insert and query operations should be parallelizable
* test membership can yield approximate results


![Uploading image.png…]()



For example, one use case of Bloom filters is the following: you have a huge list of malicious URLs. In your browser, before a user navigates to a new URL, you want to check if it’s inside the list of dangerous URLs. You can use a bloom filter to do that! It will take less space than saving the full list of URLs, and if the answer from the Bloom filter is “no” (the URL is not a malicious one), you can safely let the user visit it


## Summary
![image](https://github.com/user-attachments/assets/72efdd27-e1fb-4ed4-95ca-e876f89636d5)

Ref: https://ricardoanderegg.com/posts/bloom-filters-poster/
