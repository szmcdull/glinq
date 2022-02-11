package glinq

func First[T any](source IEnumerable[T], err *error) T {
	if getter, ok := source.(IGetAt[T]); ok {
		v, e := getter.GetAt(0)
		if e != nil && err != nil {
			*err = e
		}
		return v
	}

	iter := source.GetEnumerator()
	e := iter.MoveNext()
	var v T
	if e != nil {
		v = iter.Current()
		if err != nil {
			*err = e
		}
	}
	return v
}

func FirstWhere[T any](source IEnumerable[T], pred func(T) bool, err *error) T {
	iter := source.GetEnumerator()
	e := iter.MoveNext()
	for ; e == nil; e = iter.MoveNext() {
		x := iter.Current()
		if pred(x) {
			return x
		}
	}
	if err != nil {
		*err = e
	}
	var result T
	return result
}
