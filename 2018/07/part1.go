package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Step is a step with next steps.
type Step struct {
	id   string
	next map[string]bool
	prev map[string]bool
}

func makeStep(id string) *Step {
	s := Step{id: id}
	s.next = make(map[string]bool)
	s.prev = make(map[string]bool)
	return &s
}

func getOrMakeStep(steps map[string]*Step, id string) (*Step, map[string]*Step) {
	if _, ok := steps[id]; !ok {
		steps[id] = makeStep(id)
	}

	return steps[id], steps
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

	steps := make(map[string]*Step)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimPrefix(line, "Step ")
		line = strings.TrimSuffix(line, " can begin.")

		vals := strings.Split(line, " must be finished before step ")
		if len(vals) != 2 {
			log.Fatalf("Invalid input: %s\n", line)
		}

		var s, n *Step
		s, steps = getOrMakeStep(steps, vals[0])
		n, steps = getOrMakeStep(steps, vals[1])

		s.next[n.id] = true
		n.prev[s.id] = true
	}

	for len(steps) > 0 {
		noPrevs := make([]*Step, 0)
		for _, s := range steps {
			if len(s.prev) == 0 {
				noPrevs = append(noPrevs, s)
			}
		}

		sort.Slice(noPrevs[:], func(i, j int) bool {
			return noPrevs[i].id < noPrevs[j].id
		})

		s := noPrevs[0]
		fmt.Printf("%s", s.id)

		for k := range s.next {
			delete(steps[k].prev, s.id)
		}
		delete(steps, s.id)
	}

	fmt.Printf("\n")
}
