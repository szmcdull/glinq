package garray

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	l := []int{8, 6, 4, 2, 5, 3, 1}

	r := []int{1, 2, 3, 4, 5, 6, 8}
	sort.Sort(Sortable(l))
	for i := range l {
		if l[i] != r[i] {
			t.Errorf(`#%d = %d expected %d`, i, l[i], r[i])
		}
	}

	r2 := []int{8, 6, 5, 4, 3, 2, 1}
	sort.Sort(SortableDescending(l))
	for i := range l {
		if l[i] != r2[i] {
			t.Errorf(`#%d = %d expected %d`, i, l[i], r2[i])
		}
	}
}

func TestOrderBy(t *testing.T) {
	l := []string{`the`, `lazy`, `dog`, `jumps`, `over`, `the`, `silver`, `fox`}

	r := []string{`dog`, `fox`, `jumps`, `lazy`, `over`, `silver`, `the`, `the`}
	sort.Sort(OrderBy(l, func(x string) byte { return x[0] }))
	for i := range l {
		if l[i] != r[i] {
			t.Errorf(`#%d = %s expected %s`, i, l[i], r[i])
		}
	}

	r2 := []string{`the`, `the`, `silver`, `over`, `lazy`, `jumps`, `fox`, `dog`}
	sort.Sort(OrderByDescending(l, func(x string) byte { return x[0] }))
	for i := range l {
		if l[i] != r2[i] {
			t.Errorf(`#%d = %s expected %s`, i, l[i], r2[i])
		}
	}
}

func ExampleSort() {
	l := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	Sort(l)
	fmt.Println(l)
	// Output: [1 1 2 3 3 4 5 5 5 6 9]
}

func ExampleSortDescending() {
	l := []string{"pineapple", "apple", "banana", "pear", "cherry", "orange"}
	SortDescending(l)
	fmt.Println(l)
	// Output: [pineapple pear orange cherry banana apple]
}

func ExampleOrderBy() {
	type Person struct {
		Name string
		Age  int
	}
	l := []Person{
		{"Alice", 32},
		{"Bob", 22},
		{"Charlie", 42},
		{"David", 27},
	}
	ordered := OrderBy(l, func(p Person) int { return p.Age })
	sort.Sort(ordered)
	for _, p := range ordered.slice {
		fmt.Println(p.Name, p.Age)
	}
	// Output:
	// Bob 22
	// David 27
	// Alice 32
	// Charlie 42
}

func ExampleOrderByDescending() {
	type Product struct {
		Name  string
		Price float64
	}
	l := []Product{
		{"Laptop", 999.99},
		{"Mouse", 19.99},
		{"Keyboard", 49.99},
		{"Headphones", 79.99},
	}
	ordered := OrderByDescending(l, func(p Product) float64 { return p.Price })
	sort.Sort(ordered)
	for _, p := range ordered.slice {
		fmt.Println(p.Name, p.Price)
	}
	// Output:
	// Laptop 999.99
	// Headphones 79.99
	// Keyboard 49.99
	// Mouse 19.99
}

func ExampleSortBy() {
	type Person struct {
		Name string
		Age  int
	}
	l := []Person{
		{"Alice", 32},
		{"Bob", 22},
		{"Charlie", 42},
		{"David", 27},
	}
	SortBy(l, func(p Person) int { return p.Age })
	for _, p := range l {
		fmt.Println(p.Name, p.Age)
	}
	// Output:
	// Bob 22
	// David 27
	// Alice 32
	// Charlie 42
}

func ExampleSortByDescending() {
	type Product struct {
		Name  string
		Price float64
	}
	l := []Product{
		{"Laptop", 999.99},
		{"Mouse", 19.99},
		{"Keyboard", 49.99},
		{"Headphones", 79.99},
	}
	SortByDescending(l, func(p Product) float64 { return p.Price })
	for _, p := range l {
		fmt.Println(p.Name, p.Price)
	}
	// Output:
	// Laptop 999.99
	// Headphones 79.99
	// Keyboard 49.99
	// Mouse 19.99
}
