/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.



Example 1:

Input: root = [1,null,2,3]

Output: [1,2,3]

Explanation:

        1
         \
          2
         /
        3

Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [1,2,4,5,6,7,3,8,9]

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

import "fmt"

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

// Recursive preorder traversal: Root -> Left -> Right
func preorderTraversalRecursive(root *TreeNode) []int {
	result := []int{}
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node != nil {
			result = append(result, node.Val) // Root
			traverse(node.Left)               // Left
			traverse(node.Right)              // Right
		}
	}
	traverse(root)
	return result
}

// Iterative preorder traversal using stack
func preorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		// Pop last element
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Process root
		result = append(result, node.Val)

		// Push right first, then left to process left first
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// Helper to create pointer slice from ints (for easier tree building)
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

	fmt.Println("Recursive Preorder:")
	fmt.Println(preorderTraversalRecursive(tree1)) // [1 2 3]
	fmt.Println(preorderTraversalRecursive(tree2)) // [1 2 4 5 6 7 3 8 9]
	fmt.Println(preorderTraversalRecursive(tree3)) // []
	fmt.Println(preorderTraversalRecursive(tree4)) // [1]

	fmt.Println("Iterative Preorder:")
	fmt.Println(preorderTraversalIterative(tree1)) // [1 2 3]
	fmt.Println(preorderTraversalIterative(tree2)) // [1 2 4 5 6 7 3 8 9]
	fmt.Println(preorderTraversalIterative(tree3)) // []
	fmt.Println(preorderTraversalIterative(tree4)) // [1]
}
