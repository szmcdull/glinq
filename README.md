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
}
```

**See [glinq_test.go](https://github.com/szmcdull/glinq/blob/main/unsafe/glinq_test.go) for more examples**