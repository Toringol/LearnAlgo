/*
* Implement strStr().
*
* Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*
* Clarification:
*
* What should we return when needle is an empty string? This is a great question to ask during an interview.
*
* For the purpose of this problem, we will return 0 when needle is an empty string.
* This is consistent to C's strstr() and Java's indexOf().
*
* Example 1:
*
* Input: haystack = "hello", needle = "ll"
* Output: 2
*
* Example 2:
*
* Input: haystack = "aaaaa", needle = "bba"
* Output: -1
*
* Example 3:
*
* Input: haystack = "", needle = ""
* Output: 0
*
* Constraints:
*
* 0 <= haystack.length, needle.length <= 5 * 104
* haystack and needle consist of only lower-case English characters.
 */

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := 0

		for k := 0; k < len(needle); k++ {
			if haystack[i+j] != needle[k] {
				break
			}

			j++
		}

		if len(needle) == j {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issip"

	// haystack := "aabaaabaaac"
	// needle := "aabaaac"

	fmt.Println(strStr(haystack, needle))
}
