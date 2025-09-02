package main

import "fmt"

func fibonacciDP(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	fmt.Println(fibonacciDP(2))
	fmt.Println(fibonacciDP(3))
	fmt.Println(fibonacciDP(4))
	fmt.Println(fibonacciDP(5))
	fmt.Println(fibonacciDP(6))
	fmt.Println(fibonacciDP(7))
	fmt.Println(fibonacciDP(20))
}
