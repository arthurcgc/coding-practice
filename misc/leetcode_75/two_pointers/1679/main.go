package main

// searchFor returns the targetIndex
func targetFound(nums []int, i, j, target int, alreadyFound []bool) bool {
	for ; i <= j; i++ {
		v1 := nums[i]
		if v1 == target && !alreadyFound[i] {
			alreadyFound[i] = true
			return true
		}
		v2 := nums[j]
		if v2 == target && !alreadyFound[j] {
			alreadyFound[j] = true
			return true
		}
		j--
	}
	return false
}

func maxOperations(nums []int, k int) int {
	operations := 0
	n := len(nums)
	j := n - 1
	alreadyRemoved := make([]bool, n)
	for i := 0; i < n && j > i; {
		left := nums[i]
		right := nums[j]
		if i%2 == 0 {
			if !alreadyRemoved[i] && targetFound(nums, i+1, j, k-left, alreadyRemoved) {
				alreadyRemoved[i] = true
				operations++
			}
			i++
		} else {
			if !alreadyRemoved[j] && targetFound(nums, i, j-1, k-right, alreadyRemoved) {
				alreadyRemoved[j] = true
				operations++
			}
			j--
		}
	}
	return operations
}
