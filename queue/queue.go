package queue

import "practice/stack"

type queue struct {
	S1 []int
	S2 []int
}

func New(s1, s2 []int) queue {
	return queue{S1: s1, S2: s2}
}

func (q *queue) Push(ele int) {
	stack.Push(ele, &q.S1)
}

func (q *queue) Pop() int {
	if len(q.S2) == 0 {
		q.pushOneStackToAnother()
	}

	//   popping from stack S2
	lastElement := stack.Pop(&q.S2)
	return lastElement
}

func (q *queue) pushOneStackToAnother() {
	for len(q.S1) > 0 {
		// popping from stack s1
		lastElement := stack.Pop(&q.S1)

		//  pushing into stack s2
		stack.Push(lastElement, &q.S2)
	}
}
