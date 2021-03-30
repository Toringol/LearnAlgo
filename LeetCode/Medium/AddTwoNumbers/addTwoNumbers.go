/*
* You are given two non-empty linked lists representing two non-negative integers.
* The digits are stored in reverse order, and each of their nodes contains a single digit.
* Add the two numbers and return the sum as a linked list.
*
* You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*
* Example 1:
*
* Input: l1 = [2,4,3], l2 = [5,6,4]
* Output: [7,0,8]
* Explanation: 342 + 465 = 807.
*
* Example 2:
*
* Input: l1 = [0], l2 = [0]
* Output: [0]
*
* Example 3:
*
* Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
* Output: [8,9,9,9,0,0,0,1]
*
* Constraints:
*
* The number of nodes in each linked list is in the range [1, 100].
* 0 <= Node.val <= 9
* It is guaranteed that the list represents a number that does not have leading zeros.
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersWithCarry(val1, val2, carry int) (int, int) {
	sum := val1 + val2 + carry

	if sum > 9 {
		return sum % 10, 1
	}

	return sum, 0
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	resultList := new(ListNode)
	head := resultList
	var prevLastListNode *ListNode

	addValue := 0
	carry := 0

	for l1 != nil && l2 != nil {
		addValue, carry = addTwoNumbersWithCarry(l1.Val, l2.Val, carry)

		resultList.Val = addValue
		resultList.Next = &ListNode{}
		prevLastListNode = resultList

		resultList = resultList.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		addValue, carry = addTwoNumbersWithCarry(l1.Val, 0, carry)

		resultList.Val = addValue
		resultList.Next = &ListNode{}
		prevLastListNode = resultList

		resultList = resultList.Next

		l1 = l1.Next
	}

	for l2 != nil {
		addValue, carry = addTwoNumbersWithCarry(0, l2.Val, carry)

		resultList.Val = addValue
		resultList.Next = &ListNode{}
		prevLastListNode = resultList

		resultList = resultList.Next

		l2 = l2.Next
	}

	if carry != 0 {
		resultList.Val = carry
	} else {
		prevLastListNode.Next = nil
	}

	return head
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, "->")
		list = list.Next
	}

	fmt.Println()
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	printList(addTwoNumbers(l1, l2))

	l1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}

	l2 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	printList(addTwoNumbers(l1, l2))

	l1 = &ListNode{
		Val: 0,
	}

	l2 = &ListNode{
		Val: 0,
	}

	printList(addTwoNumbers(l1, l2))
}
