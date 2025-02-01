// Post-Order Traversal follows this order:
// 1 -> Traverse the left subtree
// 2 -> Traverse the right subtree
// 3 -> Visit the root node

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
	// 	fmt.Println(fmt.Sprint(
	// 		```
	// 							1
	// 			2                3
	// 	4      	5        6       7
	// 8		9   10 11    12 13   14 15
	// 		```
	// ))
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}},
	}

	fmt.Println("Post-Order Traversal (Recursive):")
	postOrderRecursive(root)
}

func postOrderRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	postOrderRecursive(root.Left)
	postOrderRecursive(root.Right)
	fmt.Print(root.Val, " ")
}
