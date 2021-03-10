package main

import (
	"errors"
	"fmt"
)

func binarySearch(numbers []int, searchKey int) (int, error) {
	leftBound := 0
	rightBound := len(numbers)

	for leftBound < rightBound {
		middle := (leftBound + rightBound) / 2

		curNumber := numbers[middle]

		if curNumber == searchKey {
			return middle, nil
		} else if searchKey < curNumber {
			rightBound = middle
		} else {
			leftBound = middle + 1
		}
	}

	return 0, errors.New("No such key in arr ")
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(binarySearch(numbers, 1))
	fmt.Println(binarySearch(numbers, 2))
	fmt.Println(binarySearch(numbers, 3))
	fmt.Println(binarySearch(numbers, 4))
	fmt.Println(binarySearch(numbers, 5))
	fmt.Println(binarySearch(numbers, 6))
	fmt.Println(binarySearch(numbers, -1))
}
