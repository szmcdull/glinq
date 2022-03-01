package garray

import (
	"golang.org/x/exp/constraints"
)

/*
 * Some utilities for slices
 */

type (
	Number interface {
		constraints.Integer | constraints.Float
	}
)

// Apply a function to each item of a slice
func Apply[S ~[]T, T any](l S, fun func(T)) {
	for i := range l {
		fun(l[i])
	}
}

func Average[S ~[]T, T Number](l S) T {
	result := Sum[S, T](l) / T(len(l))
	return result
}

func Sum[S ~[]T, T Number](l S) T {
	var sum T
	for _, x := range l {
		sum += x
	}
	return sum
}

func First[S ~[]T, T any](l S, pref func(T) bool) (T, bool) {
	for _, x := range l {
		if pref(x) {
			return x, true
		}
	}
	var r T
	return r, false
}

func Last[S ~[]T, T any](l S, pref func(T) bool) (T, bool) {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if pref(x) {
			return x, true
		}
	}
	var r T
	return r, false
}

func IndexOf[S ~[]T, T comparable](l S, v T) int {
	for i, x := range l {
		if x == v {
			return i
		}
	}
	return -1
}

func IndexWhere[S ~[]T, T any](l S, pref func(T) bool) int {
	for i, x := range l {
		if pref(x) {
			return i
		}
	}
	return -1
}

// use P version when T is a large struct, to improve performance
func IndexWhereP[S ~[]T, T any](l S, pref func(*T) bool) int {
	for i := range l {
		if pref(&l[i]) {
			return i
		}
	}
	return -1
}

func LastIndexOf[S ~[]T, T comparable](l S, v T) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if x == v {
			return i
		}
	}
	return -1
}

func LastIndexWhere[S ~[]T, T any](l S, pref func(T) bool) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := l[i]
		if pref(x) {
			return i
		}
	}
	return -1
}

// use P version when T is a large struct, to improve performance
func LastIndexWhereP[S ~[]T, T any](l S, pref func(*T) bool) int {
	for i := len(l) - 1; i >= 0; i-- {
		x := &l[i]
		if pref(x) {
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
