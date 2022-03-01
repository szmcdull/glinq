package main

import (
	"fmt"
	"sort"

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

	sort.Sort(garray.Sortable(l))
	fmt.Printf("%v\n", l) // [1 2 3 4 5 6 8]

	l2 := []string{`the`, `lazy`, `dog`, `jumps`, `over`, `the`, `silver`, `fox`}
	sort.Sort(garray.OrderByDescending(l2, func(x string) byte { return x[len(x)-1] })) // sort descending by the last character
	fmt.Printf("%v\n", l2)                                                              // [lazy fox jumps silver over dog the the]
}
