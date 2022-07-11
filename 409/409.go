package main

func longestPalindrome(s string) int {
	count := [128]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	ans := 0

	for i := 0; i < len(count); i++ {
		ans += (count[i] / 2) * 2
		if ans%2 == 0 && count[i]%2 == 1 {
			ans++
		}
	}
	return ans
}
