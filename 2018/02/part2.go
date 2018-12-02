package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func addWord(word string, seen map[string]bool) (bool, string) {
	fmt.Printf("Checking '%s (%d)'\n", word, len(word))

	subWords := make([]string, len(word))
	for i := 0; i < len(word); i++ {
		subWord := word[:i] + word[i+1:]
		if _, ok := seen[subWord]; ok {
			return true, subWord
		}
		subWords[i] = subWord
	}

	for _, v := range subWords {
		seen[v] = true
	}

	return false, ""
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

	seen := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()

		if found, common := addWord(word, seen); found {
			fmt.Printf("Found a common base: %s (%d)\n", common, len(common))
			break
		}
	}
}
