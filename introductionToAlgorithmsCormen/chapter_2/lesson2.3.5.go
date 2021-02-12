/*
* 2.3.5
* Возвращаясь к задаче поиска (см. упр. 2.1.3), нетрудно заметить, что если
* последовательность А отсортирована, то можно сравнить значение среднего
* элемента этой последовательности с искомым значением v и сразу исключить
* половину последовательности из дальнейшего рассмотрения. Бинарный поиск
* (binary search) - это алгоритм, в котором такая процедура повторяется
* неоднократно, что всякий раз приводит к уменьшению оставшейся части
* последовательности в два раза. Запишите псевдокод алгоритма бинарного
* поиска (либо итеративный, либо рекурсивный). Докажите, что время работы
* этого алгоритма в наихудшем случае составляет O(log n).
 */

package main

import "fmt"

func binarySearch(nums []int, searchVal int) int {
	startPos := 0
	endPos := len(nums)

	for startPos < endPos {
		middle := (startPos + endPos) / 2

		if nums[middle] == searchVal {
			return middle
		} else if nums[middle] < searchVal {
			startPos = middle + 1
		} else {
			endPos = middle
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(binarySearch(nums, 1))
}
