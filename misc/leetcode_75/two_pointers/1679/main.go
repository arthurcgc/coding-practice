package main

import "slices"

func maxOperations(nums []int, k int) int {
	n := len(nums)
	j := n - 1
	operations := 0
	slices.Sort(nums)
	for i := 0; i < n && j > i; {
		left := nums[i]
		right := nums[j]
		sum := left + right
		if sum == k {
			i++
			j--
			operations++
		} else if sum > k {
			// go left
			j--
		} else {
			// go right
			i++
		}
	}
	return operations
}
