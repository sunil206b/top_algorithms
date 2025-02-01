package main

import "fmt"

// You are given an integer array nums consisting of n elements, and an integer k.

// Find a contiguous subarray whose length is equal to k that has
// the maximum average value and return this value. Any answer
// with a calculation error less than 10-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000

func main() {
	// nums := []int{1, 12, -5, -6, 50, 3}
	// k := 4

	// nums := []int{5}
	// k := 1

	// nums := []int{-1, -12, -5, -6, -50, -3}
	// k := 2

	nums := []int{4, 4, 4, 4}
	k := 2

	fmt.Printf("Maximum average: %.5f\n", findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	currentSum := 0

	// Calculate initial sum of the first k elements.
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	// Initialize the maximum sum as the sum of the first k elements
	maxSum := currentSum

	// Slide the window through the array
	for i := k; i < n; i++ {
		// Update the sum by removing the first element of the previous window
		// and adding the new element
		currentSum = currentSum - nums[i-k] + nums[i]

		// Update the maximum sum
		if maxSum < currentSum {
			maxSum = currentSum
		}
	}

	return float64(maxSum) / float64(k)
}
