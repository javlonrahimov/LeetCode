package main

func FindAnagrams(s, p string) []int {

	n, m := len(s), len(p)
	if m > n {
		return nil
	}

	mp := make(map[byte]int)
	ms := make(map[byte]int)
	anagrams := []int{}
	for i := 0; i < m; i++ {
		mp[p[i]]++
		ms[s[i]]++
	}

	if compare(mp, ms) {
		anagrams = append(anagrams, 0)
	}

	for i := m; i < n; i++ {
		ms[s[i-m]]--
		if ms[s[i-m]] <= 0 {
			delete(ms, s[i-m])
		}

		ms[s[i]]++

		if compare(mp, ms) {
			anagrams = append(anagrams, i-m+1)
		}
	}

	return anagrams
}

func compare(mp, ms map[byte]int) bool {
	if len(mp) != len(ms) {
		return false
	}

	for k, v := range ms {
		if vp, ok := mp[k]; !ok || vp != v {
			return false
		}
	}

	return true
}
