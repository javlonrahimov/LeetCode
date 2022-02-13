package main

import "math"

func Subsets(nums []int) [][]int {
	result := make([][]int, int(math.Pow(2, float64(len(nums)))))
	helper(nums, &result, len(result)>>1)
	return result
}

func helper(nums []int, result *[][]int, flag int) {
	if len(nums) == 0 {
		(*result)[0] = []int{}
		return
	}

	helper(nums[1:], result, flag>>1)
	for i := flag; i < flag<<1; i++ {
		temp := make([]int, len((*result)[i-flag]))
		copy(temp, (*result)[i-flag])
		(*result)[i] = append(temp, nums[0])
	}
}
