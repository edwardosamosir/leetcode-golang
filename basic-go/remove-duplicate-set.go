package main

import (
	"fmt"
)

// removeDuplicatesSet menghapus duplikat dari slice dengan konsep HashSet
func removeDuplicatesSet[T comparable](arr []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(arr)) // alokasi awal sesuai panjang input

	for _, v := range arr {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{} // tambahkan ke "set"
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Contoh slice integer
	fmt.Println(removeDuplicatesSet([]int{1, 2, 2, 3, 4, 1})) // [1 2 3 4]

	// Contoh slice string
	fmt.Println(removeDuplicatesSet([]string{"a", "b", "a", "c"})) // [a b c]

	// Contoh slice float
	fmt.Println(removeDuplicatesSet([]float64{1.1, 1.1, 2.2, 3.3})) // [1.1 2.2 3.3]

	// Contoh slice boolean
	fmt.Println(removeDuplicatesSet([]bool{true, false, true})) // [true false]
}
