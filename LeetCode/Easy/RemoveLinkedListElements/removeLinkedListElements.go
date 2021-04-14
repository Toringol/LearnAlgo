/*
* Given the head of a linked list and an integer val,
* remove all the nodes of the linked list that has Node.val == val, and return the new head.
*
* Example 1:
*
* Input: head = [1,2,6,3,4,5,6], val = 6
* Output: [1,2,3,4,5]
*
* Example 2:
*
* Input: head = [], val = 1
* Output: []
*
* Example 3:
*
* Input: head = [7,7,7,7], val = 7
* Output: []
*
* Constraints:
*
* The number of nodes in the list is in the range [0, 104].
* 1 <= Node.val <= 50
* 0 <= k <= 50
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	start := head

	for head != nil && head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	if start != nil && start.Val == val {
		return start.Next
	}

	return start
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, "->")
		list = list.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}

	printList(head)

	head = removeElements(head, 6)

	printList(head)

	head = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
				},
			},
		},
	}

	head = removeElements(head, 7)

	printList(head)

	head = nil

	head = removeElements(head, 1)

	printList(head)
}
