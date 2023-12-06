package main

func swap(s string, i, j int) string {
	newS1 := s[:i]
	newS1 += string(s[j])
	newS2 := s[i+1 : j]
	newS2 += string(s[i])
	newS2 += s[j+1:]
	return newS1 + newS2
}

func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': false,
		'A': false,
		'e': false,
		'E': false,
		'i': false,
		'I': false,
		'o': false,
		'O': false,
		'u': false,
		'U': false,
	}
	i := 0
	j := len(s) - 1
	for ; i < len(s); i++ {
		v1 := rune(s[i])
		// if not vowel skip
		if _, ok := vowels[v1]; !ok {
			continue
		}
		for ; j > i; j-- {
			v2 := rune(s[j])
			if _, ok := vowels[v2]; !ok {
				continue
			}
			s = swap(s, i, j)
			j--
			break
		}
		// impossible to find another j pair for i, no need to continue
		if j == i {
			break
		}
	}
	return s
}
