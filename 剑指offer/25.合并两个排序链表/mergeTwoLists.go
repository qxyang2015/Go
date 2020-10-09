package MergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	lNext := l
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			lNext.Next = l2
			l2 = l2.Next
		} else {
			lNext.Next = l1
			l1 = l1.Next
		}
		lNext = lNext.Next
	}
	if l1 != nil {
		lNext.Next = l1
	} else {
		lNext.Next = l2
	}
	return l.Next
}
