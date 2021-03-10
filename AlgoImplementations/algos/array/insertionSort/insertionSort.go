package main

import "fmt"

func insertionSort(numbers []int) {
	if len(numbers) < 1 {
		return
	}

	for i := 1; i < len(numbers); i++ {
		for j := i; j > 0 && numbers[j-1] > numbers[j]; j-- {
			numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
		}
	}
}
func main() {
	numbers := []int{3, 1, 2, 5, 4}

	insertionSort(numbers)

	fmt.Println(numbers)
}
