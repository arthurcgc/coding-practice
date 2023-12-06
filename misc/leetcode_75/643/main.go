package main

func avg(nums []int) float64 {
	sum := float64(0)
	n := float64(len(nums))
	for _, v := range nums {
		sum += float64(v)
	}
	return sum / n
}

func findMaxAverage(nums []int, k int) float64 {
	maxAvg := float64(-9999999999)
	for i := 0; i < len(nums); i++ {
		j := i + k
		if j > len(nums) {
			return maxAvg
		}
		currAvg := avg(nums[i:j])
		if currAvg > maxAvg {
			maxAvg = currAvg
		}
	}
	return maxAvg
}
