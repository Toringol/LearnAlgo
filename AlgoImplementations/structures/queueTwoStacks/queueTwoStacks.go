package main

import (
	"errors"
	"fmt"
)

const stackLength int = 3

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

type Queue struct {
	top    *Stack
	bottom *Stack
}

func (q *Queue) Enqueue(key int) error {
	if q.top.topPointer == stackLength && q.bottom.topPointer == stackLength {
		return errors.New("Queue is full")
	}

	if q.bottom.topPointer == stackLength {
		for q.bottom.topPointer != 0 && q.top.topPointer != stackLength {
			key, err := q.bottom.Pop()
			if err != nil {
				return err
			}

			q.top.Push(key)
		}
	}

	err := q.bottom.Push(key)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.top.topPointer == 0 && q.bottom.topPointer == 0 {
		return 0, errors.New("Queue is empty")
	} else if q.top.topPointer == 0 {
		bottomStackPointer := q.bottom.topPointer

		for i := 0; i < bottomStackPointer; i++ {
			key, err := q.bottom.Pop()
			if err != nil {
				return 0, err
			}

			err = q.top.Push(key)
			if err != nil {
				return 0, err
			}
		}
	}

	key, err := q.top.Pop()
	if err != nil {
		return 0, err
	}

	return key, nil
}

func main() {
	queue := &Queue{
		top:    &Stack{},
		bottom: &Stack{},
	}

	fmt.Println(queue.Dequeue())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	fmt.Println(queue.Enqueue(7))

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	queue.Enqueue(1)
	queue.Enqueue(2)

	fmt.Println(queue.Dequeue())

	queue.Enqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
