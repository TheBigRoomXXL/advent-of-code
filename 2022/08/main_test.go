package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	in := `30373
25512
65332
33549
35390`
	want := [][]int8{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	ans := parseInput(strings.NewReader(in))
	if reflect.DeepEqual(in, want) {
		t.Errorf("\nexpected:\n %d\ngot:\n %d\n", want, ans)
	}

}

func TestIsVisible(t *testing.T) {
	var tests = []struct {
		id      int
		forest  [][]int8
		x       int
		y       int
		visible bool
	}{
		{1, [][]int8{{0, 9, 0}, {8, 5, 7}, {0, 6, 0}}, 1, 1, false},
		{2, [][]int8{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 1, 1, false},
		{3, [][]int8{{0, 9, 0, 0, 0}, {6, 5, 3, 3, 2}, {0, 9, 0, 0, 0}}, 1, 1, true},
		{3, [][]int8{{9, 9, 9}, {9, 5, 9}, {9, 1, 9}, {9, 1, 9}}, 1, 1, true},
		{4, [][]int8{{0, 9, 0}, {8, 5, 7}, {0, 6, 0}}, 0, 0, true},
		{5, [][]int8{{0, 1, 0}, {2, 8, 3}, {0, 4, 0}}, 1, 1, true},
		{6, [][]int8{{0, 1, 6, 0}, {2, 8, 7, 3}, {0, 4, 8, 0}}, 1, 1, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.id)
		t.Run(testname, func(t *testing.T) {
			ans := isVisible(tt.forest, tt.x, tt.y)
			if ans != tt.visible {
				t.Errorf("\nexpected:\n %t\ngot:\n %t\n", tt.visible, ans)
			}
		})
	}
}

func TestScenicScore(t *testing.T) {
	var tests = []struct {
		forest [][]int8
		x      int
		y      int
		score  int
	}{
		{[][]int8{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		}, 1, 2, 4},
		{[][]int8{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		}, 3, 2, 8},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.score)
		t.Run(testname, func(t *testing.T) {
			ans := scenicScore(tt.forest, tt.x, tt.y)
			if ans != tt.score {
				t.Errorf("\nexpected:\n %d\ngot:\n %d\n", tt.score, ans)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	in := `30373
25512
65332
33549
35390`
	want := 21
	ans := part1(strings.NewReader(in))
	if ans != want {
		t.Errorf("\nexpected: %d\ngot: %d\n", want, ans)
	}

}
