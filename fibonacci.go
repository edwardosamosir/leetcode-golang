package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciWithMemo(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	memo[n] = fibonacciWithMemo(n-1, memo) + fibonacciWithMemo(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
	fmt.Println(fibonacci(20))

	memo := make(map[int]int)
	fmt.Println(fibonacciWithMemo(2, memo))
	fmt.Println(fibonacciWithMemo(3, memo))
	fmt.Println(fibonacciWithMemo(4, memo))
	fmt.Println(fibonacciWithMemo(5, memo))
	fmt.Println(fibonacciWithMemo(6, memo))
	fmt.Println(fibonacciWithMemo(7, memo))
	fmt.Println(fibonacciWithMemo(20, memo))
}
