package main

import (
	"fmt"
	"math/bits"
)

const LENGTH = 256
const ROUNDS = 64
const DENSITY = 16
const ROWS = 128

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

func knotHash(input []byte) []int {
	list := makeList(0, LENGTH)

	pos, skip := 0, 0
	for i := 0; i < ROUNDS; i++ {
		pos, skip = hashRound(list, input, pos, skip)
	}

	denseHash := make([]int, LENGTH/DENSITY)
	for i := range denseHash {
		denseHash[i] = condense(list[i*DENSITY : (i+1)*DENSITY])
	}

	return denseHash
}

func bitSet(row []int, bit int) bool {
	if (row[bit/8] & (1 << uint(7-bit%8))) != 0 {
		return true
	}
	return false
}

func spreadColor(matrix, plot [][]int, i, j, color int) {
	if i < 0 || i >= len(plot) || j < 0 || j >= len(plot[i]) {
		return
	}

	if !bitSet(matrix[i], j) || plot[i][j] != 0 {
		return
	}

	plot[i][j] = color
	spreadColor(matrix, plot, i-1, j, color)
	spreadColor(matrix, plot, i+1, j, color)
	spreadColor(matrix, plot, i, j-1, color)
	spreadColor(matrix, plot, i, j+1, color)
}

func main() {
	key := []byte("xlqgujun")

	count := 0
	matrix := make([][]int, ROWS)
	for i := 0; i < ROWS; i++ {
		s := fmt.Sprintf("%s-%d", key, i)
		input := append([]byte(s), 17, 31, 73, 47, 23)

		matrix[i] = knotHash(input)
		for _, v := range matrix[i] {
			count += bits.OnesCount(uint(v))
		}
	}
	fmt.Printf("Full squares: %d\n", count)

	count = 0
	plot := make([][]int, ROWS)
	for i := 0; i < ROWS; i++ {
		plot[i] = make([]int, len(matrix[i])*8)
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < len(matrix[i])*8; j++ {
			if !bitSet(matrix[i], j) || plot[i][j] != 0 {
				continue
			}

			count++
			spreadColor(matrix, plot, i, j, count)
		}
	}

	fmt.Printf("Number of groups: %d\n", count)
}
