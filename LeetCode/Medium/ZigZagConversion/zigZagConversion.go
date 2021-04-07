/*
* The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
* (you may want to display this pattern in a fixed font for better legibility)
*
* P   A   H   N
* A P L S I I G
* Y   I   R
* And then read line by line: "PAHNAPLSIIGYIR"
*
* Write the code that will take a string and make this conversion given a number of rows:
*
* string convert(string s, int numRows);
*
* Example 1:
*
* Input: s = "PAYPALISHIRING", numRows = 3
* Output: "PAHNAPLSIIGYIR"
*
* Example 2:
*
* Input: s = "PAYPALISHIRING", numRows = 4
* Output: "PINALSIGYAHRPI"
* Explanation:
* P     I    N
* A   L S  I G
* Y A   H R
* P     I
*
* Example 3:
*
* Input: s = "A", numRows = 1
* Output: "A"
*
* Constraints:
*
* 1 <= s.length <= 1000
* s consists of English letters (lower-case and upper-case), ',' and '.'.
* 1 <= numRows <= 1000
 */

package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	zigZagMatrix := make([]strings.Builder, numRows)

	curRow := 0
	isDown := true

	for i := 0; i < len(s); i++ {
		zigZagMatrix[curRow].WriteByte(s[i])

		if isDown {
			curRow++
		} else {
			curRow--
		}

		if curRow == 0 {
			isDown = true
		}

		if curRow == numRows-1 {
			isDown = false
		}
	}

	result := new(strings.Builder)

	for _, row := range zigZagMatrix {
		result.WriteString(row.String())
	}

	return result.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
