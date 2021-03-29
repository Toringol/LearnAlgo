/*
* Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
* Return the linked list sorted as well.
*
* Example 1:
*
* Input: head = [1,1,2]
* Output: [1,2]
*
* Example 2:
*
* Input: head = [1,1,2,3,3]
* Output: [1,2,3]
*
* Constraints:
*
* The number of nodes in the list is in the range [0, 300].
* -100 <= Node.val <= 100
* The list is guaranteed to be sorted in ascending order.
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	startNode := head

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return startNode
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, "->")
		list = list.Next
	}
	fmt.Println()
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	list3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	newList1 := deleteDuplicates(list1)
	printList(newList1)

	newList2 := deleteDuplicates(list2)
	printList(newList2)

	newList3 := deleteDuplicates(list3)
	printList(newList3)
}
