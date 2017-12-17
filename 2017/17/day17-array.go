package main

import (
	"container/list"
	"fmt"
)

const insertCount = 100000
const stepsPerInsert = 301

func dumpList(l *list.List, cur *list.Element) {
	fmt.Printf("List: ")
	for e := l.Front(); e != nil; e = e.Next() {
		if e == cur {
			fmt.Printf("(%d), ", e.Value)
		} else {
			fmt.Printf("%d, ", e.Value)
		}
	}
	fmt.Printf("\n")
}

func dumpArray(a []int, cur int) {
	fmt.Printf("List: ")
	for i := range a {
		if i == cur {
			fmt.Printf("(%d), ", a[i])
		} else {
			fmt.Printf("%d, ", a[i])
		}
	}
	fmt.Printf("\n")
}

func main() {
	a := []int{0}

	cur := 0
	for i := 0; i < insertCount; i++ {
		//dumpArray(a, cur)
		cur = (cur + stepsPerInsert + 1) % len(a)
		newA := append([]int{}, a[:cur+1]...)
		newA = append(newA, i+1)
		newA = append(newA, a[cur+1:]...)
		a = newA
	}

	//dumpArray(a, cur)
	fmt.Printf("The value after the last insert is: %d\n", a[(cur+2)%len(a)])

	for i := range a {
		if a[i] == 0 {
			fmt.Printf("The value after zero is: %d\n", a[i+1])
			break
		}
	}
}
