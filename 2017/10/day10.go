package main

import "fmt"

const LENGTH = 256
const ROUNDS = 64
const DENSITY = 16

func makeList(start, end int) []int {
	list := make([]int, end-start)
	for i := range list {
		list[i] = i + start
	}

	return list
}

func offset(length, start, count int) int {
	return (start + count) % length
}

func reverse(list []int, start, count int) {
	for i := 0; i < (count / 2); i++ {
		list[offset(len(list), start, i)], list[offset(len(list), start, count-i-1)] =
			list[offset(len(list), start, count-i-1)], list[offset(len(list), start, i)]
	}
}

func hashRound(list []int, input []byte, pos, skip int) (int, int) {
	for _, len := range input {
		reverse(list, pos, int(len))
		pos += int(len) + skip
		skip++
	}
	return pos, skip
}

func condense(list []int) int {
	res := 0
	for _, v := range list {
		res ^= v
	}
	return res
}

func main() {
	input := []byte("187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216")
	input = append(input, 17, 31, 73, 47, 23)

	list := makeList(0, LENGTH)

	pos, skip := 0, 0
	for i := 0; i < ROUNDS; i++ {
		pos, skip = hashRound(list, input, pos, skip)
	}

	denseHash := make([]int, LENGTH/DENSITY)
	for i := range denseHash {
		denseHash[i] = condense(list[i*DENSITY : (i+1)*DENSITY])
	}

	fmt.Printf("Dense hash: ")
	for _, v := range denseHash {
		fmt.Printf("%02x", v)
	}
	fmt.Printf("\n")
}
