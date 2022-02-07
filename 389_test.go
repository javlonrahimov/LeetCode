package main

import (
	"fmt"
	"testing"
)

func Test389(t *testing.T) {
	cases := []struct {
		S    string
		T    string
		Want byte
	}{
		{S: "absd", T: "absde", Want: byte('e')},
		{S: "", T: "y", Want: byte('y')},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s and %s", test.S, test.T), func(t *testing.T) {
			got := FindTheDifference(test.S, test.T)
			if got != test.Want {
				t.Errorf("got %x, wnat %x", got, test.Want)
			}
		})
	}
}
