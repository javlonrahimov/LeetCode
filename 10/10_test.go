package main

import (
	"fmt"
	"testing"
)

func Test10(t *testing.T) {
	cases := []struct {
		Input   string
		Pattern string
		Want    bool
	}{
		{Input: "aa", Pattern: "a", Want: false},
		{Input: "aa", Pattern: "a*", Want: true},
		{Input: "ab", Pattern: ".*", Want: true},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("is %s matches to %s", test.Input, test.Pattern), func(t *testing.T) {
			got := IsMatch(test.Input, test.Pattern)
			if got != test.Want {
				t.Errorf("got %t, want %t", got, test.Want)
			}
		})
	}
}
