package main

import "fmt"

// Remove duplicates from int slice
func removeDuplicateInts(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

// Remove duplicates from string slice
func removeDuplicateStrings(arr []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, str := range arr {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

// Fungsi generik untuk menghapus duplikat dari slice
func removeDuplicatesGeneric[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	// Test int slices
	fmt.Println(removeDuplicateInts([]int{1, 2, 2, 3}))
	fmt.Println(removeDuplicateInts([]int{}))
	fmt.Println(removeDuplicateInts([]int{5, 5, 5}))

	// Test string slices
	fmt.Println(removeDuplicateStrings([]string{"1", "1", "2"}))
	fmt.Println(removeDuplicateStrings([]string{"a", "b", "a", "c", "b"}))

	// Test dengan slice integer
	fmt.Println(removeDuplicatesGeneric([]int{1, 2, 2, 3}))
	fmt.Println(removeDuplicatesGeneric([]int{}))
	fmt.Println(removeDuplicatesGeneric([]int{5, 5, 5}))

	// Test dengan slice string
	fmt.Println(removeDuplicatesGeneric([]string{"1", "1", "2"}))
	fmt.Println(removeDuplicatesGeneric([]string{"a", "b", "a", "c", "b"}))

	// Test dengan slice float
	fmt.Println(removeDuplicatesGeneric([]float64{1.1, 1.1, 2.2, 3.3}))

	// Test dengan slice bool
	fmt.Println(removeDuplicatesGeneric([]bool{true, false, true, true}))
}
