package garray

import (
	"fmt"
	"testing"
)

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

func TestMap(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	result := Map(l, func(i int) string {
		return fmt.Sprintf(`%d`, i)
	})
	if result[0] != `1` ||
		result[1] != `2` ||
		result[2] != `3` ||
		result[3] != `4` ||
		result[4] != `5` {

		t.Fail()
	}
}

func TestShallowCopy(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	result := ShallowCopy(l)
	if result[0] != 1 ||
		result[1] != 2 ||
		result[2] != 3 ||
		result[3] != 4 ||
		result[4] != 5 {

		t.Fail()
	}
}

func ExampleFindIf() {
	list := []int{1, 2, 3, 4, 5}
	pos := FindIf(list, func(i int) bool {
		return list[i] == 3
	})
	fmt.Println(pos)
	// output: 2
}

func ExampleRemoveIf() {
	list := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	result, deleted := RemoveIf(list, func(i int) bool {
		return list[i] == 2
	})
	fmt.Println(result, deleted)
	// output: [1 3 1 3 1 3] 3
}
