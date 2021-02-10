/*
* 2.1.4
* Рассмотрим задачу сложения двух n-битовых двоичных чисел, хранящихся
* в n-элементных массивах A и B. Сумму этих двух чисел необходимо занести
* в двоичной форме в (n + 1)-элементный массив С. Приведите строгую формулировку
* задачи и составьте псевдокод для сложения этих двух чисел.
 */

package main

import "fmt"

func bitAdd(a, b, carry byte) (byte, byte) {
	sum := a + b + carry

	if sum == 0 {
		return 0, 0
	} else if sum == 1 {
		return 1, 0
	} else if sum == 2 {
		return 0, 1
	}

	return 1, 1
}

func bitSum(bitArrA, bitArrB []byte) []byte {
	var highArr []byte
	var lowArr []byte

	if len(bitArrA) >= len(bitArrB) {
		highArr = bitArrA
		lowArr = bitArrB
	} else {
		highArr = bitArrB
		lowArr = bitArrA
	}

	bitArrC := make([]byte, len(highArr)+1)

	curAdd, curCarry := byte(0), byte(0)

	for i := len(lowArr) - 1; i >= 0; i-- {
		curAdd, curCarry = bitAdd(highArr[i+len(highArr)-len(lowArr)], lowArr[i], curCarry)

		bitArrC[i+len(bitArrC)-len(lowArr)] = curAdd
	}

	if len(highArr) != len(lowArr) {
		for i := len(highArr) - len(lowArr) - 1; i >= 0; i-- {
			curAdd, curCarry = bitAdd(highArr[i], byte(0), curCarry)

			bitArrC[i+1] = curAdd
		}
	}

	bitArrC[0] = curCarry

	return bitArrC
}

func main() {
	bitArrA := []byte{1, 1, 0, 0, 0, 0}
	bitArrB := []byte{0, 0, 0, 1, 1, 1, 1, 0}

	fmt.Println("First bit arr:", bitArrA)
	fmt.Println("First bit arr:", bitArrB)

	fmt.Println(bitSum(bitArrA, bitArrB))
}
