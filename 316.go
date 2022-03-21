package main

func removeDuplicateLetters(s string) string {
	lastInstance := [26]int{}
	hasSeen := [26]bool{}

	for i := range s {
		lastInstance[s[i]-'a'] = i
	}

	var res []rune
	for i, r := range s {
		for len(res) > 0 {
			lastRune := res[len(res)-1]
			if r <= lastRune && lastInstance[lastRune-'a'] >= i && !hasSeen[r-'a'] {
				hasSeen[lastRune-'a'] = false
				res = res[:len(res)-1]
				continue
			}
			break
		}
		if !hasSeen[r-'a'] {
			res = append(res, r)
			hasSeen[r-'a'] = true
		}
	}
	return string(res)
}
