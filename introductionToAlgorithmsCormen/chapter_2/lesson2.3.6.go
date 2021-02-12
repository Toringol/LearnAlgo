/*
* 2.3.6
* Заметим, что в цикле while в строках 5-7 процедуры INSERTION-SORT
* в разделе 2.1 для сканирования (в обратном порядке) отсортированного
* подмассива А[1..j - 1] используется линейный поиск. Можно ли использовать
* бинарный поиск (см. упр. 2.3.5) вместо линейного, чтобы время работы этого
* алгоритма в наихудшем случае улучшилось и стало равным O(n * log n).
 */

package main

import "fmt"

func binarySearch(nums []int, startPos, endPos, searchVal int) int {
	for startPos < endPos {
		middle := (startPos + endPos) / 2

		if nums[middle] < searchVal {
			startPos = middle + 1
		} else {
			endPos = middle
		}
	}

	return startPos
}

func insertionSortUsingBinarySearch(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		curEl := arr[i]

		insertPos := binarySearch(arr, 0, i, curEl)

		for j := i; j > insertPos; j-- {
			arr[j] = arr[j-1]
		}

		arr[insertPos] = curEl
	}
}

func main() {
	arr := []int{5, 1, 1, -2, 1, 3}

	fmt.Println("Start arr:", arr)

	insertionSortUsingBinarySearch(arr)

	fmt.Println("Arr after sort:", arr)
}
