/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1
*/

package main

import (
	"fmt"
	"strconv"
)

func isPalindromeStringConv(x int) bool {
	// Convert the number to string
	str := strconv.Itoa(x)

	// Reverse the string
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	// Compare original and reversed
	return str == reversed
}

func isPalindrome(x int) bool {
	// Negative numbers cannot be palindromes
	if x < 0 {
		return false
	}

	// Numbers ending in 0 cannot be palindromes unless the number is 0
	if x%10 == 0 && x != 0 {
		return false
	}

	reversed := 0
	for x > reversed {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	// Check for even-length or odd-length numbers
	return x == reversed || x == reversed/10
}

func main() {
	fmt.Println(isPalindromeStringConv(121))   // true
	fmt.Println(isPalindromeStringConv(-121))  // false
	fmt.Println(isPalindromeStringConv(10))    // false
	fmt.Println(isPalindromeStringConv(12321)) // true
	fmt.Println(isPalindromeStringConv(0))     // true

	fmt.Println(isPalindrome(121))   // true
	fmt.Println(isPalindrome(-121))  // false
	fmt.Println(isPalindrome(10))    // false
	fmt.Println(isPalindrome(12321)) // true
	fmt.Println(isPalindrome(0))     // true
}
