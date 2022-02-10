package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		Nums []int
		K    int
		Want int
	}{
		{Nums: []int{1, 1, 1}, K: 2, Want: 2},
		{Nums: []int{1, 2, 3}, K: 3, Want: 2},
		{Nums: []int{-1, -1, 1}, K: 1, Want: 1},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("list %v k %d", test.Nums, test.K), func(t *testing.T) {
			got := SubarraySum(test.Nums, test.K)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
