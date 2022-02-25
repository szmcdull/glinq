package unsafe

import (
	"errors"

	"github.com/szmcdull/glinq/galgorithm"
	"golang.org/x/exp/constraints"
)

type (
	Number interface {
		constraints.Integer | constraints.Float
	}
)

var (
	ErrEmptyEnumerable = errors.New(`IEnumerable is empty`)
)

func Min[T constraints.Ordered](q IEnumerable[T]) (result T) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	if !ok {
		return
	}
	result = iter.Current()
	if ok {
		for ok = iter.MoveNext(); ok; ok = iter.MoveNext() {
			v := iter.Current()
			if v < result {
				result = v
			}
		}
	}
	return result
}

func Max[T constraints.Ordered](q IEnumerable[T]) (result T) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	if !ok {
		return
	}
	result = iter.Current()
	if ok {
		for ok = iter.MoveNext(); ok; ok = iter.MoveNext() {
			v := iter.Current()
			if v > result {
				result = v
			}
		}
	}

	return result
}

func MinBy[T any, K constraints.Ordered](q IEnumerable[T], selector func(T) K) (result T) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	if !ok {
		return
	}
	if ok {
		result = iter.Current()
		k := selector(result)
		for ok = iter.MoveNext(); ok; ok = iter.MoveNext() {
			v := iter.Current()
			k2 := selector(v)
			if k2 < k {
				result = v
			}
		}
	}
	return result
}

func MaxBy[T any, K constraints.Ordered](q IEnumerable[T], selector func(T) K) (result T) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	if !ok {
		return
	}
	if ok {
		result = iter.Current()
		k := selector(result)
		for ok = iter.MoveNext(); ok; ok = iter.MoveNext() {
			v := iter.Current()
			k2 := selector(v)
			if k2 > k {
				result = v
			}
		}
	}
	return result
}

func Average[T Number](q IEnumerable[T]) (result T) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	if !ok {
		return
	}
	result = iter.Current()
	count := 1
	if ok {
		for ok = iter.MoveNext(); ok; ok = iter.MoveNext() {
			v := iter.Current()
			result += v
			count++
		}
	}
	return result / T(count)
}

func Sum[T Number](q IEnumerable[T]) (result T) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	if !ok {
		return
	}
	result = iter.Current()
	if ok {
		for ok = iter.MoveNext(); ok; ok = iter.MoveNext() {
			v := iter.Current()
			result += v
		}
	}
	return result
}

func Clamp[T constraints.Ordered](q IEnumerable[T], min, max T) IEnumerable[T] {
	return Select(q, func(t T) T { return galgorithm.Clamp(t, min, max) })
}

func Abs[T Number](q IEnumerable[T]) IEnumerable[T] {
	return Select(q, func(t T) T { return galgorithm.Abs(t) })
}
