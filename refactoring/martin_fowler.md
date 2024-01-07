# Refactoring
Book by Kent Beck and Martin Fowler

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
