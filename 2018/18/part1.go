package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const LumberYard = '#'
const OpenGround = '.'
const Woods = '|'

type Point struct {
	X, Y int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func loadLandscape(inputFile string) [][]byte {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	landscape := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		landscape = append(landscape, []byte(line))
	}

	return landscape
}

func dumpLandscape(landscape [][]byte) {
	for y := 0; y < len(landscape); y++ {
		fmt.Printf("%4d: %s\n", y, string(landscape[y]))
	}
}

func countTiles(scan [][]byte, t byte) int {
	count := 0
	for y := 0; y < len(scan); y++ {
		for x := 0; x < len(scan[y]); x++ {
			if scan[y][x] == t {
				count++
			}
		}
	}
	return count
}

func countNeighbors(landscape [][]byte, x, y int, t byte) int {
	count := 0

	for dy := -1; dy <= 1; dy++ {
		if y+dy < 0 || y+dy >= len(landscape) {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			if x+dx < 0 || x+dx >= len(landscape[y+dy]) {
				continue
			}
			if dx == 0 && dy == 0 {
				continue
			}
			if landscape[y+dy][x+dx] == t {
				count++
			}
		}
	}

	return count
}

func clockTick(current [][]byte) [][]byte {
	next := make([][]byte, len(current))

	for y := 0; y < len(current); y++ {
		next[y] = make([]byte, len(current[y]))
		for x := 0; x < len(current[y]); x++ {
			switch current[y][x] {
			case OpenGround:
				// An open acre will become filled with trees if three or more
				// adjacent acres contained trees. Otherwise, nothing happens.
				if countNeighbors(current, x, y, Woods) >= 3 {
					next[y][x] = Woods
				} else {
					next[y][x] = OpenGround
				}

			case Woods:
				// An acre filled with trees will become a lumberyard if
				// three or more adjacent acres were lumberyards. Otherwise,
				// nothing happens.
				if countNeighbors(current, x, y, LumberYard) >= 3 {
					next[y][x] = LumberYard
				} else {
					next[y][x] = Woods
				}

			case LumberYard:
				// An acre containing a lumberyard will remain a lumberyard
				// if it was adjacent to at least one other lumberyard and
				// at least one acre containing trees. Otherwise, it becomes open.
				if countNeighbors(current, x, y, LumberYard) >= 1 && countNeighbors(current, x, y, Woods) >= 1 {
					next[y][x] = LumberYard
				} else {
					next[y][x] = OpenGround
				}
			}
		}
	}

	return next
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s input-file rounds", os.Args[0])
	}

	rounds, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid rounds argument\n")
	}

	landscape := loadLandscape(os.Args[1])

	fmt.Printf("Starting time:\n")
	dumpLandscape(landscape)

	for round := 0; round < rounds; round++ {
		landscape = clockTick(landscape)

		fmt.Printf("\nRound %d:\n", round)
		dumpLandscape(landscape)
	}

	fmt.Printf("%d wooded acres\n", countTiles(landscape, Woods))
	fmt.Printf("%d lumberyards\n", countTiles(landscape, LumberYard))
	fmt.Printf("Total resource value: %d\n", countTiles(landscape, LumberYard)*countTiles(landscape, Woods))
}
