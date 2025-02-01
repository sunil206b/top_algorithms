package main

import "fmt"

// Given a string s, find the length of the longest substring
// without repeating characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3 ("abc")
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1 ("b")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3 ("wke")
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
}

func lengthOfLongestSubstring(s string) int {
	// Map to store last seen character and it's index
	charIndices := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		// If the character is already in the substring, update the start position
		if index, exist := charIndices[s[end]]; exist && start <= index {
			start = index + 1
		}

		// Update the character's last seen index
		charIndices[s[end]] = end

		// Update the maximum length
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}
