package main

func CheckInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	countsS1 := [26]int{}
	for _, c := range s1 {
		countsS1[c-'a']++
	}

	i := 0
	countsWin := [26]int{}
	for j := 0; j < len(s2); j++ {
		countsWin[s2[j]-'a']++

		if j-i+1 > len(s1) {
			countsWin[s2[i]-'a']--
			i++
		}

		if countsS1 == countsWin {
			return true
		}

	}
	return false
}
