package main

// Given an integer array arr, count how many elements x there are, such that x + 1 is also in arr. If there are duplicates in arr, count them separately.
func countElements(arr []int) int {
	xElements := 0
	elements := map[int]struct{}{}
	for _, num := range arr {
		elements[num] = struct{}{}
	}
	for _, num := range arr {
		if _, ok := elements[num+1]; ok {
			xElements++
		}
	}
	return xElements
}

func main() {

}
