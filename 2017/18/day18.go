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
const numProcesses = 2

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

func step(program []instruction, proc *process) bool {
	i := program[proc.pc]

	switch i.opcode {
	case "set":
		proc.setValue(i.target, proc.getValue(i.operand))

	case "add":
		proc.setValue(i.target, proc.getValue(i.target)+proc.getValue(i.operand))

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
	}

	proc.pc++
	return false
}

func main() {
	program := loadProgram()
	processes := make([]process, numProcesses)

	for i := 0; i < numProcesses; i++ {
		processes[i] = process{
			pid:       i,
			registers: make([]int, 26),
			pc:        0,
			queue:     []int{},
		}
		processes[i].setValue("p", i)
	}
	processes[0].receiver = &processes[1]
	processes[1].receiver = &processes[0]

loop:
	for {
		deadlock := true
		for i := 0; i < numProcesses; i++ {
			if step(program, &processes[i]) == false {
				deadlock = false
			}
			if processes[i].pc < 0 || processes[i].pc >= len(program) {
				fmt.Printf("[%d] PC out of bounds: %d\n", processes[i].pid, processes[i].pc)
				break loop
			}
		}
		if deadlock == true {
			fmt.Printf("Deadlocked.\n")
			break loop
		}
	}
	for i := 0; i < numProcesses; i++ {
		fmt.Printf("[%d]: pc %d, sent %d\n", i, processes[i].pc, processes[i].sendCount)
	}
}
