package main

import "testing"

func Test148(t *testing.T) {

	cases := []struct {
		Name  string
		Input *ListNode
		Want  *ListNode
	}{
		{Name: "Test 1",
			Input: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}},
			Want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := sortList(test.Input)
			if got != test.Want {
				t.Errorf("want %v, got %v", test.Want, got)
			}
		})
	}

}
