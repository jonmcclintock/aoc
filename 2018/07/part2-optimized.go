package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	if len(os.Args[1:]) != 3 {
		log.Fatalf("Usage: %s input-file worker-count task-time", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	workerCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid worker count", err)
	}

	taskTime, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid task time", err)
	}

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

	var result bytes.Buffer
	curTime := 0
	workerTimes := make([]int, workerCount)
	workerTasks := make([]*Step, workerCount)
	for {
		//fmt.Printf("Time %03d: ", curTime)
		curTime++

		busy := false
		full := true
		for i := 0; i < workerCount; i++ {
			if workerTimes[i] == 0 {
				//fmt.Printf(".    ")
				full = false
				continue
			}
			//fmt.Printf("%s    ", workerTasks[i].id)
			busy = true
			workerTimes[i]--
			if workerTimes[i] == 0 {
				s := workerTasks[i]
				result.WriteString(s.id)
				for k := range s.next {
					delete(steps[k].prev, s.id)
				}
			}
		}
		//fmt.Printf("\n")

		if !busy && len(steps) == 0 {
			break
		}

		if full {
			continue
		}

		noPrevs := make([]*Step, 0)
		for _, s := range steps {
			if len(s.prev) == 0 {
				noPrevs = append(noPrevs, s)
			}
		}

		sort.Slice(noPrevs[:], func(i, j int) bool {
			return noPrevs[i].id < noPrevs[j].id
		})

		for _, s := range noPrevs {
			worker := -1
			for j := 0; j < workerCount; j++ {
				if workerTimes[j] == 0 {
					worker = j
				}
			}
			if worker == -1 {
				break
			}
			workerTasks[worker] = s
			workerTimes[worker] = int(s.id[0]-'A') + taskTime + 1

			delete(steps, s.id)
		}
	}

	fmt.Printf("Steps to completion: %d\n", curTime-1)
	fmt.Println(result.String())
}
