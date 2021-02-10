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
	lowLen := 0
	highLen := 0
	var highArr []byte
	var lowArr []byte

	if len(bitArrA) >= len(bitArrB) {
		highLen = len(bitArrA)
		lowLen = len(bitArrB)
		highArr = bitArrA
		lowArr = bitArrB
	} else {
		highLen = len(bitArrB)
		lowLen = len(bitArrA)
		highArr = bitArrB
		lowArr = bitArrA
	}

	bitArrC := make([]byte, highLen+1)

	curAdd, curCarry := byte(0), byte(0)

	if highLen == lowLen {
		for i := highLen - 1; i >= 0; i-- {
			curAdd, curCarry = bitAdd(bitArrA[i], bitArrB[i], curCarry)

			bitArrC[i+1] = curAdd
		}
	} else {
		for i := lowLen - 1; i >= 0; i-- {
			curAdd, curCarry = bitAdd(highArr[i+highLen-lowLen], lowArr[i], curCarry)

			bitArrC[i+1] = curAdd
		}

		for i := highLen - lowLen - 1; i >= 0; i-- {
			curAdd, curCarry = bitAdd(highArr[i], byte(0), curCarry)

			bitArrC[i+1] = curAdd

			if curCarry == 1 && i == 0 {
				curCarry = 1
			} else {
				curCarry = 0
			}
		}

	}

	bitArrC[0] = curCarry

	return bitArrC
}

func main() {
	bitArrA := []byte{1, 0, 1, 1, 0, 1}
	bitArrB := []byte{1, 1, 1, 1, 1}

	fmt.Println("First bit arr:", bitArrA)
	fmt.Println("First bit arr:", bitArrB)

	fmt.Println(bitSum(bitArrA, bitArrB))
}
