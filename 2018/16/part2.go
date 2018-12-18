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

// How many registers there are.
var regCount = 4

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

func dumpOperationCount(matchCounts map[ops.Operation]map[int]int) {
	fmt.Printf("   #  ")
	for code := 0; code < 16; code++ {
		fmt.Printf("   %2d", code)
	}
	fmt.Printf("\n")

	for _, op := range ops.AllOps {
		freq := matchCounts[op]
		fmt.Printf("%s: ", op.String())
		for code := 0; code < 16; code++ {
			fmt.Printf(" %4d", freq[code])
		}
		fmt.Printf("\n")
	}
}

func runProgram(inputFile string, opcodes map[int]ops.Operation) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	regs := make([]int, regCount)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		operation := parseOperation(scanner.Text())
		regs = opcodes[operation[0]].Execute(operation[1:], regs)
	}

	fmt.Printf("After program completes, registers are: %+v\n", regs)
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s input-operations input-program", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalCount := 0
	matchCounts := make(map[ops.Operation]map[int]int, 0)

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

			if !match {
				continue
			}

			matchCount++
			code := operation[0]
			if _, ok := matchCounts[op]; !ok {
				matchCounts[op] = make(map[int]int)
			}

			if _, ok := matchCounts[op][code]; !ok {
				matchCounts[op][code] = 1
			} else {
				matchCounts[op][code]++
			}
		}

		if matchCount >= matchThreshold {
			totalCount++
		}
	}

	fmt.Printf("Total count of operations with %d or more opcode matches: %d\n", matchThreshold, totalCount)

	//dumpOperationCount(matchCounts)

	// Now try to do the mapping.
	opcodes := make(map[int]ops.Operation, 0)
	for len(matchCounts) > 0 {
		// Look for an instruction with only one count.
		found := -1
		for i, op := range ops.AllOps {
			if len(matchCounts[op]) == 1 {
				found = i
				break
			}
		}

		if found == -1 {
			log.Fatalf("Couldn't find an opcode with only 1 match")
		}

		// The first and only key will be the one opcode.
		foundOp := ops.AllOps[found]
		var code int
		for k := range matchCounts[foundOp] {
			code = k
		}
		opcodes[code] = foundOp

		delete(matchCounts, foundOp)

		// Now remove the opcode from the other lists.
		for _, op := range ops.AllOps {
			if _, ok := matchCounts[op][code]; ok {
				delete(matchCounts[op], code)
			}
		}
	}

	fmt.Printf("Found opcodes:\n")
	for i := 0; i < len(opcodes); i++ {
		fmt.Printf("- %2d: %s\n", i, opcodes[i].String())
	}

	runProgram(os.Args[2], opcodes)
}
