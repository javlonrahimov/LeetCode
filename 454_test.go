package main

import (
	"testing"
)

func Test454(t *testing.T) {

	cases := []struct {
		Nums1 []int
		Nums2 []int
		Nums3 []int
		Nums4 []int
		Want  int
	}{
		{Nums1: []int{1, 2}, Nums2: []int{-2, -1}, Nums3: []int{-1, 2}, Nums4: []int{0, 2}, Want: 2},
		{Nums1: []int{0}, Nums2: []int{0}, Nums3: []int{0}, Nums4: []int{0}, Want: 1},
	}

	t.Run("4Sum", func(t *testing.T) {
		for _, test := range cases {
			got := FourSumCount(test.Nums1, test.Nums2, test.Nums3, test.Nums4)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		}
	})

}
