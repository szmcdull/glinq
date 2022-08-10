package gmap

// ToSlice iterate through a map and generate a slice of Pair[K,V]
func ToSlice[K comparable, V any](m map[K]V) []Pair[K, V] {
	result := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, Pair[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return result
}

// Map iterate through a map and generate a R[]
func Map[K comparable, V any, R any](m map[K]V, f func(K, V) R) []R {
	result := make([]R, 0, len(m))
	for k, v := range m {
		result = append(result, f(k, v))
	}
	return result
}
