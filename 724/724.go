package main

func pivotIndex(nums []int) int {

	sum := sum(nums)
	leftsum := 0

	for index, value := range nums {
		if leftsum == (sum - leftsum - value) {
			return index
		}
		leftsum += value
	}

	return -1
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
