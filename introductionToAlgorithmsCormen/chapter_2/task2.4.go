/*
* 2.4
* Пусть А[1..n] представляет собой массив из n различных чисел.
* Если i < j и A[i] > A[j], то пара (i, j) называется инверсией А.
* Разработайте алгоритм, определяющий количество инверсий, содержащихся
* в произвольной перестановке n элементов, время работы которого в
* наихудшем случае равно O(n * log n). (Указание: модифицируйте алгоритм
* сортировки слиянием.)
 */

package main

// var inversionCounter int

func merge(lNums, rNums []int) []int {
	// fmt.Println(lNums)
	// fmt.Println(rNums)

	// innerInversionCounter := 0
	// multipleCounter := 0
	// emptyRNumArr := false

	sortedNums := make([]int, len(lNums)+len(rNums))

	for i, j, k := 0, 0, 0; k < len(lNums)+len(rNums); k++ {
		if i == len(lNums) {
			sortedNums[k] = rNums[j]
			j++
		} else if j == len(rNums) {
			sortedNums[k] = lNums[i]
			i++
			// if !emptyRNumArr {
			// 	emptyRNumArr = true
			// 	multipleCounter = innerInversionCounter
			// } else {
			// 	innerInversionCounter += multipleCounter
			// }
		} else if lNums[i] <= rNums[j] {
			sortedNums[k] = lNums[i]
			i++
		} else {
			sortedNums[k] = rNums[j]
			j++
			// innerInversionCounter++
			// fmt.Println("InnerCounter:", innerInversionCounter)
		}
	}

	// fmt.Println(innerInversionCounter)

	// inversionCounter += innerInversionCounter

	// fmt.Println(inversionCounter)

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
	nums := []int{3, 6, 5, 4, 3, 2, 1}

	mergeSort(nums)

	// fmt.Println(inversionCounter)
}
