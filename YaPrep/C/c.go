package main

import "fmt"

func scanArr(nums []int) {
	for i := range nums {
		fmt.Scan(&nums[i])
	}
}

func deleteDublicates(nums []int) {
	if len(nums) == 0 {
		return
	}

	prevEl := nums[0]

	fmt.Println(prevEl)

	for i := 1; i < len(nums); i++ {
		if prevEl != nums[i] {
			fmt.Println(nums[i])
			prevEl = nums[i]
		}
	}
}

func main() {
	var length int

	fmt.Scan(&length)

	nums := make([]int, length)

	scanArr(nums)

	deleteDublicates(nums)
}
