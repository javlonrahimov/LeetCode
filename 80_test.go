package main

import (
	"fmt"
	"testing"
)

func Test80(t *testing.T) {
	cases := []struct {
		Input []int
		Want  int
	}{
		{Input: []int{1, 1, 1, 2, 2, 3}, Want: 5},
		{Input: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, Want: 7},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("test: %d", i+1), func(t *testing.T) {
			got := RemoveDuplicates(test.Input)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
