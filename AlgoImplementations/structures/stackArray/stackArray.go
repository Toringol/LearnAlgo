package main

import (
	"errors"
	"fmt"
)

const stackLength int = 256

type Stack struct {
	array      [stackLength]int
	topPointer int
}

func (st *Stack) Push(key int) error {
	if st.topPointer == stackLength {
		return errors.New("Stack is full")
	}

	st.array[st.topPointer] = key
	st.topPointer++

	return nil
}

func (st *Stack) Pop() (int, error) {
	if st.topPointer == 0 {
		return 0, errors.New("Empty stack")
	}

	st.topPointer--

	return st.array[st.topPointer], nil
}

func main() {
	stack := &Stack{}

	for i := 0; i <= 256; i++ {
		fmt.Println(stack.Push(i))
	}

	stack = &Stack{}

	stack.Push(12)
	stack.Push(13)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
