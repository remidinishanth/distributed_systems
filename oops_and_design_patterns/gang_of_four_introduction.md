Object-oriented programs are made up of objects. An **object** packages both data and
the procedures that operate on that data. The procedures are typically called methods
or operations. An object performs an operation when it receives a request (or message)
from a client.

Requests are the only way to get an object to execute an operation. Operations are
the only way to change an object's internal data. Because of these restrictions, the
object's internal state is said to be **encapsulated**; it cannot be accessed directly, and its
representation is invisible from outside the object.

An object's implementation is defined by its **class**. The class specifies the object's internal data and
representation and defines the operations the object can perform. 

Objects are created by instantiating a class. The object is said to be an instance of the
class. The process of **instantiating** a class allocates storage for the object's internal data
(made up of instance variables) and associates the operations with these data.

#### Interface
Specifying Object Interfaces

Every operation declared by an object specifies the operation's name, the objects it
takes as parameters, and the operation's return value. This is known as the operation's
signature. The set of all signatures defined by an object's operations is called the
interface to the object. 

An object's interface characterizes the complete set of requests
that can be sent to the object. Any request that matches a signature in the object's
interface may be sent to the object.

#### Dynamic binding
The run-time association of a request to an object and one of its operations is known as
dynamic binding.

#### Inheritance
New classes can be defined in terms of existing classes using class inheritance. When
a subclass inherits from a parent class, it includes the definitions of all the data and
operations that the parent class defines. Objects that are instances of the subclass will
contain all data defined by the subclass and its parent classes, and they'll be able to
perform all operations defined by thissubclass and its parents.

#### Abstract Class
An abstract class is one whose main purpose is to define a common interface for its
subclasses. An abstract class will defer some or all of its implementation to operations
defined in subclasses; hence an abstract class cannot be instantiated. The operations
that an abstract class declares but doesn't implement are called abstract operations.

Classes that aren't abstract are called concrete classes.

### Object Modeling Technique

Our OMT-based notation depicts a class as a rectangle
with the class name in bold. Operations appear in normal type below the class name.
Any data that the class defines comes after the operations. Return typesand instance variable types are optional,since we don't assume a statically
typed implementation language. Lines separate the class name
from the operations and the operations from the data:

<img width="225" alt="image" src="https://github.com/remidinishanth/oops_and_design_patterns/assets/19663316/06f7921d-ffe4-4cc1-bb7c-85cb73c94b7f">

A dashed arrowhead line indicates a class that instantiates objects of another class. The
arrow points to the class of the instantiated objects.

<img width="407" alt="image" src="https://github.com/remidinishanth/oops_and_design_patterns/assets/19663316/d7e5dd5b-65c1-4d28-b1e4-b0fcced52035">

We indicate the subclass relationship with a vertical line and a triangle:

<img width="407" alt="image" src="https://github.com/remidinishanth/oops_and_design_patterns/assets/19663316/84332fdb-7f4c-4780-9470-4c6f6f80c606">
