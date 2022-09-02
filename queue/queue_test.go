package queue

import (
	"errors"
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
		{desc: "failure:queue capacity is equal to its maximum capacity", input: 3, expRes: stack.Stack{},
			expErr: errors.New("you can not enter element  right now...stack capacity is equal to its maximum capacity")},
	}

	for i, tc := range testcases {
		q := New(stack.Stack{}, stack.Stack{})
		if i == 1 {
			q = New(stack.Stack{S: make([]int, 0, 0)}, stack.Stack{S: []int{1}})
		}

		err := q.Push(tc.input)
		
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
		{desc: "success:stack S2 is  empty", S1: []int{2, 3}, S2: make([]int, 0, 100), expRes: 2, expErr: nil},
		{desc: "failure:stack S2 capacity is  equal to its maximum capacity", S1: []int{2, 3}, S2: make([]int, 0, 0),
			expRes: -1, expErr: errors.New("stack is empty right now")},
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
