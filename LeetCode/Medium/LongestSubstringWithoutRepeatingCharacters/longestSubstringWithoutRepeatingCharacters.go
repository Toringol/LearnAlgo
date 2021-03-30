/*
* Given a string s, find the length of the longest substring without repeating characters.
*
* Example 1:
*
* Input: s = "abcabcbb"
* Output: 3
* Explanation: The answer is "abc", with the length of 3.
*
* Example 2:
*
* Input: s = "bbbbb"
* Output: 1
* Explanation: The answer is "b", with the length of 1.
*
* Example 3:
*
* Input: s = "pwwkew"
* Output: 3
* Explanation: The answer is "wke", with the length of 3.
* Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*
* Example 4:
*
* Input: s = ""
* Output: 0
*
* Constraints:
*
* 0 <= s.length <= 5 * 104
* s consists of English letters, digits, symbols and spaces.
 */

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	longestLength := 0

	for i := 0; i < len(s); i++ {
		characters := make(map[byte]struct{}, len(s)-i)
		currentLength := 0

		for j := i; j < len(s); j++ {
			if _, isFlag := characters[s[j]]; isFlag {
				break
			}

			currentLength++
			characters[s[j]] = struct{}{}
		}

		if longestLength < currentLength {
			longestLength = currentLength
		}
	}

	return longestLength
}

func main() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))

	s = "bbbbb"

	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s))

	s = ""

	fmt.Println(lengthOfLongestSubstring(s))

	s = " "

	fmt.Println(lengthOfLongestSubstring(s))

	s = "dvdf"

	fmt.Println(lengthOfLongestSubstring(s))
}
