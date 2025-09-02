package main

import (
	"fmt"
	"math"
)

func primeFactors(n int) []int {
	factors := []int{}

	// Handle 2 as a special case (smallest prime factor), handle the number of 2s that divide n
	for n%2 == 0 {
		if len(factors) == 0 || factors[len(factors)-1] != 2 {
			factors = append(factors, 2)
		}

		n /= 2
	}

	// Check for odd factors starting from 3
	// n must be odd at this point, thus skip the even numbers and iterate only for odd
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			if len(factors) == 0 || factors[len(factors)-1] != i {
				factors = append(factors, i)
			}
			n /= i
		}
	}

	// If n is still greater than 2, then it must be prime
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	fmt.Println(primeFactors(1))
	fmt.Println(primeFactors(2))
	fmt.Println(primeFactors(3))
	fmt.Println(primeFactors(4))
	fmt.Println(primeFactors(5))
	fmt.Println(primeFactors(6))
	fmt.Println(primeFactors(7))
	fmt.Println(primeFactors(8))
	fmt.Println(primeFactors(9))
	fmt.Println(primeFactors(10))
	fmt.Println(primeFactors(9240)) // [2 3 5 7 11]
	fmt.Println(primeFactors(119))  // [7 17]
}
