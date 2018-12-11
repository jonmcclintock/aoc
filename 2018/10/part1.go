package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Point is a point.
type Point struct {
	x, y   int
	xV, yV int
}

func parseCoord(s string) (int, int) {
	parts := strings.Split(s, ", ")
	x, _ := strconv.Atoi(strings.Trim(parts[0], " "))
	y, _ := strconv.Atoi(strings.Trim(parts[1], " "))
	return x, y
}

func getBounds(points []*Point) (int, int, int, int) {
	minX, minY := 1000000000, 100000000
	maxX, maxY := -1000000000, -1000000000

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}

		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	return minX, minY, maxX, maxY
}

func updatePoints(points []*Point) {
	for _, p := range points {
		p.x = p.x + p.xV
		p.y = p.y + p.yV
	}
}

func calcArea(points []*Point) int {
	minX, minY, maxX, maxY := getBounds(points)
	return (maxX - minX) * (maxY - minY)
}

func dumpPoints(points []*Point) {
	minX, minY, maxX, maxY := getBounds(points)
	rows, cols := maxY-minY+1, maxX-minX+1
	matrix := make([][]bool, cols)
	for i := 0; i < cols; i++ {
		matrix[i] = make([]bool, rows)
	}

	for _, p := range points {
		matrix[p.x-minX][p.y-minY] = true
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if matrix[x][y] == true {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
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

	points := make([]*Point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "> velocity=<")
		posStr := strings.TrimPrefix(input[0], "position=<")
		velStr := strings.TrimSuffix(input[1], ">")

		p := Point{}
		p.x, p.y = parseCoord(posStr)
		p.xV, p.yV = parseCoord(velStr)

		points = append(points, &p)
	}

	lastArea := calcArea(points)
	fmt.Printf("Start area is %d\n", lastArea)

	for i := 0; ; i++ {
		updatePoints(points)
		newArea := calcArea(points)
		delta := newArea - lastArea

		if newArea < 2000 {
			dumpPoints(points)
		}
		fmt.Printf("%03d: New area is %d (delta: %d)\n", i, newArea, delta)

		if (delta > 0) && ((delta * delta) > (lastArea * lastArea >> 3)) {
			break
		}
		lastArea = newArea
	}

	fmt.Printf("Final area is %d\n", lastArea)
}
