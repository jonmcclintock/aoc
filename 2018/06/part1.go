package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Point is a point
type Point struct {
	x, y int
}

func (p Point) distance(o Point) int {
	dX := o.x - p.x
	dY := o.y - p.y
	if dX < 0 {
		dX = -dX
	}
	if dY < 0 {
		dY = -dY
	}

	return dX + dY
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s input-file", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := make([]Point, 0)

	minX, minY := -1, -1
	maxX, maxY := -1, -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, ", ")
		if len(vals) != 2 {
			log.Fatalf("Invalid input: %s\n", line)
		}

		p := Point{}
		p.x, _ = strconv.Atoi(vals[0])
		p.y, _ = strconv.Atoi(vals[1])

		if p.x < minX || minX == -1 {
			minX = p.x
		}
		if p.y < minY || minY == -1 {
			minY = p.y
		}
		if p.x > maxX || maxX == -1 {
			maxX = p.x
		}
		if p.y > maxY || maxY == -1 {
			maxY = p.y
		}

		points = append(points, p)
	}

	counts := make([]int, len(points))
	infinite := make([]bool, len(points))
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			curPoint := Point{x: x, y: y}
			closest := -1
			bestDistance := 1000000
			equidistant := false

			for i, p := range points {
				curDistance := p.distance(curPoint)
				if curDistance < bestDistance {
					closest = i
					bestDistance = curDistance
					equidistant = false
				} else if curDistance == bestDistance {
					equidistant = true
				}
			}

			if equidistant {
				fmt.Printf(".")
				continue
			}

			if bestDistance == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c", 'A'+closest)
			}
			counts[closest]++

			if x == minX || x == maxX || y == minY || y == maxY {
				infinite[closest] = true
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\nGrid range is (%d, %d) to (%d, %d) with %d coordinates\n", minX, minY, maxX, maxY, len(points))

	biggest := 0
	biggestArea := 0
	for i, a := range counts {
		fmt.Printf("%c (%3d,%3d): %d", 'A'+i, points[i].x, points[i].y, a)
		if infinite[i] {
			fmt.Printf(" (infinite)\n")
		} else {
			fmt.Printf("\n")
		}

		if a > biggestArea && !infinite[i] {
			biggest = i
			biggestArea = a
		}
	}

	fmt.Printf("\n%c (%d, %d) has the biggest area: %d\n", 'A'+biggest, points[biggest].x, points[biggest].y, biggestArea)
}
