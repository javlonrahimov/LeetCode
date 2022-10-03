package main

func nextPermutation(nums []int) {
	var switchPos, minPos int
	minPos = -1
	for switchPos = len(nums) - 1; switchPos > 0; switchPos-- {
		if nums[switchPos] > nums[switchPos-1] {
			break
		}
	}
	switchPos--
	if switchPos != -1 {
		for i := len(nums) - 1; i > switchPos; i-- {
			if nums[i] > nums[switchPos] && (minPos == -1 || nums[i] < nums[minPos]) {
				minPos = i
			}
		}
		nums[switchPos], nums[minPos] = nums[minPos], nums[switchPos]
	}
	sort.Ints(nums[switchPos+1:])
}
