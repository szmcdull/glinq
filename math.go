package glinq

import (
	"constraints"
	"io"

	"github.com/szmcdull/glinq/galgorithm"
)

type (
	Number interface {
		constraints.Integer | constraints.Float
	}
)

func Min[T constraints.Ordered](q IEnumerable[T]) (T, error) {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	result := iter.Current()
	if err == nil {
		for err = iter.MoveNext(); err == nil; err = iter.MoveNext() {
			v := iter.Current()
			if v < result {
				result = v
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return result, err
}

func Max[T constraints.Ordered](q IEnumerable[T]) (T, error) {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	result := iter.Current()
	if err == nil {
		for err = iter.MoveNext(); err == nil; err = iter.MoveNext() {
			v := iter.Current()
			if v > result {
				result = v
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return result, err
}

func MinBy[T any, K constraints.Ordered](q IEnumerable[T], selector func(T) K) (result T, err error) {
	iter := q.GetEnumerator()
	err = iter.MoveNext()
	if err == nil {
		result = iter.Current()
		k := selector(result)
		for err = iter.MoveNext(); err == nil; err = iter.MoveNext() {
			v := iter.Current()
			k2 := selector(v)
			if k2 < k {
				result = v
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return result, err
}

func MaxBy[T any, K constraints.Ordered](q IEnumerable[T], selector func(T) K) (result T, err error) {
	iter := q.GetEnumerator()
	err = iter.MoveNext()
	if err == nil {
		result = iter.Current()
		k := selector(result)
		for err = iter.MoveNext(); err == nil; err = iter.MoveNext() {
			v := iter.Current()
			k2 := selector(v)
			if k2 > k {
				result = v
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return result, err
}

func Average[T Number](q IEnumerable[T]) (T, error) {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	result := iter.Current()
	count := 1
	if err == nil {
		for err = iter.MoveNext(); err == nil; err = iter.MoveNext() {
			v := iter.Current()
			result += v
			count++
		}
	}
	if err == io.EOF {
		err = nil
	}
	return result / T(count), err
}

func Sum[T Number](q IEnumerable[T]) (T, error) {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	result := iter.Current()
	if err == nil {
		for err = iter.MoveNext(); err == nil; err = iter.MoveNext() {
			v := iter.Current()
			result += v
		}
	}
	if err == io.EOF {
		err = nil
	}
	return result, err
}

func Clamp[T constraints.Ordered](q IEnumerable[T], min, max T) IEnumerable[T] {
	return Select(q, func(t T) T { return galgorithm.Clamp(t, min, max) })
}

func Abs[T Number](q IEnumerable[T]) IEnumerable[T] {
	return Select(q, func(t T) T { return galgorithm.Abs(t) })
}
