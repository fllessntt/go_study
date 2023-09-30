package study

import "fmt"

func Factorial(n int) (res int) {
	if n > 0 {
		res = n * Factorial(n-1)
		return res
	}
	return 1
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func sqrtRecurse(x, guess, preGuess, epsilon float64) float64 {
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		return guess
	}
	newGuess := (guess + x/guess) / 2
	if newGuess == preGuess {
		return guess
	}
	return sqrtRecurse(x, newGuess, guess, epsilon)
}

func sqrt(x float64) float64 {
	if x < 0 {
		return -1.0
	}
	return sqrtRecurse(x, 1.0, 0.0, 1e-9)
}

func TestSqrt() {
	x := 38.0
	result := sqrt(x)
	fmt.Printf("%.2f 的算术平方根为 %.6f\n", x, result)
}
