package main

import (
	"fmt"
	"testing"
)

func Test125(t *testing.T) {
	cases := []struct {
		Input string
		Want  bool
	}{
		{Input: "A man, a plan, a canal: Panama", Want: true},
		{Input: "race a car", Want: false},
		{Input: " ", Want: true},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s gets checked weather palindrome or not", test.Input), func(t *testing.T) {
			got := IsPalindrome(test.Input)
			if got != test.Want {
				t.Errorf("got %t, want %t", got, test.Want)
			}
		})
	}
}
