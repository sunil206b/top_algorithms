// In In-Order Traversal, we visit the nodes in the following order:
// 1 -> Traverse the left subtree
// 2 -> Visit the root node
// 3 -> Traverse the right subtree
// Similar to DFS approach it uses stack

package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {
	// Constructing the tree:
	//        1
	//       / \
	//      2   3
	//     / \   \
	//    4   5   6
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}},
	}

	fmt.Println("In-Order Traversal (Recursive):")
	inOrderRecursive(root)

	fmt.Println("\nIn-Order Traversal (Iterative):")
	inOrderIterative(root)
}

func inOrderRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	inOrderRecursive(root.Left)
	fmt.Print(root.Val, " ")
	inOrderRecursive(root.Right)
}

func inOrderIterative(root *TreeNode) {
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(current.Val, " ")

		current = current.Right
	}
}
