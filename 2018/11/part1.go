package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func buildGrid(w, h, serial int) [][]int {
	grid := make([][]int, w)
	for x := 0; x < w; x++ {
		grid[x] = make([]int, h)
		for y := 0; y < h; y++ {
			grid[x][y] = computePowerLevel(x, y, serial)
		}
	}
	return grid
}

func computePowerLevel(x, y, serial int) int {
	rackID := x + 10
	basePower := rackID*y + serial
	hundreds := (basePower * rackID / 100) % 10

	return hundreds - 5
}

func sumPower(grid [][]int, x, y int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += grid[x+i][y+j]
		}
	}
	return sum
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s size serial", os.Args[0])
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid grid size", err)
	}

	serial, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid serial", err)
	}

	grid := buildGrid(size, size, serial)

	maxX, maxY, maxPower := 0, 0, 0
	for x := 0; x < size-3; x++ {
		for y := 0; y < size-3; y++ {
			power := sumPower(grid, x, y)
			if power > maxPower {
				maxX = x
				maxY = y
				maxPower = power
			}
		}
	}

	fmt.Printf("Highest power is %d at (%d, %d)\n", maxPower, maxX, maxY)
}
