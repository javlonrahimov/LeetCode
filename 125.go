package main

import(
	"strings"
	"fmt"
)

func main() {

	fmt.Println(isPalendrome("A man, a plan, a canal: Panama"))

}

func isPalindrome(s string) bool {

	s = strip_string(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	fmt.Println(s)

	reversedStr := ""
    for i := len(s)-1; i >= 0; i-- {
        reversedStr += string(s[i])
    }
    for i := range(s) {
        if s[i] != reversedStr[i] {
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