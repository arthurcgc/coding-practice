package main

import "fmt"

func validString(s string, alphabet map[rune]bool) bool {
	for _, c := range s {
		if used := alphabet[c]; used {
			return false
		}
		alphabet[c] = true
	}
	return true
}

func newAlphabet() map[rune]bool {
	alphabet := make(map[rune]bool)
	for _, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		alphabet[c] = false
	}
	return alphabet
}

func checkLongest(current, longest string) bool {
	return len(current) > len(longest)
}

func lengthOfLongestSubstring(s string) int {
	longest := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; {
			current := s[i:j]
			alphabet := newAlphabet()
			fmt.Printf("current = %s\n", current)
			if validString(current, alphabet) {
				// keep going
				j++
				if checkLongest(current, longest) {
					longest = current
				}
			} else {
				i++
			}
		}
	}
	return len(longest)
}

func main() {
	// println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring("abcabcbb"))
	// println(lengthOfLongestSubstring("abcabcbbz"))
}
