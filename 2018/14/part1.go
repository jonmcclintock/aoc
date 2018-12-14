package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func makeBatch(elves []int, recipes []int) ([]int, []int) {
	sum := 0
	for _, pos := range elves {
		sum += recipes[pos]
	}

	sumStr := strconv.Itoa(sum)
	for j := 0; j < len(sumStr); j++ {
		recipes = append(recipes, int(sumStr[j]-'0'))
	}

	for elf, pos := range elves {
		elves[elf] = (elves[elf] + recipes[pos] + 1) % len(recipes)
	}

	return elves, recipes
}

func main() {
	if len(os.Args[1:]) != 3 {
		log.Fatalf("Usage: %s input-file elves rounds predictions", os.Args[0])
	}

	elfCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid elf count")
	}

	rounds, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid round count")
	}

	predictions, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid prediction count")
	}

	elves := make([]int, elfCount)
	for i := 0; i < elfCount; i++ {
		elves[i] = i
	}

	recipes := []int{3, 7}

	for len(recipes) < rounds {
		elves, recipes = makeBatch(elves, recipes)
	}

	fmt.Printf("After %d rounds, made %d recipes\n", rounds, len(recipes))

	for len(recipes) < rounds+predictions {
		elves, recipes = makeBatch(elves, recipes)
	}

	fmt.Printf("Predictions: ")
	for i := 0; i < predictions; i++ {
		fmt.Printf("%d", recipes[rounds+i])
	}
	fmt.Printf("\n")

}
