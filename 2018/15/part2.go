package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

const baseHP = 200

// The Square of a square
type Square int

const (
	Empty Square = iota
	Elf
	Goblin
	Wall
)

var attackHP = map[Square]int{
	Elf:    3,
	Goblin: 3,
}

var enemies = map[Square]Square{
	Elf:    Goblin,
	Goblin: Elf,
}

// A Unit represents an actor on the map.
type Unit struct {
	Type Square
	X, Y int
	HP   int
}

// Reading directions.
var directions = [][]int{
	[]int{0, -1}, // Up
	[]int{-1, 0}, // Left
	[]int{1, 0},  // Right
	[]int{0, 1},  // Down
}

// Read in the map file and return an array of the map, and a list of the Units.
func readMap(fileName string) ([][]Square, []*Unit) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cave := make([][]Square, 0)
	units := make([]*Unit, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]Square, len(line))
		for i := range row {
			switch line[i] {
			case '.':
				row[i] = Empty
			case '#':
				row[i] = Wall
			case 'G':
				row[i] = Goblin
				u := Unit{
					Type: Goblin,
					X:    i,
					Y:    len(cave),
					HP:   baseHP,
				}
				units = append(units, &u)
			case 'E':
				row[i] = Elf
				u := Unit{
					Type: Elf,
					X:    i,
					Y:    len(cave),
					HP:   baseHP,
				}
				units = append(units, &u)
			}
		}

		cave = append(cave, row)
	}

	return cave, units
}

// Sort the Units in reading order.
func sortUnits(units []*Unit) {
	sort.Slice(units, func(i, j int) bool {
		if units[i].Y < units[j].Y {
			return true
		}
		if units[i].Y > units[j].Y {
			return false
		}
		return units[i].X < units[j].X
	})
}

// Get the shortest path to the target from the given position. Returns
// the x,y direction to go, and the length of the path.
func getShortestPath(sourceX, sourceY, targetX, targetY int, cave [][]Square) (int, int, int) {
	distances := make([][]int, len(cave))
	for y := range distances {
		distances[y] = make([]int, len(cave[y]))
		for x := range distances[y] {
			distances[y][x] = math.MaxInt64
		}
	}

	visited := make([][]bool, len(cave))
	for y := range visited {
		visited[y] = make([]bool, len(cave[y]))
	}

	// Start with the target, and compute the distances back to the source.
	tentative := make([]int, 1)
	tentative[0] = targetX<<16 + targetY
	distances[targetY][targetX] = 0

	//fmt.Printf("Looking for distances from (%d,%d) to (%d,%d)\n", sourceX, sourceY, targetX, targetY)

	for len(tentative) > 0 {
		// Find the tentative cell with the shortest distance
		cur := 0
		minDistance := math.MaxInt64
		for i, pos := range tentative {
			x, y := (pos >> 16), pos&0xffff
			if distances[y][x] < minDistance {
				minDistance = distances[y][x]
				cur = i
			}
		}

		// Remove that cell from the tentative list
		curX, curY := tentative[cur]>>16, tentative[cur]&0xffff
		tentative = append(tentative[:cur], tentative[cur+1:]...)
		visited[curY][curX] = true

		//fmt.Printf("Visited (%d,%d), distance is %d (queue is %d long)\n", curX, curY, distances[curY][curX], len(tentative))

		// Consider the neighbors.
		for _, move := range directions {
			newX := curX + move[0]
			newY := curY + move[1]

			// Skip if we're past the edge of the cave
			if newX < 0 || newY < 0 || newX >= len(cave[0]) || newY >= len(cave) {
				continue
			}

			// Skip if it's not empty.
			if cave[newY][newX] != Empty {
				continue
			}

			// If we haven't visited that cell, set a tentative distance
			if !visited[newY][newX] {
				visited[newY][newX] = true
				tentative = append(tentative, newX<<16+newY)
				distances[newY][newX] = distances[curY][curX] + 1
				//fmt.Printf("- First visit to (%d,%d), set distance to %d\n", newX, newY, distances[newY][newX])
			} else {
				if distances[newY][newX] > distances[curY][curX] {
					distances[newY][newX] = distances[curY][curX] + 1
					//fmt.Printf("- Revisit to (%d,%d), set distance to %d\n", newX, newY, distances[newY][newX])
				}
			}
		}
	}

	//fmt.Printf("-> Distances from (%d,%d) to (%d,%d):\n", sourceX, sourceY, targetX, targetY)
	//dumpDistances(sourceX, sourceY, targetX, targetY, distances)

	// We've got the distance, now pick a direction.
	bestDistance := math.MaxInt64
	bestDX, bestDY := 0, 0
	for _, move := range directions {
		newX := sourceX + move[0]
		newY := sourceY + move[1]

		// Skip if we're past the edge of the cave
		if newX < 0 || newY < 0 || newX >= len(cave[0]) || newY >= len(cave) {
			continue
		}

		// Skip if it's not empty.
		if cave[newY][newX] != Empty {
			continue
		}

		if distances[newY][newX] != math.MaxInt64 && distances[newY][newX]+1 < bestDistance {
			bestDistance = distances[newY][newX] + 1
			bestDX, bestDY = move[0], move[1]
		}
	}

	if bestDistance != math.MaxInt64 {
		bestDistance++
	}

	//fmt.Printf("-> Shortest path from (%d, %d) to (%d, %d) is %d long going (%d,%d)\n", sourceX, sourceY, targetX, targetY, bestDistance, bestDX, bestDY)
	return bestDX, bestDY, bestDistance
}

// Move the Unit on the map towards the closest target.
func moveUnit(u *Unit, allUnits []*Unit, cave [][]Square) {
	// fmt.Printf("\nLooking for targets from (%d,%d) (t: %d, hp: %d)\n", u.X, u.Y, u.Type, u.HP)

	// First build the list of candidate in-range squares
	candidates := make([]int, 0)
	for _, t := range allUnits {
		if t.Type == u.Type || t.HP <= 0 {
			continue
		}

		for _, move := range directions {
			testX, testY := t.X+move[0], t.Y+move[1]

			// Skip if we're past the edge of the cave
			if testX < 0 || testY < 0 || testX >= len(cave[0]) || testY >= len(cave) {
				continue
			}

			// If this unit is next to us, we don't need to move.
			if u.X == testX && u.Y == testY {
				// fmt.Printf("- Target is next to this unit at (%d,%d)\n", t.X, t.Y)
				return
			}

			// If it's not empty, skip it.
			if cave[testY][testX] != Empty {
				continue
			}

			candidates = append(candidates, testX<<16+testY)
		}
	}

	//fmt.Printf("-> %d candidate squares\n", len(candidates))

	// No candidates, don't move.
	if len(candidates) == 0 {
		return
	}

	// Now filter out unreachable candidates and build a list of shortest path candidates..
	shortestDistance := math.MaxInt64
	reachable := make([]int, 0)
	for _, pos := range candidates {
		_, _, d := getShortestPath(u.X, u.Y, pos>>16, pos&0xffff, cave)
		if d < shortestDistance {
			reachable = []int{pos}
			shortestDistance = d
		} else if d == shortestDistance {
			reachable = append(reachable, pos)
		}
	}

	//fmt.Printf("-> %d nearby reachable squares with a path of %d squares\n", len(reachable), shortestDistance)

	// No reachable candidates, don't move.
	if len(reachable) == 0 {
		return
	}

	// Sort the reachables by reading order.
	sort.Slice(reachable, func(i, j int) bool {
		iX, iY := reachable[i]>>16, reachable[i]&0xffff
		jX, jY := reachable[j]>>16, reachable[j]&0xffff
		if iY < jY {
			return true
		} else if iY > jY {
			return false
		}
		return iX < jX
	})

	//fmt.Printf("-> Nearest reachable (of %d) is (%d,%d)\n", len(reachable), reachable[0]>>16, reachable[0]&0xffff)

	// Get the distance to the position.
	dX, dY, _ := getShortestPath(u.X, u.Y, reachable[0]>>16, reachable[0]&0xffff, cave)

	//fmt.Printf("- Moving unit at %d,%d by %d,%d\n", u.X, u.Y, dX, dY)

	cave[u.Y][u.X] = Empty
	u.X += dX
	u.Y += dY
	cave[u.Y][u.X] = u.Type
}

// Find a target of the given unit and attack it.
func targetAndAttack(u *Unit, allUnits []*Unit, cave [][]Square) {
	enemy := Elf
	if u.Type == Elf {
		enemy = Goblin
	}

	// Find the weakest neighbor enemy.
	var weakest *Unit
	for _, dir := range directions {
		x, y := u.X+dir[0], u.Y+dir[1]
		if cave[y][x] == enemy {
			for _, t := range allUnits {
				if t.HP <= 0 {
					continue
				}
				if t.X == x && t.Y == y && (weakest == nil || t.HP < weakest.HP) {
					weakest = t
				}
			}
		}
	}

	if weakest == nil {
		return
	}

	weakest.HP -= attackHP[u.Type]
	fmt.Printf("Unit at (%d,%d) attacked unit at (%d,%d), %d HP remaining \n", u.X, u.Y, weakest.X, weakest.Y, weakest.HP)
	if weakest.HP <= 0 {
		fmt.Printf("Unit at (%d,%d) died!\n", weakest.X, weakest.Y)
		cave[weakest.Y][weakest.X] = Empty
	}
}

// Count the units of the given type
func countUnits(units []*Unit, t Square) int {
	count := 0
	for _, u := range units {
		if u.HP > 0 && u.Type == t {
			count++
		}
	}
	return count
}

func runSimulation(units []*Unit, cave [][]Square) (int, []*Unit) {
	rounds := 0
	for countUnits(units, Elf) > 0 && countUnits(units, Goblin) > 0 {

		// Move and attack
		complete := true
		for i, u := range units {
			if u.HP <= 0 {
				continue
			}

			moveUnit(u, units, cave)
			targetAndAttack(u, units, cave)

			if countUnits(units, Elf) == 0 || countUnits(units, Goblin) == 0 {
				if i < len(units)-1 {
					complete = false
				}
				break
			}
		}

		// Remove any dead units
		for i := 0; i < len(units); i++ {
			if units[i].HP <= 0 {
				units = append(units[:i], units[i+1:]...)
				i--
			}
		}

		sortUnits(units)

		if complete {
			rounds++
		}
		fmt.Printf("\nAfter %d rounds\n", rounds)
		dumpCave(cave)
		dumpUnits(units)
		fmt.Printf("\n---\n")
	}

	return rounds, units
}

func dumpDistances(sX, sY, tX, tY int, distances [][]int) {
	for y, row := range distances {
		for x, d := range row {
			if y == sY && x == sX {
				fmt.Printf("[")
			} else if y == tY && x == tX {
				fmt.Printf(">")
			} else {
				fmt.Printf(" ")
			}

			if d == math.MaxInt64 {
				fmt.Printf("..")
			} else {
				fmt.Printf("%2d", d)
			}

			if y == sY && x == sX {
				fmt.Printf("]")
			} else if y == tY && x == tX {
				fmt.Printf("<")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func dumpCave(cave [][]Square) {
	for y, row := range cave {
		fmt.Printf("%3d ", y)
		for x, square := range row {
			switch square {
			case Empty:
				fmt.Printf(".")
			case Wall:
				fmt.Printf("#")
			case Elf:
				fmt.Printf("E")
			case Goblin:
				fmt.Printf("G")
			}
			if x%5 == 4 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func dumpUnits(units []*Unit) {
	for i, u := range units {
		fmt.Printf("%d: (%d,%d), type %d: %d HP\n", i, u.X, u.Y, u.Type, u.HP)
	}
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s map", os.Args[0])
	}

	cave, units := readMap(os.Args[1])
	sortUnits(units)

	fmt.Printf("Starting cave:\n")
	dumpCave(cave)

	for power, dead := 3, 1; dead > 0; power++ {
		runCave := make([][]Square, len(cave))
		for i := range runCave {
			runCave[i] = make([]Square, len(cave[i]))
			copy(runCave[i], cave[i])
		}

		runUnits := make([]*Unit, len(units))
		for i, u := range units {
			newUnit := *u
			runUnits[i] = &newUnit
		}

		attackHP[Elf] = power
		rounds, resultUnits := runSimulation(runUnits, runCave)
		dead = countUnits(units, Elf) - countUnits(resultUnits, Elf)
		fmt.Printf("Combat with power %d ends after %d full rounds with %d elves dead\n", power, rounds, dead)
		winner := Elf
		winnerStr := "Elves"
		if countUnits(resultUnits, Goblin) > 0 {
			winner = Goblin
			winnerStr = "Goblins"
		}

		points := 0
		for _, u := range resultUnits {
			if u.Type == winner {
				points += u.HP
			}
		}

		fmt.Printf("%s win with %d total hit points left\n", winnerStr, points)
		fmt.Printf("Outcome: %d * %d = %d\n", rounds, points, rounds*points)
	}
}
