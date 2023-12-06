package main

import "strconv"

func hammingWeight(num uint32) int {
	numString := strconv.FormatUint(uint64(num), 2)
	ones := 0
	for _, c := range numString {
		if c == '1' {
			ones++
		}
	}
	return ones
}

func main() {
	println(hammingWeight(11))
}
