![image](https://user-images.githubusercontent.com/19663316/172449958-8975461b-ba18-4d7f-b74c-9b327c51e10c.png)

## Motivation

![image](https://user-images.githubusercontent.com/19663316/172450506-9876578a-8e87-4057-8782-90c675663693.png)

### But now we need the ducks to FLY

![image](https://user-images.githubusercontent.com/19663316/172450546-547310c9-799b-46f2-95cf-b033bdf6838a.png) 

![image](https://user-images.githubusercontent.com/19663316/172450753-5913b334-63ab-4643-b57e-b8023e983cc8.png)

![image](https://user-images.githubusercontent.com/19663316/172450865-4f1f11ec-bad9-4619-872a-2fbcac3364ef.png)


#### Problem

![image](https://user-images.githubusercontent.com/19663316/172450969-29b31625-7a7e-4281-a648-98fd1bd386d3.png)

![image](https://user-images.githubusercontent.com/19663316/172451061-a2b167ff-900f-4d61-865f-61f085f102ec.png)

![image](https://user-images.githubusercontent.com/19663316/172453211-4fc09800-05db-4e9a-a1d0-0d71683532ac.png)

![image](https://user-images.githubusercontent.com/19663316/172453303-7a1f4009-c5e8-4021-80f1-5c2e92399593.png)

![image](https://user-images.githubusercontent.com/19663316/172453241-6f55e5e6-8bbb-4222-a407-3ebda5e17752.png)

* We know that not all of the subclasses should have flying or quacking behavior, so inheritance isn’t the right answer. 
* But while having the subclasses implement Flyable and/or Quackable solves part of the problem (no inappropriately flying rubber ducks), it completely destroys code reuse for those behaviors, so it just creates a different maintenance nightmare.
* And of course there might be more than one kind of flying behavior even among the ducks that do fly...


### Solution

![image](https://user-images.githubusercontent.com/19663316/172451086-d55c9b63-8bd5-4eee-9ab5-dd31d777f608.png)

![image](https://user-images.githubusercontent.com/19663316/172451458-80e8e00a-6194-4109-b388-7c18b3181a66.png)

![image](https://user-images.githubusercontent.com/19663316/172451633-410c2335-065b-4cfd-bd67-7a6004bf2ceb.png)


### Ref: 
* https://www.oreilly.com/library/view/head-first-design/0596007124/ch01.html
* https://www.freecodecamp.org/news/the-strategy-pattern-explained-using-java-bc30542204e0/
