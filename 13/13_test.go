package main

import "testing"

func Test13(t *testing.T) {
	cases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{Name: "Test 1", Input: "III", Want: 3},
		{Name: "Test 2", Input: "LVIII", Want: 58},
		{Name: "Test 1", Input: "MCMXCIV", Want: 1994},
	}

	for _, test := range cases {
		got := romanToInt(test.Input)
		if got != test.Want {
			t.Errorf("got %d, want %d", got, test.Want)
		}
	}
}
