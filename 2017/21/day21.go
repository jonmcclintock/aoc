package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "day21-in.txt"
const seed = ".#./..#/###"
const rounds = 18

type ruleSet map[int][][]byte

func parsePattern(in string) [][]byte {
	s := strings.Split(in, "/")
	res := make([][]byte, len(s))
	for i := range res {
		res[i] = make([]byte, len(res))
	}
	for i := range res {
		for j, b := range []byte(s[i]) {
			if b == '#' {
				res[j][i] = 1
			}
		}
	}

	return res
}

func permutePattern(in [][]byte) []int {
	res := make([]int, 8)
	l := len(in)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			res[0] <<= 1
			res[0] += int(in[j][i])
			res[1] <<= 1
			res[1] += int(in[j][l-i-1])
			res[2] <<= 1
			res[2] += int(in[l-j-1][i])
			res[3] <<= 1
			res[3] += int(in[l-j-1][l-i-1])

			res[4] <<= 1
			res[4] += int(in[i][j])
			res[5] <<= 1
			res[5] += int(in[i][l-j-1])
			res[6] <<= 1
			res[6] += int(in[l-i-1][j])
			res[7] <<= 1
			res[7] += int(in[l-i-1][l-j-1])
		}
	}

	return res
}

func extractFrom(in [][]byte, x, y, size int) int {
	res := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res <<= 1
			res += int(in[j+x][i+y])
		}
	}

	return res
}

func copyInto(out, pattern [][]byte, x, y int) {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern); j++ {
			out[x+j][y+i] = pattern[j][i]
		}
	}
}

func dump(grid [][]byte) {
	fmt.Printf("Grid (%d,%d):\n", len(grid), len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func loadRules() (ruleSet, ruleSet) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules2 := ruleSet{}
	rules3 := ruleSet{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " => ")
		in := parsePattern(s[0])
		out := parsePattern(s[1])

		for _, r := range permutePattern(in) {
			if len(in) == 2 {
				rules2[r] = out
			} else {
				rules3[r] = out
			}
		}
	}

	return rules2, rules3
}

func main() {
	rules2, rules3 := loadRules()
	cur := parsePattern(seed)
	for i := 0; i < rounds; i++ {
		fmt.Printf("Round %d, input is %dx%d\n", i, len(cur), len(cur))
		//dump(cur)

		var is, os int
		var r ruleSet
		var next [][]byte
		if (len(cur) % 2) == 0 {
			is = 2
			os = 3
			r = rules2
		} else {
			is = 3
			os = 4
			r = rules3
		}
		bc := len(cur) / is

		next = make([][]byte, bc*os)
		for i := range next {
			next[i] = make([]byte, bc*os)
		}

		for i := 0; i < bc; i++ {
			for j := 0; j < bc; j++ {
				extract := extractFrom(cur, j*is, i*is, is)
				pattern, ok := r[extract]
				if !ok {
					fmt.Printf("Couldn't find a pattern for %09b\n", extract)
				}
				copyInto(next, pattern, j*os, i*os)
			}
		}

		count := 0
		for i := 0; i < bc*os; i++ {
			for j := 0; j < bc*os; j++ {
				count += int(next[i][j])
			}
		}

		fmt.Printf("Output is %dx%d, with %d lit pixels\n", len(next), len(next), count)
		cur = next
	}

	//dump(cur)
}
