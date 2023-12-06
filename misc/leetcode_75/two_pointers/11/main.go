package main

func area(h1, h2, i1, i2 int) int {
	x := i2 - i1
	height := h1
	// the minimum height determines the actual height
	if h2 < h1 {
		height = h2
	}
	return height * x
}

func maxAreaBruteForce(height []int) int {
	n := len(height)
	j := n - 1
	maxArea := 0
	for i := 0; i < n && j > i; i++ {
		for j := n - 1; j > i; j-- {
			currArea := area(height[i], height[j], i, j)
			if currArea > maxArea {
				maxArea = currArea
			}
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	n := len(height)
	j := n - 1
	maxArea := 0
	for i := 0; i < n; {
		h1 := height[i]
		h2 := height[j]
		currArea := area(h1, h2, i, j)
		if currArea > maxArea {
			maxArea = currArea
		}
		// if h1 > h2 -> decrement j (right)
		// if h1 <= h2 -> increment i (left)
		if h1 <= h2 {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
