package main

func binarySearch(target, left, right int, arr []int) int {
	if left == right || left > right {
		if target <= arr[left] {
			return left
		}
		if target > arr[left] {
			return left + 1
		}
	}
	mid := (left + right) / 2
	if target > arr[mid] {
		left := mid + 1
		return binarySearch(target, left, right, arr)
	} else if target < arr[mid] {
		right := mid - 1
		return binarySearch(target, left, right, arr)
	}
	return mid
}

func search(target int, arr []int) int {
	return binarySearch(target, 0, len(arr)-1, arr)
}

func searchInsert(nums []int, target int) int {
	return search(target, nums)
}
