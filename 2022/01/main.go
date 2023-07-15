package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	part := flag.Int("part", 0, "part number (1 or 2)")
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
	elves := parseInpout()
	max := 0
	for i := 0; i <= len(elves)-1; i++ {
		if elves[i] > max {
			max = elves[i]
		}
	}

	fmt.Println(max)
}

func part2() {
	elves := parseInpout()
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}

func parseInpout() []int {
	elves := []int{0}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			calories, err := strconv.Atoi(text)
			check(err)
			elves[len(elves)-1] += calories
		} else {
			elves = append(elves, 0)
		}
	}
	return elves
}
