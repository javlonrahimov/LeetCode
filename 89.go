package main

func grayCode(n int) []int {
	i, fac := 1, 1

	result := []int{0, 1}

	for ; i < n; i++ {
		fac = fac * 2
		l := len(result) - 1

		for ; l >= 0; l-- {
			val := result[l] + fac
			result = append(result, val)
		}
	}
	return result
}
