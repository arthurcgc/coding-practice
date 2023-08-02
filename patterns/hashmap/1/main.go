package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	// key = num in nums
	// value = index of num in nums
	hashmap := map[int]int{}
	for i, num := range nums {
		targetDiff := target - num
		if index, ok := hashmap[targetDiff]; ok {
			return []int{index, i}
		} else {
			hashmap[num] = i
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
