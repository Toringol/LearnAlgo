/*
* Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
*
* For example:
*
* A -> 1
* B -> 2
* C -> 3
* ...
* Z -> 26
* AA -> 27
* AB -> 28
* ...
*
* Example 1:
*
* Input: columnNumber = 1
* Output: "A"
*
* Example 2:
*
* Input: columnNumber = 28
* Output: "AB"
*
* Example 3:
*
* Input: columnNumber = 701
* Output: "ZY"
*
* Example 4:
*
* Input: columnNumber = 2147483647
* Output: "FXSHRXW"
*
* Constraints:
*
* 1 <= columnNumber <= 231 - 1
 */

package main

import "fmt"

func convertToTitle(columnNumber int) string {
	exelTitle := ""

	for columnNumber > 0 {
		letterOffset := columnNumber % 26
		columnNumber = columnNumber / 26

		if letterOffset == 0 {
			letterOffset = 26
			columnNumber -= 1
		}

		exelTitle = string('A'+byte(letterOffset-1)) + exelTitle
	}

	return exelTitle
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(2147483647))
}
