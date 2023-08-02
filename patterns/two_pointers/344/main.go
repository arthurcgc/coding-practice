package main

import "fmt"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Printf("%s\n", string(s))
}
