package main

import (
	"fmt"
	"math"
	"testing"
)

func Test50(t *testing.T) {
	cases := []struct {
		X    float64
		N    int
		Want float64
	}{
		{X: 2.00000, N: 10, Want: 1024.00000},
		{X: 2.10000, N: 3, Want: 9.26100},
		{X: 2.00000, N: -2, Want: 0.25000},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("pow(%f, %d)", test.X, test.N), func(t *testing.T) {
			got := MyPow(test.X, test.N)
			if diff := math.Abs(got - test.Want); diff >= 1e-9 {
				t.Errorf("got %f, want %f", got, test.Want)
			}
		})
	}
}
