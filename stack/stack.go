package stack

type Stack struct {
	S1 []int
}

func New() Stack {
	var s1 []int
	return Stack{s1}
}

func (s *Stack) Push(ele int) {
	s.S1 = append(s.S1, ele)
}

func (s *Stack) Pop() int {
	if len(s.S1) != 0 {
		lastElement := s.S1[len(s.S1)-1]
		s.S1 = s.S1[:len(s.S1)-1]
		return lastElement
	}
	return 0
}
