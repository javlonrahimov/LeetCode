package main

import "sort"

func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    
    sort.Ints(nums)
    dfs(&res, []int{}, nums)
    
    return res
}

func dfs(res *[][]int, path, nums []int) {
    if len(nums) == 0 {
        *res = append(*res, path)
    }
    
    for i, num := range nums {
        if i > 0 && num == nums[i - 1] {
            continue
        }
        
        newPath := make([]int, len(path) + 1)
        newNums := make([]int, len(nums) - 1)
        
        copy(newPath, path)
        newPath[len(newPath) - 1] = num
        
        copy(newNums, nums[:i])
        copy(newNums[i:], nums[i + 1:])
        
        dfs(res, newPath, newNums)
    }
}