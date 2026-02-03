<img width="1035" alt="image" src="https://github.com/remidinishanth/oops_and_design_patterns/assets/19663316/3424d51b-d71f-4c50-adce-89c213664f38">


### IS-A and HAS-A

Using IS-A and HAS-A:
Remember that when one class inherits from another, we say that the
subclass extends the superclass. 

When you want to know if one thing should extend another, apply the IS-A test. 

Tub extends Bathroom, sounds reasonable. Until you apply the IS-A test.

To know if you’ve designed your types correctly, ask, “Does it make sense to
say type X IS-A type Y?” If it doesn’t, you know there’s something wrong
with the design, so if we apply the IS-A test, Tub IS-A Bathroom is definitely false.

What if we reverse it to Bathroom extends Tub? That still doesn’t work,
Bathroom IS-A Tub doesn’t work. Tub and Bathroom are related, but
not through inheritance. Tub and Bathroom are joined by a HAS-A
relationship. Does it make sense to say “Bathroom HAS-A Tub”? If yes,
then it means that Bathroom has a Tub instance variable. 

In other words, Bathroom has a reference to a Tub, but
Bathroom does not extend Tub and vice-versa.

Bathroom HAS-A Tub and Tub HAS-A Bubbles.
But nobody inherits from (extends) anybody else.

The IS-A test works anywhere in the inheritance tree. If your
inheritance tree is well-designed, the IS-A test should make
sense when you ask any subclass if it IS-A any of its supertypes.

Remember, if X extends Y, X IS-A Y must make sense.

DO NOT use inheritance just so that you can reuse
code from another class, if the relationship between the
superclass and subclass violate either of the above two
rules. For example, imagine you wrote special printing
code in the Alarm class and now you need printing code
in the Piano class, so you have Piano extend Alarm so that
Piano inherits the printing code. That makes no sense! A
Piano is not a more specific type of Alarm. (So the printing
code should be in a Printer class, that all printable objects
can take advantage of via a HAS-A relationship.)

Inheritance IS-A summary:
* A subclass extends a superclass.
* A subclass inherits all public instance
variables and methods of the superclass, but
does not inherit the private instance variables
and methods of the superclass.
* Inherited methods can be overridden; instance
variables cannot be overridden (although they
can be redefined in the subclass, but that’s
not the same thing, and there’s almost never a
need to do it.)
* Use the IS-A test to verify that your
inheritance hierarchy is valid. If X extends Y,
then X IS-A Y must make sense.
* The IS-A relationship works in only one
direction. A Hippo is an Animal, but not all
Animals are Hippos.
* When a method is overridden in a subclass,
and that method is invoked on an instance of
the subclass, the overridden version of the
method is called. (The lowest one wins.)
* If class B extends A, and C extends B, class
B IS-A class A, and class C IS-A class B, and
class C also IS-A class A.
