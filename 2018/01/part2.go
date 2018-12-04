package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func runFile(infile string, cur int, seen map[int]bool) (bool, int, map[int]bool) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		offset, _ := strconv.Atoi(scanner.Text())
		cur += offset

		if _, ok := seen[cur]; ok {
			return true, cur, seen
		}

		seen[cur] = true
	}

	return false, cur, seen
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: part1.go input-file")
	}

	found := false
	cur := 0
	passes := 0
	seen := make(map[int]bool, 0)

	for {
		passes++
		found, cur, seen = runFile(os.Args[1], cur, seen)
		if found {
			fmt.Printf("Found a duplicate at pass %d: %d\n", passes, cur)
			break
		}
	}
}
