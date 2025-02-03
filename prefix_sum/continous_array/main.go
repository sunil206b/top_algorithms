package main

import "fmt"

// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

// Example 1:
// Input: nums = [0,1]
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

// Example 2:
// Input: nums = [0,1,0]
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

func findMaxLength(nums []int) int {
	one, zero := 0, 0
	maxLength := 0
	diffIndices := make(map[int]int)

	for i, num := range nums {
		if num == 0 {
			zero += 1
		} else {
			one += 1
		}

		diff := one - zero
		if _, ok := diffIndices[diff]; !ok {
			diffIndices[diff] = i
		}

		if one == zero {
			maxLength = one + zero
		} else {
			prevIndex := diffIndices[diff]
			maxLength = max(maxLength, i-prevIndex)
		}
	}
	return maxLength
}

func main() {
	nums1 := []int{0, 1}
	maxLength := findMaxLength(nums1)
	fmt.Println(maxLength)
	nums2 := []int{0, 1, 0}
	maxLength = findMaxLength(nums2)
	fmt.Println(maxLength)
	nums3 := []int{1, 1, 1, 0, 0}
	maxLength = findMaxLength(nums3)
	fmt.Println(maxLength)
}
