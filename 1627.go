package main

func MaximumWealth(accounts [][]int) int {

	maxWealth := 0
	for _, i := range accounts {
		s := 0
		for _, j := range i {
			s += j
		}
		if s > maxWealth {
			maxWealth = s
		}
	}

	return maxWealth

}
