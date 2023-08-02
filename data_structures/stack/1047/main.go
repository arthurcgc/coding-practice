package main

import "github.com/arthurcgc/stack"

func removeDuplicates(s string) string {
	myStack := stack.StringStack{}
	for _, c := range s {
		if myStack.Peek() != string(c) {
			myStack.Push(string(c))
		} else {
			myStack.Pop()
		}
	}
	res := ""
	for myStack.Peek() != "\u0000" {
		res = myStack.Pop() + res
	}
	return res
}

func main() {

}
