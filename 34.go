package main

func searchRange(nums []int, target int) []int {
	first, last := -1, -1

	for index, value := range nums {
		if value == target && first == -1 {
			first = index
			last = index
		} else if value == target {
			last = index
		}
	}

	return []int{first, last}
}
