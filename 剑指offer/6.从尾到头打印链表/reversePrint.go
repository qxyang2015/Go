package reversePrint

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(ReversePrint(head.Next), head.Val)
}
