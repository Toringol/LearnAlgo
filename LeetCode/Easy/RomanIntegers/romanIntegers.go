/*
* Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
*
* Symbol       Value
* I             1
* V             5
* X             10
* L             50
* C             100
* D             500
* M             1000
*
* For example, 2 is written as II in Roman numeral, just two one's added together.
* 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
*
* Roman numerals are usually written largest to smallest from left to right.
* However, the numeral for four is not IIII. Instead, the number four is written as IV.
* Because the one is before the five we subtract it making four.
* The same principle applies to the number nine, which is written as IX.
* There are six instances where subtraction is used:
*
* I can be placed before V (5) and X (10) to make 4 and 9.
* X can be placed before L (50) and C (100) to make 40 and 90.
* C can be placed before D (500) and M (1000) to make 400 and 900.
* Given a roman numeral, convert it to an integer.
*
* Example 1:
*
* Input: s = "III"
* Output: 3
*
* Example 2:
*
* Input: s = "IV"
* Output: 4
*
* Example 3:
*
* Input: s = "IX"
* Output: 9
*
* Example 4:
*
* Input: s = "LVIII"
* Output: 58
* Explanation: L = 50, V = 5, III = 3.
*
* Example 5:
*
* Input: s = "MCMXCIV"
* Output: 1994
* Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*
* Constraints:
*
* 1 <= s.length <= 15
* s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
* It is guaranteed that s is a valid roman numeral in the range [1, 3999].
 */

package main

import "fmt"

func romanSymToInt(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "IV":
		return 4
	case "X":
		return 10
	case "IX":
		return 9
	case "L":
		return 50
	case "XL":
		return 40
	case "C":
		return 100
	case "XC":
		return 90
	case "D":
		return 500
	case "CD":
		return 400
	case "M":
		return 1000
	case "CM":
		return 900
	}

	return 0
}

func checkSubInstance(s byte, nextSym byte) bool {
	if s == 'I' && (nextSym == 'V' || nextSym == 'X') {
		return true
	} else if s == 'X' && (nextSym == 'L' || nextSym == 'C') {
		return true
	} else if s == 'C' && (nextSym == 'D' || nextSym == 'M') {
		return true
	}

	return false
}

func romanToInt(s string) int {
	res := 0
	romanNumber := ""

	for i := 0; i < len(s); i++ {
		romanNumber = string(s[i])

		if (s[i] == 'I' || s[i] == 'X' || s[i] == 'C') &&
			i != len(s)-1 && checkSubInstance(s[i], s[i+1]) {

			romanNumber += string(s[i+1])
			i++
		}

		res += romanSymToInt(romanNumber)
	}

	return res
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
