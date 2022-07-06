package main

func romanToInt(s string) int {

	symbols := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	var result, prev int

	for i := 0; i < len(s); i++ {
		current := symbols[string(s[i])]
		result += current
		if prev < current {
			result -= 2 * prev
		}
		prev = current
	}

	return result
}
