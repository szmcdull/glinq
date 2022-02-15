package unsafe

import (
	"errors"
)

type (
	IEnumerator[T any] interface {
		Current() T
		MoveNext() bool
	}
	IEnumerable[T any] interface {
		GetEnumerator() IEnumerator[T]
		Count() int
		Any() bool
	}
	IRangeEnumerator[T any] interface {
		ICount[T]
		IGetAt[T]
		SeekOnce(i int) error // can only seek once
	}
	ICount[T any] interface {
		Count() int
	}
	IGetAt[T any] interface {
		GetAt(i int) T
	}
)

var (
	ErrInvalidState = errors.New(`Invalid state`)
)

func ToSlice[T any](me IEnumerable[T]) []T {
	i := me.GetEnumerator()
	var result []T
	ok := i.MoveNext()
	for ; ok; ok = i.MoveNext() {
		result = append(result, i.Current())
	}
	return result
}

func ToMap[T any, K comparable, V any](
	me IEnumerable[T], keyFunc func(T) K, valFunc func(T) V) map[K]V {

	i := me.GetEnumerator()
	result := make(map[K]V)
	ok := i.MoveNext()
	for ; ok; ok = i.MoveNext() {
		p := i.Current()
		k := keyFunc(p)
		v := valFunc(p)
		result[k] = v
	}
	return result
}

// Apply f() to all items in q
func Do[T any](q IEnumerable[T], f func(T)) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	for ; ok; ok = iter.MoveNext() {
		f(iter.Current())
	}
}

// Apply f() to all items in q, break when f() returns false
func Foreach[T any](q IEnumerable[T], f func(T) error) (err error) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	for ; ok; ok = iter.MoveNext() {
		if err = f(iter.Current()); err != nil {
			break
		}
	}
	return err
}

// Apply f() to all items in q, break when f() returns false. A counted index is passed to f()
func ForeachI[T any](q IEnumerable[T], f func(int, T) error) (err error) {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	i := 0
	for ; ok; ok = iter.MoveNext() {
		if err = f(i, iter.Current()); err != nil {
			break
		}
		i++
	}
	return err
}

func Contains[T comparable](q IEnumerable[T], v T) bool {
	iter := q.GetEnumerator()
	ok := iter.MoveNext()
	for ; ok; ok = iter.MoveNext() {
		if iter.Current() == v {
			return true
		}
	}
	return false
}
