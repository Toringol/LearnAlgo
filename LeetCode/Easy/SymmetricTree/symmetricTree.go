/*
* Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*
* Example 1:
*
* Input: root = [1,2,2,3,4,4,3]
* Output: true
*
* Example 2:
*
* Input: root = [1,2,2,null,3,null,3]
* Output: false
*
* Constraints:
*
* The number of nodes in the tree is in the range [1, 1000].
* -100 <= Node.val <= 100
*
* Follow up: Could you solve it both recursively and iteratively?
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return mirrorCheck(root, root)
}

func mirrorCheck(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val &&
		mirrorCheck(left.Left, right.Right) &&
		mirrorCheck(left.Right, right.Left)
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(isSymmetric(tree))

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(isSymmetric(tree))
}
