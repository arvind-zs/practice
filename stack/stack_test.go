package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expRes []int
	}{
		{desc: "success:pushing element in stack", input: 1, expRes: []int{1}},
	}

	for i, tc := range testcases {
		s := New(nil)
		s.Push(tc.input)

		if !reflect.DeepEqual(tc.expRes, s.S1) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, s.S1)
		}
	}
}

func TestPop(t *testing.T) {
	testcases := []struct {
		desc   string
		S1     []int
		expRes int
	}{
		{desc: "success:stack S1 is not empty", S1: []int{2, 3}, expRes: 3},
		{desc: "failure:stack S1 is  empty", S1: nil, expRes: -1},
	}

	for i, tc := range testcases {
		s := New(tc.S1)
		res := s.Pop()

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}
	}

}
