package gset

import (
	"fmt"
	"testing"

	"github.com/szmcdull/glinq/unsafe"
)

func Test(t *testing.T) {
	s := NewFromSlice(unsafe.ToSlice(unsafe.Range(1, 11)))
	s2 := NewFromSlice(unsafe.ToSlice(unsafe.Range(5, 15)))
	s3 := Copy(s)
	s4 := Copy(s)
	s3.Add(s2)
	s4.Sub(s2)
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
	set.AddItems(3, 2, 1, 1, 2, 3)
	fmt.Println(Sorted(set))
	// Output: [1 2 3]
}

func ExampleHashSet_Add() {
	// Initialize HashSet
	set := HashSet[string]{"apple": {}, "banana": {}}

	// Add other to set
	other := map[string]struct{}{"pear": {}, "orange": {}}
	set.Add(other)

	// Print the contents of the updated set
	for _, k := range Sorted(set) {
		fmt.Println(k)
	}

	// Output:
	// apple
	// banana
	// orange
	// pear
}

func ExampleHashSet_Sub() {
	// Initialize HashSet
	set := HashSet[string]{"apple": {}, "banana": {}, "pear": {}, "orange": {}}

	// Remove elements in other from set
	other := map[string]struct{}{"pear": {}, "orange": {}}
	set.Sub(other)

	// Print the contents of the updated set
	for _, k := range Sorted(set) {
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
	set.And(other)

	// Print the contents of the updated set
	for _, k := range Sorted(set) {
		fmt.Println(k)
	}

	// Output:
	// apple
	// pear
}

func ExampleSorted() {
	// create a new HashSet and add some elements
	set := HashSet[int]{}
	set.AddItem(5)
	set.AddItem(1)
	set.AddItem(9)

	// get the sorted list of elements
	sortedList := Sorted(set)

	// print the sorted list
	fmt.Println(sortedList)

	// Output: [1 5 9]
}
