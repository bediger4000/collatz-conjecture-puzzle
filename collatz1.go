package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	iters := collatz(n)
	fmt.Printf("%d requires %d iterations to reach 1\n", n, iters)
}

func collatz(n int64) int {
	iterations := 0
	for n != 1 {
		fmt.Printf("%5d %d\n", iterations, n)
		iterations++
		if (n & 0x01) == 0x01 {
			n = 3*n + 1
		} else {
			n /= 2
		}
	}
	fmt.Printf("%5d %d\n", iterations, n)
	return iterations
}
