/*
* Given a string s, determine if it is a palindrome,
* considering only alphanumeric characters and ignoring cases.
*
* Example 1:
*
* Input: s = "A man, a plan, a canal: Panama"
* Output: true
* Explanation: "amanaplanacanalpanama" is a palindrome.
*
* Example 2:
*
* Input: s = "race a car"
* Output: false
* Explanation: "raceacar" is not a palindrome.
*
* Constraints:
*
* 1 <= s.length <= 2 * 105
* s consists only of printable ASCII characters.
 */

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	leftBound, rightBound := 0, len(s)-1

	for leftBound < rightBound {
		if !isSymbol(s[leftBound]) {
			leftBound++
			continue
		}

		if !isSymbol(s[rightBound]) {
			rightBound--
			continue
		}

		if !strings.EqualFold(string(s[leftBound]), string(s[rightBound])) {
			return false
		}

		leftBound++
		rightBound--
	}

	return true
}

func isSymbol(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
