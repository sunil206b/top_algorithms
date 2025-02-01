package main

import "fmt"

func main() {
	x := 27
	fmt.Println(NearestIntegerCubeRoot(x))
	x1 := 30
	fmt.Println(NearestIntegerCubeRoot(x1))
}

func NearestIntegerCubeRoot(x int) int {
	low := 0
	high := x
	result := 0

	for low <= high {
		mid := low + (high-low)/2

		midCubed := mid * mid * mid
		if midCubed == x {
			return mid
		}

		if midCubed < x {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
