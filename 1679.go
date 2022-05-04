package main

func maxOperations(nums []int, k int) int {
    ans, seen := 0, make(map[int]int)
    for _, num := range nums {
        n := k - num
        if seen[n] > 0 {
            ans++
            seen[n]--
        } else {
            seen[num]++
        }
    }
    return ans
}
