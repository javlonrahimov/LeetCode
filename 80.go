package main

func RemoveDuplicates(nums []int) int {

	if nums == nil {
		return 0
	}

	if len(nums) <= 2 {
		return len(nums)
	}

	i, j := 1, 2

	for j < len(nums) {
		if nums[j] == nums[i] && nums[j] == nums[i-1] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}

	return i + 1
}
