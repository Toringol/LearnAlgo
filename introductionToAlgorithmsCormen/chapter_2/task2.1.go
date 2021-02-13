/*
* 2.1
* Несмотря на то что с увеличением количества сортируемых элементов
* время сортировки методом слияний в наихудшем случае растет как O(n * log n),
* а время сортировки вставкой - как O(n^2), благодаря постоянным множителям
* на практике для малых размеров задач на большинстве машин сортировка вставкой
* выполняется быстрее. Таким образов, есть смысл использовать сортировку встаков
* в процессе сортивки методом слияния, когда подзадачи становлятся достаточно
* маленькими. Рассмотрите модификацию алгоритма сортировки слиянием, в котором
* n/k подмассивов длиной k сортируются вставкой, после чего они объединяются с
* помощью обычного механизма слияния. Величина k должна быть найдена в процессе
* решения задачи.
 */

package main

import "fmt"

func insertionSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		curEl := nums[i]

		insertPos := binarySearch(nums, 0, i, curEl)

		for j := i; j > insertPos; j-- {
			nums[j] = nums[j-1]
		}

		nums[insertPos] = curEl
	}

	return nums
}

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

func chooseThreshold(numsLen int) int {
	if numsLen > 32 {
		return 32
	}

	return 2
}

func merge(lNums, rNums []int) []int {
	sortedNums := make([]int, len(lNums)+len(rNums))

	for i, j, k := 0, 0, 0; k < len(sortedNums); k++ {
		if i == len(lNums) {
			sortedNums[k] = rNums[j]
			j++
		} else if j == len(rNums) {
			sortedNums[k] = lNums[i]
			i++
		} else if lNums[i] <= rNums[j] {
			sortedNums[k] = lNums[i]
			i++
		} else {
			sortedNums[k] = rNums[j]
			j++
		}
	}

	return sortedNums
}

func mergeSort(nums []int, threshold int) []int {
	if len(nums) < threshold {
		return insertionSort(nums)
	}

	middle := len(nums) / 2

	lNums := mergeSort(nums[:middle], threshold)
	rNums := mergeSort(nums[middle:], threshold)

	return merge(lNums, rNums)
}

func main() {
	nums := []int{6, 2, 1, 1, -4, -7, 10, 12, 7, 5, -3, 20}

	fmt.Println("Before sort:", nums)

	nums = mergeSort(nums, chooseThreshold(len(nums)))

	fmt.Println("After sort:", nums)
}
