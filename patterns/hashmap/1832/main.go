package main

func checkIfPangram(sentence string) bool {
	n := 26
	uniqueLetters := map[rune]struct{}{}
	for _, c := range sentence {
		if _, ok := uniqueLetters[c]; ok {
			continue
		}
		uniqueLetters[c] = struct{}{}
	}
	return len(uniqueLetters) == n
}

func main() {

}
