# Daily Coding Problem: Problem #537 [Easy] 

A Collatz sequence in mathematics can be defined as follows. Starting
with any positive integer:

* if n is even, the next number in the sequence is n / 2
* if n is odd, the next number in the sequence is 3n + 1

It is conjectured that every such sequence eventually reaches the number 1.
Test this conjecture.

Bonus: What input n <= 1000000 gives the longest sequence?

Program `collatz3` solves this bonus question.

I get 524 numbers in the sequence for starting number 837799

    837799 requires 524 iterations, found in 614.514707ms

My iterations may be 1 less than number of terms in the sequence.
[This fellow](https://www.mathblog.dk/project-euler-14/) got 525
elements in the sequence for 837799.
He also has a clever memoization using an array.

## Programs

None of the programs have dependencies outside the standard library.

All of the programs can be easily built from the command line:

    $ go build collatz5.go
    $ ./collatz5 20 > 20.dot
    $ dot -Tpng -o 20.png 20.dot

* [collatz1.go](collatz1.go) - print out a collatz sequence for a single number: `./collatz1 10`
* [collatz2.go](collatz2.go) - print out length of collatz sequences for N numbers: `./collatz2 $N`
* [collatz3.go](collatz3.go) - find maximum length of collatz sequence for numbers up to 1,000,000: `./collatz3`. Memoized algorithm.
* [collatz4.go](collatz4.go) - find maximum length of collatz sequence for numbers up to 1,000,000: `./collatz4`. Straightforward algorithm.
* [collatz5.go](collatz5.go) - draw a [graphviz](https://graphviz.org/) of Collatz sequences from 1 to N: `./collatz5 $N > collatz.dot`

### Solution Design

I wrote a iterative algorithm with memoization:
if the algorithm has encountered a number before,
it can just return the number of iterations for the previously encountered number,
plus the number of iterations that have taken place already.

```go
func collatz(n int64, prev map[int64]int) int {
    iterations := 0
    for n != 1 {
        if i, ok := prev[n]; ok {
            return iterations + i
        }
        iterations++
        if (n & 0x01) == 0x01 {
            n = 3*n + 1
        } else {
            n /= 2
        }
    }
    return iterations
}
```

There's two places where memoization can occur.
First, using memos to not even invoke `func collatz`,
second inside the function, to stop iteration early.

If the Collatz Conjecture (that every input number's sequence
ends up encountering 1) is false, my memoization will break.

It turns out that at least on my laptop, in Go, the memoized code
is substantially slower:

    1201 % ./collatz3  # memoized
    837799 requires 524 iterations, found in 667.24825ms
    1202 % ./collatz4  # not memoized
    837799 requires 524 iterations, found in 279.643536ms

Maybe premature optimization is a bad idea after all.
Or maybe using a Go map is not the right approach.
It could be that a dedicated hash table somewhat optimized
for this problem would be faster.
Using a plain array would require a large array.
I tried `[50000000]int` Go arrays and still encountered sequence elements bigger than that.
Possibly some kind of partial memoization could be triggered up using arrays to store previous
sequence lengths, but it doesn't seem worth it.

The GaphViz diagramming version uses memoization slightly
differently, to avoid outputting multiple edges,
and multiple `node` directives.
Multiple edges in particular clutter up GraphViz images.
