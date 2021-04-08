/*
* Given an integer n, return the number of trailing zeroes in n!.
*
* Follow up: Could you write a solution that works in logarithmic time complexity?
*
* Example 1:
*
* Input: n = 3
* Output: 0
* Explanation: 3! = 6, no trailing zero.
*
* Example 2:
*
* Input: n = 5
* Output: 1
* Explanation: 5! = 120, one trailing zero.
*
* Example 3:
*
* Input: n = 0
* Output: 0
*
* Constraints:
*
* 0 <= n <= 104
 */

package main

import "fmt"

func trailingZeroes(n int) int {
	trailingZeroesCounter := 0

	for ; n != 0; n /= 5 {
		trailingZeroesCounter += n / 5
	}

	return trailingZeroesCounter
}

func main() {
	fmt.Println(trailingZeroes(30))
}
