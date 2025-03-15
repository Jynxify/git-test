package math

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func Modulo(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}

func Power(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func Factorial(a int) int {
	result := 1
	for i := 1; i <= a; i++ {
		result *= i
	}
	return result
}

func Fibonacci(a int) int {
	if a <= 1 {
		return a
	}
	return Fibonacci(a-1) + Fibonacci(a-2)
}
