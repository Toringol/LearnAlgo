/*
* Merge two sorted linked lists and return it as a sorted list.
* The list should be made by splicing together the nodes of the first two lists.
*
* Example 1:
*
* Input: l1 = [1,2,4], l2 = [1,3,4]
* Output: [1,1,2,3,4,4]
*
* Example 2:
*
* Input: l1 = [], l2 = []
* Output: []
*
* Example 3:
*
* Input: l1 = [], l2 = [0]
* Output: [0]
*
* Constraints:
*
* The number of nodes in both lists is in the range [0, 50].
* -100 <= Node.val <= 100
* Both l1 and l2 are sorted in non-decreasing order.
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	resultSortedList := new(ListNode)
	startListNode := resultSortedList
	var prevLastListNode *ListNode

	for l1 != nil || l2 != nil {
		if l1 == nil {
			resultSortedList.Val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			resultSortedList.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			resultSortedList.Val = l1.Val
			l1 = l1.Next
		} else {
			resultSortedList.Val = l2.Val
			l2 = l2.Next
		}

		resultSortedList.Next = new(ListNode)
		prevLastListNode = resultSortedList
		resultSortedList = resultSortedList.Next
	}

	if prevLastListNode != nil {
		prevLastListNode.Next = nil
	}

	return startListNode
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	for sortedList := mergeTwoLists(list1, list2); sortedList != nil; sortedList = sortedList.Next {
		fmt.Print(sortedList.Val, "->")
	}

}
