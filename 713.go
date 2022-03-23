package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k<= 1 {
        return 0
    }
    prod, ans, left := 1, 0, 0
    for right := 0; right < len(nums); right++ {
        prod *= nums[right]
        for prod >= k {
            prod /= nums[left]
            left++
        }
        ans += right - left + 1
    }
    return ans;
}