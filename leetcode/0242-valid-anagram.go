/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert strings to rune slices (so Unicode works correctly)
	sRunes := []rune(s)
	tRunes := []rune(t)

	// Sort both slices
	sort.Slice(sRunes, func(i, j int) bool { return sRunes[i] < sRunes[j] })
	sort.Slice(tRunes, func(i, j int) bool { return tRunes[i] < tRunes[j] })

	// Comparison
	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}

func main() {
	// ASCII test cases
	fmt.Println("ASCII tests:")
	fmt.Println(isAnagram("listen", "silent"))   // true
	fmt.Println(isAnagram("hello", "world"))     // false
	fmt.Println(isAnagram("anagram", "nagaram")) // true

	// Unicode Latin with diacritics
	fmt.Println("\nUnicode Latin tests:")
	fmt.Println(isAnagram("àéîö", "îöàé"))     // true
	fmt.Println(isAnagram("café", "éfac"))     // true
	fmt.Println(isAnagram("résumé", "émusér")) // true

	// Unicode Asian characters
	fmt.Println("\nUnicode Asian tests:")
	fmt.Println(isAnagram("漢字", "字漢"))       // true
	fmt.Println(isAnagram("漢字", "漢漢"))       // false
	fmt.Println(isAnagram("こんにちは", "はちにこん")) // true (Japanese hiragana)

	// Emoji tests
	fmt.Println("\nEmoji tests:")
	fmt.Println(isAnagram("🙂🙃", "🙃🙂"))         // true
	fmt.Println(isAnagram("🙂🙂", "🙃🙃"))         // false
	fmt.Println(isAnagram("👨‍💻👩‍🚀", "👩‍🚀👨‍💻")) // true

	// Mixed Unicode
	fmt.Println("\nMixed Unicode tests:")
	fmt.Println(isAnagram("Hello世界🌍", "🌍界世Hello")) // true
	fmt.Println(isAnagram("Café☕", "☕éfaC"))       // true
}
