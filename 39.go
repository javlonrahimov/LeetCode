package main

import "sort"

func CombinationSum(candidates []int, target int) [][]int {

	result := [][]int{}
	sub := []int{}
	sort.Ints(candidates)
	helper39(candidates, target, &result, sub)

	return result
}

func helper39(candidates []int, target int, result *[][]int, sub []int) {

	if len(candidates) == 0 {
		return
	}
	if candidates[0] == target {
		sub = append(sub, candidates[0])
		*result = append(*result, sub)
		return
	} else if candidates[0] < target {
		helper39(candidates[1:], target, result, sub)
		sub2 := make([]int, len(sub))
		copy(sub2, sub)
		sub2 = append(sub2, candidates[0])
		helper39(candidates, target-candidates[0], result, sub2)
	} else {
		return
	}
}
