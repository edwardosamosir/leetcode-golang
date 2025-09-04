/*
   You are climbing a staircase. It takes n steps to reach the top.

   Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

   Example 1:

   Input: n = 2
   Output: 2
   Explanation: There are two ways to climb to the top.
   1. 1 step + 1 step
   2. 2 steps
   Example 2:

   Input: n = 3
   Output: 3
   Explanation: There are three ways to climb to the top.
   1. 1 step + 1 step + 1 step
   2. 1 step + 2 steps
   3. 2 steps + 1 step


   Constraints:

   1 <= n <= 45
*/

package main

import (
	"fmt"
)

// Recursive approach
func climbStairsRecursive(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

// Iterative approach (bottom-up / dynamic programming)
func climbStairsIterative(n int) int {
	if n <= 2 {
		return n
	}

	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func main() {
	fmt.Println(climbStairsRecursive(5)) // 8
	fmt.Println(climbStairsIterative(5)) // 8

	fmt.Println(climbStairsRecursive(10)) // 89
	fmt.Println(climbStairsIterative(10)) // 89
}
