package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"os"
)

func main() {
	stacks, instructions := parseInput()
	part := flag.Int("part", 0, "part number (1 or 2)")
	flag.Parse()
	switch *part {
	case 1:
		apply_9000(stacks, instructions)
	case 2:
		apply_9001(stacks, instructions)
	default:
		panic("invalid part")
	}

	//Get the front of each stack
	for _, stack := range stacks {
		fmt.Printf("%c", stack.Front().Value)
	}
	fmt.Println()
}

type Instruction struct {
	quantity int
	from     int
	to       int
}

func parseInput() ([9]*list.List, []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)
	stackDiagram := []string{}

	// First parse the stack diagram
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		stackDiagram = append(stackDiagram, scanner.Text())
	}
	stacks := [9]*list.List{}
	for i, _ := range stacks {
		stacks[i] = list.New()
	}

	// Read it in reverse and skip the line with numbers
	for i := len(stackDiagram) - 2; i >= 0; i-- {
		line := stackDiagram[i]
		for j := 0; j < 9; j++ {
			if line[4*j+1] != ' ' {
				stacks[j].PushFront(line[4*j+1])
			}
		}
	}

	//Then parse the instruction
	instructions := []Instruction{}
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	return stacks, instructions
}

func parseInstruction(line string) Instruction {
	instru := Instruction{}
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &instru.quantity, &instru.from, &instru.to)
	if err != nil {
		panic(err)
	}
	// subtract one so they're zero indexed...
	instru.from--
	instru.to--
	return instru
}

func apply_9000(stacks [9]*list.List, instructions []Instruction) {
	for _, instru := range instructions {
		for i := 0; i < instru.quantity; i++ {
			crate := stacks[instru.from].Remove(stacks[instru.from].Front())
			stacks[instru.to].PushFront(crate)
		}
	}
}

func apply_9001(stacks [9]*list.List, instructions []Instruction) {
	for _, instru := range instructions {
		crates := []any{}
		for i := 0; i < instru.quantity; i++ {
			crate := stacks[instru.from].Remove(stacks[instru.from].Front())
			crates = append([]any{crate}, crates...) //prepend
		}
		for _, crate := range crates {
			stacks[instru.to].PushFront(crate)
		}
	}
}
