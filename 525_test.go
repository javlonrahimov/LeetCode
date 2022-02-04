package main

import (
	"fmt"
	"testing"
)

func Test525(t *testing.T) {
	cases := []struct {
		Input []int
		Want  int
	}{
		{Input: []int{0, 1}, Want: 2},
		{Input: []int{0, 1, 0}, Want: 2},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("contiguous array for %v", test.Input), func(t *testing.T) {
			got := FindMaxLength(test.Input)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
