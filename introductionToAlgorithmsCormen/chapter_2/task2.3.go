/*
* 2.3
* Следующий фрагмент кода реализует правило Горнера для
* вычисления полинома P(x) = E(k = 0 to n) a(k) * x^k =
* a0 + x * (a1 + x * (a2 + ... + x * (a(n - 1) + x * an)))
* для заданных коэффециентов a0, a1, ..., an и значения x.
 */

package main

import (
	"fmt"
	"math"
)

func gornerSum(coefs []int, x int) int {
	y := 0

	for _, coef := range coefs {
		y = coef + x*y
	}

	return y
}

func polinomSum(coefs []int, x int) int {
	y := 0

	for i, coef := range coefs {
		y += coef * int(math.Pow(float64(x), float64(len(coefs)-i-1)))
	}

	return y
}

func main() {
	coefs := []int{1, 2, 3, 4, 5}
	x := 3

	// 1 * 3 ^ 4 + 2 * 3 ^ 3 + 3 * 3 ^ 2 + 4 * 3 + 5 =
	// 81 + 54 + 27 + 12 + 5 =
	// 135 + 39 + 5 =
	// 179

	fmt.Println(gornerSum(coefs, x))
	fmt.Println(polinomSum(coefs, x))
}
