package main

import (
	"strings"
)

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, i := range operations {
		if strings.Contains(i, "++") {
			res++
		} else {
			res--
		}
	}

	return 0
}
