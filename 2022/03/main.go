package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	prioritiesSum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rucksacks := scanner.Text()
		sharedItem := findSharedItem(rucksacks)
		prioritiesSum += findPriority(sharedItem)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(prioritiesSum)

}

func part2() {
	prioritiesSum := 0
	var group [3]string
	i := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		group[i] = scanner.Text()
		i++
		if i == 3 {
			i = 0
			badge := findSharedBadge(group)
			prioritiesSum += findPriority(badge)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(prioritiesSum)

}

func findSharedItem(rucksacks string) rune {
	// Could use sets to avoid duplicated checks but the perfomance
	// is not needed for the input size.
	compartment1 := rucksacks[:len(rucksacks)/2]
	compartment2 := rucksacks[len(rucksacks)/2:]
	for _, item1 := range compartment1 {
		for _, item2 := range compartment2 {
			if item1 == item2 {
				return item1
			}
		}
	}
	panic("No match found")
}

func findSharedBadge(group [3]string) rune {
	// That's a lot of nesting
	for _, item1 := range group[0] {
		for _, item2 := range group[1] {
			if item1 == item2 {
				for _, item3 := range group[2] {
					if item1 == item3 {
						return item3
					}
				}
			}
		}
	}
	panic("No match found")
}

func findPriority(item rune) int {
	// A-Z is 65 to 90 in ASCII
	// a-z is 97 to 102 in ASCII
	if item < 97 {
		return int(item) - 38
	} else {
		return int(item) - 96
	}
}
