package main

import "fmt"

func compressRune(char byte, index int, chars []byte) int {
	if index >= len(chars) {
		return 0
	}
	if chars[index] != char {
		return 0
	}
	return compressRune(char, index+1, chars) + 1
}

func main() {
	fmt.Println(compressRune('c', 4, []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}
