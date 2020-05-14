package main

import "fmt"

func main() {

	prev := make(map[int64]int)

	for i := 1; i < 1000000; i++ {
		iters := collatz(i, prev)
		prev[int64(i)] = iters
	}

	maxIters := -1
	var maxItersN int64
	for n, iters := range prev {
		if iters > maxIters {
			maxIters = iters
			maxItersN = n
		}
	}

	fmt.Printf("%d requires %d iterations\n", maxItersN, maxIters)
}

func collatz(m int, prev map[int64]int) int {
	iterations := 0
	n := int64(m)
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
