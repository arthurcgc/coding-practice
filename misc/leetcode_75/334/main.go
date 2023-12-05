package main

func increasingTriplet(nums []int) bool {
	n := len(nums)
	n1 := 9999999999999
	n2 := n1
	// we have to call n1 and n2, not low and mid, because we will work with them interchangeably
	// if we update n1, n1 becomes the new mid and n2 becomes the new low
	// if we update n2, n2 stays mid
	for i := 0; i < n; i++ {
		v := nums[i]
		if v < n1 {
			n1 = v
		} else if v < n2 && v != n1 {
			n2 = v
		} else if v > n1 && v > n2 {
			return true
		}
	}
	return false
}
