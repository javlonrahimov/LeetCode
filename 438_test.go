package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test438(t *testing.T) {
	cases := []struct {
		Input   string
		Pattern string
		Want    []int
	}{
		{Input: "cbaebabacd", Pattern: "abc", Want: []int{0, 6}},
		{Input: "abab", Pattern: "ab", Want: []int{0, 1, 2}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("starting indexes of %s in %s", test.Pattern, test.Input), func(t *testing.T) {
			got := FindAnagrams(test.Input, test.Pattern)
			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
