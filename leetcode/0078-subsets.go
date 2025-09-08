/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/

package main

import "fmt"

// Iterative approach
func subsets(nums []int) [][]int {
	subsets := [][]int{{}} // start with empty subset

	for _, num := range nums {
		newSets := [][]int{}
		for _, set := range subsets {
			newSet := append([]int{}, set...) // copy current subset
			newSet = append(newSet, num)      // add current number
			newSets = append(newSets, newSet)
		}
		subsets = append(subsets, newSets...)
	}

	return subsets
}

// Backtracking approach
func subsetsBacktracking(nums []int) [][]int {
	var result [][]int

	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		tmp := append([]int{}, current...) // copy current subset
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i+1, current)
			current = current[:len(current)-1] // backtrack
		}
	}

	backtrack(0, []int{})
	return result
}

func main() {
	fmt.Println("Iterative:")
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))

	fmt.Println("Backtracking:")
	fmt.Println(subsetsBacktracking([]int{1, 2, 3}))
	fmt.Println(subsetsBacktracking([]int{0}))
}
