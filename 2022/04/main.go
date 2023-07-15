package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	result := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a1, a2 := getAssignments(scanner.Text())
		a3 := mergeSet(a1, a2)
		if len(a3) == len(a1) || len(a3) == len(a2) {
			result++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func part2() {
	result := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a1, a2 := getAssignments(scanner.Text())
		a3 := mergeSet(a1, a2)
		if len(a3) != len(a1)+len(a2) {
			result++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func getAssignments(input string) (map[int]bool, map[int]bool) {
	assignments := strings.Split(input, ",")
	a1 := makeSet(assignments[0])
	a2 := makeSet(assignments[1])
	return a1, a2

}

func makeSet(input string) map[int]bool {
	str_range := strings.Split(input, "-")
	start, err := strconv.Atoi(str_range[0])
	if err != nil {
		panic("Error during conversion")
	}
	end, err := strconv.Atoi(str_range[1])
	if err != nil {
		panic("Error during conversion")
	}
	set := map[int]bool{}
	for i := start; i <= end; i++ {
		set[i] = true
	}
	return set
}

func mergeSet(a1 map[int]bool, a2 map[int]bool) map[int]bool {
	a3 := map[int]bool{}
	for k, v := range a1 {
		a3[k] = v
	}
	for k, v := range a2 {
		a3[k] = v
	}
	return a3
}
