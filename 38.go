package main

import "strconv"

func countAndSay(n int) string {
	temp := "1"
	for i := 2; i <= n; i++ {
		temp = countNum(temp)
	}
	return temp
}
func countNum(s string) string {
	count := 0
	str := string(s[0])
	temp := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == str {
			count++

		} else {
			temp += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
		if i == len(s)-1 {
			temp += strconv.Itoa(count) + string(s[i])
		}
		str = string(s[i])
	}
	return temp
}
