package main

import (
	"fmt"
	ReverseList "github.com/qxyang2015/Go/剑指offer/24.反转链表"
)

func main() {
	l := &ReverseList.ListNode{
		Val: 1,
		Next: &ReverseList.ListNode{
			Val: 2,
			Next: &ReverseList.ListNode{
				Val: 3,
				Next: &ReverseList.ListNode{
					Val: 4,
					Next: &ReverseList.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ShowList(l)
	resL := ReverseList.ReverseList(l)
	ShowList(resL)
	fmt.Println("Done!")
}

func ShowList(l *ReverseList.ListNode) {
	node := l
	for {
		if node == nil {
			break
		}
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}
