/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

package main

import (
	"fmt"
)

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Convert array to linked list
func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

// Convert linked list to array
func linkedListToArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Pretty print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d → ", head.Val)
		head = head.Next
	}
	fmt.Println("null")
}

// Merge two sorted linked lists
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Append the remaining nodes
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func main() {
	list1 := arrayToLinkedList([]int{1, 2, 4})
	list2 := arrayToLinkedList([]int{1, 3, 4})

	merged := mergeTwoLists(list1, list2)

	// Pretty output
	printList(merged) // Output: 1 → 1 → 2 → 3 → 4 → 4 → null

	// As array
	fmt.Println(linkedListToArray(merged)) // Output: [1 1 2 3 4 4]
}
