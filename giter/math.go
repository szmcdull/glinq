//go:build go1.23
// +build go1.23

package giter

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// type (
// 	Slice[T any] []T

// 	StdIter[T any] iter.Seq[T]

// 	Iter[T any] interface {
// 		Iter() iter.Seq[T]
// 	}
// )

// func (s Slice[T]) Iter() iter.Seq[T] {
// 	return slices.Values(s)
// }

// func (i StdIter[T]) Iter() iter.Seq[T] {
// 	return iter.Seq[T](i)
// }

func Min[T constraints.Ordered](it iter.Seq[T]) T {
	var min T
	first := true

	for v := range it {
		if first || v < min {
			min = v
			first = false
		}
	}
	if first {
		panic("Min called on empty iterator")
	}
	return min
}

func Max[T constraints.Ordered](it iter.Seq[T]) T {
	var max T
	first := true
	for v := range it {
		if first || v > max {
			max = v
			first = false
		}
	}
	if first {
		panic("Max called on empty iterator")
	}
	return max
}

func MinBy[T any, V constraints.Ordered](it iter.Seq[T], selector func(T) V) T {
	var min T
	var minVal V
	first := true
	for v := range it {
		val := selector(v)
		if first || val < minVal {
			min = v
			minVal = val
			first = false
		}
	}
	if first {
		panic("MinBy called on empty iterator")
	}
	return min
}

func MaxBy[T any, V constraints.Ordered](it iter.Seq[T], selector func(T) V) T {
	var max T
	var maxVal V
	first := true
	for v := range it {
		val := selector(v)
		if first || val > maxVal {
			max = v
			maxVal = val
			first = false
		}
	}
	if first {
		panic("MaxBy called on empty iterator")
	}
	return max
}

func Sum[T constraints.Integer | constraints.Float](it iter.Seq[T]) T {
	var sum T
	for v := range it {
		sum += v
	}
	return sum
}

func Average[T constraints.Integer | constraints.Float](it iter.Seq[T]) T {
	var sum T
	var count int
	for v := range it {
		sum += v
		count++
	}
	if count == 0 {
		panic("Average called on empty iterator")
	}
	return sum / T(count)
}
