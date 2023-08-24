package garray

import (
	"fmt"
	"log"
	"strings"
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
	fmt.Println(CountIf(l, func(i int) bool { return i%2 == 0 }))
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

func ExampleFilter() {
	lst := []int{1, 2, 3, 4, 5, 6}
	isEven := func(n int) bool { return n%2 == 0 } // expected filtered array of even numbers

	evens := Filter(lst, isEven)
	fmt.Println(evens)
	// Output: [2 4 6]
}

func ExampleFilterI() {
	lst := []string{"apple", "banana", "cherry", "grape"}
	firstTwoItems := func(i int) bool {
		return i < 2 // expected Filtered array of first two items from the list
	}

	result := FilterI(lst, firstTwoItems)
	fmt.Println(result)
	// Output: [apple banana]
}

func ExampleFilterIE() {
	nums := []int{1, 2, 3, 4, 5}

	evenNums := func(i int) (bool, error) {
		return nums[i]%2 == 0, nil // expected filterd array with only even numbers
	}

	filteredNums, _ := FilterIE(nums, evenNums)
	fmt.Println(filteredNums)
	// Output: [2 4]
}

func ExampleApply() {
	numList := []int{1, 2, 3, 4, 5, 6}
	double := func(n int) error {
		fmt.Print(n*2, " ")
		return nil
	}

	err := Apply(numList, double)
	if err != nil {
		fmt.Println("Error Occured")
	}
	// expected output to be array elements multiplied by 2

	// Output:2 4 6 8 10 12
}

func ExampleApplyI() {
	words := []string{"bar", "foo"}
	var capitalize = func(i int) error {
		words[i] = strings.ToUpper(words[i])
		return nil
	}

	err := ApplyI(words, capitalize) // all elements capitalized
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", words)
	// Output:[BAR FOO]
}

func ExampleSumBy() {
	type student struct {
		name string
		age  int
	}
	students := []student{{`Jack`, 10}, {`Mike`, 20}, {`Rose`, 15}}

	totalAge := SumBy(students, func(s student) int {
		return s.age
	})
	fmt.Println(totalAge)
	// Output: 45
}

func ExampleSumByP() {
	type student struct {
		name string
		age  int
	}
	students := []student{{`Jack`, 10}, {`Mike`, 20}, {`Rose`, 15}}

	totalAge := SumByP(students, func(s *student) int {
		return s.age
	})
	fmt.Println(totalAge)
	// Output: 45
}

func ExampleSumByI() {
	type student struct {
		name string
		age  int
	}
	students := []student{{`Jack`, 10}, {`Mike`, 20}, {`Rose`, 15}}

	totalAge := SumByI(students, func(i int) int {
		return students[i].age
	})
	fmt.Println(totalAge)
	// Output: 45
}
