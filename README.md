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
- Skip(While)
- Take(While)
- Contains
- Any
- Join

(`OrderBy` And `GroupBy` will always access all items. For performance and ease of use reason I'll not implement them in IEnumerable way. Please use `glinq/garray` package instead)

Adapters:
- FromSlice
- FromMap
- ReadLines	(wrapping bufio.NewScanner)
- sqlxq.Queryx (wrapping sqlx.Queryx)

- ToSlice
- ToMap

Also some similar utilities directly for slices in the `glinq/garray` package. These are more handful, without FromSlice/ToSlice conversion.
- Sort(Descending)
- SortBy(Descending)
- Map
- ToMap
- Filter
- Apply
- RemoveIf
- Contains
- (Last)IndexOf
- (Last)IndexWhere
- First, Last
- Concat
- ShallowCopy
- Sum
- Average
- GroupBy
- Count
- Reverse

Also a generic SyncMap type (should I separate it in another repository?)
- Load
- Store
- Range(E)
- Delete(If)
- LoadAndDelete
- LoadOrStore
- LoadOrNew
- LoadAndUpdate
- ToSlice
- Keys
- Len
- Clear
- PopAll

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

**More examples:**
- [garray doc](https://pkg.go.dev/github.com/szmcdull/glinq/garray)
- [glinq_test.go](https://github.com/szmcdull/glinq/blob/main/unsafe/glinq_test.go)
- [garray_test.go](https://github.com/szmcdull/glinq/blob/main/garray/garray_test.go)
