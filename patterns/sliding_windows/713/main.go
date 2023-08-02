package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	left := 0
	ans := 0
	prod := 1
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod = prod / nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
