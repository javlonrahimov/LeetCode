package main

import "testing"

func Test102(t *testing.T) {
	cases := []struct {
		Name string
		Num  string
		K    int
		Want string
	}{
		{Name: "1432219", Num: "1432219", K: 3, Want: "1219"},
		{Name: "10200", Num: "10200", K: 1, Want: "200"},
		{Name: "10", Num: "10", K: 2, Want: "0"},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := removeKDigits(test.Num, test.K)

			if got != test.Want {
				t.Errorf("want %s, got %s", test.Want, got)
			}
		})
	}
}
