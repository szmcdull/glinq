package garray

import (
	"time"

	"golang.org/x/exp/constraints"
)

/*
 * Some utilities for slices
 */

type (
	Number interface {
		constraints.Integer | constraints.Float | time.Duration
	}
)

// like LINQ.Select, map a slice to another slice
func Map[Src any, Dst any](l []Src, f func(Src) Dst) []Dst {
	result := make([]Dst, len(l))
	for i, src := range l {
		result[i] = f(src)
	}
	return result
}

func MapI[Src any, Dst any](l []Src, f func(i int) Dst) []Dst {
	result := make([]Dst, len(l))
	for i := range l {
		result[i] = f(i)
	}
	return result
}

func MapIE[Src any, Dst any](l []Src, f func(i int) (Dst, error)) (result []Dst, err error) {
	result = make([]Dst, len(l))
	for i := range l {
		result[i], err = f(i)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Filter selects items that f() returns true
func Filter[Src any](l []Src, f func(Src) bool) (result []Src) {
	result = make([]Src, 0, len(l)/2)
	for _, x := range l {
		if f(x) {
			result = append(result, x)
		}
	}
	return
}

// FilterI selects items that f() returns true
func FilterI[Src any](l []Src, f func(i int) bool) (result []Src) {
	result = make([]Src, 0, len(l)/2)
	for i := range l {
		if f(i) {
			result = append(result, l[i])
		}
	}
	return
}

// FilterIE selects items that f() returns true
func FilterIE[Src any](l []Src, f func(i int) (bool, error)) (result []Src, err error) {
	result = make([]Src, 0, len(l)/2)
	for i := range l {
		if r, err := f(i); err == nil {
			if r {
				result = append(result, l[i])
			}
		} else {
			return nil, err
		}
	}
	return result, nil
}

// Apply a function to each item of a slice
func Apply[S ~[]T, T any](l S, fun func(T) error) error {
	for i := range l {
		if err := fun(l[i]); err != nil {
			return err
		}
	}
	return nil
}

func ApplyI[S ~[]T, T any](l S, fun func(int) error) error {
	for i := range l {
		if err := fun(i); err != nil {
			return err
		}
	}
	return nil
}

// Calculate the average of a slice of numbers
func Average[S ~[]T, T Number](l S) T {
	result := Sum[S, T](l) / T(len(l))
	return result
}

// Sum a slice of numbers
func Sum[S ~[]T, T Number](l S) T {
	var sum T
	for _, x := range l {
		sum += x
	}
	return sum
}

// find first matching element
func First[S ~[]T, T any](l S, pred func(T) bool) (T, bool) {
	for _, x := range l {
		if pred(x) {
			return x, true
		}
	}
	var r T
	return r, false
}

// find last matching element
func Last[S ~[]T, T any](l S, pred func(T) bool) (T, bool) {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if pred(x) {
			return x, true
		}
	}
	var r T
	return r, false
}

// find the position of an element
func IndexOf[S ~[]T, T comparable](l S, v T) int {
	for i, x := range l {
		if x == v {
			return i
		}
	}
	return -1
}

// FindIf finds the position of an matching element.
// Its use case is the same as IndexWhere, except that you don't have to specify the type of the elements.
// (sometimes the type is very long, or it is an anonymous/temporary type that you have to repeat the definition.)
// It passes each index of l to pred() and returns the first i which pred(i) is true.
// I find it interestingly handful :) It opens a new way to generics programming that does not support lambda expression.
func FindIf[S ~[]T, T any](l S, pred func(i int) bool) int {
	for i := range l {
		if pred(i) {
			return i
		}
	}
	return -1
}

// RemoveIf removes matching items.
func RemoveIf[S ~[]T, T any](l S, pred func(i int) bool) (result S, deleted int) {
	result = make(S, 0, len(l))
	for i := range l {
		if !pred(i) {
			result = append(result, l[i])
		} else {
			deleted++
		}
	}
	if deleted == 0 {
		return l, 0
	}
	return result, deleted
}

// find the position of an matching element
func IndexWhere[S ~[]T, T any](l S, pred func(T) bool) int {
	for i, x := range l {
		if pred(x) {
			return i
		}
	}
	return -1
}

// like IndexWhere but parses arguments using pointers, to avoid copying large structs and improve performance
func IndexWhereP[S ~[]T, T any](l S, pred func(*T) bool) int {
	for i := range l {
		if pred(&l[i]) {
			return i
		}
	}
	return -1
}

// find the last position of an element
func LastIndexOf[S ~[]T, T comparable](l S, v T) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if x == v {
			return i
		}
	}
	return -1
}

// find the last position of an matching element
func LastIndexWhere[S ~[]T, T any](l S, pred func(T) bool) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if pred(x) {
			return i
		}
	}
	return -1
}

// like LastIndexWhere but parses arguments using pointers, to avoid copying large structs and improve performance
func LastIndexWhereP[S ~[]T, T any](l S, pred func(*T) bool) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := &l[i]
		if pred(x) {
			return i
		}
	}
	return -1
}

// copy the contents of all slices_, returns a new slice
func Concat[S ~[]T, T any](slices_ ...S) S {
	c := 0
	for i := range slices_ {
		c += len(slices_[i])
	}
	if len(slices_) == 0 || c == 0 {
		return nil
	}

	result := make([]T, c)
	p := 0
	for _, s := range slices_ {
		p += copy(result[p:], s)
	}
	return result
}

func ShallowCopy[T any](s slice[T]) slice[T] {
	result := make(slice[T], len(s))
	copy(result, s)
	return result
}
