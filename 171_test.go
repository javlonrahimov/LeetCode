package main

import "testing"

func Test171(t *testing.T) {
	cases := []struct {
		Input string
		Want  int
	}{
		{Input: "A", Want: 1},
		{Input: "AB", Want: 28},
		{Input: "ZY", Want: 701},
		{Input: "I", Want: 9},
	}

	for _, test := range cases {
		t.Run(test.Input, func(t *testing.T) {
			got := TitleToNumber(test.Input)

			if got != test.Want {
				t.Errorf("want %d, got %d", test.Want, got)
			}
		})
	}
}
