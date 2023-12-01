package main

import "strings"

func isWord(s string) bool {
	return len(s) > 0
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	newS := ""
	for i := len(ss) - 1; i >= 0; i-- {
		current := ss[i]
		if isWord(current) {
			if len(newS) == 0 {
				newS += current
			} else {
				newS += " " + current
			}
		}
	}
	return newS
}
