package study

func Factorial(n int) (res int) {
	if n > 0 {
		res = n * Factorial(n-1)
		return res
	}
	return 1
}
