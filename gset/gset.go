package gset

// Experimental. Use at your own risk

func FromSlice[T comparable](l []T) map[T]struct{} {
	result := make(map[T]struct{}, len(l))
	for _, v := range l {
		result[v] = struct{}{}
	}
	return result
}

func ToSlice[T comparable](m map[T]struct{}) []T {
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// Add all in B to A
func Add[T comparable](A, B map[T]struct{}) {
	for k := range B {
		A[k] = struct{}{}
	}
}

// Remove all in B from A
func Sub[T comparable](A, B map[T]struct{}) {
	for k := range A {
		if _, ok := B[k]; ok {
			delete(A, k)
		}
	}
}

// Remove anything not in B from A
func And[T comparable](A, B map[T]struct{}) {
	for k := range A {
		if _, ok := B[k]; !ok {
			delete(A, k)
		}
	}
}

// Shallow copy
func Copy[T comparable](other map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{}, len(other))
	for k := range other {
		result[k] = struct{}{}
	}
	return result
}
