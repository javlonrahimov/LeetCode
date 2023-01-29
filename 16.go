package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ans int
	diff := math.MaxInt

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		h := len(nums) - 1

		for l < h {
			if nums[i]+nums[l]+nums[h] == target {
				ans = target
				return ans
			} else if int(math.Abs(float64(nums[i]+nums[l]+nums[h]-target))) < diff {
				diff = int(math.Abs(float64(nums[i] + nums[l] + nums[h] - target)))
				ans = nums[i] + nums[l] + nums[h]
			}

			if nums[i]+nums[l]+nums[h] > target {
				h--
			} else {
				l++
			}
		}
	}
	return ans
}
