/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.



Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).


Constraints:

3 <= nums.length <= 500
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Sort the array to use two-pointer technique
	sort.Ints(nums)

	// Initialize closest with the sum of the first three elements
	closest := nums[0] + nums[1] + nums[2]

	// Loop for the first element
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		// Two-pointer search for the other two elements
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Update closest if current sum is nearer to target
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}

			// Move pointers depending on sum vs target
			if sum < target {
				left++ // Need a larger sum
			} else {
				right-- // Need a smaller sum
			}
		}
	}

	return closest
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // Output: 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      // Output: 0
	fmt.Println(threeSumClosest([]int{1, 1, 1, 1}, 3))   // Output: 3
}
