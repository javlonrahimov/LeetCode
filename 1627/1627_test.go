package main

import (
	"fmt"
	"testing"
)

func Test1627(t *testing.T) {
	cases := []struct {
		Input [][]int
		Want  int
	}{
		{Input: [][]int{{1, 2, 3}, {3, 2, 1}}, Want: 6},
		{Input: [][]int{{1, 5}, {7, 3}, {3, 5}}, Want: 10},
		{Input: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, Want: 17},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%v gets for max wealth", test.Input), func(t *testing.T) {
			got := MaximumWealth(test.Input)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
