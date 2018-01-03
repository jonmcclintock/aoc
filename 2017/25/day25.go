package main

import "fmt"

const startState = "A"

const stepsToRun = 12459852

//const stepsToRun = 6

type update struct {
	writeValue int
	moveRight  bool
	nextState  string
}

type stateTransition map[int]update

type program map[string]stateTransition

func main() {
	/*p := program{
		"A": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "B",
			},
			1: update{
				writeValue: 0,
				moveRight:  false,
				nextState:  "B",
			},
		},
		"B": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  false,
				nextState:  "A",
			},
			1: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "A",
			},
		},
	}*/
	p := program{
		"A": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "B",
			},
			1: update{
				writeValue: 1,
				moveRight:  false,
				nextState:  "E",
			},
		},
		"B": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "C",
			},
			1: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "F",
			},
		},
		"C": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  false,
				nextState:  "D",
			},
			1: update{
				writeValue: 0,
				moveRight:  true,
				nextState:  "B",
			},
		},
		"D": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "E",
			},
			1: update{
				writeValue: 0,
				moveRight:  false,
				nextState:  "C",
			},
		},
		"E": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  false,
				nextState:  "A",
			},
			1: update{
				writeValue: 0,
				moveRight:  true,
				nextState:  "D",
			},
		},
		"F": stateTransition{
			0: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "A",
			},
			1: update{
				writeValue: 1,
				moveRight:  true,
				nextState:  "C",
			},
		},
	}

	cursor := 0
	state := startState
	tape := map[int]int{}
	minCursor := 0
	maxCursor := 0
	for i := 0; i < stepsToRun; i++ {
		v, ok := tape[cursor]
		if !ok {
			v = 0
		}

		u := p[state][v]
		tape[cursor] = u.writeValue
		if u.moveRight {
			cursor++
			if cursor > maxCursor {
				maxCursor = cursor
			}
		} else {
			cursor--
			if cursor < minCursor {
				minCursor = cursor
			}
		}
		state = u.nextState

		/*		fmt.Printf("Tape: ")
				for i := minCursor; i <= maxCursor; i++ {
					fmt.Printf("%d ", tape[i])
				}
				fmt.Printf("\n")*/
	}

	checksum := 0
	for _, v := range tape {
		checksum += v
	}

	fmt.Printf("Finished, checksum is: %d\n", checksum)
}
