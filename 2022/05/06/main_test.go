package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		in, out string
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "7"},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "5"},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "11"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.in, tt.out)
		t.Run(testname, func(t *testing.T) {
			ans := part1(strings.NewReader(tt.in))
			if ans != tt.out {
				t.Errorf("got %s, want %s", ans, tt.out)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		in, out string
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "19"},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "23"},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "26"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.in, tt.out)
		t.Run(testname, func(t *testing.T) {
			ans := part2(strings.NewReader(tt.in))
			if ans != tt.out {
				t.Errorf("got %s, want %s", ans, tt.out)
			}
		})
	}
}
