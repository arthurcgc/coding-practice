package main

func isPrefixCandidate(cand string, strs []string) bool {
	n := len(cand)
	for _, s := range strs {
		if n > len(s) {
			return false
		}
		if cand != s[:n] {
			return false
		}
	}
	return true
}

func _longestCommonPrefix(strs []string) string {
	longest := ""
	for _, s := range strs {
		// l = len(s) because in Go s[:l] is non inclusive
		for l := len(s); l >= 1; l-- {
			candidate := s[:l]
			if isPrefixCandidate(candidate, strs) {
				if len(candidate) > len(longest) {
					longest = candidate
				}
			}
		}
	}
	return longest
}
