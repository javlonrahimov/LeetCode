package main

func removeKDigits(num string, k int) string {
	stack := make([]byte, 0)

	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	var zeros int
	for i := range stack {
		if stack[i] != '0' {
			break
		}
		zeros++
	}
    if len(stack) == 0 || zeros == len(stack) {
		return "0"
	}
	return string(stack[zeros:])
}