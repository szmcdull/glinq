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

func ExampleMapIE() {
	type student struct {
		name string
		age  int
	}

	students := []student{{`Jack`, 10}, {`Mike`, 20}, {`Rose`, 15}}

	studentAges, err := ToMapIE(students,
		func(i int) (string, error) { return students[i].name, nil },
		func(i int) (int, error) { return students[i].age, nil })
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(studentAges)
	}

	// output: map[Jack:10 Mike:20 Rose:15]
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
	if &l[0] == &result[0] {
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

func ExampleCount() {
	l := []int{1, 0, 1, 0, 1}
	fmt.Println(Count(l, 1))
	// output: 3
}

func ExampleCountIf() {
	l := []int{1, 2, 3, 4, 5}
	fmt.Println(CountIf(l, func(i int) bool { return i/2*2 == i }))
	// output: 2
}

func ExampleReverse() {
	l := []int{1, 2, 3, 4}
	Reverse(l)
	fmt.Println(l)
	l = []int{1, 2, 3}
	Reverse(l)
	fmt.Println(l)
	// output: [4 3 2 1]
	// [3 2 1]
}
