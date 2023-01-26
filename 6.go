package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	answer := ""
	n := len(s)
	charsInSection := 2 * (numRows - 1)

	for currRow := 0; currRow < numRows; currRow++ {
		index := currRow

		for index < n {
			answer += string(s[index])

			if currRow != 0 && currRow != numRows-1 {
				charsInBetween := charsInSection - 2*currRow
				secondIndex := index + charsInBetween

				if secondIndex < n {
					answer += string(s[secondIndex])
				}
			}
			index += charsInSection
		}
	}

	return answer
}
