package main

import "fmt"

func longestOnes(nums []int, k int) int {
	zeros := 0
	arr := []int{}
	maxLen := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		arr = append(arr, nums[right])
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
			arr = nums[left : right+1]
		}
		if len(arr) > maxLen {
			maxLen = len(arr)
		}
	}
	return maxLen
}

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println(longestOnes(arr, k))
	arr = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k = 3
	fmt.Println(longestOnes(arr, k))
}
