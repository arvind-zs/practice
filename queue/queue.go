package queue

import (
	"log"

	"practice/stack"
)

type queue struct {
	s1 stack.Stack
	s2 stack.Stack
}

func New(s1, s2 stack.Stack) queue {
	return queue{s1: s1, s2: s2}
}

func (q *queue) Push(ele int) error {
	if q.s1.Len() == 0 && q.s2.Len() == 0 {
		q.s1 = stack.New()
		q.s2 = stack.New()
	}

	return q.s1.Push(ele)
}

func (q *queue) Pop() (int, error) {
	if q.s2.Len() == 0 {
		q.pushOneStackToAnother()
	}

	//   popping from stack S2
	return q.s2.Pop()
}

func (q *queue) pushOneStackToAnother() {
	for q.s1.Len() > 0 {
		// popping from stack s1
		lastElement, _ := q.s1.Pop()

		//  pushing into stack s2
		err := q.s2.Push(lastElement)
		if err != nil {
			log.Println(err.Error())

			return
		}
	}
}
