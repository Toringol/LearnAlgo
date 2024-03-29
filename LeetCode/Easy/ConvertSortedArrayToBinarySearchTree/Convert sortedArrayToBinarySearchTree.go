/*
* Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
*
* A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
*
* Example 1:
*
* Input: nums = [-10,-3,0,5,9]
* Output: [0,-3,9,-10,null,5]
* Explanation: [0,-10,5,null,-3,null,9] is also accepted:
*
* Example 2:
*
* Input: nums = [1,3]
* Output: [3,1]
* Explanation: [1,3] and [3,1] are both a height-balanced BSTs.
*
* Constraints:
*
* 1 <= nums.length <= 104
* -104 <= nums[i] <= 104
* nums is sorted in a strictly increasing order.
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

func InorderWalk(tree *TreeNode) {
	if tree.Left != nil {
		InorderWalk(tree.Left)
	}

	fmt.Println(tree.Val)

	if tree.Right != nil {
		InorderWalk(tree.Right)
	}
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}

	tree := sortedArrayToBST(nums)

	InorderWalk(tree)
}
