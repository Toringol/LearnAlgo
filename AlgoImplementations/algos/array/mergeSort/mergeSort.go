package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middleBound := len(nums) / 2

	return merge(mergeSort(nums[:middleBound]), mergeSort(nums[middleBound:]))
}

func merge(leftArr, rightArr []int) []int {
	leftArrIter, rightArrIter := 0, 0

	mergedArr := make([]int, len(leftArr)+len(rightArr))

	for i := 0; i < len(mergedArr); i++ {
		if leftArrIter == len(leftArr) {
			mergedArr[i] = rightArr[rightArrIter]
			rightArrIter++
		} else if rightArrIter == len(rightArr) {
			mergedArr[i] = leftArr[leftArrIter]
			leftArrIter++
		} else if leftArr[leftArrIter] <= rightArr[rightArrIter] {
			mergedArr[i] = leftArr[leftArrIter]
			leftArrIter++
		} else {
			mergedArr[i] = rightArr[rightArrIter]
			rightArrIter++
		}
	}

	return mergedArr
}

func main() {
	nums := []int{3, 2, 1, 8, 7, -2, -5}

	fmt.Println(mergeSort(nums))
}
