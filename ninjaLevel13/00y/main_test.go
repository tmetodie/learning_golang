package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{5, 13}, 18},
		test{[]int{-2, 3}, 1},
		test{[]int{-2, -5}, -7},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected:", v.answer, ", Got:", x)
		}
	}
}
