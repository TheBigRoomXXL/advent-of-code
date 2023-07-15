package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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

	// Apply the list of instruction on the stack (the apply mecanism change if it's part 1 or 2)
	part := flag.Int("part", 0, "part number (1 or 2)")
	flag.Parse()
	switch *part {
	case 1:
		for _, instruct := range instructions {
			apply_9000(instruct, stacks)
		}
	case 2:
		for _, instruct := range instructions {
			apply_9001(instruct, stacks)
		}
	default:
		panic("invalid part")
	}

	//Get the front if each stack
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

func parseInstruction(line string) Instruction {
	cleanup := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(line, "move ", ""),
			"from ", ""),
		"to ", "")
	data := strings.Split(cleanup, " ")
	quantity, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	from, err := strconv.Atoi(data[1])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(data[2])
	if err != nil {
		panic(err)
	}
	return Instruction{quantity: quantity, from: from - 1, to: to - 1}
}

func apply_9000(instruct Instruction, stacks [9]*list.List) {
	for i := 0; i < instruct.quantity; i++ {
		crate := stacks[instruct.from].Remove(stacks[instruct.from].Front())
		stacks[instruct.to].PushFront(crate)
	}
}

func apply_9001(instruct Instruction, stacks [9]*list.List) {
	crates := []any{}
	for i := 0; i < instruct.quantity; i++ {
		crate := stacks[instruct.from].Remove(stacks[instruct.from].Front())
		crates = append([]any{crate}, crates...) //prepend
	}
	for _, crate := range crates {
		stacks[instruct.to].PushFront(crate)
	}
}
