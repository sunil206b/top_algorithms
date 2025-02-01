// Pre-order traversal follows this order:
// 1 -> Visit the root
// 2 -> Traverse the left subtree
// 3 -> Traverse the right subtree
// Similar to DFS approach it uses stack

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

	fmt.Println("Pre-Order Traversal (Recursive):")
	preOrderRecursive(root)

	fmt.Println("\nPre-Order Traversal (Iterative):")
	preOrderIterative(root)
}

func preOrderRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val, " ")
	preOrderRecursive(root.Left)
	preOrderRecursive(root.Right)
}

func preOrderIterative(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(node.Val, " ")

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}
