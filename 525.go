package main

func FindMaxLength(nums []int) int {

	hashmap := map[int]int{}

	hashmap[0] = -1

	maxLen, count := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count = count + 1
		} else {
			count = count - 1
		}

		if _, ok := hashmap[count]; ok {
			if maxLen < (i - hashmap[count]) {
				maxLen = (i - hashmap[count])
			}
		} else {
			hashmap[count] = i
		}
	}

	return maxLen
}
