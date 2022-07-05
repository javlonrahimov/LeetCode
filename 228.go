package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	start := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 && nums[i+1]-nums[i] != -1 {
			if start != nums[i] {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
			} else {
				result = append(result, fmt.Sprintf("%d", start))
			}
			start = nums[i+1]
			i++
		}
	}
	return result
}
