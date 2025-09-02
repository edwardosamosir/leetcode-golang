package main

import (
	"fmt"
	"math"
)

func isPrimeSqrtWhileLoop(n int) bool {
	if n <= 1 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	i := 2

	for i <= limit {
		if n%i == 0 {
			return false
		}
		i++
	}

	return true
}

func main() {
	// Prime numbers
	fmt.Println(isPrimeSqrtWhileLoop(2))
	fmt.Println(isPrimeSqrtWhileLoop(3))
	fmt.Println(isPrimeSqrtWhileLoop(5))
	fmt.Println(isPrimeSqrtWhileLoop(7))
	fmt.Println(isPrimeSqrtWhileLoop(11))
	fmt.Println(isPrimeSqrtWhileLoop(13))
	fmt.Println(isPrimeSqrtWhileLoop(17))
	fmt.Println(isPrimeSqrtWhileLoop(19))
	fmt.Println(isPrimeSqrtWhileLoop(23))
	fmt.Println(isPrimeSqrtWhileLoop(29))
	fmt.Println(isPrimeSqrtWhileLoop(41))
	fmt.Println(isPrimeSqrtWhileLoop(97))

	// Not prime numbers
	fmt.Println(isPrimeSqrtWhileLoop(4))
	fmt.Println(isPrimeSqrtWhileLoop(6))
	fmt.Println(isPrimeSqrtWhileLoop(8))
	fmt.Println(isPrimeSqrtWhileLoop(9))
	fmt.Println(isPrimeSqrtWhileLoop(12))
	fmt.Println(isPrimeSqrtWhileLoop(14))
	fmt.Println(isPrimeSqrtWhileLoop(18))
	fmt.Println(isPrimeSqrtWhileLoop(20))
	fmt.Println(isPrimeSqrtWhileLoop(24))
	fmt.Println(isPrimeSqrtWhileLoop(30))
	fmt.Println(isPrimeSqrtWhileLoop(40))
	fmt.Println(isPrimeSqrtWhileLoop(99))
}
