/*
* 2.1.2
* Перепишите процедуру INSERTION-SORT для сортировки в
* невозрастающем порядке вместо неубывающего.
 */

package main

import "fmt"

func insertionSortDecrease(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		curEl := arr[i]

		j := i - 1
		for ; j >= 0 && arr[j] < curEl; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = curEl
	}
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}

	fmt.Println("Start arr:", arr)

	insertionSortDecrease(arr)

	fmt.Println("Arr after sort:", arr)
}
