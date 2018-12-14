package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s input-file elves target", os.Args[0])
	}

	elfCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid elf count")
	}

	targetStr := os.Args[2]
	target := make([]byte, len(targetStr))
	for i := 0; i < len(targetStr); i++ {
		target[i] = targetStr[i] - '0'
	}

	elves := make([]int, elfCount)
	for i := 0; i < elfCount; i++ {
		elves[i] = i
	}

	recipes := []byte{3, 7}

loop:
	for {
		sum := 0
		for _, pos := range elves {
			sum += int(recipes[pos])
		}

		sumStr := strconv.Itoa(sum)
		for j := 0; j < len(sumStr); j++ {
			recipes = append(recipes, sumStr[j]-'0')

			if len(recipes) >= len(target) && bytes.Compare(target, recipes[len(recipes)-len(target):]) == 0 {
				break loop
			}
		}

		for elf, pos := range elves {
			elves[elf] = (elves[elf] + int(recipes[pos]) + 1) % len(recipes)
		}
	}

	fmt.Printf("%+v first appears after %d recipes\n", target, len(recipes)-len(target))

}
