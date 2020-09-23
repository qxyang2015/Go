package queue

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (cq *CQueue) AppendTail(value int) {
	cq.stack1.PushBack(value)
}

func (cq *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if cq.stack2.Len() == 0 {
		for cq.stack1.Len() > 0 {
			cq.stack2.PushBack(cq.stack1.Remove(cq.stack1.Back()))
		}
	}
	if cq.stack2.Len() != 0 {
		e := cq.stack2.Back()
		cq.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
