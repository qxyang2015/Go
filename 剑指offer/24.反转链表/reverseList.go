package ReverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preNode := head
	node := head.Next
	for node != nil {
		nexNode := node.Next

		node.Next = preNode

		preNode, node = node, nexNode
	}
	head.Next = nil
	return preNode
}
