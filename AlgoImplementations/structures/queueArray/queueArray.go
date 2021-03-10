package main

import (
	"errors"
	"fmt"
)

const queueLength int = 3

type Queue struct {
	array       [queueLength]int
	headPointer int
	tailPointer int
	count       int
}

func (q *Queue) Enqueue(key int) error {
	if q.headPointer == q.tailPointer && q.count != 0 {
		return errors.New("Queue is full")
	}

	q.array[q.tailPointer] = key
	q.tailPointer = (q.tailPointer + 1) % queueLength
	q.count++

	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.count == 0 {
		return 0, errors.New("Queue is empty")
	}

	key := q.array[q.headPointer]
	q.headPointer = (q.headPointer + 1) % queueLength
	q.count--

	return key, nil
}

func main() {
	queue := &Queue{}

	fmt.Println(queue.Dequeue())

	fmt.Println(queue.Enqueue(1))
	fmt.Println(queue.Enqueue(2))
	fmt.Println(queue.Enqueue(3))
	fmt.Println(queue.Enqueue(4))

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
