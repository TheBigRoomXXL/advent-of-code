package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var result string
	part := flag.Int("part", 0, "part number (1 or 2)")
	flag.Parse()
	switch *part {
	case 1:
		result = part1(os.Stdin)
	case 2:
		result = part2(os.Stdin)
	default:
		panic("invalid part")
	}
	fmt.Println(result)
}

func part1(input io.Reader) string {
	instrus := parseInput(input)
	n := 0
	x := 1
	total := 0
	hasProccessed := false
	for len(instrus) > 0 {
		n++
		if n%40 == 20 {
			total += x * n
		}
		if instrus[0].cmd == "noop" {
			instrus = instrus[1:]
		} else {
			if hasProccessed {
				x += instrus[0].val
				hasProccessed = false
				instrus = instrus[1:]
			} else {
				hasProccessed = true
			}
		}
	}
	return fmt.Sprint(total)
}

func part2(input io.Reader) string {
	instrus := parseInput(input)
	n := 0
	x := 1
	hasProccessed := false
	for len(instrus) > 0 {
		n++
		draw(n, x)

		if instrus[0].cmd == "noop" {
			instrus = instrus[1:]
		} else {
			if hasProccessed {
				x += instrus[0].val
				hasProccessed = false
				instrus = instrus[1:]
			} else {
				hasProccessed = true
			}
		}
		if n%40 == 0 {
			fmt.Println()
			n = 0
		}
	}
	return ""
}

func draw(n int, x int) {
	if n >= x && n <= x+2 { // ⚠ n start at 1, not 0
		fmt.Print("#")
	} else {
		fmt.Print(".")
		// fmt.Printf("%d %d\n", n, x)
	}
}

type Instruction struct {
	cmd string
	val int
}

func parseInput(input io.Reader) []Instruction {
	scanner := bufio.NewScanner(input)
	commands := []Instruction{}
	for scanner.Scan() {
		var val int
		cmd := scanner.Text()[:4]
		if cmd == "addx" {
			var err error
			val, err = strconv.Atoi(scanner.Text()[5:])
			if err != nil {
				panic("fuck")
			}
		}
		commands = append(commands, Instruction{cmd, val})
	}
	return commands
}
