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
	var result int
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

func part1(input io.Reader) int {
	history := map[string]Rope{"[0,0]": {0, 0, nil}}
	motions := parseInput(input)
	tail := Rope{0, 0, nil}
	head := Rope{0, 0, &tail}
	for _, motion := range motions {
		for i := 0; i < motion.dist; i++ {
			head.move(motion)
			history[fmt.Sprintf("[%d,%d]", tail.x, tail.y)] = tail
		}
	}
	return len(history)
}

func part2(input io.Reader) int {
	history := map[string]Rope{"[0,0]": {0, 0, nil}}
	motions := parseInput(input)
	ropes := []*Rope{{0, 0, nil}}
	for i := 0; i < 9; i++ {
		ropes = append(ropes, &Rope{0, 0, ropes[len(ropes)-1]})
	}
	head := ropes[len(ropes)-1]
	for _, motion := range motions {
		for i := 0; i < motion.dist; i++ {
			head.move(motion)
			history[fmt.Sprintf("[%d,%d]", ropes[0].x, ropes[0].y)] = *ropes[0]
		}
	}
	return len(history)
}

type Motion struct {
	dir  string
	dist int
}

type Rope struct {
	x    int
	y    int
	tail *Rope
}

func (r *Rope) move(m Motion) {
	switch m.dir {
	case "L":
		r.x--
	case "R":
		r.x++
	case "U":
		r.y--
	case "D":
		r.y++
	case "LU":
		r.x--
		r.y--
	case "RU":
		r.x++
		r.y--
	case "LD":
		r.x--
		r.y++
	case "RD":
		r.x++
		r.y++
	}
	if r.tail != nil {
		r.tail.follow(r, m)
	}
}

func (r1 *Rope) follow(r2 *Rope, m Motion) {
	followTable := map[[2]int]string{
		{-2, 0}:  "L",
		{-1, 2}:  "LD",
		{-2, 1}:  "LD",
		{-2, 2}:  "LD",
		{-2, -1}: "LU",
		{-1, -2}: "LU",
		{-2, -2}: "LU",
		{2, 0}:   "R",
		{2, 1}:   "RD",
		{2, 2}:   "RD",
		{1, 2}:   "RD",
		{1, -2}:  "RU",
		{2, -1}:  "RU",
		{2, -2}:  "RU",
		{0, -2}:  "U",
		{0, 2}:   "D",
	}
	xGap := r2.x - r1.x
	yGap := r2.y - r1.y
	dir := followTable[[2]int{xGap, yGap}]
	if dir != "" {
		r1.move(Motion{dir, 1})
	}
}

func parseInput(input io.Reader) []Motion {
	scanner := bufio.NewScanner(input)
	motions := []Motion{}
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			panic(err)
		}
		motions = append(motions, Motion{direction, distance})
	}
	return motions
}
