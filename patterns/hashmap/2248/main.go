package main

import (
	"fmt"
	"sort"
)

// Given a 2D integer array nums where nums[i] is a non-empty array of distinct positive integers, return the list of integers that are present in each array of nums sorted in ascending order.

func intersection(nums [][]int) []int {
	// use a hashmap to count each integer, than compare it to the length of nums, since every array in nums has distinct integers
	hashmap := map[int]int{}
	for _, arr := range nums {
		for _, num := range arr {
			hashmap[num]++
		}
	}
	n := len(nums)
	result := []int{}
	for k, v := range hashmap {
		if v == n {
			result = append(result, k)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	nums := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}
	fmt.Println(intersection(nums))
}
