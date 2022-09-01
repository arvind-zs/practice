package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		s      []int
		expRes []int
	}{
		{desc: "success:pushing element in stack", input: 1, s: []int{3, 4}, expRes: []int{3, 4, 1}},
	}

	for i, tc := range testcases {

		Push(tc.input, &tc.s)

		if !reflect.DeepEqual(tc.expRes, tc.s) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, tc.s)
		}
	}
}

func TestPop(t *testing.T) {
	testcases := []struct {
		desc   string
		S      []int
		expRes int
	}{
		{desc: "success:stack S1 is not empty", S: []int{2, 3}, expRes: 3},
		{desc: "failure:stack S1 is  empty", S: nil, expRes: -1},
	}

	for i, tc := range testcases {

		res := Pop(&tc.S)

		if !reflect.DeepEqual(tc.expRes, res) {
			t.Errorf("testcases %d failed expected %v got %v", i+1, tc.expRes, res)
		}
	}

}
