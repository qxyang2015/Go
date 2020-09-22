package NumWays

func NumWays(n int) int {
	if n < 2 {
		return 1
	}
	step1 := 1
	step2 := 1
	for i := 2; i <= n; i++ {
		step1, step2 = step2, (step1+step2)%1000000007
	}
	return step2
}
