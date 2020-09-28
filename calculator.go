// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes a variatic parameter of numbers and returns the result of multiplying them.
func Multiply(nums ...float64) float64 {

	if len(nums) == 0 {
		return 0.0
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		total := nums[0]
		for _, num := range nums[1:] {
			total = total * num
		}
		return total
	}
}

// Divide returns the result of dividing b from a
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero error")
	}
	return a / b, nil
}

// Sqrt return the square root of a number
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("number cannot be less than zero")
	}
	return math.Sqrt(a), nil
}
