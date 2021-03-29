/*
* Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
*
* The number of elements initialized in nums1 and nums2 are m and n respectively.
* You may assume that nums1 has a size equal to m + n such that
* it has enough space to hold additional elements from nums2.
*
* Example 1:
*
* Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
* Output: [1,2,2,3,5,6]
*
* Example 2:
*
* Input: nums1 = [1], m = 1, nums2 = [], n = 0
* Output: [1]
*
* Constraints:
*
* nums1.length == m + n
* nums2.length == n
* 0 <= m, n <= 200
* 1 <= m + n <= 200
* -109 <= nums1[i], nums2[i] <= 109
 */

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	iterNums1 := 0
	iterNums2 := 0

	nums1Cap := m + n

	for iterNums1 != nums1Cap {
		fmt.Println(iterNums1, iterNums2)
		if nums2[iterNums2] < nums1[iterNums1] || iterNums1 == m {
			for i := iterNums1; i < 2*m-iterNums1-1; i++ {
				nums1[i+1] = nums1[i]
			}

			nums1[iterNums1] = nums2[iterNums2]
			iterNums2++
			m++
		}

		iterNums1++
	}
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m, n := 2, 1

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}
