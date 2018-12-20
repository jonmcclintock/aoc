package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const Sand = '.'
const Clay = '#'
const Water = '~'
const Trace = '|'
const Spring = '+'

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

// Parse the value range (or single value) and return the range.
func parseRange(str string) (int, int) {
	if !strings.Contains(str, "..") {
		value, err := strconv.Atoi(str[2:])
		if err != nil {
			log.Fatalf("Invalid value '%s': %+v", str, err)
		}
		return value, value
	}

	rangeStrs := strings.Split(str, "..")

	var start, end int
	var err error
	start, err = strconv.Atoi(rangeStrs[0][2:])
	if err != nil {
		log.Fatalf("Invalid value '%s': %+v", rangeStrs[0], err)
	}
	end, err = strconv.Atoi(rangeStrs[1])
	if err != nil {
		log.Fatalf("Invalid value '%s': %+v", rangeStrs[1], err)
	}

	return start, end
}

func loadScan(inputFile string) ([][]byte, Point) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	min := Point{X: math.MaxInt64, Y: math.MaxInt64}
	max := Point{X: math.MinInt64, Y: math.MinInt64}

	points := make([]Point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, ", ")
		startX, startY, endX, endY := 0, 0, 0, 0

		for _, val := range vals {
			if val[0] == 'x' {
				startX, endX = parseRange(val)
			} else {
				startY, endY = parseRange(val)
			}
		}

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				points = append(points, Point{X: x, Y: y})
			}
		}

		min.X = minInt(min.X, startX)
		min.Y = minInt(min.Y, startY)
		max.X = maxInt(max.X, endX)
		max.Y = maxInt(max.Y, endY)
	}

	fmt.Printf("Read %d points, ranging from (%d,%d) to (%d,%d)\n", len(points), min.X, min.Y, max.X, max.Y)

	origin := Point{
		X: maxInt(min.X-1, 0),
		Y: 0,
	}
	size := Point{
		X: max.X - origin.X + 2,
		Y: max.Y + 1,
	}

	scan := make([][]byte, size.Y)
	for y := 0; y < size.Y; y++ {
		scan[y] = make([]byte, size.X)
		for x := 0; x < size.X; x++ {
			scan[y][x] = Sand
		}
	}

	for _, p := range points {
		setPoint(scan, origin, p, Clay)
	}

	return scan, origin
}

func getPoint(scan [][]byte, origin, p Point) byte {
	return scan[p.Y-origin.Y][p.X-origin.X]
}

func setPoint(scan [][]byte, origin, p Point, v byte) {
	scan[p.Y-origin.Y][p.X-origin.X] = v
}

func dumpScan(scan [][]byte, origin Point) {
	for y := 0; y < len(scan); y++ {
		fmt.Printf("%4d: %s\n", y+origin.Y, string(scan[y]))
	}
}

func countTiles(scan [][]byte, t byte) int {
	count := 0
	for y := 0; y < len(scan)-1; y++ {
		for x := 0; x < len(scan[y]); x++ {
			if scan[y][x] == t {
				count++
			}
		}
	}
	return count
}

// Flow water down below the current point, which we assume is already full
// of water or a trace.
func flowDown(scan [][]byte, p Point) int {
	row := p.Y + 1

	// If the cell below us is Sand, flow a Trace into it.
	if scan[row][p.X] == Sand {
		scan[row][p.X] = Trace
		return 1
	}

	return 0
}

// Find the boundary or edge going in the given direction. Returns true to
// indicate that you're bounded on this side, as well as the boundary edge.
func findBoundary(scan [][]byte, p Point, dir int) (bool, int) {
	var i int
	bounded := false
	for i = p.X; i >= 0 && i < len(scan[p.Y]); i += dir {
		// If we've reached the left side, stop here.
		if scan[p.Y][i] == Clay {
			bounded = true
			i -= dir
			break
		}

		// If we've fallen off an edge and we don't have water or clay below us, stop.
		if scan[p.Y+1][i] != Water && scan[p.Y+1][i] != Clay {
			break
		}
	}

	return bounded, i
}

// Flow horizontally from the current point, which we assume to be full of water.
func flowAcross(scan [][]byte, p Point) int {
	// Find the left and right bounds of this clay
	leftBounded, leftBounds := findBoundary(scan, p, -1)
	rightBounded, rightBounds := findBoundary(scan, p, 1)

	fillType := byte(Trace)
	if leftBounded && rightBounded {
		fillType = Water
	}

	count := 0
	for i := leftBounds; i <= rightBounds; i++ {
		if scan[p.Y][i] != fillType {
			scan[p.Y][i] = fillType
			count++
		}
	}

	return count
}

// Run a single pass at flowing water through the scan.
func flowWater(scan [][]byte) int {
	count := 0

	// Water flows down. We don't flow past the bottom row. For each row, we'll
	// look at the contents of the cell below it and decide what to do.
	for y := 0; y < len(scan)-1; y++ {
		for x := 0; x < len(scan[y]); x++ {
			if scan[y][x] == Spring || scan[y][x] == Trace {
				count += flowDown(scan, Point{X: x, Y: y})
			}
			if scan[y][x] == Trace {
				count += flowAcross(scan, Point{X: x, Y: y})
			}
		}
	}

	return count
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s input-file", os.Args[0])
	}

	scan, origin := loadScan(os.Args[1])
	setPoint(scan, origin, Point{X: 500, Y: 0}, Spring)

	fmt.Printf("Starting scan:\n")
	dumpScan(scan, origin)

	fillCount := math.MaxInt64
	for round := 1; fillCount > 0; round++ {
		fillCount = flowWater(scan)

		fmt.Printf("\nRound %d, filled %d cells:\n", round, fillCount)
		dumpScan(scan, origin)
	}

	fmt.Printf("%d wet tiles\n", countTiles(scan, Trace)+countTiles(scan, Water))
	fmt.Printf("%d water tiles\n", countTiles(scan, Water))
}
