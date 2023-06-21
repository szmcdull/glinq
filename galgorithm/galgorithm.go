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

// Abs returns the absolute value of n.
func Abs[T Number](n T) T {
	if n < 0 {
		n = -n
	}
	return n
}

// Clamp returns n if it is in [min, max], otherwise min or max.
func Clamp[T constraints.Ordered](n, min, max T) T {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}

// Min returns the minimum value of the given values.
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

// Max returns the maximum value in values.
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

// InRange returns true if v is in [min, max).
func InRange[T constraints.Ordered](v, min, max T) bool {
	return v >= min && v < max
}
