package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	in := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	want := []Motion{
		{"R", 4},
		{"U", 4},
		{"L", 3},
		{"D", 1},
		{"R", 4},
		{"D", 1},
		{"L", 5},
		{"R", 2},
	}

	ans := parseInput(strings.NewReader(in))
	if reflect.DeepEqual(in, want) {
		t.Errorf("\nexpected:\n %v\ngot:\n %v\n", want, ans)
	}

}

func TestPart1(t *testing.T) {
	in := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	want := 13

	ans := part1(strings.NewReader(in))
	if ans != want {
		t.Errorf("\nexpected:\n %d\ngot:\n %d\n", want, ans)
	}

}

func TestPart2A(t *testing.T) {
	in := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	want := 1
	ans := part2(strings.NewReader(in))
	if ans != want {
		t.Errorf("\nexpected:\n %d\ngot:\n %d\n", want, ans)
	}

}

func TestPart2B(t *testing.T) {
	in := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	want := 36
	ans := part2(strings.NewReader(in))
	if ans != want {
		t.Errorf("\nexpected:\n %d\ngot:\n %d\n", want, ans)
	}

}
