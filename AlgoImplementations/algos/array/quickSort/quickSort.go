package main

import (
	"fmt"
	"math/rand"
)

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left, right := 0, len(nums)-1

	pivot := rand.Int() % len(nums)

	nums[pivot], nums[right] = nums[right], nums[pivot]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	quickSort(nums[:left])
	quickSort(nums[left+1:])
}

func main() {
	nums := []int{3, 2, 1, 8, 7, -2, -5}

	quickSort(nums)

	fmt.Println(nums)
}
