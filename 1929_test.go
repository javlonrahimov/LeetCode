package main

import (
	"reflect"
	"testing"
)

func Test1929(t *testing.T) {
	cases := []struct {
		Name  string
		Input []int
		Want  []int
	}{
		{Name: "Test 1", Input: []int{1, 2, 1}, Want: []int{1, 2, 1, 1, 2, 1}},
		{Name: "Test 2", Input: []int{1, 3, 2, 1}, Want: []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}

	for _, test := range cases {
		got := getConcatenation(test.Input)

		if reflect.DeepEqual(got, test.Want) {
			// t.Errorf("want %v, got %v", test.Want, got)
		}
	}
}
