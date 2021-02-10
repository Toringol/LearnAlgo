/*
* 2.1.3
* Рассмотрим задачу поиска.
* Вход. Последовательность из n чисел A = <a1, a2, ..., an> и значение v.
* Выход. Индекс i, такой, что v = A[i], или специальное значение NIL, если
* v в A отсутствует.
*
* Составьте псевдокод линейного поиска, при работе которого выполняется
* сканирование последовательности в поисках значения v.
 */

package main

import "fmt"

func linearSearch(arr []int, v int) (int, bool) {
	for i, value := range arr {
		if value == v {
			return i, true
		}
	}

	return 0, false
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}

	index, ok := linearSearch(arr, 4)
	if ok {
		fmt.Println("Correct search:", index)
	} else {
		fmt.Println("Seach is working wrong")
	}

	index, ok = linearSearch(arr, 10)
	if ok {
		fmt.Println("Seach is working wrong")
	} else {
		fmt.Println("Correct search", ok)
	}
}
