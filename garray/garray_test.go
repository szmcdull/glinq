package garray

import "testing"

func TestAverage(t *testing.T) {
	if Average([]int{2, 4}) != 3 {
		t.Fail()
	}
}
