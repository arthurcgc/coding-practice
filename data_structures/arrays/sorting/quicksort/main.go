package main

import "fmt"

func swap(i, j int, arr []int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func partition(low, high int, arr []int) int {
	pivot := arr[high] // We're choosing the last element as pivot

	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			swap(i, j, arr)
		}
	}
	swap(i+1, high, arr)
	return i + 1
}

func quicksort(low, high int, arr []int) {
	if low >= high || low < 0 {
		return
	}

	pivot := partition(low, high, arr)

	quicksort(low, pivot-1, arr)  // left side of pivot
	quicksort(pivot+1, high, arr) // right side of pivot
}

func main() {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	quicksort(0, len(arr)-1, arr)
	fmt.Printf("%v\n", arr)
}
