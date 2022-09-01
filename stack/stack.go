package stack

func Push(ele int, s *[]int) {
	*s = append(*s, ele)
}

func Pop(s *[]int) int {
	if len(*s) != 0 {
		lastElement := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return lastElement
	}
	return -1
}
