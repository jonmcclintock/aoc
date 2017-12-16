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
const rounds = 1000000000

type move struct {
	stepType byte
	params   []byte
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
			m := move{stepType: v[0]}
			switch m.stepType {
			case 's':
				n, _ := strconv.Atoi(v[1:])
				m.params = []byte{byte(n)}
			case 'x':
				s := strings.Split(v[1:], "/")
				a, _ := strconv.Atoi(s[0])
				b, _ := strconv.Atoi(s[1])
				m.params = []byte{byte(a), byte(b)}
			case 'p':
				s := strings.Split(v[1:], "/")
				m.params = []byte{s[0][0], s[1][0]}
			}

			moves = append(moves, m)
		}
	}

	return moves
}

func spin(positions []byte, params []byte) {
	n := int(params[0])
	//fmt.Printf("Spin %d dancers\n", n)
	newPos := append(positions[len(positions)-n:], positions[:len(positions)-n]...)
	for i := range positions {
		positions[i] = newPos[i]
	}
}

func exchange(positions []byte, params []byte) {
	a := params[0]
	b := params[1]
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

func partner(positions []byte, params []byte) {
	a := find(positions, params[0])
	b := find(positions, params[1])
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
	history := map[string]string{}

	positions := make([]byte, numPositions)
	for i := 0; i < len(positions); i++ {
		positions[i] = 'a' + byte(i)
	}

	for i := 0; i < rounds; i++ {
		old := string(positions)
		if n, ok := history[old]; ok {
			positions = []byte(n)
			continue
		}

		for _, m := range moves {
			//dump(positions)
			switch m.stepType {
			case 's':
				spin(positions, m.params)
			case 'x':
				exchange(positions, m.params)
			case 'p':
				partner(positions, m.params)
			}
		}
		history[old] = string(positions)
	}
	fmt.Printf("History has %d entries\n", len(history))
	dump(positions)

}
