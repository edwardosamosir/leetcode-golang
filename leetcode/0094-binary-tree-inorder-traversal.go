/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]

Explanation:

        1
         \
          2
         /
        3

Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [4,2,6,5,7,1,3,9,8]

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



Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?
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

// Recursive inorder traversal
func inorderTraversalRecursive(root *TreeNode) []int {
	result := []int{}
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node != nil {
			traverse(node.Left)               // Left
			result = append(result, node.Val) // Root
			traverse(node.Right)              // Right
		}
	}
	traverse(root)
	return result
}

// Iterative inorder traversal using stack
func inorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// Go as left as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Pop from stack and process
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Move to the right node
		current = current.Right
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

	fmt.Println("Recursive:")
	fmt.Println(inorderTraversalRecursive(tree1)) // [1 3 2]
	fmt.Println(inorderTraversalRecursive(tree2)) // [4 2 6 5 7 1 3 9 8]
	fmt.Println(inorderTraversalRecursive(tree3)) // []
	fmt.Println(inorderTraversalRecursive(tree4)) // [1]

	fmt.Println("Iterative:")
	fmt.Println(inorderTraversalIterative(tree1)) // [1 3 2]
	fmt.Println(inorderTraversalIterative(tree2)) // [4 2 6 5 7 1 3 9 8]
	fmt.Println(inorderTraversalIterative(tree3)) // []
	fmt.Println(inorderTraversalIterative(tree4)) // [1]
}
