package queue

import "practice/stack"

type queue struct {
	s1 stack.Stack
	s2 stack.Stack
}

func New(s1, s2 stack.Stack) queue {
	return queue{s1: s1, s2: s2}
}

func (q *queue) Push(ele int) {
	if len(q.s1.S) == 0 && len(q.s2.S) == 0 {
		q.s1 = stack.New()
		q.s2 = stack.New()
	}

	q.s1.Push(ele)
}

func (q *queue) Pop() int {
	if len(q.s2.S) == 0 {
		q.pushOneStackToAnother()
	}

	//   popping from stack S2
	lastElement := q.s2.Pop()
	return lastElement
}

func (q *queue) pushOneStackToAnother() {
	for len(q.s1.S) > 0 {
		// popping from stack s1
		lastElement := q.s1.Pop()

		//  pushing into stack s2
		q.s2.Push(lastElement)
	}
}
