package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type cart struct {
	x, y    int
	d       byte
	t       int
	crashed bool
}

var turns = []map[byte]byte{
	map[byte]byte{ // Turn left
		'>': '^',
		'^': '<',
		'<': 'v',
		'v': '>',
	},
	map[byte]byte{ // Go straight
		'>': '>',
		'^': '^',
		'<': '<',
		'v': 'v',
	},
	map[byte]byte{ // Turn right
		'>': 'v',
		'^': '>',
		'<': '^',
		'v': '<',
	},
}

// Maps current direction to a map of current character to a delta cart
var moves = map[byte]map[byte]cart{
	'>': map[byte]cart{
		'-': cart{
			x: 1,
			y: 0,
			d: '>',
		},
		'+': cart{
			x: 1,
			y: 0,
			d: '>',
		},
		'/': cart{
			x: 1,
			y: 0,
			d: '^',
		},
		'\\': cart{
			x: 1,
			y: 0,
			d: 'v',
		},
	},
	'<': map[byte]cart{
		'-': cart{
			x: -1,
			y: 0,
			d: '<',
		},
		'+': cart{
			x: -1,
			y: 0,
			d: '<',
		},
		'/': cart{
			x: -1,
			y: 0,
			d: 'v',
		},
		'\\': cart{
			x: -1,
			y: 0,
			d: '^',
		},
	},
	'v': map[byte]cart{
		'|': cart{
			x: 0,
			y: 1,
			d: 'v',
		},
		'+': cart{
			x: 0,
			y: 1,
			d: 'v',
		},
		'/': cart{
			x: 0,
			y: 1,
			d: '<',
		},
		'\\': cart{
			x: 0,
			y: 1,
			d: '>',
		},
	},
	'^': map[byte]cart{
		'|': cart{
			x: 0,
			y: -1,
			d: '^',
		},
		'+': cart{
			x: 0,
			y: -1,
			d: '^',
		},
		'/': cart{
			x: 0,
			y: -1,
			d: '>',
		},
		'\\': cart{
			x: 0,
			y: -1,
			d: '<',
		},
	},
}

func dumpCarts(carts []*cart) {
	fmt.Printf("Carts:\n")
	for i, _ := range carts {
		fmt.Printf("%+v\n", carts[i])
	}
	fmt.Printf("---\n")
}

func dumpTrack(trackMap [][]byte, carts []*cart) {
	count := 0
	for y, row := range trackMap {
		fmt.Printf("%03d #", y)
	cell:
		for x, col := range row {
			for _, c := range carts {
				if c.crashed {
					continue
				}
				if c.x == x && c.y == y {
					fmt.Printf("%c", c.d)
					count++
					continue cell
				}
			}
			fmt.Printf("%c", col)
		}
		fmt.Printf("#\n")
	}
	fmt.Printf("%d carts left\n", count)
	fmt.Printf("\n")
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

	width := 0
	carts := make([]*cart, 0)
	trackMap := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		if width == 0 {
			width = len(row)
		} else if len(row) != width {
			log.Fatalf("Rows aren't equal width (got %d, previously %d)\n", len(row), width)
		}

		for i := 0; i < len(row); i++ {
			if row[i] == '>' || row[i] == '<' || row[i] == '^' || row[i] == 'v' {
				c := cart{
					x: i,
					y: len(trackMap),
					d: row[i],
				}
				carts = append(carts, &c)

				if row[i] == '<' || row[i] == '>' {
					row[i] = '-'
				} else if row[i] == 'v' || row[i] == '^' {
					row[i] = '|'
				}
			}
		}
		trackMap = append(trackMap, row)

	}

	dumpCarts(carts)
	for cartCount := len(carts); cartCount > 1; {
		//dumpTrack(trackMap, carts)

		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y < carts[j].y {
				return true
			}
			if carts[i].y > carts[j].y {
				return false
			}
			return carts[i].x < carts[j].x
		})

		for _, c := range carts {
			if c.crashed {
				continue
			}

			m, ok := moves[c.d]
			if !ok {
				log.Fatalf("No move for cart char '%c'\n", c.d)
			}

			delta, ok := m[trackMap[c.y][c.x]]
			if !ok {
				log.Fatalf("No move for map char '%c'\n", trackMap[c.y][c.x])
			}

			c.x += delta.x
			c.y += delta.y

			if trackMap[c.y][c.x] == '+' {
				c.d = turns[c.t][c.d]
				c.t = (c.t + 1) % len(turns)
			} else {
				delta, ok = m[trackMap[c.y][c.x]]
				if !ok {
					log.Fatalf("No move for map char '%c'\n", trackMap[c.y][c.x])
				}

				c.d = delta.d
			}

			for i := range carts {
				if carts[i].crashed {
					continue
				}
				for j := i + 1; j < len(carts); j++ {
					if carts[j].crashed {
						continue
					}
					if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
						fmt.Printf("Got a collision at %d,%d\n", carts[i].x, carts[i].y)
						carts[i].crashed = true
						carts[j].crashed = true
						cartCount -= 2
						//dumpTrack(trackMap, carts)
					}
				}
			}
		}
	}
	dumpTrack(trackMap, carts)

	var last *cart
	for _, c := range carts {
		if !c.crashed {
			last = c
			break
		}
	}
	fmt.Printf("Last cart is at %d,%d, facing %c\n", last.x, last.y, last.d)
}
