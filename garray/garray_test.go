package garray

import "testing"

func TestAverage(t *testing.T) {
	if Average([]int{2, 4}) != 3 {
		t.Fail()
	}
}

func TestConcat(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := Concat(a, b)
	if len(c) != 6 || c[0] != 1 || c[1] != 2 || c[2] != 3 || c[3] != 4 || c[4] != 5 || c[5] != 6 {
		t.Fail()
	}
}
