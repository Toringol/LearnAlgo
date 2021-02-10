package main

import "fmt"

func insertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		curEl := arr[i]

		j := i - 1
		for ; j >= 0 && curEl < arr[j]; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = curEl
	}
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}

	fmt.Println("Start arr:", arr)

	insertionSort(arr)

	fmt.Println("Arr after sort:", arr)
}
