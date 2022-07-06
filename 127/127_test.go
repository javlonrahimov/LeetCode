package main

import (
	"fmt"
	"testing"
)

func Test127(t *testing.T) {
	cases := []struct {
		BeginWord string
		EndWord   string
		WordList  []string
		Want      int
	}{
		{BeginWord: "hit", EndWord: "cog", WordList: []string{"hot", "dot", "dog", "lot", "log", "cog"}, Want: 5},
		{BeginWord: "hit", EndWord: "cog", WordList: []string{"hot", "dot", "dog", "lot", "log"}, Want: 0},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Begin %s, end %s", test.BeginWord, test.EndWord), func(t *testing.T) {
			got := LadderLength(test.BeginWord, test.EndWord, test.WordList)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
