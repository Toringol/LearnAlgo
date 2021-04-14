/*
* Given an integer array nums, return all the triplets
* [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
*
* Notice that the solution set must not contain duplicate triplets.
*
* Example 1:
*
* Input: nums = [-1,0,1,2,-1,-4]
* Output: [[-1,-1,2],[-1,0,1]]
*
* Example 2:
*
* Input: nums = []
* Output: []
*
* Example 3:
*
* Input: nums = [0]
* Output: []
*
* Constraints:
*
* 0 <= nums.length <= 3000
* -105 <= nums[i] <= 105
 */

package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		leftBound, rightBound := i+1, len(nums)-1

		for leftBound < rightBound {
			sum := nums[leftBound] + nums[i] + nums[rightBound]

			switch {
			case sum > 0:
				rightBound--
			case sum < 0:
				leftBound++
			default:
				result = append(result, []int{nums[i], nums[leftBound], nums[rightBound]})
				leftBound, rightBound = nextBounds(nums, leftBound, rightBound)
			}
		}
	}

	return result
}

func nextBounds(nums []int, leftBound, rightBound int) (int, int) {
	for leftBound < rightBound {
		switch {
		case nums[leftBound] == nums[leftBound+1]:
			leftBound++
		case nums[rightBound] == nums[rightBound-1]:
			rightBound--
		default:
			leftBound++
			rightBound--
			return leftBound, rightBound
		}
	}

	return leftBound, rightBound
}

func main() {

}
