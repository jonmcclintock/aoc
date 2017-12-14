package main

import "fmt"

const target = 325489

type direction struct {
	X int // X increment per step
	Y int // Y increment per step
	D int // Distance increment per direction
}

type memory struct {
	NN int // (-x, -y)
	NP int // (-x, +y)
	PN int // (+x, -y)
	PP int // (+x, +y)
}

func getMemory(m [][]memory, x, y int) int {
	ax := x
	ay := y
	if ax < 0 {
		ax = -ax
	}
	if ay < 0 {
		ay = -ay
	}

	if ax >= len(m) {
		return 0
	}
	if ay >= len(m[ax]) {
		return 0
	}

	if x < 0 {
		if y < 0 {
			return m[ax][ay].NN
		}
		return m[ax][ay].NP
	}
	if y < 0 {
		return m[ax][ay].PN
	}
	return m[ax][ay].PP
}

func setMemory(m [][]memory, x, y, v int) [][]memory {
	ax := x
	ay := y
	if ax < 0 {
		ax = -ax
	}
	if ay < 0 {
		ay = -ay
	}

	fmt.Printf("Setting cell %d, %d to %d\n", x, y, v)
	res := m
	if ax >= len(res) {
		res = append(res, make([][]memory, ax-len(res)+1)...)
	}
	if ay >= len(res[ax]) {
		res[ax] = append(res[ax], make([]memory, ay-len(res[ax])+1)...)
	}

	if x < 0 {
		if y < 0 {
			res[ax][ay].NN = v
			return res
		}
		res[ax][ay].NP = v
		return res
	}
	if y < 0 {
		res[ax][ay].PN = v
		return res
	}
	res[ax][ay].PP = v
	return res
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
	m := [][]memory{
		{
			memory{
				NN: 1,
				NP: 1,
				PN: 1,
				PP: 1,
			},
		},
	}

	for i := 0; i < 1000000; i++ {
		//fmt.Printf("Before:  I: %d, X: %d, Y: %d, S: %d, D: %d\n", i, d.X, d.Y, s, d.D)
		d.X += sequence[s].X
		d.Y += sequence[s].Y

		v := 0
		fmt.Printf("Neighbors of %d,%d: ", d.X, d.Y)
		for j := 0; j <= 2; j++ {
			for k := 0; k <= 2; k++ {
				fmt.Printf("%d(%d,%d) ", getMemory(m, d.X+j-1, d.Y+k-1), d.X+j-1, d.Y+k-1)
				v += getMemory(m, d.X+j-1, d.Y+k-1)
			}
		}
		fmt.Printf("\n")

		m = setMemory(m, d.X, d.Y, v)
		if v > target {
			fmt.Printf("X: %d, Y: %d, V: %d\n", d.X, d.Y, v)
			return
		}

		c++
		if c == d.D {
			d.D += sequence[s].D
			s = (s + 1) % len(sequence)
			c = 0
		}
	}
}
