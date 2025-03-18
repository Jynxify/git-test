package maths

import (
	"fmt"
	"testing"
)

// Function to test addition
func TestAddition(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{0, 0, 0},
		{-1, -1, -2},
		{1000, 2000, 3000},
		{-1000, 1000, 0},
	}

	for _, test := range tests {
		result := test.a + test.b
		if result != test.expected {
			t.Errorf("Expected %d + %d to equal %d, but got %d", test.a, test.b, test.expected, result)
		}
	}
}

// Function to test subtraction
func TestSubtraction(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 0},
		{0, 0, 0},
		{-1, -1, 0},
		{1000, 2000, -1000},
		{-1000, 1000, -2000},
	}

	for _, test := range tests {
		result := test.a - test.b
		if result != test.expected {
			t.Errorf("Expected %d - %d to equal %d, but got %d", test.a, test.b, test.expected, result)
		}
	}
}

// Function to test multiplication
func TestMultiplication(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 1},
		{0, 0, 0},
		{-1, -1, 1},
		{1000, 2000, 2000000},
		{-1000, 1000, -1000000},
	}

	for _, test := range tests {
		result := test.a * test.b
		if result != test.expected {
			t.Errorf("Expected %d * %d to equal %d, but got %d", test.a, test.b, test.expected, result)
		}
	}
}

// Function to test division
func TestDivision(t *testing.T) {
	tests := []struct {
		a, b, expected int
		errExpected    bool
	}{
		{1, 1, 1, false},
		{0, 1, 0, false},
		{-1, -1, 1, false},
		{1000, 2000, 0, false},
		{-1000, 1000, -1, false},
		{1, 0, 0, true}, // Division by zero
	}

	for _, test := range tests {
		result, err := safeDivide(test.a, test.b)
		if test.errExpected {
			if err == nil {
				t.Errorf("Expected error for %d / %d, but got none", test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect error for %d / %d, but got %v", test.a, test.b, err)
			}
			if result != test.expected {
				t.Errorf("Expected %d / %d to equal %d, but got %d", test.a, test.b, test.expected, result)
			}
		}
	}
}

// Helper function to safely divide two integers
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
