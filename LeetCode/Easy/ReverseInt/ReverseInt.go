/*
* Given a signed 32-bit integer x, return x with its digits reversed.
* If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
*
* Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*
* Example 1:
*
* Input: x = 123
* Output: 321
*
* Example 2:
*
* Input: x = -123
* Output: -321
*
* Example 3:
*
* Input: x = 120
* Output: 21
*
* Example 4:
*
* Input: x = 0
* Output: 0
*
* Constraints:
*
* -231 <= x <= 231 - 1
 */

package main

import "fmt"

const maxInt32 = 1<<31 - 1
const minInt32 = -(1 << 31)

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

func main() {
	fmt.Println(reverseInt(321))
}
