package unsafe

import "io"

func Last[T any](source IEnumerable[T], found *bool) T {
	return LastWhere(source, func(T) bool { return true }, found)
}

func LastWhere[T any](source IEnumerable[T], pred func(T) bool, retFound *bool) T {
	if getter := source.(IRangeEnumerator[T]); getter != nil {
		v := getter.GetAt(getter.Count() - 1)
		if retFound != nil {
			*retFound = true
		}
		return v
	}

	var result T
	found := false
	Foreach(source, func(x T) error {
		if pred(x) {
			result = x
			found = true
			return io.EOF
		}
		return nil
	})
	if retFound != nil {
		*retFound = found
	}
	return result
}
