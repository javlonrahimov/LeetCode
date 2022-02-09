package main

import (
	"fmt"
	"testing"
)

func Test532(t *testing.T) {
	cases := []struct {
		Input []int
		K     int
		Want  int
	}{
		{Input: []int{3, 1, 4, 1, 5}, K: 2, Want: 2},
		{Input: []int{1, 2, 3, 4, 5}, K: 1, Want: 4},
		{Input: []int{1, 3, 1, 5, 4}, K: 0, Want: 1},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%v and %d", test.Input, test.K), func(t *testing.T) {
			got := FindPairs(test.Input, test.K)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
