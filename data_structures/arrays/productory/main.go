package main

import "fmt"

func prod(nums []int, index int) int {
	// base case
	if len(nums) == 0 {
		return 1
	}
	// second base case
	if len(nums) == 1 {
		return nums[0]
	}
	fmt.Printf("nums = %v\n", nums)
	return nums[index] * prod(nums[:index], index-1)
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(prod(nums, len(nums)-1))
	nums = []int{-1, 1, -3, 3}
	fmt.Println(prod(nums, len(nums)-1))
}
