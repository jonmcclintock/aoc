package main

import "fmt"

const target = 325489

type direction struct {
	X int // X increment per step
	Y int // Y increment per step
	D int // Distance increment per direction
}

func main() {
	sequence := []direction{
		direction{
			X: 1,
			Y: 0,
			D: 0,
		},
		direction{
			X: 0,
			Y: 1,
			D: 1,
		},
		direction{
			X: -1,
			Y: 0,
			D: 0,
		},
		direction{
			X: 0,
			Y: -1,
			D: 1,
		},
	}

	s := 0
	c := 0
	d := direction{
		X: 0,
		Y: 0,
		D: 1,
	}
	for i := 0; i < target-1; i++ {
		//fmt.Printf("Before:  I: %d, X: %d, Y: %d, S: %d, D: %d\n", i, d.X, d.Y, s, d.D)
		d.X += sequence[s].X
		d.Y += sequence[s].Y
		c++
		if c == d.D {
			d.D += sequence[s].D
			s = (s + 1) % len(sequence)
			c = 0
		}
	}

	fmt.Printf("X: %d, Y: %d\n", d.X, d.Y)
}
