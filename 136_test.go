package main

import (
	"testing"
)

func Test136(t *testing.T) {
	cases := []struct {
		Name  string
		Input []int
		Want  int
	}{
		{Name: "Test 1", Input: []int{2, 2, 1}, Want: 1},
		{Name: "Test 2", Input: []int{4, 1, 2, 1, 2}, Want: 4},
		{Name: "Size 1", Input: []int{1}, Want: 1},
	}

	for _, test := range cases {
		got := SingleNumber(test.Input)
		if got != test.Want {
			t.Errorf("got %d, want %d", got, test.Want)
		}
	}
}
