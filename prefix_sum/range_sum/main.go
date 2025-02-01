package main

import "fmt"

// Range Sum Query - Immutable

// Given an integer array nums, handle multiple queries of the following type:

// Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the NumArray class:

// NumArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

// Example 1:
// Input
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// Output
// [null, 1, -1, -3]

// NumArray struct to store the prefix sum array
type NumArray struct {
	prefix []int
}

// Constructor initializes the NumArray object and computes the prefix sum array
func Constructor(nums []int) *NumArray {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	return &NumArray{prefix: prefix}
}

// SumRange calculates the sum of elements between indices left and right inclusive
func (na *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return na.prefix[right]
	}
	return na.prefix[right] - na.prefix[left-1]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)

	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
