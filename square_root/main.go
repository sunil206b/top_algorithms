package main

import "fmt"

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
// The returned integer should be non-negative as well.

// You must not use any built-in exponent function or operator.

// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

// Example 1:
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.

// Example 2:
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

// Constraints:
// 0 <= x <= 231 - 1

func main() {
	x := 4
	fmt.Println(NearestIntegerSquareRoot(x))
	x1 := 8
	fmt.Println(NearestIntegerSquareRoot(x1))
}

func NearestIntegerSquareRoot(x int) int {
	low := 0
	high := x
	result := 0

	for low <= high {
		mid := low + (high-low)/2

		midSquared := mid * mid
		if midSquared == x {
			return mid
		}

		if midSquared < x {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}
