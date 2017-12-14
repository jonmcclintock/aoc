package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day8-in.txt"

type instruction struct {
	target     string
	increment  bool
	incValue   int
	operand    string
	comparator string
	compValue  int
}

func parseInstruction(line string) instruction {
	fields := strings.Fields(line)

	increment := true
	if fields[1] == "dec" {
		increment = false
	}

	incValue, _ := strconv.Atoi(fields[2])
	compValue, _ := strconv.Atoi(fields[6])
	return instruction{
		target:     fields[0],
		increment:  increment,
		incValue:   incValue,
		operand:    fields[4],
		comparator: fields[5],
		compValue:  compValue,
	}
}

func loadInstructions() []instruction {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := make([]instruction, 0)
	for scanner.Scan() {
		inst := parseInstruction(scanner.Text())
		instructions = append(instructions, inst)
	}

	return instructions
}

func getRegister(regs map[string]int, name string) int {
	val, ok := regs[name]
	if !ok {
		return 0
	}
	return val
}

func setRegister(regs map[string]int, name string, val int) {
	regs[name] = val
}

func main() {
	instructions := loadInstructions()
	regs := map[string]int{}

	overallMaxReg := ""
	overallMaxValue := 0
	for _, inst := range instructions {
		opVal := getRegister(regs, inst.operand)
		compResult := false
		switch inst.comparator {
		case "<":
			compResult = opVal < inst.compValue
		case ">":
			compResult = opVal > inst.compValue
		case "==":
			compResult = opVal == inst.compValue
		case "!=":
			compResult = opVal != inst.compValue
		case "<=":
			compResult = opVal <= inst.compValue
		case ">=":
			compResult = opVal >= inst.compValue
		default:
			fmt.Printf("Unknown comparitor '%s'\n", inst.comparator)
		}

		if !compResult {
			continue
		}

		value := getRegister(regs, inst.target)
		if inst.increment {
			value = value + inst.incValue
		} else {
			value = value - inst.incValue
		}
		setRegister(regs, inst.target, value)

		if value > overallMaxValue {
			overallMaxValue = value
			overallMaxReg = inst.target
		}
	}

	maxReg := ""
	maxValue := 0
	for k, v := range regs {
		if v > maxValue {
			maxReg = k
			maxValue = v
		}
	}

	fmt.Printf("Max value is %d in register %s\n", maxValue, maxReg)
	fmt.Printf("Overall max value was %d in register %s\n", overallMaxValue, overallMaxReg)
}
