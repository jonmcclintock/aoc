package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day23-in.txt"

type instruction struct {
	opcode  string
	target  string
	operand string
}

type process struct {
	pid       int
	pc        int
	registers []int
	queue     []int
	receiver  *process
	sendCount int
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
			target: s[1],
		}
		if len(s) >= 3 {
			i.operand = s[2]
		}

		program = append(program, i)
	}

	return program
}

func (p process) getValue(s string) int {
	if s[0] >= 'a' && s[0] <= 'z' {
		return p.registers[s[0]-'a']
	}
	val, _ := strconv.Atoi(s)
	return val
}

func (p *process) setValue(s string, v int) {
	if s[0] >= 'a' && s[0] <= 'z' {
		p.registers[s[0]-'a'] = v
	}
}

func step(program []instruction, proc *process, stats map[string]int) bool {
	i := program[proc.pc]

	if _, ok := stats[i.opcode]; !ok {
		stats[i.opcode] = 1
	} else {
		stats[i.opcode]++
	}

	switch i.opcode {
	case "set":
		proc.setValue(i.target, proc.getValue(i.operand))

	case "add":
		proc.setValue(i.target, proc.getValue(i.target)+proc.getValue(i.operand))

	case "sub":
		proc.setValue(i.target, proc.getValue(i.target)-proc.getValue(i.operand))

	case "mul":
		proc.setValue(i.target, proc.getValue(i.target)*proc.getValue(i.operand))

	case "mod":
		proc.setValue(i.target, proc.getValue(i.target)%proc.getValue(i.operand))

	case "snd":
		proc.sendCount++
		proc.receiver.queue = append(proc.receiver.queue, proc.getValue(i.target))

	case "rcv":
		if len(proc.queue) == 0 {
			return true
		}
		proc.setValue(i.target, proc.queue[0])
		proc.queue = proc.queue[1:]

	case "jgz":
		if proc.getValue(i.target) > 0 {
			proc.pc += proc.getValue(i.operand)
			return false
		}

	case "jnz":
		if proc.getValue(i.target) != 0 {
			proc.pc += proc.getValue(i.operand)
			return false
		}
	}

	proc.pc++
	return false
}

func main() {
	program := loadProgram()
	proc := process{
		pid:       0,
		registers: make([]int, 8),
		pc:        0,
		queue:     []int{},
	}
	stats := map[string]int{}
	for {
		step(program, &proc, stats)
		if proc.pc < 0 || proc.pc >= len(program) {
			fmt.Printf("[%d] PC out of bounds: %d\n", proc.pid, proc.pc)
			break
		}
	}

	fmt.Printf("Stats: %v\n", stats)
}
