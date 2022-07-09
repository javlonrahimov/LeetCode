package main

func isIsomorphic(s string, t string) bool {

	charMap := map[string]string{}

	for index, value := range s {
		if _, ok := charMap[string(value)]; ok {
			if charMap[string(value)] != string(t[index]) {
				return false
			}
		} else if checkForValue(string(t[index]), charMap) {
			return false
		} else {
			charMap[string(value)] = string(t[index])
		}
	}

	return true

}

func checkForValue(userValue string, students map[string]string) bool {
	for _, value := range students {
		if value == userValue {
			return true
		}
	}
	return false
}
