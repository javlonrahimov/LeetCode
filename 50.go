package main

func MyPow(x float64, n int) float64 {

	switch {
	case n == 0:
		return float64(1)
	case n == 1:
		return x
	case n < 0:
		return 1 / MyPow(x, -n)
	}
	if n&1 == 0 {
		return MyPow(x*x, n>>1)
	} else {
		return MyPow(x*x, n>>1) * x
	}
}
