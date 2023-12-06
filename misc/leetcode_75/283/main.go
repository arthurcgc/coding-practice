package main

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func sendToBack(nums []int, index int) {
	for i := index; i+1 < len(nums); i++ {
		swap(nums, i, i+1)
	}
}

func moveZeroes(nums []int) {
	realCounter := 0
	for i := 0; i < len(nums) && realCounter < len(nums); i++ {
		v := nums[i]
		if v == 0 {
			sendToBack(nums, i)
			i--
		}
		realCounter++
	}
}
