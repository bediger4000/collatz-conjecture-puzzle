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

	appears := make(map[int64]int64)
	appears[1] = 0
	fmt.Printf("N1 [label=\"1\"];\n")

	for i := 1; i < max; i++ {
		fmt.Printf("/* i %d */\n", i)
		if _, ok := appears[int64(i)]; !ok {
			fmt.Printf("/* iterating %d */\n", i)
			collatz(i, appears)
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func collatz(m int, appears map[int64]int64) {
	n := int64(m)
	for n != 1 {
		last := n
		if (n & 0x01) == 0x01 {
			n = 3*n + 1
		} else {
			n /= 2
		}
		already(n, last, appears)
	}
}

func already(n, last int64, appears map[int64]int64) {
	if _, ok := appears[n]; !ok {
		fmt.Printf("N%d [label=\"%d\"];\n", n, n)
		fmt.Printf("N%d -> N%d;\n", last, n)
		fmt.Printf("/* Inserting appears[%d] = %d */\n", n, last)
		appears[n] = last
	}
}
