/*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.



Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]


Constraints:

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

*/

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int

	// Step 1: Sort the array to make two-pointer approach possible
	sort.Ints(nums)

	n := len(nums)

	// Step 2: Loop for the first number (i)
	for i := 0; i < n-3; i++ {
		// Skip duplicates for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Step 3: Loop for the second number (j)
		for j := i + 1; j < n-2; j++ {
			// Skip duplicates for the second number
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Step 4: Two-pointer initialization
			low := j + 1
			high := n - 1

			// Step 5: Find remaining two numbers that sum up to target
			for low < high {
				sum := nums[i] + nums[j] + nums[low] + nums[high]

				if sum == target {
					// Step 6: Found quadruplet
					result = append(result, []int{nums[i], nums[j], nums[low], nums[high]})

					// Skip duplicates for low pointer
					for low < high && nums[low] == nums[low+1] {
						low++
					}

					// Skip duplicates for high pointer
					for low < high && nums[high] == nums[high-1] {
						high--
					}

					// Move both pointers to continue search
					low++
					high--
				} else if sum < target {
					// Step 7: Sum too small, move low pointer to right
					low++
				} else {
					// Step 8: Sum too large, move high pointer to left
					high--
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // [[2,2,2,2]]
}
