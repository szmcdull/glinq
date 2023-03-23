package gset

import (
	"fmt"

	"github.com/szmcdull/glinq/garray"
	"github.com/szmcdull/glinq/gmap"
	"golang.org/x/exp/constraints"
)

// Experimental. Use at your own risk

type (
	HashSet[T comparable] map[T]struct{}
)

func FromSlice[S ~[]T, T comparable](l S) map[T]struct{} {
	result := make(map[T]struct{}, len(l))
	for _, v := range l {
		result[v] = struct{}{}
	}
	return result
}

func ToSlice[M ~map[T]struct{}, T comparable](m M) []T {
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Add all in B to A
func Add[M ~map[T]struct{}, T comparable](A, B M) {
	for k := range B {
		A[k] = struct{}{}
	}
}

func AddItems[M ~map[T]struct{}, T comparable](s M, v ...T) {
	for _, vv := range v {
		s[vv] = struct{}{}
	}
}

// Remove all in B from A
func Sub[M ~map[T]struct{}, T comparable](A, B M) {
	for k := range A {
		if _, ok := B[k]; ok {
			delete(A, k)
		}
	}
}

// Remove anything not in B from A
func And[M ~map[T]struct{}, T comparable](A, B M) {
	for k := range A {
		if _, ok := B[k]; !ok {
			delete(A, k)
		}
	}
}

// Shallow copy
func Copy[M ~map[T]struct{}, T comparable](other M) M {
	result := make(map[T]struct{}, len(other))
	for k := range other {
		result[k] = struct{}{}
	}
	return result
}

func NewFromSlice[S ~[]T, T comparable](source S) HashSet[T] {
	return HashSet[T](FromSlice(source))
}

func (s HashSet[T]) ToSlice() []T {
	return gmap.Keys(s)
}

func (s HashSet[T]) Add(other map[T]struct{}) HashSet[T] {
	result := HashSet[T](Copy(s))
	Add(result, other)
	return result
}

func (s HashSet[T]) Sub(other map[T]struct{}) HashSet[T] {
	result := HashSet[T](Copy(s))
	Sub(result, other)
	return result
}

func (s HashSet[T]) And(other map[T]struct{}) HashSet[T] {
	result := HashSet[T](Copy(s))
	And(result, other)
	return result
}

func (s HashSet[T]) String() string {
	b := make([]byte, 0, len(s)*5)
	b = append(b, []byte(`Set[`)...)
	for k := range s {
		b = append(b, []byte(fmt.Sprintf(`%#v `, k))...)
	}
	b[len(b)-1] = ']'
	return string(b)
}

func (s HashSet[T]) Contains(s2 HashSet[T]) bool {
	for v2 := range s2 {
		_, ok := s[v2]
		if !ok {
			return false
		}
	}
	return true
}

func (s HashSet[T]) ContainsItem(v T) bool {
	_, ok := s[v]
	return ok
}

func (s HashSet[T]) AddItemChecked(v T) bool {
	if _, ok := s[v]; ok {
		return false
	}
	s[v] = struct{}{}
	return true
}

func (s HashSet[T]) RemoveItem(v T) {
	delete(s, v)
}

func (s HashSet[T]) RemoveItemChecked(v T) bool {
	if _, ok := s[v]; !ok {
		return false
	}
	delete(s, v)
	return true
}

// Sorted returns a sorted set as []T
func Sorted[T constraints.Ordered](s HashSet[T]) []T {
	result := gmap.Keys(s)
	garray.Sort(result)
	return result
}
