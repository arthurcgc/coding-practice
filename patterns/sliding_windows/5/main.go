package main

func isPalindrome(s string) bool {
	n := len(s)
	j := n - 1
	for i := 0; i < n; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	longest := ""
	for i := 0; i < n; i++ {
		for j := n; j >= i; j-- {
			current := s[i:j]
			if len(current) > len(longest) {
				if isPalindrome(current) {
					longest = current
				}
			}
		}
	}
	return longest
}

func main() {
	println(longestPalindrome("babad"))
	println(longestPalindrome("cbbd"))
}
