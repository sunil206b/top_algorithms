package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "racecar"
	s2 := "Radar"
	s3 := "notpalindrome"

	fmt.Println("'racecar' is palindrome: ", isPalindrome(s1))
	fmt.Println("'Radar' is palindrome: ", isPalindrome(s2))
	fmt.Println("'notpalindrome' is palindrome: ", isPalindrome(s3))
}

func isPalindrome(s string) bool {
	tempS := strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		if tempS[left] != tempS[right] {
			return false
		}

		left++
		right--
	}
	return true
}
