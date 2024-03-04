package garray

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// sortable a generics slice which implements sort.Interface
// inspired by https://github.com/amtoaer/generic-sort
type (
	slice[T any]                    []T
	sortable[T constraints.Ordered] struct {
		slice[T]
	}
	sortableDescending[T constraints.Ordered] struct {
		sortable[T]
	}
	sortableBy[T any, V constraints.Ordered] struct {
		slice[T]
		selector func(T) V
	}
	sortableByDescending[T any, V constraints.Ordered] struct {
		sortableBy[T, V]
	}
	sortableByComparer[T any] struct {
		slice[T]
		less func(T, T) bool
	}
)

func Sortable[T constraints.Ordered](l []T) sortable[T] {
	return sortable[T]{l}
}

func SortableDescending[T constraints.Ordered](l []T) sortableDescending[T] {
	return sortableDescending[T]{sortable[T]{l}}
}

func Sort[T constraints.Ordered](l []T) {
	sort.Sort(Sortable(l))
}

func SortDescending[T constraints.Ordered](l []T) {
	sort.Sort(SortableDescending(l))
}

func OrderBy[T any, V constraints.Ordered](l []T, selector func(T) V) *sortableBy[T, V] {
	return &sortableBy[T, V]{slice[T](l), selector}
}

func OrderByDescending[T any, V constraints.Ordered](l []T, selector func(T) V) *sortableByDescending[T, V] {
	return &sortableByDescending[T, V]{sortableBy[T, V]{slice[T](l), selector}}
}

func SortBy[T any, V constraints.Ordered](l []T, selector func(T) V) {
	sort.Sort(OrderBy(l, selector))
}

func SortByDescending[T any, V constraints.Ordered](l []T, selector func(T) V) {
	sort.Sort(OrderByDescending(l, selector))
}

func SortByComparer[T any](l []T, less func(T, T) bool) {
	sort.Sort(sortableByComparer[T]{slice[T](l), less})
}

func (s slice[T]) Len() int {
	return len(s)
}

func (s slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortable[T]) Less(i, j int) bool {
	return s.slice[i] < s.slice[j]
}

func (s sortableDescending[T]) Less(i, j int) bool {
	return s.sortable.Less(j, i)
}

func (s sortableBy[T, V]) Less(i, j int) bool {
	return s.selector(s.slice[i]) < s.selector(s.slice[j])
}

func (s sortableByDescending[T, V]) Less(i, j int) bool {
	return s.selector(s.slice[i]) > s.selector(s.slice[j])
}

func (s sortableByComparer[T]) Less(i, j int) bool {
	return s.less(s.slice[i], s.slice[j])
}
