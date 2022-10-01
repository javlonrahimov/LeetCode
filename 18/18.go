package main

import "sort"

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    l := len(nums)
    set := make(map[[4]int]struct{})
    
    var res [][]int
    for i1 := 0; i1 < l; i1++ {
        for i2 := i1 + 1; i2 < l; i2++ {
            arr := [4]int{nums[i1], nums[i2], 0, 0}
            i3, i4 := i2 + 1, l - 1
            for i3 < i4 {
                sum := nums[i1] + nums[i2] + nums[i3] + nums[i4]
                if sum < target {
                    i3++
                } else if sum > target {
                    i4--
                } else {
                    arr[2], arr[3] = nums[i3], nums[i4]
                    if _, ok := set[arr]; !ok {
                        set[arr] = struct{}{}
                        res = append(res, []int{arr[0], arr[1], arr[2], arr[3]})
                    }
                    i3++
                    for i3 < i4 && nums[i3] == nums[i3 - 1] {
                        i3++
                    }
                }
            }
        }
    }
    return res
}
