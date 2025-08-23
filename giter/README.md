# GITER - Go 1.23 Iterator LINQ

A LINQ-style operation library based on the newly introduced `iter.Seq[T]` in Go 1.23.

## Requirements

- Go 1.23 or higher


## Features

### Basic Operations
- `Where(predicate)` — Filter elements
- `Select(selector)` — Projection transformation
- `SelectMany(selector)` — Flattening projection
- `Take(count)` / `TakeWhile(predicate)` — Take first N/elements satisfying condition
- `Skip(count)` / `SkipWhile(predicate)` — Skip first N/elements satisfying condition
- `Concat()` — Concatenate sequences
- `Chunk(size)` — Split into chunks
- `Append()` / `Prepend()` — Append/Prepend elements
- `DefaultIfEmpty()` — Provide default if sequence is empty

### Aggregation & Search
- `Aggregate()` — Aggregate values
- `Count()` / `CountWhere(predicate)` — Count elements
- `Any(predicate)` / `All(predicate)` — Existence check
- `First()` / `FirstOrDefault()` — Get first element
- `FirstWhere()` / `FirstWhereOrDefault()` — Get first matching element
- `Last()` / `LastOrDefault()` — Get last element
- `LastWhere()` / `LastWhereOrDefault()` — Get last matching element
- `Single()` / `SingleOrDefault()` — Get unique element
- `ElementAt()` / `ElementAtOrDefault()` — Get element at index
- `Contains(value)` — Contains check
- `SequenceEqual()` — Sequence equality

### Math Operations (math.go)
- `Min()` / `Max()` — Minimum/Maximum values
- `MinBy(selector)` / `MaxBy(selector)` — Minimum/Maximum by condition
- `Sum()` / `Average()` — Sum/Average

### Set Operations
- `Distinct()` / `DistinctBy(keySelector)` — Remove duplicates
- `Union()` / `Intersect()` / `Except()` — Set operations

### Grouping and Sorting
- `GroupBy(keySelector)` — Group by key (copies all items)
- `GroupByWithValue(keySelector, valueSelector)` — Group and transform values
- `OrderBy(keySelector, less)` / `OrderByDescending()` / `ThenBy()` — Sorting (see order.go)

### Joining and Generation
- `Zip(other, resultSelector)` — Zip operation
- `Range(start, count)` — Generate number sequence
- `Repeat(value, count)` — Repeat value sequence
- `Empty()` — Empty sequence

### Type & Utility
- `OfType[T, R]()` — Filter by type (using type assertion)
- `Cast[T, R]()` — Cast elements to another type
- `ToSlice()` — Convert to slice
- `TryEach()` — Safe iteration (error handling)
- `ForEach()` — Simple iteration

> Note: `Reverse()` is commented out due to performance concerns (see code for details).

## Usage Examples

```go
package main

import (
    "fmt"
    "slices"
    "github.com/szmcdull/glinq/giter"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Chain operations: filter even numbers, square them, take first 3
    result := giter.Take(
        giter.Select(
            giter.Where(slices.Values(numbers), func(x int) bool { return x%2 == 0 }),
            func(x int) int { return x * x },
        ),
        3,
    )
    
    fmt.Println(giter.ToSlice(result)) // [4, 16, 36]
    
    // Grouping example
    words := []string{"apple", "apricot", "banana", "cherry"}
    groups := giter.GroupBy(slices.Values(words), func(s string) byte { return s[0] })
    
    for group := range groups {
        fmt.Printf("%c: %v\n", group.Key, giter.ToSlice(group.Values))
    }
    // a: [apple apricot]
    // b: [banana]  
    // c: [cherry]
}
```


## Differences from .NET LINQ

- `Reverse()` is not implemented due to performance concerns.
- Sorting (`OrderBy`, `OrderByDescending`, `ThenBy`) is implemented in `order.go` and may differ in usage.
- Some methods (e.g., `GroupBy`) require full traversal and buffer all input.

## Notes

1. Most operations are lazily evaluated (streaming), but some (like `GroupBy`, `OrderBy`, etc.) require full traversal and buffer all input.
2. Use `TryEach` to safely handle iterators that may panic.
3. `Chunk` operation uses `iter.Pull` for efficient chunking.
