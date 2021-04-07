/*
* Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
*
* In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
*
* Example 1:
*
* Input: rowIndex = 3
* Output: [1,3,3,1]
*
* Example 2:
*
* Input: rowIndex = 0
* Output: [1]
*
* Example 3:
*
* Input: rowIndex = 1
* Output: [1,1]
*
* Constraints:
*
* 0 <= rowIndex <= 33
*
* Follow up: Could you optimize your algorithm to use only O(rowIndex) extra space?
 */

package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	pascalTriangleRow := make([]int, 2)
	pascalTriangleRow[0] = 1
	pascalTriangleRow[1] = 1

	for i := 1; i < rowIndex; i++ {
		pascalIndexRow := make([]int, 0)

		pascalIndexRow = append(pascalIndexRow, 1)

		for j := 1; j < len(pascalTriangleRow); j++ {
			pascalIndexRow = append(pascalIndexRow, pascalTriangleRow[j]+pascalTriangleRow[j-1])
		}

		pascalIndexRow = append(pascalIndexRow, 1)

		pascalTriangleRow = pascalIndexRow
	}

	return pascalTriangleRow
}

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(6))
}
