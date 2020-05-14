package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	prev := make(map[int64]int)

	for i := 1; i < n; i++ {
		iters := collatz(i, prev)
		prev[int64(i)] = iters
		fmt.Printf("%10d %10d\n", i, iters)
	}
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
