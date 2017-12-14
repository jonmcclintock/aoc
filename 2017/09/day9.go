package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "day9-in.txt"

type State int

const (
	Normal State = iota
	Garbage
	Escape
)

func calculateScore(line string) (int, int) {
	total := 0
	depth := 0
	state := Normal
	garbage := 0

	for i := 0; i < len(line); i++ {
		switch state {
		case Normal:
			switch line[i] {
			case '{':
				depth++
			case '}':
				total += depth
				depth--
			case '<':
				state = Garbage
			}

		case Garbage:
			switch line[i] {
			case '!':
				state = Escape
			case '>':
				state = Normal
			default:
				garbage++
			}

		case Escape:
			state = Garbage

		}
	}

	return total, garbage
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		score, garbage := calculateScore(scanner.Text())
		fmt.Printf("Input file score: %d\n", score)
		fmt.Printf("Garbage chars: %d\n", garbage)
	}

	tests := map[string][]int{
		"{}":                            []int{1, 0},
		"{{{}}}":                        []int{6, 0},
		"{{},{}}":                       []int{5, 0},
		"{{{},{},{{}}}}":                []int{16, 0},
		"{<a>,<a>,<a>,<a>}":             []int{1, 4},
		"{{<ab>},{<ab>},{<ab>},{<ab>}}": []int{9, 8},
		"{{<!!>},{<!!>},{<!!>},{<!!>}}": []int{9, 0},
		"{{<a!>},{<a!>},{<a!>},{<ab>}}": []int{3, 17},
	}
	for line, expected := range tests {
		score, garbage := calculateScore(line)
		fmt.Printf("Expected score %d, actual %d: %s\n", expected[0], score, line)
		fmt.Printf("-> Expected garbage %d, actual %d\n", expected[1], garbage)
	}
}
