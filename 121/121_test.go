package main

import (
	"fmt"
	"testing"
)

func Test121(t *testing.T) {
	cases := []struct {
		Input []int
		Want  int
	}{
		{Input: []int{7, 1, 5, 3, 6, 4}, Want: 5},
		{Input: []int{7, 6, 4, 3, 1}, Want: 0},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("checks max profit form pirces %v", test.Input), func(t *testing.T) {
			got := MaxProfit(test.Input)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
