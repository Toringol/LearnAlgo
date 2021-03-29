/*
* You are climbing a staircase. It takes n steps to reach the top.
*
* Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*
* Example 1:
*
* Input: n = 2
* Output: 2
* Explanation: There are two ways to climb to the top.
* 1. 1 step + 1 step
* 2. 2 steps
*
* Example 2:
*
* Input: n = 3
* Output: 3
* Explanation: There are three ways to climb to the top.
* 1. 1 step + 1 step + 1 step
* 2. 1 step + 2 steps
* 3. 2 steps + 1 step
*
* Constraints:
* 1 <= n <= 45
 */

package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	climbArr := make([]int, n+1)
	climbArr[1] = 1
	climbArr[2] = 2

	for i := 3; i <= n; i++ {
		climbArr[i] = climbArr[i-1] + climbArr[i-2]
	}

	return climbArr[n]
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
}
