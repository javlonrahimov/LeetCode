package main

import (
	"fmt"
	"testing"
)

func Test567(t *testing.T) {
	cases := []struct {
		S1   string
		S2   string
		Want bool
	}{
		{S1: "ab", S2: "eidbaooo", Want: true},
		{S1: "ab", S2: "eidbosoo", Want: false},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("S1 = %s, S2 = %s", test.S1, test.S2), func(t *testing.T) {
			got := CheckInclusion(test.S1, test.S2)
			if got != test.Want {
				t.Errorf("got %t, want %t", got, test.Want)
			}
		})
	}
}
