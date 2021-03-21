package main

import "fmt"

func fib(number int) int {
	if number <= 1 {
		return number
	}

	return fib(number-1) + fib(number-2)
}

func fibFast(number int) int {
	fibNumbers := make(map[int]int)

	for i := 0; i <= number; i++ {
		var curVal int

		if i <= 2 {
			curVal = 1
		} else {
			curVal = fibNumbers[i-1] + fibNumbers[i-2]
		}

		fibNumbers[i] = curVal
	}

	return fibNumbers[number]
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(fibFast(5))
}

// 5 3 2 1 4
// arr[left] = {3, 2, 1}
// arr [right] = {5}
// resultArr = arr[left] + pivot + arr[right]
