/*
* 2.3.7 *
* Разработайте алгоритм со временем работы O(n * log n), который для
* заданного множества S из n целых чисел и другого целого числа x
* определяет, имеются ли в множестве S два элемента, сумма которых
* равна x.
 */

package main

import "fmt"

func twoSum(nums []int, searchSum int) bool {
	numTargetSum := make(map[int]int)

	for _, num := range nums {
		targetSum := searchSum - num
		if _, ok := numTargetSum[targetSum]; ok {
			return true
		}
		numTargetSum[num] = targetSum
	}

	return false
}

func main() {
	nums := []int{3, 3, 1, 7}
	searchSum := 9

	fmt.Println(twoSum(nums, searchSum))
}
