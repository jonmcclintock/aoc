package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// CanvasSize is how big the canvas is on each side.
const CanvasSize = 1000

// Claim is a single claim for a piece of the canvas.
type Claim struct {
	id   int
	x, y int
	w, h int
}

func parseClaim(in string) Claim {
	c := Claim{}
	_, err := fmt.Sscanf(in, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func makeCanvas(size int) [][]int {
	canvas := make([][]int, size)

	for i := 0; i < size; i++ {
		canvas[i] = make([]int, size)
	}

	return canvas
}

func fillCanvas(canvas [][]int, c Claim) {
	for x := c.x; x < (c.x + c.w); x++ {
		for y := c.y; y < (c.y + c.h); y++ {
			canvas[x][y]++
		}
	}
}

func countOverlap(canvas [][]int) int {
	count := 0
	for x := 0; x < len(canvas); x++ {
		for y := 0; y < len(canvas[x]); y++ {
			if canvas[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatal(fmt.Errorf("Usage: %s input-file", os.Args[0]))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	canvas := makeCanvas(CanvasSize)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := parseClaim(scanner.Text())
		fmt.Printf("Claim %d: %dx%d at %d,%d\n", c.id, c.w, c.h, c.x, c.y)
		fillCanvas(canvas, c)
	}

	fmt.Printf("Inches with overlap: %d\n", countOverlap(canvas))
}
