package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestSum(t *testing.T) {

	info := []test{

		{[]int{22, 23}, 45},
		{[]int{22, 3}, 25},
		{[]int{22, 100}, 122},
	}
	for _, v := range info {

		val := v
		x := foo(val.data...)
		if x != val.answer {
			t.Error("expected ", val.answer, "got", x)
		}

	}

}
