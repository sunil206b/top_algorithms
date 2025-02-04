// Given a sorted array of unique integers and a target integer,
// return true if there exists a pair of numbers that sum to target,
// false otherwise. This problem is similar to Two Sum.
// (In Two Sum, the input is not sorted).

// For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13,
// return true because 4 + 9 = 13.

package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 6, 8, 9, 14, 15}
	target1 := 13
	target2 := 43

	fmt.Println("Has target sum pairs found: ", hasPairWithSum(nums, target1))
	fmt.Println("Has target sum pairs found: ", hasPairWithSum(nums, target2))
}

func hasPairWithSum(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
