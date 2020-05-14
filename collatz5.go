package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type CollatzSequence map[int64]int64

func main() {

	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var seq CollatzSequence
	seq = make(map[int64]int64)
	seq.add(1, 0)

	fmt.Println("digraph g {")
	printNode(1)

	for i := 2; i <= max; i++ {
		i := int64(i)
		if !seq.appears(i) {
			collatz(i, seq)
			fmt.Println()
		}
	}

	fmt.Println("}")
}

func collatz(n int64, seq CollatzSequence) {
	for n != 1 {
		if seq.appears(n) {
			break
		}
		printNode(n)
		last := n
		if (n & 0x01) == 0x01 {
			n = 3*n + 1
		} else {
			n /= 2
		}
		if !seq.appears(n) {
			printNode(n)
		}
		printEdge(last, n)
		seq.add(last, n)
	}
}

func (cs CollatzSequence) appears(n int64) bool {
	_, ok := cs[n]
	return ok
}

func (cs CollatzSequence) add(n, m int64) {
	cs[n] = m
}

func printNode(n int64) {
	fmt.Printf("N%d [label=\"%d\"];\n", n, n)
}
func printEdge(from, to int64) {
	fmt.Printf("N%d -> N%d;\n", from, to)
}
