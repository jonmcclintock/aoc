package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day16-in.txt"
const numPositions = 16
const rounds = 1000

type move struct {
	stepType string
	params   string
}

func loadMoves() []move {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moves := []move{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		for _, v := range s {
			if len(v) < 1 {
				continue
			}
			moves = append(moves, move{stepType: string(v[0]), params: v[1:]})
		}
	}

	return moves
}

func spin(positions []byte, params string) {
	n, _ := strconv.Atoi(params)
	//fmt.Printf("Spin %d dancers\n", n)
	newPos := append(positions[len(positions)-n:], positions[:len(positions)-n]...)
	for i := range positions {
		positions[i] = newPos[i]
	}
}

func exchange(positions []byte, params string) {
	s := strings.Split(params, "/")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	//fmt.Printf("Exchange %d with %d\n", a, b)
	positions[a], positions[b] = positions[b], positions[a]
}

func find(positions []byte, dancer byte) int {
	for i := 0; i < len(positions); i++ {
		if positions[i] == dancer {
			return i
		}
	}
	return 0
}

func partner(positions []byte, params string) {
	s := strings.Split(params, "/")
	a := find(positions, s[0][0])
	b := find(positions, s[1][0])
	//fmt.Printf("Partner %c (%d) with %c (%d)\n", s[0][0], a, s[1][0], b)
	positions[a], positions[b] = positions[b], positions[a]
}

func dump(positions []byte) {
	fmt.Printf("Positions: ")
	for i := 0; i < len(positions); i++ {
		fmt.Printf("%c", positions[i])
	}
	fmt.Printf("\n")
}

func main() {
	moves := loadMoves()

	positions := make([]byte, numPositions)
	for i := 0; i < len(positions); i++ {
		positions[i] = 'a' + byte(i)
	}

	for i := 0; i < rounds; i++ {
		for _, m := range moves {
			//dump(positions)
			switch m.stepType {
			case "s":
				spin(positions, m.params)
			case "x":
				exchange(positions, m.params)
			case "p":
				partner(positions, m.params)
			}
		}
	}
	dump(positions)

}
