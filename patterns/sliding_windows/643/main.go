package main

func findMaxAverage(nums []int, k int) float64 {
	left := 0
	sum := 0
	maxArr := []int{}
	maxSum := -999999
	maxAvg := float64(0)
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		maxArr = append(maxArr, nums[right])
		for len(maxArr) > k {
			sum -= nums[left]
			left++
			maxArr = nums[left : right+1]
		}
		if len(maxArr) == k {
			if maxSum < sum {
				maxSum = sum
				maxAvg = float64(float64(maxSum) / float64(k))
			}
		}
	}
	return maxAvg
}
