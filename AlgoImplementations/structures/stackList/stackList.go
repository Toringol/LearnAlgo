package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

type Stack struct {
	head *ListNode
}

func (st *Stack) Push(key int) {
	if st.head == nil {
		st.head = &ListNode{
			value: key,
			next:  nil,
		}

		return
	}

	newHead := &ListNode{
		value: key,
		next:  st.head,
	}

	st.head = newHead
}

func (st *Stack) Pop() (int, error) {
	if st.head == nil {
		return 0, errors.New("Stack is empty")
	}

	newHead := st.head.next
	value := st.head.value
	st.head = newHead

	return value, nil
}

func main() {
	stack := &Stack{}

	fmt.Println(stack.Pop())

	stack.Push(12)
	stack.Push(13)
	stack.Push(14)
	stack.Push(15)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
