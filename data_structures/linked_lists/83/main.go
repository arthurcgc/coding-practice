package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := p1.Next
	for p1 != nil {
		for p2 != nil && p2.Val == p1.Val {
			p2 = p2.Next
		}
		p1.Next = p2
		p1 = p1.Next
	}
	return head
}

func main() {

}
