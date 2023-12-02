package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	part := flag.Int("p", 0, "part number (1 or 2)")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	default:
		panic("invalid part")
	}
}

func part1() {
	scores := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result += scores[scanner.Text()]
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func part2() {
	scores := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result += scores[scanner.Text()]
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(result)

}
