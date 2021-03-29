/*
* Given a string s consists of some words separated by spaces,
* return the length of the last word in the string. If the last word does not exist, return 0.
*
* A word is a maximal substring consisting of non-space characters only.
*
* Example 1:
*
* Input: s = "Hello World"
* Output: 5
*
* Example 2:
*
* Input: s = " "
* Output: 0
*
* Constraints:
*
* 1 <= s.length <= 104
* s consists of only English letters and spaces ' '.
 */

package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	var tailSpacesFlag bool

	if len(s) > 0 && s[len(s)-1] == ' ' {
		tailSpacesFlag = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && tailSpacesFlag {
			continue
		} else {
			tailSpacesFlag = false
		}

		if s[i] == ' ' {
			return length
		}

		length++
	}

	return length
}

func main() {
	fmt.Println(lengthOfLastWord("aaa aa "))
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   "))

}
