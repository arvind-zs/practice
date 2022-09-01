package stack

import (
	"errors"
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expRes []int
		expErr error
	}{
		{desc: "success:pushing element in stack", input: 1, expRes: []int{1}, expErr: nil},
	}

	for i, tc := range testcases {
		s := New()
		err := s.Push(tc.input)

		if !reflect.DeepEqual(tc.expRes, s.S) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, s.S)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestPop(t *testing.T) {
	testcases := []struct {
		desc   string
		st     []int
		expRes int
		expErr error
	}{
		{desc: "success:stack st is not empty", st: []int{2, 3}, expRes: 3, expErr: nil},
		{desc: "failure:stack st is  empty", st: nil, expRes: -1, expErr: errors.New("stack is empty right now")},
	}

	for i, tc := range testcases {
		s := New()
		s.S = tc.st
		res, err := s.Pop()

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}

}
