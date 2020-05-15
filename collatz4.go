package main

import (
	"fmt"
	"time"
)

func main() {

	maxIters := -1
	var i, maxItersN int64

	begin := time.Now()
	for i = 1; i <= 1000000; i++ {
		iters := collatz(i)
		if iters > maxIters {
			maxIters = iters
			maxItersN = i
		}
	}
	et := time.Since(begin)

	fmt.Printf("%d requires %d iterations, found in %v\n", maxItersN, maxIters, et)
}

func collatz(n int64) int {
	iterations := 0
	for n != 1 {
		iterations++
		if (n & 0x01) == 0x01 {
			n = 3*n + 1
		} else {
			n /= 2
		}
	}
	return iterations
}
