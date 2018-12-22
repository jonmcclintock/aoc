package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A Room in the building
type Room struct {
	x     int
	y     int
	doors map[byte]*Room
}

var opposites = map[byte]byte{
	'N': 'S',
	'E': 'W',
	'W': 'E',
	'S': 'N',
}

var directions = map[byte][]int{
	'N': []int{0, -1},
	'E': []int{1, 0},
	'W': []int{-1, 0},
	'S': []int{0, 1},
}

// Read in the map file and return the current Room.
func readMap(fileName string) (map[int]*Room, *Room) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	here := &Room{
		x:     0,
		y:     0,
		doors: make(map[byte]*Room),
	}
	rooms := map[int]*Room{
		getCoord(0, 0): here,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Add the line to the room, trimming the carat and dollar on the ends.
		addRegexToRoom(rooms, here, line[1:len(line)-1], 0)
	}

	return rooms, here
}

// Combine an X,Y coordinate into a single integer value.
func getCoord(x, y int) int {
	if x > 2<<16 || x < -1*(2<<16) {
		log.Fatalf("Fell off the map, X coord of %d\n", x)
	}
	if y > 2<<16 || y < -1*(2<<16) {
		log.Fatalf("Fell off the map, Y coord of %d\n", y)
	}

	x += 2 << 16
	y += 2 << 16

	return y<<32 + x
}

// Add a door to the current room and return the new room.
func addDoor(rooms map[int]*Room, here *Room, dir byte) *Room {
	x := here.x + directions[dir][0]
	y := here.y + directions[dir][1]
	coord := getCoord(x, y)

	// Create the new room if it doesn't exist.
	if _, ok := rooms[coord]; !ok {
		rooms[coord] = &Room{
			x:     x,
			y:     y,
			doors: make(map[byte]*Room, 0),
		}
	}

	// Connect the rooms
	here.doors[dir] = rooms[coord]
	rooms[coord].doors[opposites[dir]] = here

	fmt.Printf("    -> Added door %c of (%2d,%2d) into (%2d,%2d)\n", dir, here.x, here.y, x, y)
	//dumpMap(rooms, rooms[coord])

	return rooms[coord]
}

// Return the position of the first unmatched closing paren after the current posision.
func skipToClosingParen(reg string, pos int) int {
	depth := 1
	for ; depth > 0 && pos < len(reg); pos++ {
		if reg[pos] == '(' {
			depth++
		}
		if reg[pos] == ')' {
			depth--
		}
	}

	if pos == len(reg) {
		return pos
	}

	return pos - 1
}

// Add the given regex to the room specified.
func addRegexToRoom(rooms map[int]*Room, here *Room, reg string, depth int) {
	cur := here
	pos := 0

	fmt.Printf("\n%2d: Starting add to room from (%d,%d): %s\n", depth, here.x, here.y, reg)

	// Consume the leading plain characters
	for pos = 0; pos < len(reg) && (reg[pos] == 'N' || reg[pos] == 'E' || reg[pos] == 'W' || reg[pos] == 'S'); pos++ {
		cur = addDoor(rooms, cur, reg[pos])
	}

	// We're at the end, we're done.
	if pos == len(reg) {
		fmt.Printf("%2d: Reached the end of the string after consuming normal chars (%d,%d)\n\n", depth, cur.x, cur.y)
		return
	}

	// We're at the end of this option, skip to the end of the branch.
	if reg[pos] == '|' {
		fmt.Printf("%2d: - Found a separator at %d, skipping to closing paren '%s'\n", depth, pos, reg)
		pos = skipToClosingParen(reg, pos+1)
		if pos == len(reg) {
			fmt.Printf("%2d: -> Skipped to end from separator.\n", depth)
			return
		}
		fmt.Printf("%2d: - Now at %d: %s\n", depth, pos, reg[pos:])
	}

	// We're at the end of this branch, continue normal processing from the next character.
	if reg[pos] == ')' {
		fmt.Printf("%2d: - At the end of the branch (%d), continuing from (%d,%d) at (%d,%d)\n", depth, pos, here.x, here.y, cur.x, cur.y)
		addRegexToRoom(rooms, cur, reg[pos+1:], depth+1)
		fmt.Printf("%2d: -> Done processing branch from (%d,%d), at (%d,%d)\n\n", depth, here.x, here.y, cur.x, cur.y)
		return
	}

	if reg[pos] != '(' {
		log.Fatalf("Unexpected character '%c' at %d, rest of pattern is: %s\n", reg[pos], pos, reg[pos:])
	}

	// We have a new branch, recurse down for each option.
	pos++
	for pos < len(reg) {
		// Parse the current option.
		fmt.Printf("%2d: - Recursing to parse an option at %d (%d,%d)\n", depth, pos, cur.x, cur.y)
		addRegexToRoom(rooms, cur, reg[pos:], depth+1)

		// Skip ahead to the next option.
		for ; pos < len(reg) && reg[pos] != '|' && reg[pos] != ')'; pos++ {
			// If we hit another branch, skip to the char after its closing paren.
			if reg[pos] == '(' {
				pos = skipToClosingParen(reg, pos+1)
			}
		}

		if pos < len(reg) && reg[pos] == ')' {
			break
		}

		// Increment past the separator.
		pos++
	}

	fmt.Printf("%2d: -> Done processing from (%d,%d), at (%d,%d)\n\n", depth, here.x, here.y, cur.x, cur.y)
}

// Get the shortest path to the target from the given position.
func getShortestPath(source, target *Room, rooms map[int]*Room, visited map[int]bool) int {
	cur := getCoord(source.x, source.y)

	// If we're here, it's zero doors away.
	if source == target {
		return 0
	}

	// We've already visited this node, it's a dead end.
	if visited[cur] {
		return -1
	}

	minDistance := -1
	visited[cur] = true
	for _, neighbor := range source.doors {
		distance := getShortestPath(neighbor, target, rooms, visited)
		if distance == -1 {
			continue
		}
		distance++
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
		}
	}

	visited[cur] = false
	return minDistance
}

// Return the distance of the room with the farthest optimally efficient path to it.
func getFarthestRoomDistance(source *Room, rooms map[int]*Room) int {
	maxDistance := 0
	minX, minY, maxX, maxY := getRoomRange(rooms)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			guess, ok := rooms[getCoord(x, y)]
			if !ok {
				continue
			}
			if guess == source {
				continue
			}
			visited := make(map[int]bool, 0)
			distance := getShortestPath(source, guess, rooms, visited)
			if distance != -1 && distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

// Get the min and max X and Y coordinates in the map.
func getRoomRange(rooms map[int]*Room) (int, int, int, int) {
	minX, minY := 0, 0
	maxX, maxY := 0, 0
	for _, room := range rooms {
		if room.x < minX {
			minX = room.x
		}
		if room.y < minY {
			minY = room.y
		}
		if room.x > maxX {
			maxX = room.x
		}
		if room.y > maxY {
			maxY = room.y
		}
	}

	return minX, minY, maxX, maxY
}

func dumpRooms(rooms map[int]*Room) {
	fmt.Printf("Rooms:\n")
	for _, room := range rooms {
		fmt.Printf("- (%2d,%2d): ", room.x, room.y)
		for dir := range room.doors {
			fmt.Printf("%c, ", dir)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func dumpMap(rooms map[int]*Room, here *Room) {
	// Scan through the rooms, find the top left and bottom right.

	minX, minY, maxX, maxY := getRoomRange(rooms)
	for y := minY; y <= maxY; y++ {
		// First print the vertical movement directions.
		for x := minX; x <= maxX; x++ {
			room, ok := rooms[getCoord(x, y)]
			fmt.Printf("#")
			if ok {
				if _, ok = room.doors['N']; ok {
					fmt.Printf("-")
				} else {
					fmt.Printf("#")
				}
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("#\n")

		// Now print the horizontal movement directions.
		for x := minX; x <= maxX; x++ {
			room, ok := rooms[getCoord(x, y)]
			if ok {
				if _, ok = room.doors['W']; ok {
					fmt.Printf("|")
				} else {
					fmt.Printf("#")
				}
			} else {
				fmt.Printf("#")
			}

			// Print an X for the origin.
			if x == 0 && y == 0 {
				fmt.Printf("X")
			} else if here != nil && here.x == x && here.y == y {
				fmt.Printf("@")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("#\n")
	}
	for x := minX; x <= maxX; x++ {
		fmt.Printf("##")
	}
	fmt.Printf("#\n\n")
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s map", os.Args[0])
	}

	rooms, _ := readMap(os.Args[1])

	//dumpRooms(rooms)

	fmt.Printf("Map:\n")
	dumpMap(rooms, nil)

	farthestRoomDistance := getFarthestRoomDistance(rooms[getCoord(0, 0)], rooms)
	fmt.Printf("Farthest away room is %d doors away.\n", farthestRoomDistance)
}
