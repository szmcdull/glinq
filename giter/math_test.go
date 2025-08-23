//go:build go1.23
// +build go1.23

package giter

import (
	"slices"
	"testing"

	"github.com/szmcdull/quantainer"
)

func TestMin(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if Min(slices.Values(s)) != 1 {
		t.Fail()
	}
}

func TestMaxBy(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if MaxBy(slices.Values(s), func(x int) int {
		return -x
	}) != 1 {
		t.Fail()
	}
}

func TestExpand(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	i := 1
	slices.Values(s)(func(x int) bool {
		if x != i {
			t.Fail()
		}
		i++
		return true
	})
}

func TestIter(t *testing.T) {
	rb := quantainer.NewRingBuffer[int](5)
	rb.AddLast(1)
	rb.AddLast(2)
	rb.AddLast(3)
	rb.AddLast(4)
	rb.AddLast(5)
	i := 1
	for v := range rb.Values() {
		if v != i {
			t.Fail()
		}
		i++
	}
}

func TestIterMaxBy(t *testing.T) {
	rb := quantainer.NewRingBuffer[int](5)
	rb.AddLast(1)
	rb.AddLast(2)
	rb.AddLast(3)
	rb.AddLast(4)
	rb.AddLast(5)
	if MaxBy(rb.Values(), func(x int) int {
		return -x
	}) != 1 {
		t.Fail()
	}
	if Max(rb.Values()) != 5 {
		t.Fail()
	}
}
