package main

import (
	"fmt"
	"math"
)

func findMaxSubArray(nums []int, lowBound, highBound int) (int, int, int) {
	if lowBound == highBound {
		return lowBound, highBound, nums[lowBound]
	}

	midBound := lowBound + (highBound-lowBound)/2

	leftLow, leftHigh, leftSum := findMaxSubArray(nums, lowBound, midBound)
	rightLow, rightHigh, rightSum := findMaxSubArray(nums, midBound+1, highBound)
	crossLow, crossHigh, crossSum := findMaxCrossSubArray(nums, lowBound, midBound, highBound)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func findMaxCrossSubArray(nums []int, lowBound, midBound, highBound int) (int, int, int) {
	leftSum, rightSum := math.MinInt64, math.MinInt64
	actualSum := 0
	maxLeft, maxRight := midBound, midBound+1

	for i := midBound; i >= lowBound; i-- {
		actualSum += nums[i]
		if actualSum > leftSum {
			leftSum = actualSum
			maxLeft = i
		}
	}

	actualSum = 0

	for i := midBound + 1; i <= highBound; i++ {
		actualSum += nums[i]
		if actualSum > rightSum {
			rightSum = actualSum
			maxRight = i
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}

func main() {
	nums := []int{13, -3, -25, 20, -3, -16, -23,
		18, 20, -7, 12, -5, -22, 15, -4, 7}

	fmt.Println(findMaxSubArray(nums, 0, len(nums)-1))
}
