package main

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	println(result[0], result[1])
}
