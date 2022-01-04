# oops_and_design_patterns

REF: http://people.eecs.berkeley.edu/~jrs/61b/

THE LANGUAGE OF OBJECT-ORIENTED PROGRAMMING
===========================================
* Object: An object is a repository of data. For example, if MyList is a
 ShoppingList object, MyList might record your shopping list.
* Class: A class is a type of object. A class is a template for creating objects. 
  Many objects of the same class might exist; 
  for instance, MyList and YourList may both be ShoppingList objects.
* Method: A procedure or function that operates on an object or a class.
 A method is associated with a particular class. For instance, addItem might
 be a method that adds an item to any ShoppingList object. Sometimes a method
 is associated with a family of classes. For instance, addItem might operate
 on any List, of which a ShoppingList is just one type.
* Inheritance: A class may inherit properties from a more general class. For
 example, the ShoppingList class inherits from the List class the property of
 storing a sequence of items.
* Polymorphism: The ability to have one method call work on several different
 classes of objects, even if those classes need different implementations of
 the method call. For example, one line of code might be able to call the
 "addItem" method on _every_ kind of List, even though adding an item to a
 ShoppingList is completely different from adding an item to a ShoppingCart.
* Object-Oriented: Each object knows its own class and which methods manipulate
 objects in that class. Each ShoppingList and each ShoppingCart knows which
 implementation of addItem applies to it.

In this list, the one thing that truly distinguishes object-oriented languages
from procedural languages (C, Fortran, Basic, Pascal) is polymorphism.


### The four pillars of object-oriented programming are:

* Abstraction
* Encapsulation
* Inheritance
* Polymorphism
