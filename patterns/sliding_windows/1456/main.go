package main

func isVowel(letter rune) bool {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	_, ok := vowels[letter]
	return ok
}

// algorithm: we start with an array (of letters) of size 0
// with each iteration we:
// 1: check the size of the array, if it's < k
// 1.1: if it is, we check if the new letter is a vowel
// 1.2: if it is, we increment the currentVowel count
// 1.2.1: since we increase currentVowel count, we check if we have to update the maxVowel count
// 2: if the size of the array is > k we increment the left side (i)
// 2.1: we have to check before going forward if we lost a vowel and decrement currVowel count accordingly

func maxVowels(s string, k int) int {
	i := 0
	j := 0
	max := 0
	currVowels := 0
	for i <= j && j < len(s) {
		subArr := s[i : j+1]
		if len(subArr) > k {
			if isVowel(rune(s[i])) {
				currVowels--
			}
			i++
		} else {
			if isVowel(rune(s[j])) {
				currVowels++
				if currVowels > max {
					max = currVowels
					if max == k {
						return max
					}
				}
			}
			j++
		}
	}
	return max
}
