package main

import "fmt"

// find average given function parameters
// no return statments allowed
func average(a, b, c int) {

	average := (a + b + c) / 3
	fmt.Println("Average = ", average)

}

// find min and maxmimum from the given function parameters
func min_max(a, b, c int) (int, int) {

	min := a
	max := a

	// Find the minimum value
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	// Find the maximum value
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	return min, max

}
