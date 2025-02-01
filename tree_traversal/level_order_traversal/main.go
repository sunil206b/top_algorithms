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
	fmt.Println("Level-Order Traversal (Iterative):")
	levelOrderIterative(root)
}

func levelOrderIterative(root *TreeNode) {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
