package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int
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

func part1(input io.Reader) int {
	mokeys := parseInput(input)
	fmt.Println(mokeys)
	return 2069
}

func part2(input io.Reader) int {
	return 2069
}

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int
}

func parseInput(input io.Reader) []Monkey {
	scanner := bufio.NewScanner(input)
	scanner.Split(ScanMonkeys)
	monkeys := []Monkey{}
	for scanner.Scan() {
		monkeys = append(monkeys, parseMonkey(scanner.Text()))
	}
	return monkeys
}

func parseMonkey(monkeystr string) Monkey {
	lines := strings.Split(monkeystr, "/n")
	items := parseItems(lines[0])
	operation := parseOperation(lines[1])
	test := parseTest(lines[2:5])
	return Monkey{items, operation, test}
}

func parseItems(line string) []int {
	itemsstr := strings.Split(line[18:], ",")
	items := []int{}
	for _, itemstr := range itemsstr {
		item := AtoiOrPanic(itemstr)
		items = append(items, item)
	}
	return items
}

func parseOperation(line string) func(int) int {
	a, b := "", ""
	n, err := fmt.Sscanf(line, "  Operation: new = old %s %s", &a, &b)
	if err != nil || n != 2 {
		panic(fmt.Sprintf("bad patern matching for operation %s", line))
	}
	if b == "old" {
		switch a {
		case "+":
			return func(x int) int { return x + x }
		case "*":
			return func(x int) int { return x * x }
		}
	} else {
		operand := AtoiOrPanic(b)
		switch a {
		case "+":
			return func(x int) int { return x + operand }
		case "*":
			return func(x int) int { return x * operand }
		}
	}
	panic("Could not return operation")
}

func parseTest(lines []string) func(int) int {
	if len(lines) != 3 {
		panic("There should be 3 lines to parse dummy")
	}
	div := AtoiOrPanic(lines[0][21:])
	fmt.Println(lines[1][29:])
	m1 := AtoiOrPanic(lines[1][29:])
	m2 := AtoiOrPanic(lines[2][30:])
	return func(x int) int {
		if x%div == 0 {
			return m1
		}
		return m2
	}

}

func AtoiOrPanic(str string) int {
	str = strings.Trim(str, " ")
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Could not convert %s to int", str))
	}
	return n
}

// Copy from buffio.ScanLines() and addapted for "\n\n" instead of "\n"
func ScanMonkeys(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
