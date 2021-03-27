/*
* Write a function to find the longest common prefix string amongst an array of strings.
*
* If there is no common prefix, return an empty string "".
*
*
*
* Example 1:
*
* Input: strs = ["flower","flow","flight"]
* Output: "fl"
*
* Example 2:
*
* Input: strs = ["dog","racecar","car"]
* Output: ""
* Explanation: There is no common prefix among the input strings.
*
* Constraints:
*
* 0 <= strs.length <= 200
* 0 <= strs[i].length <= 200
* strs[i] consists of only lower-case English letters.
 */

package main

import "fmt"

func indexMinStr(strs []string) int {
	minLen := len(strs[0])
	indexMinStr := 0

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
			indexMinStr = i
		}
	}

	return indexMinStr
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	indexEndLongestPrefix := 0
	indexMinStr := indexMinStr(strs)

	for i := range strs[indexMinStr] {
		for j := 0; j < len(strs); j++ {
			if strs[indexMinStr][i] != strs[j][i] {
				return strs[indexMinStr][:indexEndLongestPrefix]
			}
		}

		indexEndLongestPrefix++
	}

	return strs[indexMinStr]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
