package main

func repeatedCharacter(s string) byte {
	hashmap := map[rune]struct{}{}
	for _, c := range s {
		if _, ok := hashmap[c]; ok {
			return byte(c)
		}
		hashmap[c] = struct{}{}
	}
	return byte('\u0000')
}

func main() {

}
