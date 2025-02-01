package main

import "fmt"

// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,1,1], k = 2
// Output: 2

// Example 2:
// Input: nums = [1,2,3], k = 3
// Output: 2

func main() {
	nums := []int{1, 1, 1}
	k := 2

	// nums := []int{1, 2, 3}
	// k := 3

	fmt.Println("Number of subarrays:", subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	// Map to store prefix sums
	prefixSums := make(map[int]int)
	prefixSums[0] = 1 // Where prefix sum itself equals k

	count := 0
	currentSum := 0

	for _, num := range nums {
		currentSum += num

		if freq, exist := prefixSums[currentSum-k]; exist {
			count += freq
		}
		prefixSums[currentSum]++
	}
	return count
}
