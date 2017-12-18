package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day18-in.txt"

type instruction struct {
	opcode  string
	target  byte
	operand byte
	value   int
}

func loadProgram() []instruction {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	program := []instruction{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		i := instruction{
			opcode: s[0],
			target: s[1][0],
		}

		if len(s) == 3 {
			if s[2][0] >= 'a' && s[2][0] <= 'z' {
				i.operand = s[2][0]
			} else {
				i.value, _ = strconv.Atoi(s[2])
			}
		}

		program = append(program, i)
	}

	return program
}

func getOperand(i instruction, registers []int) int {
	val := i.value
	if i.operand != 0 {
		val = registers[i.operand]
	}
	return val
}

func main() {
	program := loadProgram()
	registers := make([]int, 256)
	pc := 0
	freq := 0

loop:
	for pc >= 0 && pc < len(program) {
		i := program[pc]

		switch i.opcode {
		case "snd":
			freq = registers[i.target]
			fmt.Printf("snd %c, freq: %d\n", i.target, freq)

		case "set":
			registers[i.target] = getOperand(i, registers)
			fmt.Printf("set %c to %d\n", i.target, registers[i.target])

		case "add":
			registers[i.target] += getOperand(i, registers)
			fmt.Printf("add %c to get %d\n", i.target, registers[i.target])

		case "mul":
			registers[i.target] *= getOperand(i, registers)
			fmt.Printf("mul %c to get %d\n", i.target, registers[i.target])

		case "mod":
			registers[i.target] %= getOperand(i, registers)
			fmt.Printf("mod %c to get %d\n", i.target, registers[i.target])

		case "rcv":
			if registers[i.target] != 0 {
				fmt.Printf("recv %c with %d, breaking\n", i.target, registers[i.target])
				break loop
			}
			fmt.Printf("recv %c is zero, skipping\n", i.target)

		case "jgz":
			if registers[i.target] > 0 {
				pc += getOperand(i, registers)
				fmt.Printf("jgz %c is %d, setting pc to %d\n", i.target, registers[i.target], pc)
				continue
			}
			fmt.Printf("jgz %c is %d, skipping\n", i.target, registers[i.target])
		}

		pc++
	}

	fmt.Printf("Finished with PC %d, freq is %d\n", pc, freq)
}
