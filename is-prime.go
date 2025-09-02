// Prime Numbers : 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97

package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Prime numbers
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(5))
	fmt.Println(isPrime(7))
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(13))
	fmt.Println(isPrime(17))
	fmt.Println(isPrime(19))
	fmt.Println(isPrime(23))
	fmt.Println(isPrime(29))
	fmt.Println(isPrime(41))
	fmt.Println(isPrime(97))

	// Not prime numbers
	fmt.Println(isPrime(4))
	fmt.Println(isPrime(6))
	fmt.Println(isPrime(8))
	fmt.Println(isPrime(9))
	fmt.Println(isPrime(12))
	fmt.Println(isPrime(14))
	fmt.Println(isPrime(18))
	fmt.Println(isPrime(20))
	fmt.Println(isPrime(24))
	fmt.Println(isPrime(30))
	fmt.Println(isPrime(40))
	fmt.Println(isPrime(99))
}
