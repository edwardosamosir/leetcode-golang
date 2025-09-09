/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int

	// Step 1: Sort the array first
	sort.Ints(nums)

	// Step 2: Iterate each element as the first element of triplet
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1          // Left pointer
		right := len(nums) - 1 // Right pointer

		// Step 3: Two-pointer scan
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// Found a triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < 0 {
				// Need a larger number
				left++
			} else {
				// Need a smaller number
				right--
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0,0,0]]
}
