// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "()[temp]{}"
// Output: true

// Constraints:
// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.

// Hint 1:
// Use a stack of characters.

// Hint 2:
// When you encounter an opening bracket, push it to the top of the stack.

// Hint 3:
// When you encounter a closing bracket, check if the top of the stack was the opening for it. If yes, pop it from the stack. Otherwise, return false.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("()[temp]{}"))
}

// If it is an opening bracket, we push it onto the stack.
//
// If it is a closing bracket, we check if the stack is empty or
// if the top element of the stack matches the corresponding opening bracket.
// If it does, we pop the top element from the stack.
//
// After iterating through the entire string, the string is considered
// valid if the stack is empty (all opening brackets have been matched
// and closed in the correct order). Otherwise, it is considered invalid.
func isValid(s string) bool {
	stack := make([]rune, 0)

	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		if isOpeningBracket(ch) {
			stack = append(stack, ch)
		} else if isClosingBracket(ch) {
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop the first element
		}
	}
	return len(stack) == 0
}

func isOpeningBracket(ch rune) bool {
	return ch == '(' || ch == '{' || ch == '['
}

func isClosingBracket(ch rune) bool {
	return ch == ')' || ch == '}' || ch == ']'
}
