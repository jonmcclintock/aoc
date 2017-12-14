package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day12-in.txt"

type node struct {
	id      int
	peers   map[int]*node
	visited bool
}

func parseNode(line string) (int, []int) {
	s1 := strings.Split(line, " <-> ")
	id, _ := strconv.Atoi(s1[0])

	s2 := strings.Split(s1[1], ", ")
	peers := make([]int, len(s2))
	for i, peer := range s2 {
		peers[i], _ = strconv.Atoi(peer)
	}

	return id, peers
}

func makeNode(nodes map[int]*node, id int) {
	if _, ok := nodes[id]; !ok {
		nodes[id] = &node{
			id:    id,
			peers: map[int]*node{},
		}
	}
}

func setPeer(nodes map[int]*node, a, b int) {
	makeNode(nodes, a)
	makeNode(nodes, b)
	nodes[a].peers[b] = nodes[b]
	nodes[b].peers[a] = nodes[a]
}

func loadNodes() map[int]*node {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := map[int]*node{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id, peers := parseNode(scanner.Text())
		for _, peer := range peers {
			setPeer(nodes, id, peer)
		}
	}

	return nodes
}

func countReachableNodes(cur *node) int {
	if cur.visited {
		return 0
	}
	cur.visited = true

	sum := 1
	for _, n := range cur.peers {
		sum += countReachableNodes(n)
	}

	return sum
}

func main() {
	nodes := loadNodes()

	groupCount := 0
	for _, cur := range nodes {
		if cur.visited {
			continue
		}
		groupCount++
		n := countReachableNodes(cur)
		fmt.Printf("Reachable nodes from %d: %d\n", cur.id, n)
	}

	fmt.Printf("%d groups found\n", groupCount)
}
