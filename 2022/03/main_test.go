package main

import (
	"fmt"
	"testing"
)

func TestFindPriority(t *testing.T) {
	var tests = []struct {
		item rune
		want int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%b,%d", tt.item, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := findPriority(tt.item)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestFindSharedBadge(t *testing.T) {
	group := [3]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}
	want := 'r'
	ans := findSharedBadge(group)
	if ans != want {
		t.Errorf("findSharedBadge(group) = %c; want r", ans)
	}
}
