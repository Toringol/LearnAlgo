/*
* Given two binary strings a and b, return their sum as a binary string.
*
* Example 1:
*
* Input: a = "11", b = "1"
* Output: "100"
*
* Example 2:
*
* Input: a = "1010", b = "1011"
* Output: "10101"
*
* Constraints:
*
* 1 <= a.length, b.length <= 104
* a and b consist only of '0' or '1' characters.
* Each string does not contain leading zeros except for the zero itself.
 */

package main

import (
	"fmt"
)

func bitAdd(a, b, carry byte) (byte, byte) {
	// byte(a | b = 0) = 48, byte(a | b = 1) = 49, carry(0 | 1) = 0 | 1
	// byte(1) + byte(1) + carry(1) = 147 (3)
	// byte(1) + byte(1) + carry(0) = 146 (2)
	// byte(1) + byte(0) + carry(1) = 146 (2)
	// byte(1) + byte(0) + carry(0) = 145 (1)
	// byte(0) + byte(0) + carry(1) = 145 (1)
	// byte(0) + byte(0) + carry(0) = 144 (0)
	fmt.Println(a, b, carry)
	sum := a + b + carry

	fmt.Println(sum)

	if sum == 147 {
		return 49, 49
	} else if sum == 146 {
		return 48, 49
	} else if sum == 145 {
		return 49, 48
	} else {
		return 48, 48
	}
}

func addBinary(a string, b string) string {
	highStr := ""
	lowStr := ""

	if len(a) > len(b) {
		highStr = a
		lowStr = b
	} else {
		highStr = b
		lowStr = a
	}

	resultStr := make([]byte, len(highStr)+1)

	carry := byte(48)

	for i := len(lowStr) - 1; i >= 0; i-- {
		resultStr[i+len(resultStr)-len(lowStr)], carry =
			bitAdd(lowStr[i], highStr[i+len(highStr)-len(lowStr)], carry)

	}

	for i := len(highStr) - len(lowStr) - 1; i >= 0; i-- {
		resultStr[i+1], carry = bitAdd(highStr[i], 48, carry)
	}

	if carry == 49 {
		resultStr[0] = 49
	} else {
		resultStr = resultStr[1:]
	}

	return string(resultStr)
}

func main() {
	result := addBinary("1010", "1011")

	fmt.Println(result)
}
