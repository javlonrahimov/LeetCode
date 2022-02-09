package main

import (
	"fmt"
	"testing"
)

func Test258(t *testing.T) {
	cases := []struct {
		Input int
		Want  int
	}{
		{Input: 38, Want: 2},
		{Input: 0, Want: 0},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d number", test.Input), func(t *testing.T) {
			got := AddDigits(test.Input)

			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
