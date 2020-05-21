package main

import (
	"fmt"
	"time"
)

/*
 * Caching as https://www.mathblog.dk/project-euler-14/ does it.
 * Extremely clever.
 */

const (
	maxstart  = 1000000
	cachesize = 1000001
)

func main() {

	maxIters := -1
	var i, maxItersN int64
	var cache [cachesize]int
	cache[1] = 1

	begin := time.Now()
	for i = 2; i <= maxstart; i++ {
		iterations := 0
		n := i
		for n != 1 && n >= i {
			iterations++
			if (n & 0x01) == 0x01 {
				n = 3*n + 1
			} else {
				n /= 2
			}
		}

		cache[i] = iterations + cache[n]

		if cache[i] > maxIters {
			maxIters = cache[i]
			maxItersN = i
		}
	}
	et := time.Since(begin)

	fmt.Printf("%d requires %d iterations, found in %v\n", maxItersN, maxIters, et)
}
