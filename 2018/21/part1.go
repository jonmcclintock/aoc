package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jonmcclintock/aoc/2018/19/ops"
)

const registerCount = 6

type Instruction struct {
	operation ops.Operation
	params    []int
}

func parseInstructionPointer(line string) int {
	ip, err := strconv.Atoi(strings.TrimPrefix(line, "#ip "))
	if err != nil {
		log.Fatalf("Invalid instruction pointer: %s\n", ip)
	}
	return ip
}

func parseOperation(line string) Instruction {
	strs := strings.Split(line, " ")
	operation, ok := ops.OpCodeNames[strs[0]]
	if !ok {
		log.Fatalf("Invalid opcode '%s'\n", strs[0])
	}

	params := make([]int, len(strs)-1)
	for i, str := range strs[1:] {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failure parsing register value '%s': %v", str, err)
		}

		params[i] = val
	}

	return Instruction{operation: operation, params: params}
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s input-file reg-0-value", os.Args[0])
	}

	reg0, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ipRegister := parseInstructionPointer(scanner.Text())

	program := make([]Instruction, 0)
	for scanner.Scan() {
		program = append(program, parseOperation(scanner.Text()))
	}

	registers := make([]int, registerCount)
	registers[0] = reg0

	ip := 0
	for ip < len(program) {
		registers[ipRegister] = ip
		fmt.Printf("IP %3d %v: %s %+v\n", ip, registers, program[ip].operation.String(), program[ip].params)
		registers = program[ip].operation.Execute(program[ip].params, registers)
		ip = registers[ipRegister] + 1
	}

	fmt.Printf("Program ended with instruction pointer at %d\n", ip)
	fmt.Printf("Registers: %+v\n", registers)
}
