package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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
	forest := parseInput(input)
	result := 0
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[0]); x++ {
			if isVisible(forest, y, x) {
				result++
			}
		}
	}
	return result
}

func part2(input io.Reader) int {
	forest := parseInput(input)
	max := 0
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[0]); x++ {
			score := scenicScore(forest, y, x)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func parseInput(input io.Reader) [][]int8 {
	forest := [][]int8{{}}
	br := bufio.NewReader(input)
	y := 0
	for {
		b, err := br.ReadByte()
		if err != nil {
			break
		}

		if b == '\n' {
			y++
			forest = append(forest, []int8{})
		} else {
			forest[y] = append(forest[y], int8(b)-48)
		}
	}
	return forest
}

func isVisible(forest [][]int8, y int, x int) bool {
	h := forest[y][x]
	// Check the edge
	if x == 0 || x == len(forest[0])-1 || y == 0 || y == len(forest)-1 {
		return true
	}
	// Check left
	for l := x - 1; l >= 0; l-- {
		if forest[y][l] >= h {
			break
		}
		if l == 0 {
			return true
		}
	}
	// Check right
	for r := x + 1; r < len(forest[0]); r++ {
		if forest[y][r] >= h {
			break
		}
		if r == len(forest[0])-1 {
			return true
		}
	}
	// Check top
	for t := y - 1; t >= 0; t-- {
		if forest[t][x] >= h {
			break
		}
		if t == 0 {
			return true
		}
	}
	// Check bottom
	for b := y + 1; b < len(forest); b++ {
		if forest[b][x] >= h {
			break
		}
		if b == len(forest)-1 {
			return true
		}
	}
	return false
}

func scenicScore(forest [][]int8, y int, x int) int {
	h := forest[y][x]
	left, rigth, top, bottom := 0, 0, 0, 0
	if x == 0 || x == len(forest[0])-1 || y == 0 || y == len(forest)-1 {
		return 0
	}
	//Left
	for i := 1; i <= x; i++ {
		if forest[y][x-i] >= h || x == i {
			left = i
			break
		}
	}
	//Right
	for i := 1; i < len(forest[0])-x; i++ {
		if forest[y][x+i] >= h || x+i == len(forest[0])-1 {
			rigth = i
			break
		}
	}
	//Top
	for i := 1; i <= y; i++ {
		if forest[y-i][x] >= h || y == i {
			top = i
			break
		}
	}
	//Bottom
	for i := 1; i < len(forest)-y; i++ {
		if forest[y+i][x] >= h || y+i == len(forest)-1 {
			bottom = i
			break
		}
	}
	return left * rigth * top * bottom
}
