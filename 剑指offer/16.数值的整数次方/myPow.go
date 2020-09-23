package myPow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return pow(x, n)
	} else {
		return 1 / pow(x, -n)
	}
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	v := myPow(x, n/2)
	res := v * v
	if n%2 == 1 {
		res *= x
	}
	return res
}
