package garray

// group a slice of T by a key, resulting groups of slice of R
func GroupBy[T any, Slice ~[]T, R any, Key comparable](a Slice, keyF func(v T) Key, resultF func(T) R) map[Key][]R {
	m := make(map[Key][]R)
	for _, v := range a {
		key := keyF(v)
		a := m[key]
		a = append(a, resultF(v))
		m[key] = a
	}
	return m
}

// group a slice of T by a key, resulting groups of slice of R
func GroupByP[T any, Slice ~[]T, R any, Key comparable](a Slice, keyF func(v *T) Key, resultF func(*T) R) map[Key][]R {
	m := make(map[Key][]R)
	for i := range a {
		v := &a[i]
		key := keyF(v)
		a := m[key]
		a = append(a, resultF(v))
		m[key] = a
	}
	return m
}

// todo: incomparable group by, using a customized comparer
