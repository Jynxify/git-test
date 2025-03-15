package math

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-2, -3, -5},
		{0, 0, 0},
		{1000, 2000, 3000},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{-5, -3, -2},
		{0, 0, 0},
		{1000, 500, 500},
	}

	for _, tt := range tests {
		result := Subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, -3, 6},
		{0, 5, 0},
		{100, 200, 20000},
	}

	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{6, 3, 2},
		{6, 0, 0},
		{-6, -3, 2},
		{100, 10, 10},
	}

	for _, tt := range tests {
		result := Divide(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestModulo(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{5, 0, 0},
		{-5, 3, -2},
		{100, 10, 0},
	}

	for _, tt := range tests {
		result := Modulo(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Modulo(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 8},
		{2, 0, 1},
		{0, 5, 0},
		{5, 2, 25},
	}

	for _, tt := range tests {
		result := Power(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Power(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		a, expected int
	}{
		{5, 120},
		{0, 1},
		{1, 1},
		{10, 3628800},
	}

	for _, tt := range tests {
		result := Factorial(tt.a)
		if result != tt.expected {
			t.Errorf("Factorial(%d) = %d; want %d", tt.a, result, tt.expected)
		}
	}
}
