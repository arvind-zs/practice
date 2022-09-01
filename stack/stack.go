package stack

type Stack struct {
	S []int
}

func New() Stack {
	var s []int
	return Stack{s}
}

func (st *Stack) Push(ele int) {
	if len(st.S) < 100 {
		st.S = append(st.S, ele)
	}
}

func (st *Stack) Len() int {
	return len(st.S)
}

func (st *Stack) Pop() int {
	if len(st.S) != 0 {
		lastElement := st.S[len(st.S)-1]
		st.S = st.S[:len(st.S)-1]
		return lastElement
	}
	return -1
}
