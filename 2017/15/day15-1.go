package main

import "fmt"

const startA = int64(699)
const startB = int64(124)
const rounds = 40000000

const seedA = int64(16807)
const seedB = int64(48271)
const factor = int64(2147483647)

func generate(seed, previous int64) int64 {
	return (seed * previous) % factor
}

func main() {
	matches := 0
	a, b := startA, startB
	for i := 0; i < rounds; i++ {
		a = generate(seedA, a)
		b = generate(seedB, b)

		if (a & 0xffff) == (b & 0xffff) {
			matches++
		}

		//fmt.Printf("A: %d\tB: %d\n", a, b)
	}

	fmt.Printf("%d matches\n", matches)
}
