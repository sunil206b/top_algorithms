package main

import "fmt"

// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in the list that can be
// reached again by continuously following the next pointer. Internally, pos is used
// to denote the index of the node that tail's next pointer is connected to.
// Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.

// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

// Example 2:
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

// Example 3:
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create a linked list with a cycle
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Cycle here
	fmt.Println("Has Cycle:", hasCycle(node1))

	// Create a linked list without a cycle
	nodeA := &ListNode{Val: 1}
	nodeB := &ListNode{Val: 2}
	nodeA.Next = nodeB

	fmt.Println("Has Cycle:", hasCycle(nodeA))
}

func hasCycle(head *ListNode) bool {
	// Use the two pointer: slow (tortoise) and fast (hare)
	slow, fast := head, head

	// Iterate over the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer by 1 step
		fast = fast.Next.Next // Move fast pointer by 2 steps

		// If slow and fast meet, there is a cycle
		if slow == fast {
			return true
		}
	}

	// If we reach end of the list, then there is not cycle
	return false
}
