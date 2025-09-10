/*
Given the root of a binary tree, return the postorder traversal of its nodes' values.



Example 1:

Input: root = [1,null,2,3]

Output: [3,2,1]

Explanation:

        1
         \
          2
         /
        3


Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [4,6,7,5,2,9,8,3,1]

Explanation:

          1
       /      \
      2        3
     / \        \
    4   5        8
       / \      /
      6  7     9

Example 3:

Input: root = []

Output: []

Example 4:

Input: root = [1]

Output: [1]
*/

package main

import (
	"fmt"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree builds a binary tree from level-order slice
func buildTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		current := queue[0]
		queue = queue[1:]

		// Set left child if exists
		if i < len(arr) && arr[i] != nil {
			current.Left = &TreeNode{Val: *arr[i]}
			queue = append(queue, current.Left)
		}
		i++

		// Set right child if exists
		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: *arr[i]}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// Recursive postorder traversal: Left -> Right -> Root
func postorderTraversalRecursive(root *TreeNode) []int {
	result := []int{}
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node != nil {
			traverse(node.Left)               // Left
			traverse(node.Right)              // Right
			result = append(result, node.Val) // Root
		}
	}
	traverse(root)
	return result
}

// Iterative postorder traversal using stack
func postorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Push value first
		result = append(result, node.Val)

		// Push left then right (so right processed before left after reverse)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// Reverse result to get correct postorder: Left -> Right -> Root
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Helper to convert mixed int/nil slice to []*int
func intPtrSlice(vals []interface{}) []*int {
	slice := make([]*int, len(vals))
	for i, v := range vals {
		if v == nil {
			slice[i] = nil
		} else {
			val := v.(int)
			slice[i] = &val
		}
	}
	return slice
}

func main() {
	tree1 := buildTree(intPtrSlice([]interface{}{1, nil, 2, 3}))
	tree2 := buildTree(intPtrSlice([]interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}))
	tree3 := buildTree(intPtrSlice([]interface{}{}))
	tree4 := buildTree(intPtrSlice([]interface{}{1}))

	fmt.Println("Recursive Postorder:")
	fmt.Println(postorderTraversalRecursive(tree1)) // [3 2 1]
	fmt.Println(postorderTraversalRecursive(tree2)) // [4 6 7 5 2 9 8 3 1]
	fmt.Println(postorderTraversalRecursive(tree3)) // []
	fmt.Println(postorderTraversalRecursive(tree4)) // [1]

	fmt.Println("Iterative Postorder:")
	fmt.Println(postorderTraversalIterative(tree1)) // [3 2 1]
	fmt.Println(postorderTraversalIterative(tree2)) // [4 6 7 5 2 9 8 3 1]
	fmt.Println(postorderTraversalIterative(tree3)) // []
	fmt.Println(postorderTraversalIterative(tree4)) // [1]
}
