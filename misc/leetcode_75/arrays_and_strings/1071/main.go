package main

func defineLargeAndSmall(s1, s2 string) (string, string) {
	if s2 > s1 {
		return s2, s1
	}
	return s1, s2
}

func divides(target, candidate string) bool {
	step := len(candidate)
	steps := 0
	for i := 0; i+step <= len(target); i += step {
		steps = i + step
		if candidate != target[i:i+step] {
			return false
		}
	}
	return steps == len(target)
}

func gcdOfStrings(str1 string, str2 string) string {
	large, small := defineLargeAndSmall(str1, str2)
	longest := ""
	for i := len(small); i >= 1; i-- {
		current := small[:i]
		if divides(small, current) && divides(large, current) {
			if len(longest) < len(current) {
				longest = current
			}
		}
	}
	return longest
}
