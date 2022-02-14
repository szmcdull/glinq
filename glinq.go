package glinq

import (
	"errors"
	"io"
)

type (
	IEnumerator[T any] interface {
		Current() T
		MoveNext() error // 返回nil=成功, io.EOF=已到最后一个, 其他=错误
		//Reset() error
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
		GetAt(i int) (T, error)
	}
)

var (
	ErrInvalidState = errors.New(`Invalid state`)
)

func ToSlice[T any](me IEnumerable[T]) ([]T, error) {
	i := me.GetEnumerator()
	var result []T
	err := i.MoveNext()
	for ; err == nil; err = i.MoveNext() {
		result = append(result, i.Current())
	}
	if err == io.EOF {
		return result, nil
	}
	return result, err
}

func ToMap[T any, K comparable, V any](
	me IEnumerable[T], keyFunc func(T) K, valFunc func(T) V) (map[K]V, error) {

	i := me.GetEnumerator()
	result := make(map[K]V)
	err := i.MoveNext()
	for ; err == nil; err = i.MoveNext() {
		p := i.Current()
		k := keyFunc(p)
		v := valFunc(p)
		result[k] = v
	}
	if err == io.EOF {
		return result, nil
	}
	return result, err
}

// Apply f() to all items in q
func Do[T any](q IEnumerable[T], f func(T)) error {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	for ; err == nil; err = iter.MoveNext() {
		f(iter.Current())
	}
	if err == io.EOF {
		return nil
	}
	return err
}

// Apply f() to all items in q, break when f() returns false
func Foreach[T any](q IEnumerable[T], f func(T) bool) error {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	for ; err == nil; err = iter.MoveNext() {
		if !f(iter.Current()) {
			break
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

// Apply f() to all items in q, break when f() returns false. A counted index is passed to f()
func ForeachI[T any](q IEnumerable[T], f func(int, T) bool) error {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	i := 0
	for ; err == nil; err = iter.MoveNext() {
		if !f(i, iter.Current()) {
			break
		}
		i++
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func Contains[T comparable](q IEnumerable[T], v T) (bool, error) {
	iter := q.GetEnumerator()
	err := iter.MoveNext()
	for ; err == nil; err = iter.MoveNext() {
		if iter.Current() == v {
			return true, nil
		}
	}
	if err == io.EOF {
		err = nil
	}
	return false, err
}
