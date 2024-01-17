package main

func minSubArrayLen(target int, nums []int) int {
    i := 0
    j := 0
    leastLen := 0
    currSum := 0
    for ; i < j + 1 && j < len(nums); {
        if i == 0 && j == 0 {
            currSum = nums[0]
        }
        subArr := nums[i:j+1]
        if currSum < target {
            j++
            if j < len(nums) {
                currSum += nums[j]
            }
        } else {
            n := len(subArr)
            if n < leastLen || leastLen == 0 {
                leastLen = n
            }
            currSum -= nums[i]
            i++
        }
    }
    return leastLen
}
