package main

import (
	"fmt"
	// "reflect"
	"testing"
)

func Test78(t *testing.T) {
	cases := []struct {
		Nums []int
		Want [][]int
	}{
		{Nums: []int{1, 2, 3}, Want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{Nums: []int{0}, Want: [][]int{{}, {0}}},
	}

	for index, test := range cases {
		t.Run(fmt.Sprintf("Test number %d : %v", index+1, test.Nums), func(t *testing.T) {
			// got := Subsets(test.Nums)
			// if !reflect.DeepEqual(got, test.Want) {
			// 	t.Errorf("got %v, want %v", got, test.Want)
			// }
		})
	}
}
