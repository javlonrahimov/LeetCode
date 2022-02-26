package main

func buildArray(nums []int) []int {
	ans := []int{}
	for _, value := range nums {
		ans = append(ans, nums[value])
	}
	return ans
}
