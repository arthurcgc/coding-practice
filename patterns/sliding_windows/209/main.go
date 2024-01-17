package main

func minSubArrayLen(target int, nums []int) int {
	i := 0
	j := 0
	nullValue := 999999999999
	leastLen := nullValue
	currSum := 0
	for i < j+1 && j < len(nums) {
		subArr := nums[i : j+1]
		if i == 0 && j == 0 {
			currSum = nums[0]
		}
		if currSum < target {
			j++
			if j < len(nums) {
				currSum += nums[j]
			}
		} else {
			if len(subArr) < leastLen {
				leastLen = len(subArr)
			}
			currSum -= nums[i]
			i++
		}
	}
	if leastLen == nullValue {
		return 0
	}
	return leastLen
}
