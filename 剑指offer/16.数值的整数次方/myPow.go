package myPow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return pow(x, n)
	} else {
		return pow(1/x, -n)
	}
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n&1 == 1 {
		return pow(x, n-1) * x
	} else {
		return pow(x*x, n>>1)
	}
}
