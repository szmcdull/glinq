
# Style

There are two style of set operations: in-place and chain.

## In-place operation:

A in-place operation is applied to the first argument.

```go
	A := make(map[int]struct{})
	B := map[int]struct{}{1: {}, 2: {}}
	gset.Add(A, B)
	fmt.Println(A) // Output: map[1:{} 2:{}]
```

## Chain operation:

A chain operation returns a new set.

```go
	// Initialize HashSets
	set1 := gset.HashSet[string]{"apple": {}, "banana": {}}
	set2 := gset.HashSet[string]{"pear": {}, "orange": {}}
	set3 := gset.HashSet[string]{"banana": {}, "durian": {}}

	// calculate the union of them
	result := set1.
		Add(set2).
		Add(set3)

	// Print the contents of the updated set
	for _, k := range gset.Sorted(result) {
		fmt.Println(k)
	}

	// Output:
	// apple
	// banana
	// durian
	// orange
	// pear
```

### Documents and examples: https://pkg.go.dev/github.com/szmcdull/glinq/gset
