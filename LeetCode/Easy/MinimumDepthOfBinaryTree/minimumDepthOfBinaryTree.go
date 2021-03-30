/*
* Given a binary tree, find its minimum depth.
*
* The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
*
* Note: A leaf is a node with no children.
*
* Example 1:
*
* Input: root = [3,9,20,null,null,15,7]
* Output: 2
*
* Example 2:
*
* Input: root = [2,null,3,null,4,null,5,null,6]
* Output: 5
*
* Constraints:
*
* The number of nodes in the tree is in the range [0, 105].
* -1000 <= Node.val <= 1000
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	nodes := []*TreeNode{root}

	for {
		newNodes := []*TreeNode{}

		for _, treeNode := range nodes {
			if treeNode.Left == nil && treeNode.Right == nil {
				return depth
			}

			if treeNode.Left != nil {
				newNodes = append(newNodes, treeNode.Left)
			}

			if treeNode.Right != nil {
				newNodes = append(newNodes, treeNode.Right)
			}
		}

		depth++
		nodes = newNodes
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

	fmt.Println(minDepth(tree))
}
