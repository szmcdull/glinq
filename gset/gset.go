package gset

import (
	"fmt"
)

// Experimental. Use at your own risk

type (
	HashSet[T comparable] map[T]struct{}
)

func FromSlice[T comparable](l []T) map[T]struct{} {
	result := make(map[T]struct{}, len(l))
	for _, v := range l {
		result[v] = struct{}{}
	}
	return result
}

func ToSlice[T comparable](m map[T]struct{}) []T {
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Add all in B to A
func Add[T comparable](A, B map[T]struct{}) {
	for k := range B {
		A[k] = struct{}{}
	}
}

// Remove all in B from A
func Sub[T comparable](A, B map[T]struct{}) {
	for k := range A {
		if _, ok := B[k]; ok {
			delete(A, k)
		}
	}
}

// Remove anything not in B from A
func And[T comparable](A, B map[T]struct{}) {
	for k := range A {
		if _, ok := B[k]; !ok {
			delete(A, k)
		}
	}
}

// Shallow copy
func Copy[T comparable](other map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{}, len(other))
	for k := range other {
		result[k] = struct{}{}
	}
	return result
}

func NewFromSlice[T comparable](source []T) HashSet[T] {
	return HashSet[T](FromSlice(source))
}

func (s HashSet[T]) ToSlice() []T {
	return s.ToSlice()
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

func (s HashSet[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}
