package main

import "math"

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	var max float64 = -1
	i, j := 0, len(height) - 1
	
	for i <= j {
		if height[i] < height[j] {
			max = math.Max(max, float64(height[i]*(j-i)))
			i++
		} else {
			max = math.Max(max, float64(height[j]*(j-i)))
			j--
		}
	}
	return int(max)
}