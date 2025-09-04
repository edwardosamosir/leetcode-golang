/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([])"
Output: true



Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}

	// Map for matching closing brackets to opening brackets
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, isClose := bracketMap[char]; isClose {
			// If it's a closing bracket, check last element in stack
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			// It's an opening bracket, push to stack
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("]["))     // false
}
