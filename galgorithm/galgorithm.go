package galgorithm

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type (
	Number interface {
		constraints.Integer | constraints.Float
	}

	// IntFloat[T Number] struct {
	// 	// embedded field type cannot be a (pointer to a) type parameter
	// 	// https://github.com/golang/go/issues/49030
	// 	T
	// }

	INumber interface {
		Add(other INumber) INumber
		Sub(other INumber) INumber
		Mul(other INumber) INumber
		Div(other INumber) INumber
		Comp(other INumber) int
		Eq(other INumber) bool
	}
)

var (
	ErrNoValue = errors.New(`No values provided`)
)

// func (me IntFloat[T]) Add(other IntFloat[T]) IntFloat[T] {
// 	return me.T + other.T
// }

func Abs[T Number](n T) T {
	if n < 0 {
		n = -n
	}
	return n
}

func Clamp[T constraints.Ordered](n, min, max T) T {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

func Min[T constraints.Ordered](values ...T) T {
	l := len(values)
	if l <= 0 {
		panic(ErrNoValue)
	}
	result := values[0]
	for i := 1; i < l; i++ {
		x := values[i]
		if x < result {
			result = x
		}
	}
	return result
}

func Max[T constraints.Ordered](values ...T) T {
	l := len(values)
	if l <= 0 {
		panic(ErrNoValue)
	}
	result := values[0]
	for i := 1; i < l; i++ {
		x := values[i]
		if x > result {
			result = x
		}
	}
	return result
}
