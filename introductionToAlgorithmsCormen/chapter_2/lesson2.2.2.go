/*
* 2.2.2
* Рассмотрим сортировку элементов массива А, которая выполняется следующим
* образом. Сначала определяется наименьший элемент массива А, который
* становится на место элемента А[1]. Затем производится поиск второго
* наименьшего элемента массива А, который становится на место А[2]. Этот
* процесс продолжается для первых n - 1 элементов массива А. Запишите
* псевдокод этого алгоритма, известного как сортировка выбором (selection sort).
* Какой инвариант цикла сохраняется для этого алгоритма? Почему его достаточно
* выполнить для первых n - 1 элементов, а не для всех n элементов? Определите
* время работы алгоритма в наилучшем и наихудшем случаях и запишите его в
* О-обозначениях.
 */

package main

import "fmt"

func searchIdxMinElem(nums []int, startPos int) int {
	minIdxElem := startPos

	for i := startPos + 1; i < len(nums); i++ {
		if nums[i] < nums[minIdxElem] {
			minIdxElem = i
		}
	}

	return minIdxElem
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdxElem := searchIdxMinElem(nums, i)
		nums[i], nums[minIdxElem] = nums[minIdxElem], nums[i]
	}
}

func main() {
	nums := []int{5, 1, 3, 5, 6, 4, -2}

	fmt.Println("Start arr:", nums)

	selectionSort(nums)

	fmt.Println("After sort:", nums)
}
