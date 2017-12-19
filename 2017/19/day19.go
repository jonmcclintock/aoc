package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "day19-in.txt"

func loadDiagram() [][]byte {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	diagram := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		if len(row) > 0 {
			diagram = append(diagram, row)
		}
	}

	return diagram
}

func nextStep(diagram [][]byte, x, y int, dir byte) (int, int, byte) {
	switch dir {
	case 'u':
		if diagram[y][x] == '+' {
			if diagram[y][x+1] != ' ' {
				return x + 1, y, 'r'
			}
			return x - 1, y, 'l'
		}
		if y == 0 {
			return x, y, '!'
		}
		return x, y - 1, 'u'

	case 'd':
		if diagram[y][x] == '+' {
			if diagram[y][x+1] != ' ' {
				return x + 1, y, 'r'
			}
			return x - 1, y, 'l'
		}
		if y >= (len(diagram) - 1) {
			return x, y, '!'
		}
		return x, y + 1, 'd'

	case 'l':
		if diagram[y][x] == '+' {
			if diagram[y+1][x] != ' ' {
				return x, y + 1, 'd'
			}
			return x, y - 1, 'u'
		}
		if x == 0 {
			return x, y, '!'
		}
		return x - 1, y, 'l'

	case 'r':
		if diagram[y][x] == '+' {
			if y < (len(diagram)-1) && diagram[y+1][x] != ' ' {
				return x, y + 1, 'd'
			}
			return x, y - 1, 'u'
		}
		if x >= (len(diagram[y]) - 1) {
			return x, y, '!'
		}
		return x + 1, y, 'r'
	}

	return x, y, '?'
}

func main() {
	diagram := loadDiagram()

	x, y := 0, 0
	for i := range diagram[0] {
		if diagram[0][i] == '|' {
			x = i
			break
		}
	}

	path := []byte{}
	dir := byte('d')
	count := 0
	fmt.Printf("Starting at (%d, %d: '%c') going '%c'\n", x, y, diagram[y][x], dir)
	for {
		x, y, dir = nextStep(diagram, x, y, dir)
		fmt.Printf("Stepped to (%d, %d: '%c') going '%c'\n", x, y, diagram[y][x], dir)
		count++
		if dir == '!' {
			fmt.Printf("Ran off the map at (%d, %d), stopping\n", x, y)
			break
		}

		if diagram[y][x] == ' ' {
			fmt.Printf("Ran out of path at (%d, %d), stopping\n", x, y)
			break
		}
		if diagram[y][x] >= 'A' && diagram[y][x] <= 'Z' {
			path = append(path, diagram[y][x])
		}

	}

	fmt.Printf("Took %d steps\n", count)
	fmt.Printf("Path is: %s\n", string(path))
}
