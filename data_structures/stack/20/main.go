package main

import "github.com/arthurcgc/stack"

func isOpenBracket(c string) bool {
	if c == "{" || c == "[" || c == "(" {
		return true
	}
	return false
}

func isValid(s string) bool {
	myStack := stack.StringStack{}
	for _, c := range s {
		if isOpenBracket(string(c)) {
			myStack.Push(string(c))
		} else {
			// is closing bracket, c = } or ] or )
			matchingBracket := myStack.Pop()
			if c == '}' && matchingBracket != "{" {
				return false
			}
			if c == ']' && matchingBracket != "[" {
				return false
			}
			if c == ')' && matchingBracket != "(" {
				return false
			}
		}
	}
	return myStack.Len() == 0
}

func main() {

}
