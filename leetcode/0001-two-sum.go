/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

*/

package main

import (
	"fmt"
	"sort"
)

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumTwoPointersWithPairs(nums []int, target int) []int {
	numsWithIndex := make([][2]int, len(nums))
	for i, num := range nums {
		numsWithIndex[i] = [2]int{num, i}
	}

	sort.Slice(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i][0] < numsWithIndex[j][0]
	})

	left, right := 0, len(numsWithIndex)-1

	for left < right {
		sum := numsWithIndex[left][0] + numsWithIndex[right][0]
		if sum == target {
			return []int{numsWithIndex[left][1], numsWithIndex[right][1]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func twoSumTwoPointersWithMaps(nums []int, target int) []int {
	// Step 1: Build array of maps [{ "num": x, "index": i }]
	numsWithIndex := []map[string]int{}
	for i, num := range nums {
		numsWithIndex = append(numsWithIndex, map[string]int{"num": num, "index": i})
	}

	// Step 2: Sort by "num"
	sort.Slice(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i]["num"] < numsWithIndex[j]["num"]
	})

	// Step 3: Two pointer search
	left, right := 0, len(numsWithIndex)-1
	for left < right {
		sum := numsWithIndex[left]["num"] + numsWithIndex[right]["num"]

		if sum == target {
			return []int{numsWithIndex[left]["index"], numsWithIndex[right]["index"]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func twoSumHashMap(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, found := numToIndex[complement]; found {
			return []int{idx, i}
		}

		numToIndex[num] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSumBruteForce([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumBruteForce([]int{3, 2, 4}, 6))
	fmt.Println(twoSumBruteForce([]int{3, 3}, 9))

	fmt.Println(twoSumTwoPointersWithPairs([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumTwoPointersWithPairs([]int{3, 2, 4}, 6))
	fmt.Println(twoSumTwoPointersWithPairs([]int{3, 3}, 9))

	fmt.Println(twoSumTwoPointersWithMaps([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumTwoPointersWithMaps([]int{3, 2, 4}, 6))
	fmt.Println(twoSumTwoPointersWithMaps([]int{3, 3}, 9))

	fmt.Println(twoSumHashMap([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumHashMap([]int{3, 2, 4}, 6))
	fmt.Println(twoSumHashMap([]int{3, 3}, 9))
}
