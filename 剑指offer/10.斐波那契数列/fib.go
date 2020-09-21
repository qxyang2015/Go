package fib

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return (Fib(n-1) + Fib(n-2)) % 1000000007
}

func Fib2(n int) int {
	if n < 2 {
		return n
	}
	var res int
	step1 := 0
	step2 := 1
	for i := 2; i <= n; i++ {
		res = (step1 + step2) % 1000000007
		step1 = step2
		step2 = res
	}
	return res
}
