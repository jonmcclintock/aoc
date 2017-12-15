package main

import "fmt"

const startA = int64(699)
const startB = int64(124)

const rounds = 5000000

const seedA = int64(16807)
const seedB = int64(48271)
const factor = int64(2147483647)

const criteriaA = int64(4)
const criteriaB = int64(8)

func generate(seed, criteria, previous int64) int64 {
	res := previous
	for {
		res = (seed * res) % factor
		if (res % criteria) == 0 {
			return res
		}
	}
}

func main() {
	matches := 0
	a, b := startA, startB
	for i := 0; i < rounds; i++ {
		a = generate(seedA, criteriaA, a)
		b = generate(seedB, criteriaB, b)

		if (a & 0xffff) == (b & 0xffff) {
			matches++
		}
	}

	fmt.Printf("%d matches\n", matches)
}
