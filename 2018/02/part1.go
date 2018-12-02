package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkForLetterCount(word string, target int) bool {
	counts := make(map[rune]int)
	for _, char := range word {
		counts[char]++
	}

	for _, v := range counts {
		if v == target {
			return true
		}
	}

	return false
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatal(fmt.Errorf("Usage: part1.go input-file"))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count2 := 0
	count3 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()

		if checkForLetterCount(word, 2) {
			count2++
		}

		if checkForLetterCount(word, 3) {
			count3++
		}
	}

	fmt.Printf("Counts: %d 2-letter, %d 3-letter\n", count2, count3)
	fmt.Printf("Product: %d\n", count2*count3)
}
