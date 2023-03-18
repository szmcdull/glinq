package gset

import (
	"fmt"
	"testing"

	"github.com/szmcdull/glinq/garray"
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
	set.AddItems(3, 2, 1, 1, 2, 3)
	slice := set.ToSlice()
	garray.Sort(slice)
	fmt.Println(slice)
	// Output: [1 2 3]
}
