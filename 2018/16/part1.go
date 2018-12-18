package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jonmcclintock/aoc/2018/16/ops"
)

// How many operations have to match to count.
var matchThreshold = 3

func parseRegisters(line string) []int {
	line = strings.TrimPrefix(line, "Before: [")
	line = strings.TrimPrefix(line, "After:  [")
	line = strings.TrimSuffix(line, "]")

	regStrs := strings.Split(line, ", ")
	regs := make([]int, len(regStrs))
	for i, str := range regStrs {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failure parsing register value '%s': %v", str, err)
		}

		regs[i] = val
	}

	return regs
}

func parseOperation(line string) []int {
	strs := strings.Split(line, " ")
	operation := make([]int, len(strs))
	for i, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failure parsing register value '%s': %v", str, err)
		}

		operation[i] = val
	}

	return operation
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

	totalCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		beforeRegs := parseRegisters(scanner.Text())
		scanner.Scan()
		operation := parseOperation(scanner.Text())
		scanner.Scan()
		afterRegs := parseRegisters(scanner.Text())
		scanner.Scan()

		matchCount := 0
		for _, op := range ops.AllOps {
			resultRegs := op.Execute(operation[1:], beforeRegs)

			match := true
			for i := range resultRegs {
				if resultRegs[i] != afterRegs[i] {
					match = false
				}
			}

			if match {
				matchCount++
			}
		}

		if matchCount >= matchThreshold {
			totalCount++
		}
	}

	fmt.Printf("Total count of operations with %d or more opcode matches: %d\n", matchThreshold, totalCount)
}
