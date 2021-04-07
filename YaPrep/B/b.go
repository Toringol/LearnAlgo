package main

import "fmt"

func scanArr(nums []int) {
	for i := range nums {
		fmt.Scan(&nums[i])
	}
}

func longestOneSequence(nums []int) int {
	longestLen, currentLen := 0, 0

	for _, num := range nums {
		if num == 1 {
			currentLen++
		} else {
			if currentLen > longestLen {
				longestLen = currentLen
			}
			currentLen = 0
		}
	}

	if currentLen > longestLen {
		return currentLen
	}

	return longestLen
}

func main() {
	var length int

	fmt.Scan(&length)

	nums := make([]int, length)

	scanArr(nums)

	fmt.Println(longestOneSequence(nums))
}
