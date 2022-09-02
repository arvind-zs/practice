package stack

import "errors"

type Stack struct {
	S []int
}

func New() Stack {
	s := make([]int, 0, 100)
	return Stack{s}
}

func (st *Stack) Push(ele int) error {
	if st.Len() < cap(st.S) {
		st.S = append(st.S, ele)
		return nil
	}

	return errors.New("you can not enter element  right now...stack capacity is equal to its maximum capacity")
}

func (st *Stack) Len() int {
	return len(st.S)
}

func (st *Stack) Pop() (int, error) {
	if len(st.S) != 0 {
		lastElement := st.S[len(st.S)-1]
		st.S = st.S[:len(st.S)-1]
		return lastElement, nil
	}

	return -1, errors.New("stack is empty right now")
}
