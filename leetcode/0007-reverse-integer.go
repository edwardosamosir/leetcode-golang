/*
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21


Constraints:

-2^31 <= x <= 2^31 - 1
*/

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x != 0 {
		pop := x % 10
		x /= 10

		// Check for overflow before multiplying or adding
		if result > (math.MaxInt32-pop)/10 {
			return 0
		}

		result = result*10 + pop
	}

	return result * sign
}

func main() {
	fmt.Println(reverse(123))        // 321
	fmt.Println(reverse(-123))       // -321
	fmt.Println(reverse(120))        // 21
	fmt.Println(reverse(1534236469)) // 0 (overflow)
}
