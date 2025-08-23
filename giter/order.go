//go:build go1.23
// +build go1.23

package giter

import (
	"iter"
	"sort"
)

// OrderBy sorts the sequence using a single less function, similar to sort.SliceStable.
// Example:
//
//	OrderBy(seq, func(a, b Person) bool {
//	    if a.Age != b.Age { return a.Age < b.Age }
//	    return a.Name < b.Name
//	})
func OrderBy[T any](it iter.Seq[T], less func(a, b T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		items := ToSlice(it)
		sort.SliceStable(items, func(i, j int) bool {
			return less(items[i], items[j])
		})
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

// OrderSpec represents a multi-level ordering specification that can be built
// using a fluent API for complex sorting scenarios.
//
// Example usage:
//
//	Order(seq, By(keyFunc, lessFunc).Then(keyFunc2, lessFunc2).ThenDescending(keyFunc3, lessFunc3))
//
// Or step by step:
//
//	spec := By(keyFunc, lessFunc)
//	spec = Then(spec, keyFunc2, lessFunc2)
//	spec = ThenDescending(spec, keyFunc3, lessFunc3)
//	result := Order(seq, spec)
type OrderSpec[T any] struct {
	comparers []func(T, T) int
}

// By creates a new OrderSpec with the first ordering criteria in ascending order.
// This is the starting point for building multi-level sort specifications.
//
// Example:
//
//	spec := By(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b })
func By[T any, K any](keySelector func(T) K, less func(K, K) bool) OrderSpec[T] {
	comparer := func(a, b T) int {
		ka, kb := keySelector(a), keySelector(b)
		if less(ka, kb) {
			return -1
		}
		if less(kb, ka) {
			return 1
		}
		return 0
	}

	return OrderSpec[T]{
		comparers: []func(T, T) int{comparer},
	}
}

// ByDescending creates a new OrderSpec with the first ordering criteria in descending order.
//
// Example:
//
//	spec := ByDescending(func(p Person) int { return p.Age }, func(a, b int) bool { return a < b })
func ByDescending[T any, K any](keySelector func(T) K, less func(K, K) bool) OrderSpec[T] {
	comparer := func(a, b T) int {
		ka, kb := keySelector(a), keySelector(b)
		if less(ka, kb) {
			return 1 // Reversed for descending order
		}
		if less(kb, ka) {
			return -1 // Reversed for descending order
		}
		return 0
	}

	return OrderSpec[T]{
		comparers: []func(T, T) int{comparer},
	}
}

// Then adds a secondary ordering criteria in ascending order.
// This allows chaining multiple sort criteria where subsequent criteria
// are used to break ties from previous criteria.
//
// Example:
//
//	spec = Then(spec, func(p Person) string { return p.Name }, func(a, b string) bool { return a < b })
func Then[T any, K any](spec OrderSpec[T], keySelector func(T) K, less func(K, K) bool) OrderSpec[T] {
	comparer := func(a, b T) int {
		ka, kb := keySelector(a), keySelector(b)
		if less(ka, kb) {
			return -1
		}
		if less(kb, ka) {
			return 1
		}
		return 0
	}

	return OrderSpec[T]{
		comparers: append(spec.comparers, comparer),
	}
}

// ThenDescending adds a secondary ordering criteria in descending order.
//
// Example:
//
//	spec = ThenDescending(spec, func(p Person) int { return p.Score }, func(a, b int) bool { return a < b })
func ThenDescending[T any, K any](spec OrderSpec[T], keySelector func(T) K, less func(K, K) bool) OrderSpec[T] {
	comparer := func(a, b T) int {
		ka, kb := keySelector(a), keySelector(b)
		if less(ka, kb) {
			return 1 // Reversed for descending order
		}
		if less(kb, ka) {
			return -1 // Reversed for descending order
		}
		return 0
	}

	return OrderSpec[T]{
		comparers: append(spec.comparers, comparer),
	}
}

// Order sorts the sequence using the specified multi-level ordering criteria.
// The OrderSpec parameter should be built using By() and Then() functions.
//
// This function materializes the entire sequence into a slice, sorts it using
// Go's stable sort, and then yields the sorted elements.
//
// Example:
//
//	sorted := Order(seq, By(ageSelector, intLess).Then(nameSelector, stringLess))
//
// Multi-level example:
//
//	sorted := Order(seq,
//	  Then(
//	    ThenDescending(
//	      By(gradeSelector, intLess),           // Primary: grade ascending
//	      scoreSelector, intLess),              // Secondary: score descending
//	    nameSelector, stringLess))              // Tertiary: name ascending
func Order[T any](it iter.Seq[T], spec OrderSpec[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		items := ToSlice(it)

		sort.SliceStable(items, func(i, j int) bool {
			a, b := items[i], items[j]

			// Apply all comparers in order until we find a difference
			for _, comparer := range spec.comparers {
				result := comparer(a, b)
				if result < 0 {
					return true
				}
				if result > 0 {
					return false
				}
				// Equal according to this comparer, try next one
			}

			return false // All comparers consider them equal
		})

		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}
