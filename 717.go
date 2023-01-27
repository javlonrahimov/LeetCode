package main

func isOneBitCharacter(bits []int) bool {
	last := 0

	for i := 0; i < len(bits); i++ {
		if bits[i] == 0 {
			last = 0
		} else if bits[i] == 1 {
			last = 1
			i += 1
		}
	}

	return last == 0
}
