package main

import (
	"fmt"
	"math"
)

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	l1 := head
	l2 := head.Next
	// find out l2 length we know l1Length = 1
	l2Length := 0
	for l2 != nil {
		l2Length++
		l2 = l2.Next
	}
	length = 1 + l2Length
	var middle int
	if length%2 > 0 {
		middle = int(math.Floor(float64(length) / float64(2)))
	} else {
		middle = length / 2
	}
	for steps := 0; steps < middle; steps++ {
		l1 = l1.Next
	}
	return l1
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Printf("%d\n", middleNode(l1).Val)
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Printf("%d\n", middleNode(l2).Val)
}
