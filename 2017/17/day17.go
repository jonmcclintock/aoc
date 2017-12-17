package main

import (
	"container/list"
	"fmt"
)

const insertCount = 50000000
const stepsPerInsert = 301

func dump(l *list.List, cur *list.Element) {
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

func main() {
	l := list.New()
	l.PushBack(0)

	cur := l.Front()
	for i := 0; i < insertCount; i++ {
		//dump(l, cur)
		for j := 0; j < stepsPerInsert; j++ {
			cur = cur.Next()
			if cur == nil {
				cur = l.Front()
			}
		}

		cur = l.InsertAfter(i+1, cur)
	}

	last := cur.Next()
	if last == nil {
		last = l.Front()
	}
	fmt.Printf("The value after the last insert is: %d\n", last.Value)

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 0 {
			fmt.Printf("The value after zero is: %d\n", e.Next().Value)
			break
		}
	}
}
