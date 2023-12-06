package main

func prod(nums []int) int {
    if len(nums) == 0 {
        return 1
    }
	if len(nums) == 1 {
		return nums[0]
	}
	current := len(nums) - 1
	return nums[current] * prod(nums[:current])
}

func prodIt(nums []int) int {
    p := 1
    for _, v := range nums {
        p *= v
    }
    return p
}

func productExceptSelf(nums []int) []int {
    ps := []int{}
    leftProduct := []int{}
    for i := range nums {
        var left, right int
        if i == 0 {
            left = 1
        } else {
            n := len(leftProduct)
            left = leftProduct[n-1] * nums[i-1]
        }
        leftProduct = append(leftProduct, left)

        rightArr := nums[i+1:]
        right = prodIt(rightArr)
        
		ps = append(ps, left*right)
    }
    return ps
}