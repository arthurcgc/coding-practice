package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	output := make([]int, len(nums))
	for k := len(nums) - 1; k >= 0; k-- {
		leftNum := math.Abs(float64(nums[i]))
		rightNum := math.Abs(float64(nums[j]))
		if leftNum > rightNum {
			output[k] = int(leftNum * leftNum)
			i++
		} else {
			output[k] = int(rightNum * rightNum)
			j--
		}
	}
	return output
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(arr))
}
