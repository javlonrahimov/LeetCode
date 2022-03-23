package main

import "math"

func minSubArrayLen(target int, nums []int) int {
    n, ans, left, sum := len(nums), math.MaxInt32, 0, 0
    for i := 0; i < n; i++ {
        sum += nums[i]
        for sum >= target {
            if ans > i + 1 - left {
                ans = i + 1 - left
            }
            sum -= nums[left]
            left++
        }
    }
    if ans != math.MaxInt32 {
        return ans
    } else {
        return 0
    }
}
