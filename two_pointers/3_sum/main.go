package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	// nums := []int{1, 2, 3}

	// nums := []int{-1, -1, -1, 2, 2, 2, 0, 0, 0}

	// nums := []int{0, 0, 0, 0}

	fmt.Println("Triplets:", threeSum(nums))
}

func threeSum(nums []int) [][]int {
	// List to store the each sub array that satisfies the given condition.
	res := [][]int{}
	// Sort the array
	sort.Ints(nums)

	n := len(nums)

	// Iterate over the number until len(nums)-2, because we need triplets
	for i := 0; i < n-2; i++ {
		// In the sorted list, if the nums[i] and nums[i-1] are same continue
		// the loop to avoid duplicates.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Now use the Two pointers approach to find the other two numbers
		l, r := i+1, n-1
		for l < r {
			// Calculate the sum of triplets
			sum := nums[i] + nums[l] + nums[r]

			// If sum equals to zero, then add triplets list to the response list.
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				// There might be more than two duplicates in the nums list -3, -3, -3.
				// In the top iteration we skipped only first duplicate, this will skip
				// if there are any other duplicates.
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
