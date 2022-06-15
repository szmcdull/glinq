# glinq
Go port of DotNet LINQ using generics introduced in Go 1.18

Currently implemented
- Where
- Select
- Range
- Foreach
- Min(By)
- Max(By)
- Average
- Sum
- First
- Last
- Skip
- Take
- Contains
- Any
- Join

Also some similar utilities directly for slices in the glinq/garray package. These are more handful, without FromSlice/ToSlice conversion.
- Sort(Descending)
- SortBy(Descending)
- Map
- Apply
- (Last)IndexOf
- (Last)IndexWhere
- First, Last
- Concat
- ShallowCopy
- Sum
- Average

And more to come...


# usage example

```go
package main

import (
	"fmt"

	"github.com/szmcdull/glinq/garray"
	. "github.com/szmcdull/glinq/unsafe"
)

func main() {
	numbers := garray.Concat([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10})
	numbersToRemove := Where(FromSlice(numbers), func(x int) bool {
		return x > 5
	})
	Do(numbersToRemove, func(x int) { fmt.Printf(`%d `, x) }) // 6 7 8 9 10
	fmt.Println(``)
	numbers = ToSlice(Where(FromSlice(numbers), func(x int) bool {
		return !Contains(numbersToRemove, x)
	}))
	fmt.Println(Any(Where(FromSlice(numbers), func(x int) bool { return x > 5 }))) // false
	Do(FromSlice(numbers), func(x int) { fmt.Printf(`%d `, x) })                   // 1 2 3 4 5
	fmt.Println(``)

	// Sorting
	l := []int{8, 6, 4, 2, 5, 3, 1}

	garray.Sort(l)
	fmt.Printf("%v\n", l) // [1 2 3 4 5 6 8]

	l2 := []string{`the`, `lazy`, `dog`, `jumps`, `over`, `the`, `silver`, `fox`}
	garray.SortByDescending(l2, func(x string) byte { return x[len(x)-1] }) // sort descending by the last character
	fmt.Printf("%v\n", l2)                                                  // [lazy fox jumps silver over dog the the]
}
```

**See [glinq_test.go](https://github.com/szmcdull/glinq/blob/main/unsafe/glinq_test.go) for more examples**