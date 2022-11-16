package gmap

import "testing"

func TestShallowCopy(t *testing.T) {
	l := map[int]int{1: 2, 2: 4, 3: 6}
	result := ShallowCopy(l)
	if result[1] != 2 ||
		result[2] != 4 ||
		result[3] != 6 {

		t.Fail()
	}
}
