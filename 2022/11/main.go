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
	monkeys := parseInput(input)
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.OperateOverItems(monkeys, 3)
		}
	}
	max1 := 0
	max2 := 0
	for _, monkey := range monkeys {
		if monkey.inspectCount >= max1 {
			max2 = max1
			max1 = monkey.inspectCount
		}
	}
	return max1 * max2
}

func part2(input io.Reader) int {
	monkeys := parseInput(input)

	// zen calculation from https://github.com/alexchao26/advent-of-code-go/blob/main/2022/day11/main.go#L107
	zen := 1
	for _, m := range monkeys {
		zen *= m.divisiblebY
	}
	for i := 0; i < 10_000; i++ {
		fmt.Println("round", i+1)
		for j, monkey := range monkeys {
			monkey.OperateOverItems(monkeys, zen)
			fmt.Println("mokey", j, "has inpected", monkey.inspectCount, "time")

		}
	}
	max1 := 0
	max2 := 0
	for _, monkey := range monkeys {
		if monkey.inspectCount >= max1 {
			max2 = max1
			max1 = monkey.inspectCount
		} else if monkey.inspectCount > max2 {
			max2 = monkey.inspectCount
		}
	}
	return max1 * max2
}

type Monkey struct {
	items        []int
	operation    func(int) int
	test         func(int) int
	divisiblebY  int
	inspectCount int
}

func (m *Monkey) OperateOverItems(monkeys []*Monkey, zen int) {
	for _, old := range m.items {
		new := m.operation(old) % zen
		m.inspectCount++
		// fmt.Println("    item ", old, "has become", new)
		next := m.test(new)
		monkeys[next].items = append(monkeys[next].items, new)
	}
	m.items = []int{}
}

func parseInput(input io.Reader) []*Monkey {
	scanner := bufio.NewScanner(input)
	scanner.Split(ScanMonkeys)
	monkeys := []*Monkey{}
	for scanner.Scan() {
		monkeys = append(monkeys, parseMonkey(scanner.Text()))
	}
	return monkeys
}

func parseMonkey(monkeystr string) *Monkey {
	lines := strings.Split(monkeystr, "\n")
	items := parseItems(lines[1])
	operation := parseOperation(lines[2])
	test, div := parseTest(lines[3:6])

	return &Monkey{items, operation, test, div, 0}
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

func parseTest(lines []string) (func(int) int, int) {
	if len(lines) != 3 {
		panic("There should be 3 lines to parse dummy")
	}
	div := AtoiOrPanic(lines[0][21:])
	m1 := AtoiOrPanic(lines[1][29:])
	m2 := AtoiOrPanic(lines[2][30:])
	return func(x int) int {
		if x%div == 0 {
			return m1
		}
		return m2
	}, div

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
