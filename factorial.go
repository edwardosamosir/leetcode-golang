package main

import "fmt"

func factorialRecursive(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func main() {
	fmt.Println(factorialRecursive(4))
	fmt.Println(factorialRecursive(5))
	fmt.Println(factorialRecursive(6))
	fmt.Println(factorialRecursive(7))
	fmt.Println(factorialRecursive(8))

	fmt.Println(factorialIterative(4))
	fmt.Println(factorialIterative(5))
	fmt.Println(factorialIterative(6))
	fmt.Println(factorialIterative(7))
	fmt.Println(factorialIterative(8))
}
