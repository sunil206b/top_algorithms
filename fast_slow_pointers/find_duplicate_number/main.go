package main

import "fmt"

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.

// You must solve the problem without modifying the array nums and using only constant extra space.

// Example 1:
// Input: nums = [1,3,4,2,2]
// Output: 2

// Example 2:
// Input: nums = [3,1,3,4,2]
// Output: 3

// Example 3:
// Input: nums = [3,3,3,3,3]
// Output: 3

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{1, 1}))
	fmt.Println(findDuplicate([]int{1, 2, 2, 2}))
}

func findDuplicate(nums []int) int {
	// Use Floyd's Cycle Detection Algorithm
	slow, fast := 0, 0

	// Step 1: Find the intersection point of the cycle
	for {
		slow = nums[slow]       // Move slow pointer by 1 step
		fast = nums[nums[fast]] // Move fast pointer by 2 steps

		if slow == fast {
			break
		}
	}

	// Step 2: Find the entrance to the cycle (duplicate number)
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]

		if slow == slow2 {
			break
		}
	}
	return slow
}
