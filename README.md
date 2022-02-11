# glinq
Go port of DotNet LINQ using generics introduced in Go 1.18

Currently implemented
- Where
- Select
- Range
- Min(By)
- Max(By)
- Average
- First
- Last
- Skip
- Take

More to come...

# usage example

```go
import (
	"fmt"
	"testing"

    . "github.com/szmcdull/glinq"
)

func TestWhereSelect(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	q := Where(sl, func(x int) bool { return x%2 == 0 })
	q2 := Select(q, func(x int) string { return fmt.Sprintf(`%v`, x) })
	sl2, err := ToSlice(q2)
	if err != nil || len(sl2) != 3 {
		t.Errorf(`err=%v len=%d expected 3`, err, len(sl2))
		return
	}
	for i, v := range sl2 {
		if fmt.Sprintf(`%d`, i*2) != v {
			t.Errorf(`%d != %s`, i*2, v)
		}
	}
}

func TestForeach(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	ForeachI(sl, func(i, x int) bool {
		if i != x {
			t.Errorf(`%d expected %d`, x, i)
		}
		return true
	})
}
```

**See [glinq_test.go](https://github.com/szmcdull/glinq/blob/main/glinq_test.go) for more examples**