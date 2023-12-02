package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

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
	score := 0
	i := 0
	for scanner.Scan() {
		i++
		var game string
		if i < 10 {
			game = scanner.Text()[7:]
		} else if i < 100 {
			game = scanner.Text()[8:]
		} else if i < 1000 {
			game = scanner.Text()[9:]
		} else {
			game = scanner.Text()[10:]
		}
		sets := strings.Split(game, ";")
		if validateSets(sets) {
			score += i
		}

	}
	return fmt.Sprint(score)
}

func part2(input io.Reader) string {
	return ""
}

func ParseDraw(draw string) (string, int) {
	var color string
	var n int
	_, err := fmt.Sscanf(draw, " %d %s", &n, &color)
	if err != nil {
		panic(err)
	}
	return color, n
}

func ParseSet(set string) map[string]int {
	result := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	draws := strings.Split(set, ",")
	for _, draw := range draws {
		color, n := ParseDraw(draw)
		result[color] += n
	}
	return result
}

func isResultPossible(result map[string]int) bool {
	if result["red"] > maxRed || result["green"] > maxGreen || result["blue"] > maxBlue {
		return false
	}
	return true
}

func validateSets(sets []string) bool {
	valid := 0
	for _, set := range sets {
		result := ParseSet(set)
		if isResultPossible(result) {
			valid++
		}
	}
	return valid == len(sets)
}
