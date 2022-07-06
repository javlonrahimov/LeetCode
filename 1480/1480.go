package main

func runningSum(nums []int) []int {
	runningSum := 0
	runningSums := []int{}
	for _, value := range nums {
		runningSum += value
		runningSums = append(runningSums, runningSum)
	}
	return runningSums
}
