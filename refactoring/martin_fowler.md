# Refactoring
Book by Kent Beck and Martin Fowler

The term “refactoring” can be used either as a noun or a verb. 
* The noun’s definition is: Refactoring (noun): a change made to the internal structure of software to make it
easier to understand and cheaper to modify without changing its observable behavior.
* The verb’s definition is: Refactoring (verb): to restructure software by applying a series of refactorings
without changing its observable behavior.

## Chapter 1

* Breaking down complicated chunks into small pieces is important, as is
naming things well.
  -  Naming is both important and tricky. Breaking a large function into smaller ones only
adds value if the names are good. With good names, I don’t have to read the body of the
function to see what it does.

* Here are the four steps, each followed by compiling, testing, and committing to local source code repository:
  - **Split Loop** to isolate the accumulation
  - **Slide Statements** to bring the initializing code next to the accumulation
  - **Extract Function** to create a function for calculating the total
  - **Inline Variable** to remove the variable completely
 
* Prefer to treat data as immutable as much as I can—mutable state quickly becomes something rotten.

* SPLITTING THE PHASES OF CALCULATION AND FORMATTING

* Replace Conditional with Polymorphism

> There were three major stages to this refactoring episode: decomposing the original
function into a set of nested functions, using Split Phase (154) to separate the
calculation and printing code, and finally introducing a polymorphic calculator for the
calculation logic. Each of these added structure to the code, enabling me to better
communicate what the code was doing.

* Code should be obvious: When someone needs to make a change, they
should be able to find the code to be changed easily and to make the change quickly
without introducing any errors.

* The key to effective refactoring is recognizing that you go
faster when you take tiny steps, the code is never broken, and you can compose those
small steps into substantial changes. Remember that—and the rest is silence.

### Bulb moments

* When you have to add a feature to a program but the code is not structured in
a convenient way, first refactor the program to make it easy to add the feature, then
add the feature.
* Before you start refactoring, make sure you have a solid suite of tests. These
tests must be self­checking.
* Refactoring changes the programs in small steps, so if you make a mistake, it
is easy to find where the bug is.
* Any fool can write code that a computer can understand. Good programmers
write code that humans can understand.
* When programming, follow the camping rule: Always leave the code base healthier than when you found it.
* The true test of good code is how easy it is to change it.

## Chapter 2

* The two hats: Adding functionality and Refactoring
  - Adding functionality - doesn't change existing code, adding new capabilities and tests.
  - Refactoring - not adding functionality, only restructing code, don't need to add new tests(unless some cases are missing). May be change tests to accomodate interface
