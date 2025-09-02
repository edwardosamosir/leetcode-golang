// Prime Numbers : 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97

package main

import (
	"fmt"
	"math"
)

func isPrimeSqrt(n int) bool {
	if n <= 1 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Prime numbers
	fmt.Println(isPrimeSqrt(2))
	fmt.Println(isPrimeSqrt(3))
	fmt.Println(isPrimeSqrt(5))
	fmt.Println(isPrimeSqrt(7))
	fmt.Println(isPrimeSqrt(11))
	fmt.Println(isPrimeSqrt(13))
	fmt.Println(isPrimeSqrt(17))
	fmt.Println(isPrimeSqrt(19))
	fmt.Println(isPrimeSqrt(23))
	fmt.Println(isPrimeSqrt(29))
	fmt.Println(isPrimeSqrt(41))
	fmt.Println(isPrimeSqrt(97))

	// Not prime numbers
	fmt.Println(isPrimeSqrt(4))
	fmt.Println(isPrimeSqrt(6))
	fmt.Println(isPrimeSqrt(8))
	fmt.Println(isPrimeSqrt(9))
	fmt.Println(isPrimeSqrt(12))
	fmt.Println(isPrimeSqrt(14))
	fmt.Println(isPrimeSqrt(18))
	fmt.Println(isPrimeSqrt(20))
	fmt.Println(isPrimeSqrt(24))
	fmt.Println(isPrimeSqrt(30))
	fmt.Println(isPrimeSqrt(40))
	fmt.Println(isPrimeSqrt(99))
}
