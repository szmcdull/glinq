package gset

import (
	"testing"

	"github.com/szmcdull/glinq/unsafe"
)

func Test(t *testing.T) {
	s := NewFromSlice(unsafe.ToSlice(unsafe.Range(1, 11)))
	s2 := NewFromSlice(unsafe.ToSlice(unsafe.Range(5, 15)))
	s3 := s.Add(s2)
	s4 := s.Sub(s2)
	if len(s3) != 14 || !s3.Contains(1) || !s3.Contains(2) || !s3.Contains(3) || !s3.Contains(4) || !s3.Contains(5) || !s3.Contains(6) || !s3.Contains(7) || !s3.Contains(8) || !s3.Contains(9) || !s3.Contains(10) || !s3.Contains(11) || !s3.Contains(12) || !s3.Contains(13) || !s3.Contains(14) {
		t.Errorf(`str3=%v expected Set[1 2 3 4 5 6 7 8 9 10 11 12 13 14]`, s3.String())
	}
	if len(s4) != 4 || !s4.Contains(1) || !s4.Contains(2) || !s4.Contains(3) || !s4.Contains(4) {
		t.Errorf(`str4=%v expected Set[1 2 3 4]`, s4.String())
	}
}
