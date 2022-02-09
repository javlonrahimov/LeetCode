package main

func AddDigits(num int) int {
	seen := make(map[int]bool)
	for {
		if _, ok := seen[num]; !ok {
			seen[num] = true
			num = digitSum(num)
		} else {
			break
		}
	}

	return num
}

func digitSum(num int) int {
	total := 0
	for num > 0 {
		total += num % 10
		num /= 10
	}
	return total
}
