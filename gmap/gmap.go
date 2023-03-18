package gmap

import (
	"github.com/szmcdull/glinq/garray"
	"golang.org/x/exp/constraints"
)

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

// Keys returns the keys of a map as a slice
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	garray.Sort(keys)
	return keys
}

// Map iterate through a map and generate a R[]
func Map[K comparable, V any, R any](m map[K]V, f func(K, V) R) []R {
	result := make([]R, 0, len(m))
	for k, v := range m {
		result = append(result, f(k, v))
	}
	return result
}

func ShallowCopy[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
