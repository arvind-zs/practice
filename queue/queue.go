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
	s := stack.New(q.S1)
	s.Push(ele)
	q.S1 = s.S1
}

func (q *queue) Pop() int {
	if len(q.S2) == 0 {
		q.pushOneStackToAnother(q.S1, q.S2)
	}

	//   popping from stack S2
	s := stack.New(q.S2)
	lastElement := s.Pop()
	q.S2 = s.S1
	return lastElement
}

func (q *queue) pushOneStackToAnother(s1, s2 []int) {
	for len(s1) > 0 {
		// popping from stack s1
		s := stack.New(s1)
		lastElement := s.Pop()
		s1 = s.S1

		//  pushing into stack s2
		s = stack.New(s2)
		s.Push(lastElement)
		s2 = s.S1
	}

	q.S1 = s1
	q.S2 = s2
}
