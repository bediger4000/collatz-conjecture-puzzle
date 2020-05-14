package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("digraph g {")

	prev := make(map[int64]int)
	appears := make(map[int64]bool)
	appears[1] = true

	for i := 1; i < max; i++ {
		if !appears[int64(i)] {
			iters := collatz(i, prev, appears)
			prev[int64(i)] = iters
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func collatz(m int, prev map[int64]int, appears map[int64]bool) int {
	iterations := 0
	n := int64(m)
	var last int64
	for n != 1 {
		iterations++
		last = n
		if (n & 0x01) == 0x01 {
			n = 3*n + 1
		} else {
			n /= 2
		}
		already(n, last, appears)
		if i, ok := prev[n]; ok {
			return iterations + i
		}
	}
	return iterations
}

func already(n, last int64, appears map[int64]bool) {
	if !appears[n] {
		fmt.Printf("N%d [label=\"%d\"];\n", n, n)
		fmt.Printf("N%d -> N%d;\n", last, n)
		appears[n] = true
	}
}
