/*
* Given an integer numRows, return the first numRows of Pascal's triangle.
*
* In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
*
* Example 1:
*
* Input: numRows = 5
* Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
*
* Example 2:
*
* Input: numRows = 1
* Output: [[1]]
*
* Constraints:
*
* 1 <= numRows <= 30
 */

package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	pascalTriangle := make([][]int, numRows)

	for i := 0; i < len(pascalTriangle); i++ {
		pascalTriangle[i] = make([]int, i+1)
		pascalTriangle[i][0] = 1
		pascalTriangle[i][len(pascalTriangle[i])-1] = 1
	}

	for i := 2; i < len(pascalTriangle); i++ {
		for j := 1; j < len(pascalTriangle[i])-1; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
	}

	return pascalTriangle
}

func main() {
	fmt.Println(generate(7))
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
}
