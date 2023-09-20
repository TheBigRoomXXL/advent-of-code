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
	maze, _, s, e := parseInput(input)

	seen := explore(maze, s, e)
	printSeen(seen)
	return seen[e.y][e.x]
}

func part2(input io.Reader) int {
	maze, lows, _, e := parseInput(input)
	paths := []int{}
	for _, low := range lows {
		paths = append(paths, explore(maze, low, e)[e.y][e.x])
	}
	min := 9999999999999
	for _, path := range paths {
		if path <= min {
			min = path
		}
	}

	return min
}

type Point struct {
	x int
	y int
}

type PointQueue struct {
	val []Point
}

func (q *PointQueue) Enqueue(p Point) {
	q.val = append(q.val, p)
}

func (q *PointQueue) Dequeue() Point {
	if q.Len() == 0 {
		return Point{}
	}
	p := q.val[0]
	q.val[0] = Point{}
	q.val = q.val[1:]
	return p
}

func (q *PointQueue) Len() int {
	return len(q.val)
}

func explore(maze [][]int, start Point, end Point) [][]int {
	endReached := false
	queue := PointQueue{[]Point{start}}
	seen := [][]int{}

	for i := range maze {
		seen = append(seen, []int{})
		for range maze[0] {
			seen[i] = append(seen[i], 0)
		}
	}

	for !(queue.Len() == 0 || endReached) {
		endReached = walk(maze, queue.Dequeue(), end, seen, &queue)
	}

	if !endReached { // part2
		seen[end.y][end.x] = 69420
	}

	return seen
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func walk(maze [][]int, curr Point, end Point, seen [][]int, queue *PointQueue) bool {
	for _, dir := range dirs {
		next := Point{curr.x + dir[0], curr.y + dir[1]}
		if next.x < 0 || next.x >= len(maze[0]) || next.y < 0 || next.y >= len(maze) {
			continue
		}

		if seen[next.y][next.x] != 0 {
			continue
		}

		if maze[curr.y][curr.x]+1 < maze[next.y][next.x] {
			continue
		}

		seen[next.y][next.x] = seen[curr.y][curr.x] + 1
		queue.Enqueue(next)

		if next.x == end.x && next.y == end.y {
			return true
		}
	}
	return false
}

func parseInput(input io.Reader) ([][]int, []Point, Point, Point) {
	scanner := bufio.NewScanner(input)
	maze := [][]int{}
	lows := []Point{} // part2
	var s Point
	var e Point
	y := 0

	for scanner.Scan() {
		line := []int{}
		for x, char := range scanner.Text() {
			if char == 'a' {
				lows = append(lows, Point{x, y})
			}
			if char == 'S' {
				s = Point{x, y}
			}
			if char == 'E' {
				e = Point{x, y}
			}
			line = append(line, runeToInt(char))
		}
		maze = append(maze, line)
		y++
	}

	return maze, lows, s, e
}

func runeToInt(char rune) int {
	// a-z is 97 to 122 in ASCII
	if char == 'S' {
		return 1
	} else if char == 'E' {
		return 26
	}
	return int(char) - 96
}

func printSeen(seen [][]int) {
	for i := range seen {
		for j := range seen[i] {
			fmt.Printf("|%03d", seen[i][j])
		}
		fmt.Println("|")
	}
}
