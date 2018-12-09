package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func addMarble(circle *list.List, curMarble *list.Element, marble int) (*list.Element, int) {
	if marble%23 != 0 {
		curMarble = curMarble.Next()
		if curMarble == nil {
			curMarble = circle.Front()
		}
		curMarble = circle.InsertAfter(marble, curMarble)
		return curMarble, 0
	}

	score := marble
	target := curMarble
	for i := 0; i < 7; i++ {
		target = target.Prev()
		if target == nil {
			target = circle.Back()
		}
	}
	curMarble = target.Next()
	score += target.Value.(int)
	circle.Remove(target)

	return curMarble, score
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s players marbles", os.Args[0])
	}

	playerCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid player", err)
	}

	marbleCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid marble count", err)
	}

	scores := make([]int, playerCount)
	circle := list.New()

	curMarble := circle.PushBack(0)
	for i := 0; i < marbleCount; i++ {
		newPos, score := addMarble(circle, curMarble, i+1)
		scores[i%playerCount] += score
		curMarble = newPos
	}

	highScore := 0
	for _, v := range scores {
		if v > highScore {
			highScore = v
		}
	}

	fmt.Printf("Highest score is: %d\n", highScore)
}
