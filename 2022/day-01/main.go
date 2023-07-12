package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
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
	if err := scanner.Err(); err != nil {
		check(err)
	}

	max := 0
	for i := 0; i <= len(elves)-1; i++ {
		if elves[i] > max {
			max = elves[i]
		}
	}

	fmt.Println(max)

}
