/*

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.


Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100

*/

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Build a binary tree from a slice of *int (nil represents null)
func buildTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))
	for i, val := range arr {
		if val != nil {
			nodes[i] = &TreeNode{Val: *val}
		}
	}

	j := 1
	for i := 0; i < len(arr) && j < len(arr); i++ {
		if nodes[i] != nil {
			if j < len(arr) {
				nodes[i].Left = nodes[j]
				j++
			}
			if j < len(arr) {
				nodes[i].Right = nodes[j]
				j++
			}
		}
	}

	return nodes[0]
}

// Calculate max depth using math.Max
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + int(math.Max(float64(left), float64(right)))
}

func main() {
	// Helper to create pointer easily
	p := func(x int) *int { return &x }

	arr := []*int{
		p(3), p(9), p(20), nil, nil, p(15), p(7),
	}

	root := buildTree(arr)
	fmt.Println(maxDepth(root)) // Output: 3

	arr2 := []*int{
		p(1), nil, p(2),
	}
	root2 := buildTree(arr2)
	fmt.Println(maxDepth(root2)) // Output: 2
}
