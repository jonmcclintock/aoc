package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day13-in.txt"

func loadLayers() []int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	layers := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ": ")
		d, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])
		if d >= len(layers) {
			layers = append(layers, make([]int, d-len(layers)+1)...)
		}
		layers[d] = r
	}

	return layers
}

func stepPositions(layers, positions []int) {
	for j := range positions {
		if layers[j] == 0 {
			continue
		}
		if positions[j] >= layers[j]-1 {
			positions[j] = 1 - layers[j]
		}
		positions[j]++
	}
}

func main() {
	layers := loadLayers()
	positions := make([]int, len(layers))
	sums := make([]int, len(layers))
	caught := make([]bool, len(layers))

	step := 0
	for {
		for i := 0; i < len(layers) && i <= step; i++ {
			pos := (step - i) % len(layers)
			if layers[i] != 0 && positions[i] == 0 {
				sums[pos] += i * layers[i]
				caught[pos] = true
			}
		}

		step++
		if step >= len(layers) {
			if !caught[step%len(caught)] {
				fmt.Printf("Got all the way through with delay %d\n", step-len(layers))
				break
			}
			sums[step%len(layers)] = 0
			caught[step%len(layers)] = false
		}

		stepPositions(layers, positions)
	}
}
