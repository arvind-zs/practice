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
		s := New()
		s.Push(tc.input)

		if !reflect.DeepEqual(tc.expRes, s.S) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, s.S)
		}
	}
}

func TestPop(t *testing.T) {
	testcases := []struct {
		desc   string
		st     []int
		expRes int
	}{
		{desc: "success:stack st is not empty", st: []int{2, 3}, expRes: 3},
		{desc: "failure:stack st is  empty", st: nil, expRes: -1},
	}

	for i, tc := range testcases {
		s := New()
		s.S = tc.st
		res := s.Pop()

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}
	}

}
