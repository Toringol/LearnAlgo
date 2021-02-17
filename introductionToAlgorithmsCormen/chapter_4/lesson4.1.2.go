/*
* 4.1.2
* Напишите псевдокод для прешения задачи поиска максимального подмассива
* методом грубой силы. Ваша процедура должна выполняться за вермя O(n^2).
 */

package main

import "fmt"

func findMaxSubArrayBruteForce(nums []int) (int, int, int) {
	leftMaxIndex, rightMaxIndex, maxSumArrSum := 0, 0, 0

	actualSum := 0
	for i := range nums {
		actualSum = 0
		for j := i; j < len(nums); j++ {
			actualSum += nums[j]
			if actualSum > maxSumArrSum {
				maxSumArrSum = actualSum
				leftMaxIndex = i
				rightMaxIndex = j
			}
		}
	}

	return leftMaxIndex, rightMaxIndex, maxSumArrSum
}

func main() {
	nums := []int{13, -3, -25, 20, -3, -16, -23,
		18, 20, -7, 12, -5, -22, 15, -4, 7}

	fmt.Println(findMaxSubArrayBruteForce(nums))
}
