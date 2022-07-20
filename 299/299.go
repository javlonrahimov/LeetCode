package main

import "fmt"

func getHint(secret string, guess string) string {
	cow, bull := 0, 0
	count := make(map[rune]int)
	for _, ch := range secret {
		count[ch]++
	}
	for i, ch := range guess {
		if count[ch] > 0 {
			cow++
			count[ch]--
		}
		if secret[i] == byte(ch) {
			bull++
			cow--
		}
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}
