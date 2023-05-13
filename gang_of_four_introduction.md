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

**Interface**
Specifying Object Interfaces

Every operation declared by an object specifies the operation's name, the objects it
takes as parameters, and the operation's return value. This is known as the operation's
signature. The set of all signatures defined by an object's operations is called the
interface to the object. 

An object's interface characterizes the complete set of requests
that can be sent to the object. Any request that matches a signature in the object's
interface may be sent to the object.

**Dynamic binding**:
The run-time association of a request to an object and one of its operations is known as
dynamic binding.
