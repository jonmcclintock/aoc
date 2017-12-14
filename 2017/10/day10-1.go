package main

import "fmt"

const LENGTH = 256

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

func main() {
	input := []int{
		187, 254, 0, 81, 169, 219, 1, 190, 19, 102, 255, 56, 46, 32, 2, 216,
		//3, 4, 1, 5,
	}

	list := makeList(0, LENGTH)

	pos := 0
	skip := 0
	for _, len := range input {
		reverse(list, pos, len)
		pos += len + skip
		skip++
	}

	fmt.Printf("List:\n%v\n", list)
}
