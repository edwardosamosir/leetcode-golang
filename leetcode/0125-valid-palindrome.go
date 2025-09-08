/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindromeReverse(s string) bool {
	// Buat regex untuk menghapus semua karakter non-alfanumerik
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// Ubah ke lowercase dan bersihkan
	normalized := reg.ReplaceAllString(strings.ToLower(s), "")

	// Balik string
	reversed := reverseString(normalized)

	return normalized == reversed
}

// Fungsi bantu membalik string
func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func isPalindromeTwoPointers(s string) bool {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	normalized := reg.ReplaceAllString(strings.ToLower(s), "")

	left, right := 0, len(normalized)-1

	for left < right {
		if normalized[left] != normalized[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isPalindromeForLoop(s string) bool {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	normalized := reg.ReplaceAllString(strings.ToLower(s), "")

	n := len(normalized)
	for i := 0; i < n/2; i++ {
		if normalized[i] != normalized[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Reverse approach:")
	fmt.Println(isPalindromeReverse("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindromeReverse("race a car"))                     // false
	fmt.Println(isPalindromeReverse(" "))                              // true

	fmt.Println("\nTwo pointers approach:")
	fmt.Println(isPalindromeTwoPointers("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindromeTwoPointers("race a car"))
	fmt.Println(isPalindromeTwoPointers(" "))

	fmt.Println("\nFor loop approach:")
	fmt.Println(isPalindromeForLoop("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindromeForLoop("race a car"))
	fmt.Println(isPalindromeForLoop(" "))
}
