/*
* 2.2
* Пузырьковая сортировка представляет собой популярный, но не
* эффективный алгоритм сортировки. В его основе лежит многократная
* перестановка соседних элементов, нарушающих порядок сортировки.
 */

package main

import "fmt"

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func main() {
	nums := []int{6, 2, 1, 1, -4, -7, 10, 12, 7, 5, -3, 20}

	fmt.Println("Before sort:", nums)

	bubbleSort(nums)

	fmt.Println("After sort:", nums)
}
