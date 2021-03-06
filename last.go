package glinq

import (
	"io"
)

func Last[T any](source IEnumerable[T], err *error) T {
	return LastWhere(source, func(T) bool { return true }, err)
}

func LastWhere[T any](source IEnumerable[T], pred func(T) bool, err *error) T {
	if getter := source.(IRangeEnumerator[T]); getter != nil {
		v, e := getter.GetAt(getter.Count() - 1)
		if e != nil && err != nil {
			*err = e
		}
		return v
	}

	var result T
	found := false
	e := Foreach(source, func(x T) error {
		if pred(x) {
			result = x
			found = true
			return io.EOF
		}
		return nil
	})
	if e == nil {
		if !found {
			e = io.EOF
		}
	}
	if e != nil {
		if err != nil {
			*err = e
		}
	}
	return result
}
