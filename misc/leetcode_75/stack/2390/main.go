package main

import "strings"

func removeStars(s string) string {
	stack := []string{}
	for _, r := range s {
		if r == '*' {
			// pop
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(r))
		}
	}

	return strings.Join(stack, "")
}
