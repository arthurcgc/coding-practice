package main

import "fmt"

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.
func runningSum(nums []int) []int {
	prefixSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefixSum = append(prefixSum, nums[i]+prefixSum[i-1])
	}
	return prefixSum
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
