# Daily Coding Problem: Problem #537 [Easy] 

A Collatz sequence in mathematics can be defined as follows. Starting
with any positive integer:

* if n is even, the next number in the sequence is n / 2
* if n is odd, the next number in the sequence is 3n + 1

It is conjectured that every such sequence eventually reaches the number 1.
Test this conjecture.

Bonus: What input n <= 1000000 gives the longest sequence?

I get 524 numbers in the sequence for starting number 837799

## Programs

None of the programs have dependencies outside the standard library.

All of the programs can be easily built from the command line:

    $ go build collatz5.go
    $ ./collatz5 20 > 20.dot
    $ dot -Tpng -o 20.png 20.dot

* [collatz1.go](collatz1.go) - print out a collatz sequence for a single number: `./collatz1 10`
* [collatz2.go](collatz2.go) - print out length of collatz sequences for N numbers: `./collatz2 $N`
* [collatz3.go](collatz3.go) - find maximum length of collatz sequence for numbers up to 1,000,000: `./collatz3`
* [collatz5.go](collatz5.go) - draw a [graphviz](https://graphviz.org/) of Collatz sequences from 1 to N: `./collatz5 $N > collatz.dot`
