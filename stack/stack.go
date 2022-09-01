package stack

type stack struct {
	S1 []int
}

func New(s []int) stack {
	return stack{s}
}

func (s *stack) Push(ele int) {
	s.S1 = append(s.S1, ele)
}

func (s *stack) Pop() int {
	if len(s.S1) != 0 {
		lastElement := s.S1[len(s.S1)-1]
		s.S1 = s.S1[:len(s.S1)-1]
		return lastElement
	}
	return -1
}
