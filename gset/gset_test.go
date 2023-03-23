package gset

import (
	"fmt"
	"testing"

	"github.com/szmcdull/glinq/unsafe"
)

func Test(t *testing.T) {
	s := NewFromSlice(unsafe.ToSlice(unsafe.Range(1, 11)))
	s2 := NewFromSlice(unsafe.ToSlice(unsafe.Range(5, 15)))
	s3 := s.Add(s2)
	s4 := s.Sub(s2)
	if len(s3) != 14 || !s3.ContainsItem(1) || !s3.ContainsItem(2) || !s3.ContainsItem(3) || !s3.ContainsItem(4) || !s3.ContainsItem(5) || !s3.ContainsItem(6) || !s3.ContainsItem(7) || !s3.ContainsItem(8) || !s3.ContainsItem(9) || !s3.ContainsItem(10) || !s3.ContainsItem(11) || !s3.ContainsItem(12) || !s3.ContainsItem(13) || !s3.ContainsItem(14) {
		t.Errorf(`str3=%v expected Set[1 2 3 4 5 6 7 8 9 10 11 12 13 14]`, s3.String())
	}
	if len(s4) != 4 || !s4.ContainsItem(1) || !s4.ContainsItem(2) || !s4.ContainsItem(3) || !s4.ContainsItem(4) {
		t.Errorf(`str4=%v expected Set[1 2 3 4]`, s4.String())
	}
	if !s3.Contains(s2) {
		t.Fail()
	}
	if s2.Contains(s3) {
		t.Fail()
	}
}

func ExampleHashSet_ToSlice() {
	set := HashSet[int]{}
	AddItems(set, 3, 2, 1, 1, 2, 3)
	fmt.Println(Sorted(set))
	// Output: [1 2 3]
}

func ExampleHashSet_Add() {
	// Initialize HashSets
	set1 := HashSet[string]{"apple": {}, "banana": {}}
	set2 := HashSet[string]{"pear": {}, "orange": {}}
	set3 := HashSet[string]{"banana": {}, "durian": {}}

	// calculate the union of them
	result := set1.
		Add(set2).
		Add(set3)

	// Print the contents of the updated set
	for _, k := range Sorted(result) {
		fmt.Println(k)
	}

	// Output:
	// apple
	// banana
	// durian
	// orange
	// pear
}

func ExampleHashSet_Sub() {
	// Initialize HashSet
	set := HashSet[string]{"apple": {}, "banana": {}, "pear": {}, "orange": {}}

	// Remove elements in other from set
	other := map[string]struct{}{"pear": {}, "orange": {}}
	result := set.Sub(other)

	// Print the contents of the updated set
	for _, k := range Sorted(result) {
		fmt.Println(k)
	}

	// Output:
	// apple
	// banana
}

func ExampleHashSet_And() {
	// Initialize HashSet
	set := HashSet[string]{"apple": {}, "banana": {}, "pear": {}, "orange": {}}

	// Keep only elements also in other
	other := HashSet[string]{"apple": {}, "pear": {}, "grape": {}}
	result := set.And(other)

	// Print the contents of the updated set
	for _, k := range Sorted(result) {
		fmt.Println(k)
	}

	// Output:
	// apple
	// pear
}

func ExampleAddItems() {
	s := make(map[int]struct{})
	AddItems(s, 1, 2, 3)
	fmt.Println(s) // Output: map[1:{} 2:{} 3:{}]
}

func ExampleAdd() {
	A := make(map[int]struct{})
	B := map[int]struct{}{1: {}, 2: {}}
	Add(A, B)
	fmt.Println(A) // Output: map[1:{} 2:{}]
}

func ExampleSub() {
	A := map[int]struct{}{1: {}, 2: {}, 3: {}}
	B := map[int]struct{}{2: {}}
	Sub(A, B)
	fmt.Println(A) // Output: map[1:{} 3:{}]
}

func ExampleAnd() {
	A := map[int]struct{}{1: {}, 2: {}, 3: {}}
	B := map[int]struct{}{2: {}, 3: {}, 4: {}}
	And(A, B)
	fmt.Println(A) // Output: map[2:{} 3:{}]
}

func ExampleSorted() {
	// create a new HashSet and add some elements
	set := HashSet[int]{}
	AddItems(set, 5, 1, 9)

	// get the sorted list of elements
	sortedList := Sorted(set)

	// print the sorted list
	fmt.Println(sortedList)

	// Output: [1 5 9]
}
