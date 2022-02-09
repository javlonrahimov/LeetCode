package main

func FindPairs(nums []int, k int) int {
	ans, tmp := 0, make(map[int]int)
	for _, v := range nums {
		tmp[v]++
	}
	for i, v := range tmp {
		_, ok := tmp[i+k]
		if k > 0 && ok || k == 0 && v >= 2 {
			ans++
		}
	}
	return ans
}
