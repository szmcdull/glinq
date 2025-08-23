
# garray - Go Slice Utilities

High-performance, generic utilities for direct operations on Go slices. Provides LINQ-like functionality with zero conversion overhead and maximum performance. All operations are in-place unless otherwise specified.

## Features

- **Mapping**:  
    - `Map`, `MapE` (with error), `MapI` (with index), `MapIE` (index+error)
- **Filtering**:  
    - `Filter`, `FilterI` (with index), `FilterIE` (index+error)
- **Apply**:  
    - `Apply`, `ApplyI` (with index)
- **Aggregation**:  
    - `Sum`, `Average`, `SumBy`, `SumByP`, `SumByI`
- **Min/Max**:  
    - `Min`, `Max`, `MinBy`, `MaxBy`, `MinIndex`, `MaxIndex`
- **Element Search**:  
    - `First`, `Last`, `IndexOf`, `LastIndexOf`, `IndexWhere`, `IndexWhereP`, `LastIndexWhere`, `LastIndexWhereP`, `FindIf`
- **Remove/Count**:  
    - `RemoveIf`, `Remove`, `Count`, `CountIf`
- **Contains**:  
    - `Contains`
- **Concat/Copy**:  
    - `Concat`, `ShallowCopy`, `Repeat`
- **Push/Pop**:  
    - `Push`, `Pop`, `PushHead`, `PopHead`
- **Sorting**:  
    - `Sort`, `SortDescending`, `SortBy`, `SortByDescending`, `SortByIndex`, `SortByIndexDescending`, `SortByComparer`, `OrderBy`, `OrderByDescending`
- **Grouping**:  
    - `GroupBy`, `GroupByP`
- **Other**:  
    - `ToMapI`, `ToMapIE`, `Paginate`, `Reverse`

## Usage Example

```go
package main

import (
    "fmt"
    "github.com/szmcdull/glinq/garray"
)

func main() {
    nums := []int{5, 2, 8, 3, 1}
    garray.Sort(nums)
    fmt.Println(nums) // [1 2 3 5 8]

    odds := garray.Filter(nums, func(x int) bool { return x%2 == 1 })
    fmt.Println(odds) // [1 3 5]

    squares := garray.Map(nums, func(x int) int { return x * x })
    fmt.Println(squares) // [1 4 9 25 64]
}
```


## Documentation

- [API Reference (pkg.go.dev)](https://pkg.go.dev/github.com/szmcdull/glinq/garray)
- [garray_test.go](https://github.com/szmcdull/glinq/blob/main/garray/garray_test.go)

## Notes
- All operations are performed in-place unless otherwise specified.
- Many functions have index or error variants; see API docs for details.
- For streaming/iterator pipelines, see the `giter` package.
