package main

import "fmt"

func fibonacciIter(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}

	return curr
}

func fibonacciIterFull(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fibs := []int{0, 1}

	for i := 2; i <= n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}

	return fibs[n]
}

func main() {
	fmt.Println(fibonacciIter(2))
	fmt.Println(fibonacciIter(3))
	fmt.Println(fibonacciIter(4))
	fmt.Println(fibonacciIter(5))
	fmt.Println(fibonacciIter(6))
	fmt.Println(fibonacciIter(7))
	fmt.Println(fibonacciIter(20))

	fmt.Println(fibonacciIterFull(2))
	fmt.Println(fibonacciIterFull(3))
	fmt.Println(fibonacciIterFull(4))
	fmt.Println(fibonacciIterFull(5))
	fmt.Println(fibonacciIterFull(6))
	fmt.Println(fibonacciIterFull(7))
	fmt.Println(fibonacciIterFull(20))
}
