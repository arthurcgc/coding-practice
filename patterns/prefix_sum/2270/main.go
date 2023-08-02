package main

import "fmt"

// You are given a 0-indexed integer array nums of length n.
//
// nums contains a valid split at index i if the following are true:
//
//     The sum of the first i + 1 elements is greater than or equal to the sum of the last n - i - 1 elements.
//     There is at least one element to the right of i. That is, 0 <= i < n - 1.
//
// Return the number of valid splits in nums.

func waysToSplitArray(nums []int) int {
	prefixSum := []int{nums[0]}
	n := len(nums)
	for i := 1; i < n; i++ {
		prefixSum = append(prefixSum, nums[i]+prefixSum[i-1])
	}
	// calculate the prefix sum from 0..i and compare it with the prefix sum of the rest of the array (i+1..n)
	count := 0
	for i := 0; i < n-1; i++ {
		leftSum := prefixSum[i]
		rightSum := (prefixSum[n-1] - prefixSum[i])
		if leftSum >= rightSum {
			count++
		}
	}
	return count
}

func main() {
	nums := []int{10, 4, -8, 7}
	fmt.Println(waysToSplitArray(nums))
}
