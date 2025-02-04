// Given two sorted integer arrays arr1 and arr2, return a new array
// that combines both of them and is also sorted.

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6, 8, 9, 14, 15}
	arr2 := []int{3, 5, 7, 8}

	fmt.Println("Combined sorted array: ", combineArrays(arr1, arr2))
}

func combineArrays(arr1 []int, arr2 []int) []int {
	n1 := len(arr1)
	n2 := len(arr2)

	result := []int{}

	i, j := 0, 0

	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < n1 {
		result = append(result, arr1[i])
		i++
	}

	for j < n2 {
		result = append(result, arr2[j])
		j++
	}

	return result
}
