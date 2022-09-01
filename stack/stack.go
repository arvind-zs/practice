package stack

type Stack struct {
	S []int
}

func New() Stack {
	var s []int
	return Stack{s}
}

func (st *Stack) Push(ele int) {
	st.S = append(st.S, ele)
}

func (st *Stack) Pop() int {
	if len(st.S) != 0 {
		lastElement := st.S[len(st.S)-1]
		st.S = st.S[:len(st.S)-1]
		return lastElement
	}
	return -1
}
