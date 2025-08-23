//go:build go1.23
// +build go1.23

package giter

import (
	"fmt"
	"iter"
)

func TryEach[T any](it iter.Seq[T], yield func(T) bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = fmt.Errorf(`SafeRange: %w`, v)
			case string:
				err = fmt.Errorf(`SafeRange: %s`, v)
			default:
				err = fmt.Errorf(`SafeRange: %v`, r)
			}
		}
	}()
	for v := range it {
		if !yield(v) {
			break
		}
	}
	return
}

func ForEach[T any](it iter.Seq[T], yield func(T) bool) {
	for v := range it {
		if !yield(v) {
			break
		}
	}
}

// go 1.23 slices equivalent 	=> LINQ
//  slices.Backward				=> Reverse
//  slices.Collect				=> ToArray
//  slices.Sorted				=> Order
//  slices.SortedFunc			=> OrderBy
//  slices.Chunk				=> Chunk

// Additional implementations

func Aggregate[T any, R any](it iter.Seq[T], initial R, accumulator func(R, T) R) R {
	result := initial
	for v := range it {
		result = accumulator(result, v)
	}
	return result
}

// All checks if all elements in the sequence satisfy the predicate.
func All[T any](it iter.Seq[T], predicate func(T) bool) bool {
	for v := range it {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func Any[T any](it iter.Seq[T], predicate func(T) bool) bool {
	for v := range it {
		if predicate(v) {
			return true
		}
	}
	return false
}

func Append[T any](it iter.Seq[T], value T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				return
			}
		}
		yield(value)
	}
}

func Concat[T any](its ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, it := range its {
			for v := range it {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// There is no way to chunk a push iter. A poor way is to create a goproc to redirect the iter to push to a chan, then read the chan chunk by chunk in current goproc.
// iter.Pull makes it better by introducing a new lightweight coroutine (a stack-less state machine without preemption etc.)
func Chunk[T any](source iter.Seq[T], chunkSize int) iter.Seq[iter.Seq[T]] {
	if chunkSize <= 0 {
		panic("chunkSize must be positive")
	}

	return func(yield func(iter.Seq[T]) bool) {
		next, stop := iter.Pull(source)
		defer stop()

		for {
			// Pre-check if there are more elements (avoid creating empty iterators)
			v, ok := next()
			if !ok {
				break // Source sequence has ended
			}

			// Create current chunk iterator
			chunkIter := func(chunkYield func(T) bool) {
				// First yield the pre-read first element
				if !chunkYield(v) {
					return
				}
				// Then yield remaining elements
				for i := 1; i < chunkSize; i++ {
					v, ok := next()
					if !ok {
						return // Source sequence has ended
					}
					if !chunkYield(v) {
						return
					}
				}
			}

			if !yield(chunkIter) {
				return
			}
		}
	}
}

func Distinct[T comparable](it iter.Seq[T]) iter.Seq[T] {
	seen := make(map[T]struct{})
	return func(yield func(T) bool) {
		for v := range it {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}

// DistinctBy returns distinct elements based on a key selector function
func DistinctBy[T any, K comparable](it iter.Seq[T], keySelector func(T) K) iter.Seq[T] {
	seen := make(map[K]struct{})
	return func(yield func(T) bool) {
		for v := range it {
			key := keySelector(v)
			if _, exists := seen[key]; !exists {
				seen[key] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Where filters elements based on a predicate
func Where[T any](it iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Select projects each element to a new form
func Select[T, R any](it iter.Seq[T], selector func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			if !yield(selector(v)) {
				return
			}
		}
	}
}

// SelectMany flattens sequences
func SelectMany[T, R any](it iter.Seq[T], selector func(T) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			for r := range selector(v) {
				if !yield(r) {
					return
				}
			}
		}
	}
}

// Take returns the first n elements
func Take[T any](it iter.Seq[T], count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range it {
			if i >= count {
				return
			}
			if !yield(v) {
				return
			}
			i++
		}
	}
}

// TakeWhile returns elements while the predicate is true
func TakeWhile[T any](it iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !predicate(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Skip skips the first n elements
func Skip[T any](it iter.Seq[T], count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range it {
			if i >= count {
				if !yield(v) {
					return
				}
			}
			i++
		}
	}
}

// SkipWhile skips elements while the predicate is true
func SkipWhile[T any](it iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		skipping := true
		for v := range it {
			if skipping && predicate(v) {
				continue
			}
			skipping = false
			if !yield(v) {
				return
			}
		}
	}
}

// First returns the first element or panics if empty
func First[T any](it iter.Seq[T]) T {
	for v := range it {
		return v
	}
	panic("First called on empty iterator")
}

// FirstOrDefault returns the first element or zero value if empty
func FirstOrDefault[T any](it iter.Seq[T]) T {
	for v := range it {
		return v
	}
	var zero T
	return zero
}

// FirstWhere returns the first element that satisfies the predicate
func FirstWhere[T any](it iter.Seq[T], predicate func(T) bool) T {
	for v := range it {
		if predicate(v) {
			return v
		}
	}
	panic("FirstWhere: no element satisfies the predicate")
}

// FirstWhereOrDefault returns the first element that satisfies the predicate or zero value
func FirstWhereOrDefault[T any](it iter.Seq[T], predicate func(T) bool) T {
	for v := range it {
		if predicate(v) {
			return v
		}
	}
	var zero T
	return zero
}

// Last returns the last element or panics if empty
func Last[T any](it iter.Seq[T]) T {
	var last T
	found := false
	for v := range it {
		last = v
		found = true
	}
	if !found {
		panic("Last called on empty iterator")
	}
	return last
}

// LastOrDefault returns the last element or zero value if empty
func LastOrDefault[T any](it iter.Seq[T]) T {
	var last T
	for v := range it {
		last = v
	}
	return last
}

// LastWhere returns the last element that satisfies the predicate
func LastWhere[T any](it iter.Seq[T], predicate func(T) bool) T {
	var last T
	found := false
	for v := range it {
		if predicate(v) {
			last = v
			found = true
		}
	}
	if !found {
		panic("LastWhere: no element satisfies the predicate")
	}
	return last
}

// LastWhereOrDefault returns the last element that satisfies the predicate or zero value
func LastWhereOrDefault[T any](it iter.Seq[T], predicate func(T) bool) T {
	var last T
	for v := range it {
		if predicate(v) {
			last = v
		}
	}
	return last
}

// Single returns the only element or panics
func Single[T any](it iter.Seq[T]) T {
	var result T
	found := false
	for v := range it {
		if found {
			panic("Single: sequence contains more than one element")
		}
		result = v
		found = true
	}
	if !found {
		panic("Single: sequence is empty")
	}
	return result
}

// SingleOrDefault returns the only element or zero value
func SingleOrDefault[T any](it iter.Seq[T]) T {
	var result T
	found := false
	for v := range it {
		if found {
			panic("SingleOrDefault: sequence contains more than one element")
		}
		result = v
		found = true
	}
	return result
}

// Count returns the number of elements
func Count[T any](it iter.Seq[T]) int {
	count := 0
	for range it {
		count++
	}
	return count
}

// CountWhere returns the number of elements that satisfy the predicate
func CountWhere[T any](it iter.Seq[T], predicate func(T) bool) int {
	count := 0
	for v := range it {
		if predicate(v) {
			count++
		}
	}
	return count
}

// Contains checks if the sequence contains the specified value
func Contains[T comparable](it iter.Seq[T], value T) bool {
	for v := range it {
		if v == value {
			return true
		}
	}
	return false
}

// ElementAt returns the element at the specified index
func ElementAt[T any](it iter.Seq[T], index int) T {
	if index < 0 {
		panic("ElementAt: index cannot be negative")
	}
	i := 0
	for v := range it {
		if i == index {
			return v
		}
		i++
	}
	panic("ElementAt: index is out of range")
}

// ElementAtOrDefault returns the element at the specified index or zero value
func ElementAtOrDefault[T any](it iter.Seq[T], index int) T {
	if index < 0 {
		var zero T
		return zero
	}
	i := 0
	for v := range it {
		if i == index {
			return v
		}
		i++
	}
	var zero T
	return zero
}

// ToSlice converts the iterator to a slice
func ToSlice[T any](it iter.Seq[T]) []T {
	var result []T
	for v := range it {
		result = append(result, v)
	}
	return result
}

// Reverse commented out due to performance issues
// // Reverse returns elements in reverse order
// func Reverse[T any](it iter.Seq[T]) iter.Seq[T] {
// 	return func(yield func(T) bool) {
// 		slice := ToSlice(it)
// 		for i := len(slice) - 1; i >= 0; i-- {
// 			if !yield(slice[i]) {
// 				return
// 			}
// 		}
// 	}
// }

// Prepend adds an element to the beginning of the sequence
func Prepend[T any](it iter.Seq[T], value T) iter.Seq[T] {
	return func(yield func(T) bool) {
		if !yield(value) {
			return
		}
		for v := range it {
			if !yield(v) {
				return
			}
		}
	}
}

// DefaultIfEmpty returns the sequence or a sequence with a single default value if empty
func DefaultIfEmpty[T any](it iter.Seq[T], defaultValue T) iter.Seq[T] {
	return func(yield func(T) bool) {
		empty := true
		for v := range it {
			empty = false
			if !yield(v) {
				return
			}
		}
		if empty {
			yield(defaultValue)
		}
	}
}

// Union returns the set union of two sequences (removes duplicates)
func Union[T comparable](first, second iter.Seq[T]) iter.Seq[T] {
	return Distinct(Concat(first, second))
}

// Intersect returns the set intersection of two sequences
func Intersect[T comparable](first, second iter.Seq[T]) iter.Seq[T] {
	secondSet := make(map[T]struct{})
	for v := range second {
		secondSet[v] = struct{}{}
	}

	return Distinct(func(yield func(T) bool) {
		for v := range first {
			if _, exists := secondSet[v]; exists {
				if !yield(v) {
					return
				}
			}
		}
	})
}

// Except returns the set difference of two sequences
func Except[T comparable](first, second iter.Seq[T]) iter.Seq[T] {
	secondSet := make(map[T]struct{})
	for v := range second {
		secondSet[v] = struct{}{}
	}

	return Distinct(func(yield func(T) bool) {
		for v := range first {
			if _, exists := secondSet[v]; !exists {
				if !yield(v) {
					return
				}
			}
		}
	})
}

// GroupBy groups elements by a key selector
type Grouping[K comparable, V any] struct {
	Key    K
	Values iter.Seq[V]
}

// GroupBy groups elements by a key selector. It has to COPY ALL items.
func GroupBy[T any, K comparable](it iter.Seq[T], keySelector func(T) K) iter.Seq[Grouping[K, T]] {
	return func(yield func(Grouping[K, T]) bool) {
		groups := make(map[K][]T)

		// Collect all items into groups
		for item := range it {
			key := keySelector(item)
			groups[key] = append(groups[key], item)
		}

		// Yield each group
		for key, items := range groups {
			group := Grouping[K, T]{
				Key: key,
				Values: func(yield func(T) bool) {
					for _, item := range items {
						if !yield(item) {
							return
						}
					}
				},
			}
			if !yield(group) {
				return
			}
		}
	}
}

// GroupByWithValue groups elements by a key selector and transforms values. It has to COPY ALL items.
func GroupByWithValue[T any, K comparable, V any](it iter.Seq[T], keySelector func(T) K, valueSelector func(T) V) iter.Seq[Grouping[K, V]] {
	return func(yield func(Grouping[K, V]) bool) {
		groups := make(map[K][]V)

		// Collect all items into groups
		for item := range it {
			key := keySelector(item)
			value := valueSelector(item)
			groups[key] = append(groups[key], value)
		}

		// Yield each group
		for key, values := range groups {
			group := Grouping[K, V]{
				Key: key,
				Values: func(yield func(V) bool) {
					for _, value := range values {
						if !yield(value) {
							return
						}
					}
				},
			}
			if !yield(group) {
				return
			}
		}
	}
}

// Zip combines two sequences into one
func Zip[T, U, R any](first iter.Seq[T], second iter.Seq[U], resultSelector func(T, U) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		next1, stop1 := iter.Pull(first)
		defer stop1()
		next2, stop2 := iter.Pull(second)
		defer stop2()

		for {
			v1, ok1 := next1()
			v2, ok2 := next2()

			if !ok1 || !ok2 {
				return
			}

			if !yield(resultSelector(v1, v2)) {
				return
			}
		}
	}
}

// SequenceEqual determines whether two sequences are equal
func SequenceEqual[T comparable](first, second iter.Seq[T]) bool {
	next1, stop1 := iter.Pull(first)
	defer stop1()
	next2, stop2 := iter.Pull(second)
	defer stop2()

	for {
		v1, ok1 := next1()
		v2, ok2 := next2()

		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			return true // Both sequences ended at the same time
		}

		if v1 != v2 {
			return false
		}
	}
}

// Range generates a sequence of integers
func Range(start, count int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < count; i++ {
			if !yield(start + i) {
				return
			}
		}
	}
}

// Repeat generates a sequence with a repeated value
func Repeat[T any](value T, count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

// Empty returns an empty sequence
func Empty[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Do nothing - empty sequence
	}
}

// OfType filters elements to those of a specific type (using type assertion)
// This is a simplified version since Go doesn't have runtime type information like .NET
func OfType[T, R any](it iter.Seq[T]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			if r, ok := any(v).(R); ok {
				if !yield(r) {
					return
				}
			}
		}
	}
}

// Cast attempts to cast each element to the specified type
func Cast[T, R any](it iter.Seq[T]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			if !yield(any(v).(R)) {
				return
			}
		}
	}
}
