package galgorithm

import "constraints"

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
