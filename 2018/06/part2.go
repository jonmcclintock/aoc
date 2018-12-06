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

	regionSize := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			curPoint := Point{x: x, y: y}
			total := 0

			for _, p := range points {
				total += p.distance(curPoint)
			}

			if total < 10000 {
				regionSize++
			}
		}
	}

	fmt.Printf("Region size: %d\n", regionSize)
}
