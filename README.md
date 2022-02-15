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

Also some similar utilities for slices in the glinq/garray package.

And more to come...


# usage example

```go
import (
	"fmt"
	"sort"

	"github.com/szmcdull/glinq/garray"
	. "github.com/szmcdull/glinq/unsafe"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersToRemove := Where(FromSlice(numbers), func(x int) bool {
		return x > 5
	})
	Do(numbersToRemove, func(x int) { fmt.Printf(`%d `, x) }) // 6 7 8 9 10
	fmt.Println(``)
	numbers = ToSlice(Where(FromSlice(numbers), func(x int) bool {
		return !Contains(numbersToRemove, x)
	}))
	Do(FromSlice(numbers), func(x int) { fmt.Printf(`%d `, x) }) // 1 2 4 5 6
	fmt.Println(``)

	// Sorting
	l := []int{8, 6, 4, 2, 5, 3, 1}

	sort.Sort(garray.Sortable(l))
	fmt.Printf("%v\n", l) // [1 2 3 4 5 6 8]

	sort.Sort(garray.SortableDescending(l))
	fmt.Printf("%v\n", l) // [8 6 5 4 3 2 1]
}
```

**See [glinq_test.go](https://github.com/szmcdull/glinq/blob/main/unsafe/glinq_test.go) for more examples**