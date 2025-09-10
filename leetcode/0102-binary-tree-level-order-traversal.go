/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:

      3
     / \
    9  20
       / \
      15   7
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000


*/

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

		if i < len(arr) && arr[i] != nil {
			current.Left = &TreeNode{Val: *arr[i]}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: *arr[i]}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// Recursive DFS-based level order
func levelOrderDFS(root *TreeNode) [][]int {
	var result [][]int

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(result) == level {
			result = append(result, []int{})
		}

		result[level] = append(result[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}

// Iterative BFS-based level order
func levelOrderIter(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
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
	root1 := buildTree(intPtrSlice([]interface{}{3, 9, 20, nil, nil, 15, 7}))
	fmt.Println(levelOrderDFS(root1))  // ➤ [[3],[9,20],[15,7]]
	fmt.Println(levelOrderIter(root1)) // ➤ [[3],[9,20],[15,7]]

	root2 := buildTree(intPtrSlice([]interface{}{1}))
	fmt.Println(levelOrderDFS(root2))  // ➤ [[1]]
	fmt.Println(levelOrderIter(root2)) // ➤ [[1]]

	root3 := buildTree(intPtrSlice([]interface{}{}))
	fmt.Println(levelOrderDFS(root3))  // ➤ []
	fmt.Println(levelOrderIter(root3)) // ➤ []
}
