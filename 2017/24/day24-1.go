package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day24-in.txt"

type component struct {
	a int
	b int
}

type componentList map[int][]component

func loadComponents() componentList {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	comps := componentList{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "/")
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])
		c := component{a: a, b: b}
		if _, ok := comps[a]; ok {
			comps[a] = append(comps[a], c)
		} else {
			comps[a] = []component{c}
		}
		if _, ok := comps[b]; ok {
			comps[b] = append(comps[b], c)
		} else {
			comps[b] = []component{c}
		}
	}

	return comps
}

func findMaxBridge(comps componentList, pinCount int, curPath []component) ([]component, int) {
	used := map[component]bool{}
	for _, c := range curPath {
		used[c] = true
	}

	testPath := make([]component, len(curPath)+1)
	copy(testPath, curPath)

	maxPath := curPath
	maxStrength := bridgeStrength(curPath)

	//dumpBridge(curPath, maxStrength)

	for _, c := range comps[pinCount] {
		if skip, ok := used[c]; ok && skip {
			continue
		}

		var testCount int
		if c.a == pinCount {
			testCount = c.b
		} else {
			testCount = c.a
		}
		testPath[len(curPath)] = c
		path, strength := findMaxBridge(comps, testCount, testPath)
		if strength > maxStrength {
			maxStrength = strength
			maxPath = path
		}

	}

	return maxPath, maxStrength
}

func bridgeStrength(bridge []component) int {
	strength := 0
	for _, c := range bridge {
		strength += c.a + c.b
	}
	return strength
}

func dumpBridge(bridge []component, strength int) {
	fmt.Printf("Bridge has strength %d and is %d components long: ", strength, len(bridge))
	for _, c := range bridge {
		fmt.Printf("%d/%d, ", c.a, c.b)
	}
	fmt.Printf("\n")
}

func main() {
	comps := loadComponents()
	fmt.Printf("Comps: %v\n", comps)

	bridge, strength := findMaxBridge(comps, 0, []component{})

	fmt.Printf("Done!\n\n")
	dumpBridge(bridge, strength)
}
