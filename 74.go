package main

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, value := range row {
			if target < value {
				return false
			}
			if target == value {
				return true
			}
		}
	}
	return false
}
