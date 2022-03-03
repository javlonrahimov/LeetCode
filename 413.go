package main

func numberOfArithmeticSlices(nums []int) int {
	ans, count := 0, 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			count += 1
			ans += count
		} else {
			count = 0
		}
	}
	return ans
}
