package main

func FindTheDifference(s, t string) byte {
	var result byte

	for _, c := range []byte(s) {
		result ^= c
	}

	for _, c := range []byte(t) {
		result ^= c
	}

	return result
}
