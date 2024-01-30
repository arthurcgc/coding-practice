package main

// This function finds the length of the longest subarray in the given slice of integers (nums)
// such that the subarray can have at most one zero. The function uses a sliding window approach
// to iterate through the array and keep track of the number of zeroes encountered.
// The variables i and j represent the start and end indices of the current subarray.
// The variable zeroes keeps track of the number of zeroes encountered so far.
// The variable lastZeroIndex stores the index of the last zero encountered.
// The variable max stores the length of the longest subarray found so far.
// The variable curr stores the length of the current subarray (excluding zeroes).
// The algorithm iterates through the array using the indices i and j.
// If a zero is encountered, the algorithm updates the indices and zeroes count accordingly.
// If a non-zero element is encountered, the algorithm calculates the length of the current subarray
// (excluding zeroes) and updates the max length if necessary.
// Finally, the algorithm returns the length of the longest subarray found.
func longestSubarray(nums []int) int {
	i := 0
	j := 0
	zeroes := 0
	lastZeroIndex := -1
	max := 0
	curr := 0
	for i <= j && j < len(nums) {
		if nums[j] == 0 {
			zeroes++
			if zeroes > 1 {
				i = lastZeroIndex + 1
				zeroes--
				lastZeroIndex = j
			} else {
				lastZeroIndex = j
			}
		} else {
			curr = len(nums[i:j+1]) - zeroes
			if curr > max {
				max = curr
			}
		}
		j++
	}

	if lastZeroIndex == -1 {
		return max - 1
	}

	return max
}
