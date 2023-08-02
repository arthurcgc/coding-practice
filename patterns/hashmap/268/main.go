package main

func missingNumber(nums []int) int {
	countNumbers := map[int]struct{}{}
	for _, num := range nums {
		countNumbers[num] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := countNumbers[i]; !ok {
			return i
		}
	}
	return -1
}

func main() {

}
