/*
* Given an integer x, return true if x is palindrome integer.
*
* An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
*
* Example 1:
*
* Input: x = 121
* Output: true
*
* Example 2:
*
* Input: x = -121
* Output: false
* Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
*
* Example 3:
*
* Input: x = 10
* Output: false
* Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*
* Example 4:
*
* Input: x = -101
* Output: false
*
* Constraints:
*
* -231 <= x <= 231 - 1
 */

package main

import "fmt"

const maxInt32 int = 1<<31 - 1
const minInt32 int = -(1 << 31)

func reverseInt(x int) int {
	res, rem := 0, 0

	for x != 0 {
		rem = x % 10
		res = rem + res*10
		x /= 10
	}

	if res > maxInt32 || res < minInt32 {
		return 0
	}

	return res
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverseInt(x)
}

func main() {
	fmt.Println(isPalindrome(10))
}
