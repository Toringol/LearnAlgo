/*
* Given a binary tree, determine if it is height-balanced.
*
* For this problem, a height-balanced binary tree is defined as:
*
* a binary tree in which the left and right subtrees of
* every node differ in height by no more than 1.
*
* Example 1:
*
* Input: root = [3,9,20,null,null,15,7]
* Output: true
*
* Example 2:
*
* Input: root = [1,2,2,3,3,null,null,4,4]
* Output: false
*
* Example 3:
*
* Input: root = []
* Output: true
*
* Constraints:
*
* The number of nodes in the tree is in the range [0, 5000].
* -104 <= Node.val <= 104
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftDepth := checkDepths(root.Left)
	rightDepth := checkDepths(root.Right)

	if leftDepth > rightDepth+1 || rightDepth > leftDepth+1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func checkDepths(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := checkDepths(root.Left)
	rightDepth := checkDepths(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isBalanced(tree))
}
