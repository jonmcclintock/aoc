package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "day22-in.txt"
const burstCount = 10000000

type state int

const (
	clean state = iota
	weakened
	infected
	flagged
)

type grid struct {
	cells                  map[uint64]state
	minX, maxX, minY, maxY int
}

func (g grid) get(p position) state {
	if val, ok := g.cells[p.toUint64()]; ok {
		return val
	}
	return clean
}

func (g *grid) set(p position, s state) {
	if p.x < g.minX {
		g.minX = p.x
	}
	if p.y < g.minY {
		g.minY = p.y
	}
	if p.x > g.maxX {
		g.maxX = p.x
	}
	if p.y > g.maxY {
		g.maxY = p.y
	}

	g.cells[p.toUint64()] = s
}

func (g grid) dump(cur position) {
	fmt.Printf("Grid ranges from (%d, %d) to (%d, %d)\n", g.minX, g.minY, g.maxX, g.maxY)
	var p position
	for p.y = g.minY; p.y <= g.maxY; p.y++ {
		for p.x = g.minX; p.x <= g.maxX; p.x++ {
			if p.x == cur.x && p.y == cur.y {
				fmt.Print("[")
			} else {
				fmt.Print(" ")
			}
			switch g.get(p) {
			case clean:
				fmt.Print(".")
			case weakened:
				fmt.Print("W")
			case infected:
				fmt.Print("#")
			case flagged:
				fmt.Print("F")
			}
			if p.x == cur.x && p.y == cur.y {
				fmt.Print("]")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func loadMap() grid {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := grid{
		cells: map[uint64]state{},
	}
	cur := position{x: 0, y: 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		for cur.x = 0; cur.x < len(row); cur.x++ {
			if row[cur.x] == '#' {
				res.set(cur, infected)
			} else {
				res.set(cur, clean)
			}
		}
		cur.y++
	}

	return res
}

type position struct {
	x, y int
}

func (p *position) add(otherP position) {
	p.x += otherP.x
	p.y += otherP.y
}

func (p position) toUint64() uint64 {
	return (uint64(p.x) << 32) | uint64(p.y)&0xffffffff
}

func main() {
	directions := []position{
		{x: 0, y: -1}, // North   /|\             |
		{x: 1, y: 0},  // East     |  Turn        |  Turn
		{x: 0, y: 1},  // South    |  Left        |  Right
		{x: -1, y: 0}, // West     |             \|/
	}

	g := loadMap()

	curP := position{
		x: (g.maxX - g.minX) / 2,
		y: (g.maxY - g.minY) / 2,
	}
	curD := 0

	countC := 0
	countW := 0
	countI := 0
	countF := 0
	for i := 0; i < burstCount; i++ {
		switch g.get(curP) {
		case clean:
			// Turn left, set cur to weakened
			curD = curD - 1
			if curD < 0 {
				curD = len(directions) - 1
			}
			g.set(curP, weakened)
			countW++
		case weakened:
			// Don't turn, infect cur
			g.set(curP, infected)
			countI++
		case infected:
			// Turn right, set cur to flagged
			curD = (curD + 1) % len(directions)
			g.set(curP, flagged)
			countF++
		case flagged:
			// Reverse, clean
			curD = (curD + 2) % len(directions)
			g.set(curP, clean)
			countC++
		}

		// Now step in the current direction.
		curP.add(directions[curD])
	}

	fmt.Printf("After %d steps, %d cells were infected, %d were cleaned\n", burstCount, countI, countC)
	g.dump(curP)
}
