package garray

import "golang.org/x/exp/constraints"

/*
 * Some utilities for slices
 */

type (
	Number interface {
		constraints.Integer | constraints.Float
	}
)

// Apply a function to each item of a slice
func Apply[T any](l []T, fun func(T)) {
	for i := range l {
		fun(l[i])
	}
}

func Average[T Number](l []T) T {
	result := Sum(l) / T(len(l))
	return result
}

func Sum[T Number](l []T) T {
	var sum T
	for _, x := range l {
		sum += x
	}
	return sum
}

func First[T any](l []T, pref func(T) bool) (T, bool) {
	for _, x := range l {
		if pref(x) {
			return x, true
		}
	}
	var r T
	return r, false
}

func Last[T any](l []T, pref func(T) bool) (T, bool) {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if pref(x) {
			return x, true
		}
	}
	var r T
	return r, false
}

func IndexOf[T comparable](l []T, v T) int {
	for i, x := range l {
		if x == v {
			return i
		}
	}
	return -1
}

func IndexWhere[T any](l []T, pref func(T) bool) int {
	for i, x := range l {
		if pref(x) {
			return i
		}
	}
	return -1
}

func LastIndexOf[T comparable](l []T, v T) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if x == v {
			return i
		}
	}
	return -1
}

func LastIndexWhere[T any](l []T, pref func(T) bool) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if pref(x) {
			return i
		}
	}
	return -1
}
