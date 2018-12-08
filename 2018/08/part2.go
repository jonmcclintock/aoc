package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Node is a node.
type Node struct {
	id       int
	children []*Node
	metadata []int
}

func makeNode(input []int, pos int, allNodes map[int]*Node) (*Node, int) {
	if len(input) < pos+1 {
		log.Fatalf("Tried to parse a node starting at %d when input only has %d", pos, len(input))
	}

	cCount, mCount := input[pos], input[pos+1]
	pos += 2

	n := &Node{
		id:       len(allNodes),
		children: make([]*Node, cCount),
		metadata: make([]int, mCount),
	}
	allNodes[n.id] = n

	for i := 0; i < cCount; i++ {
		n.children[i], pos = makeNode(input, pos, allNodes)
	}

	if pos+mCount > len(input) {
		log.Fatalf("Not enough room to read metadata (need %d, have %d)\n", pos+mCount, len(input))
	}

	for i := 0; i < mCount; i++ {
		n.metadata[i] = input[pos]
		pos++
	}

	return n, pos
}

func calculateValue(n *Node) int {
	value := 0

	if len(n.children) == 0 {
		for _, v := range n.metadata {
			value += v
		}
		return value
	}

	for _, v := range n.metadata {
		if v > len(n.children) {
			continue
		}
		value += calculateValue(n.children[v-1])
	}

	return value
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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputStrings := strings.Split(scanner.Text(), " ")
	input := make([]int, len(inputStrings))
	for i, s := range inputStrings {
		input[i], _ = strconv.Atoi(s)
	}

	nodes := make(map[int]*Node)
	root, _ := makeNode(input, 0, nodes)

	mSum := 0
	for _, n := range nodes {
		for _, v := range n.metadata {
			mSum += v
		}
	}

	fmt.Printf("%d nodes found, metadata sum is %d\n", len(nodes), mSum)
	fmt.Printf("Value of root is %d\n", calculateValue(root))
}
