package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatal(fmt.Errorf("Usage: part1.go input-file"))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cur_freq := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		offset, _ := strconv.Atoi(scanner.Text())
		cur_freq += offset
	}

	fmt.Printf("End frequency: %d\n", cur_freq)
}
