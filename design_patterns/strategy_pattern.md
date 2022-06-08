![image](https://user-images.githubusercontent.com/19663316/172449958-8975461b-ba18-4d7f-b74c-9b327c51e10c.png)

## Motivation

**Problem:**   Started with a simple SimUDuck App.

* Joe's company makes a duck pond simulation game, SimUDuck, The game can show a large variety of duck species swimming and making quacking sounds.

![image](https://user-images.githubusercontent.com/19663316/172450506-9876578a-8e87-4057-8782-90c675663693.png)

### But now we need the ducks to FLY

![image](https://user-images.githubusercontent.com/19663316/172450546-547310c9-799b-46f2-95cf-b033bdf6838a.png) 

### First Design:
* Joe add a method `fly()` into the Duck class. It seem to work, but something went horribly wrong because not all ducks can fly.

![image](https://user-images.githubusercontent.com/19663316/172450753-5913b334-63ab-4643-b57e-b8023e983cc8.png)

What we got, Rubber ducks are also able to fly

![image](https://user-images.githubusercontent.com/19663316/172679433-3659275b-b1d9-4af4-b0ec-f0917cfa43e6.png)

![image](https://user-images.githubusercontent.com/19663316/172450865-4f1f11ec-bad9-4619-872a-2fbcac3364ef.png)


**Drawbacks**
* A localized update to the code caused a nonlocal side effect (flying rubber ducks)!


![image](https://user-images.githubusercontent.com/19663316/172450969-29b31625-7a7e-4281-a648-98fd1bd386d3.png)

### Second Design:
Using inheritance and polymorphism
* Always override the `fly()` mehtod in the subclass where needed.

![image](https://user-images.githubusercontent.com/19663316/172451061-a2b167ff-900f-4d61-865f-61f085f102ec.png)

**Drawbacks:** 
* Everytime a new duck is added, you will be forced to look at and possibly override `fly()` and `quack()`. 
* So is there a cleaner way of having only some of the duck types fly or quack?

### Third Idea

How about an interface?

![image](https://user-images.githubusercontent.com/19663316/172453211-4fc09800-05db-4e9a-a1d0-0d71683532ac.png)

![image](https://user-images.githubusercontent.com/19663316/172453303-7a1f4009-c5e8-4021-80f1-5c2e92399593.png)

**Drawbacks:** It completely destroy code reuse for those behaviors.

![image](https://user-images.githubusercontent.com/19663316/172453241-6f55e5e6-8bbb-4222-a407-3ebda5e17752.png)


### Solution

1. Design Principles: **Identify the aspects of your application that vary and separate them form what stays the same!**
                               which means Encapsulate the parts that vary!
2. Design Principles: **Program to an interface, not an implementation!** (interface here means supertype! including interface
                               and abstract class!.. making use of the polymorphism functionality).
3. Design Principles:  **Favor composition over interface! **


* We know that not all of the subclasses should have flying or quacking behavior, so inheritance isn’t the right answer. 
* But while having the subclasses implement Flyable and/or Quackable solves part of the problem (no inappropriately flying rubber ducks), it completely destroys code reuse for those behaviors, so it just creates a different maintenance nightmare.
* And of course there might be more than one kind of flying behavior even among the ducks that do fly...


### Strategy Pattern

![image](https://user-images.githubusercontent.com/19663316/172451086-d55c9b63-8bd5-4eee-9ab5-dd31d777f608.png)

![image](https://user-images.githubusercontent.com/19663316/172451458-80e8e00a-6194-4109-b388-7c18b3181a66.png)

![image](https://user-images.githubusercontent.com/19663316/172451633-410c2335-065b-4cfd-bd67-7a6004bf2ceb.png)


### Ref: 
* https://www.oreilly.com/library/view/head-first-design/0596007124/ch01.html
* https://www.freecodecamp.org/news/the-strategy-pattern-explained-using-java-bc30542204e0/
