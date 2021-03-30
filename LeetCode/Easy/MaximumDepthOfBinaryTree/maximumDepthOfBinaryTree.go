/*
* Given the root of a binary tree, return its maximum depth.
*
* A binary tree's maximum depth is the number of nodes along
* the longest path from the root node down to the farthest leaf node.
*
* Example 1:
*
* Input: root = [3,9,20,null,null,15,7]
* Output: 3
*
* Example 2:
*
* Input: root = [1,null,2]
* Output: 2
*
* Example 3:
*
* Input: root = []
* Output: 0
*
* Example 4:
*
* Input: root = [0]
* Output: 1
*
* Constraints:
*
* The number of nodes in the tree is in the range [0, 104].
* -100 <= Node.val <= 100
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSize := maxDepth(root.Left)
	rightSize := maxDepth(root.Right)

	if leftSize > rightSize {
		return leftSize + 1
	} else {
		return rightSize + 1
	}
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(maxDepth(tree))

	tree = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(maxDepth(tree))

	tree = nil

	fmt.Println(maxDepth(tree))

	tree = &TreeNode{
		Val: 0,
	}

	fmt.Println(maxDepth(tree))
}
