package main

import (
	"testing"
)

func Test24(t *testing.T) {
	cases := []struct {
		Name string
		Input *ListNode
		Want *ListNode
	} {
		{
			Name: "Normal",
			Input: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
			Want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}},
		},
		{Name: "Just head", Input: &ListNode{Val: 1, Next: nil}, Want: &ListNode{Val: 1, Next: nil}},
	}

	for _, test := range cases {
		got := SwapPairs(test.Input)

		if got != test.Want {
			t.Errorf("want %v, got %v", test.Want, got)
		}
	}
}