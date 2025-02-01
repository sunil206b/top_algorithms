package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{10, 20, 15, 25, 30, 5}
	fmt.Println(BinarySearch(25, nums1))
	nums2 := []int{6, 9, 3, 1, 15, 12, 18}
	fmt.Println(BinarySearch(15, nums2))
	nums3 := []int{6, 9, 3, 1, 15, 12, 18}
	fmt.Println(BinarySearch(10, nums3))
}

func BinarySearch(target int, nums []int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
