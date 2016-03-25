package math

func Factorial(i int) int {
	factorial_memo := 1
	for i > 1 {
		factorial_memo = factorial_memo * i
		i -= 1
	}
	return factorial_memo
}

func Fibonacci(i int) int {
	if i <= 1 {
		return i
	} else {
		return Fibonacci(i-1) + Fibonacci(i-2)
	}
}
