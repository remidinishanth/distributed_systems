### IS-A and HAS-A

Using IS-A and HAS-A:
Remember that when one class inherits from another, we say that the
subclass extends the superclass. When you want to know if one thing should
extend another, apply the IS-A test. 

Tub extends Bathroom, sounds reasonable.

Until you apply the IS-A test.

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
