package main

import (
	"fmt"
	"testing"
)

func Test105(t *testing.T) {
	cases := []struct {
		Input TreeNode
		Want  int
	}{
		{Input: TreeNode{Val: 1, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}, Want: 3},
		{Input: TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}, Want: 2},
	}

	for index, test := range cases {
		t.Run(fmt.Sprintf("Test %d", index), func(t *testing.T) {
			got := MaxDepth(&test.Input)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
