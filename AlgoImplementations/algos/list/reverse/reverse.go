package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

// func (l *ListNode) Reverse() *ListNode {
// 	curr := l
// 	var top *ListNode = nil

// 	for curr != nil {
// 		tmp := curr.next
// 		curr.next = top
// 		top = curr
// 		curr = tmp
// 	}

// 	return top
// }

// a -> b -> c
// curr = a top = nil tmp = b a.next = nil top = a curr = b
// tmp = c curr.next = nil top = b curr = c
//

// func (l *ListNode) Reverse() *ListNode {
// 	curr := l
// 	var newNext *ListNode = nil

// 	for curr != nil {
// 		tmp := curr.next
// 		curr.next = newNext
// 		newNext = curr
// 		curr = tmp
// 	}

// 	return newNext
// }

func (l *ListNode) Reverse() *ListNode {
	curNode := l
	var nextNode *ListNode

	for curNode != nil {
		curNode, nextNode, curNode.next = curNode.next, curNode, nextNode
	}

	return nextNode
}

func (l *ListNode) PrintList() {
	for l != nil {
		fmt.Println(l.value)
		l = l.next
	}
}

func main() {
	root := &ListNode{}

	root.value = 1
	root.next = &ListNode{
		value: 2,
		next: &ListNode{
			value: 3,
			next:  nil,
		},
	}

	root.PrintList()

	root = root.Reverse()

	root.PrintList()

}
