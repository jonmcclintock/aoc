package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stateIndices = map[string]int{
	".....": 0x0,
	"....#": 0x1,
	"...#.": 0x2,
	"...##": 0x3,
	"..#..": 0x4,
	"..#.#": 0x5,
	"..##.": 0x6,
	"..###": 0x7,
	".#...": 0x8,
	".#..#": 0x9,
	".#.#.": 0xa,
	".#.##": 0xb,
	".##..": 0xc,
	".##.#": 0xd,
	".###.": 0xe,
	".####": 0xf,
	"#....": 0x10,
	"#...#": 0x11,
	"#..#.": 0x12,
	"#..##": 0x13,
	"#.#..": 0x14,
	"#.#.#": 0x15,
	"#.##.": 0x16,
	"#.###": 0x17,
	"##...": 0x18,
	"##..#": 0x19,
	"##.#.": 0x1a,
	"##.##": 0x1b,
	"###..": 0x1c,
	"###.#": 0x1d,
	"####.": 0x1e,
	"#####": 0x1f,
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s input-file cycles", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cycles, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid cycle count", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	curState := ".." + strings.TrimPrefix(scanner.Text(), "initial state: ") + ".."
	scanner.Scan()
	scanner.Text()

	rules := make(map[int]bool, 32)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " => ")
		if len(input) != 2 {
			fmt.Printf("Got invalid input: %v\n", input)
		}

		pattern := stateIndices[input[0]]
		result := false
		if input[1] == "#" {
			result = true
		}
		rules[pattern] = result
	}

	fmt.Printf("Rules: %v\n", rules)

	zero := 2
	for i := 0; i < cycles; i++ {
		fmt.Printf("Round %d: %s  %s\n", i, curState[0:zero], curState[zero:])
		newState := make([]byte, len(curState))

		for j := 0; j < len(curState); j++ {
			s := make([]byte, 5)
			for k := 0; k < 5; k++ {
				if j-2+k < 0 {
					s[k] = '.'
				} else if j-2+k >= len(curState) {
					s[k] = '.'
				} else {
					s[k] = curState[j-2+k]
				}
			}
			if rules[stateIndices[string(s)]] {
				newState[j] = '#'
			} else {
				newState[j] = '.'
			}
		}

		if newState[1] == '#' {
			newState = append([]byte{'.'}, newState...)
			zero++
		}
		if newState[1] == '#' {
			newState = append([]byte{'.'}, newState...)
			zero++
		}

		if newState[len(newState)-2] == '#' {
			newState = append(newState, '.')
		}
		if newState[len(newState)-2] == '#' {
			newState = append(newState, '.')
		}

		curState = string(newState)
	}

	sum := 0
	for i := 0; i < len(curState); i++ {
		if curState[i] == '#' {
			sum += i - zero
		}
	}

	fmt.Printf("After %d rounds, sum is %d\n", cycles, sum)
}
