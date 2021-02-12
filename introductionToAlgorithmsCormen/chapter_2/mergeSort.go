package main

import (
	"fmt"
)

func merge(lNums, rNums []int) []int {
	sortedNums := make([]int, len(lNums)+len(rNums))

	for i, j, k := 0, 0, 0; k < len(lNums)+len(rNums); k++ {
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

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	part := len(nums) / 2

	lNums := mergeSort(nums[:part])
	rNums := mergeSort(nums[part:])

	return merge(lNums, rNums)
}

func main() {
	nums := []int{5, 3, 2, 1, 6, 7, 2, 1, -6}

	fmt.Println("Start arr:", nums)

	nums = mergeSort(nums)

	fmt.Println("After sort:", nums)
}
