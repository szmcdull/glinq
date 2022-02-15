package unsafe

func First[T any](source IEnumerable[T], ok *bool) T {
	iter := source.GetEnumerator()
	e := iter.MoveNext()
	var v T
	if e {
		v = iter.Current()
	}
	if ok != nil {
		*ok = e
	}
	return v
}

func FirstWhere[T any](source IEnumerable[T], pred func(T) bool, found *bool) T {
	iter := source.GetEnumerator()
	e := iter.MoveNext()
	for ; e; e = iter.MoveNext() {
		x := iter.Current()
		if pred(x) {
			return x
		}
	}
	if found != nil {
		*found = e
	}
	var result T
	return result
}
