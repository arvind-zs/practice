package queue

import (
	"reflect"
	"testing"

	"practice/stack"
)

func TestPush(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expRes stack.Stack
		expErr error
	}{
		{desc: "success:pushing element in queue", input: 1, expRes: stack.Stack{S: []int{1}}, expErr: nil},
	}

	for i, tc := range testcases {
		q := New(stack.Stack{}, stack.Stack{})
		err := q.Push(tc.input)

		if !reflect.DeepEqual(tc.expRes, q.s1) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, q.s1)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}
}

func TestPop(t *testing.T) {
	testcases := []struct {
		desc   string
		S1     []int
		S2     []int
		expRes int
		expErr error
	}{
		{desc: "success:stack S2 is not empty", S1: []int{2, 3}, S2: []int{1}, expRes: 1, expErr: nil},
	}

	for i, tc := range testcases {
		q := New(stack.Stack{S: tc.S1}, stack.Stack{S: tc.S2})
		res, err := q.Pop()

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}

		if !reflect.DeepEqual(tc.expErr, err) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expErr, err)
		}
	}

}
