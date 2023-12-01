package main

func insertWord(start int, word string, currentWord string) string {
	for i := start; i < len(word); i++ {
		currentWord += string(word[i])
	}
	return currentWord
}

func mergeAlternately(word1 string, word2 string) string {
	n := len(word1)
	m := len(word2)
	newWord := ""
	counter1 := 0
	counter2 := 0

	for i := 0; i < n+m; i++ {
		if i%2 == 0 {
			if counter1 >= len(word1) {
				return insertWord(counter2, word2, newWord)
			}
			newWord += string(word1[counter1])
			counter1++
		} else {
			if counter2 >= len(word2) {
				return insertWord(counter1, word1, newWord)
			}
			newWord += string(word2[counter2])
			counter2++
		}
	}

	return newWord
}

func main() {
	println(mergeAlternately("abc", "pqr"))
}
