package main

import "sort"

func threeSum(nums []int) [][]int {
  res := [][]int{}
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i + 1, n - 1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[l], nums[r], nums[i]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}