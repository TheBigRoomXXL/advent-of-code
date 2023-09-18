package main

import (
	"reflect"
	"testing"
)

func TestParseItems(t *testing.T) {
	in := "  Starting items: 79, 98"
	want := []int{79, 98}

	ans := parseItems(in)
	if !reflect.DeepEqual(ans, want) {
		t.Errorf("\nWanted %v\nGot:%v\n", want, ans)
	}
}

func TestParseOperation(t *testing.T) {
	var tests = []struct {
		in   string
		want func(int) int
	}{
		{"  Operation: new = old * 19", func(x int) int { return x * 19 }},
		{"  Operation: new = old + 6", func(x int) int { return x + 6 }},
		{"  Operation: new = old * old", func(x int) int { return x * x }},
		{"  Operation: new = old + 3", func(x int) int { return x + 3 }},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			ans := parseOperation(tt.in)
			if tt.want(69) != ans(69) {
				t.Errorf("\nWanted %v\nGot:%v\n", tt.want(69), ans(69))
			}
		})
	}
}

func TestParseTest(t *testing.T) {
	in := []string{
		"  Test: divisible by 13",
		"    If true: throw to monkey 1",
		"    If false: throw to monkey 3",
	}
	want := func(x int) int {
		if x%13 == 0 {
			return 1
		}
		return 3
	}

	ans, _ := parseTest(in)
	if ans(13) != want(13) {
		t.Errorf("\nWanted %v\nGot:%v\n", want(13), ans(13))
	}
	if ans(69) != want(69) {
		t.Errorf("\nWanted %v\nGot:%v\n", want(69), ans(69))
	}
}
