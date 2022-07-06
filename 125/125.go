package main

import (
	"strings"
)

func main() {

}

func IsPalindrome(s string) bool {

	s = strip_string(s)
	s = strings.ReplaceAll(s, " ", "")
	processedString := strings.ToLower(s)

	for i := 0; i < len(processedString)/2; i++ {
		if processedString[i] != processedString[len(processedString)-1-i] {
			return false
		}
	}
	return true
}

func strip_string(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}
