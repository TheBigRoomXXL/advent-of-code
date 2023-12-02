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
	part := flag.Int("p", 0, "part number (1 or 2)")
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
	scanner := bufio.NewScanner(input)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := []rune{}
		for _, r := range line {
			if isDigit(r) {
				digits = append(digits, r)
			}
		}
		calibration_str := string(digits[0]) + string(digits[len(digits)-1])
		calibration, err := strconv.Atoi(calibration_str)
		if err != nil {
			panic("bad conversion")
		}
		result += calibration
	}

	return fmt.Sprint(result)
}

func part2(input io.Reader) string {
	return ""
}

func isDigit(r rune) bool {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, v := range digits {
		if r == v {
			return true
		}
	}
	return false
}
