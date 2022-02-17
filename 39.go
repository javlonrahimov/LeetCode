package main

import (
	"reflect"
	"testing"
)

func Test39(t *testing.T) {
	cases := []struct {
		Name   string
		Input  []int
		Target int
		Want   [][]int
	}{}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := CombinationSum(test.Input, test.Target)

			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("want %v, got %v", test.Want, got)
			}
		})
	}
}
