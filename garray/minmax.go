package garray

func Max[S ~[]T, T Number](s S) T {
	length := len(s)
	if length == 0 {
		panic(`Max: empty slice`)
	}

	max := s[0]
	for i := 1; i < length; i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	return max
}

func Min[S ~[]T, T Number](s S) T {
	length := len(s)
	if length == 0 {
		panic(`Min: empty slice`)
	}

	min := s[0]
	for i := 1; i < length; i++ {
		if s[i] < min {
			min = s[i]
		}
	}

	return min
}

func MaxBy[S ~[]T, T any, N Number](s S, selector func(T) N) N {
	length := len(s)
	if length == 0 {
		panic(`MaxBy: empty slice`)
	}

	max := selector(s[0])
	for i := 1; i < length; i++ {
		current := selector(s[i])
		if current > max {
			max = current
		}
	}

	return max
}

func MinBy[S ~[]T, T any, N Number](s S, selector func(T) N) N {
	length := len(s)
	if length == 0 {
		panic(`MinBy: empty slice`)
	}

	min := selector(s[0])
	for i := 1; i < length; i++ {
		current := selector(s[i])
		if current < min {
			min = current
		}
	}

	return min
}
