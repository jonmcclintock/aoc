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

func testLayers(layers []int, delay int) (bool, int) {
	positions := make([]int, len(layers))

	for i := 0; i < delay; i++ {
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

	sum := 0
	caught := false
	for i := 0; i < len(layers); i++ {
		if layers[i] != 0 && positions[i] == 0 {
			sum += i * layers[i]
			fmt.Printf("Caught on layer %d, sum is now %d\n", i, sum)
			caught = true
		}
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

	return caught, sum
}

func main() {
	layers := loadLayers()

	delay := 0
	for {
		caught, sum := testLayers(layers, delay)
		fmt.Printf("Total severity with delay %d is %d\n", delay, sum)
		if !caught {
			break
		}
		delay++
	}
}
