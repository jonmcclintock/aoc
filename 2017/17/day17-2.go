package main

import "fmt"

const insertCount = 50000000
const stepsPerInsert = 301

func main() {

	cur := 0
	afterZero := 0
	for i := 1; i <= insertCount; i++ {
		cur = (cur + stepsPerInsert) % i
		if cur == 0 {
			afterZero = i
		}
		cur = cur + 1
	}

	fmt.Printf("After zero is %d\n", afterZero)
}
