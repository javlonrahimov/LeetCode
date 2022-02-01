package main

func MaxProfit(prices []int) int {

	maxProfit := 0

	if len(prices) <= 1 {
		return maxProfit
	}
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		temp := prices[i] - min
		if temp > maxProfit {
			maxProfit = temp
		}
		if min > prices[i] {
			min = prices[i]
		}
	}

	return maxProfit
}
