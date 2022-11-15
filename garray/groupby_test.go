package garray

import (
	"fmt"
	"strings"
)

func ExampleGroupBy() {
	a := []string{`Jack`, `Mike`, `Jones`, `Janes`, `Tom`, `Terry`}
	m := GroupBy(a, func(s string) string { return string(s[0:1]) }, func(s string) string { return strings.ToLower(s) })
	fmt.Print(m)
	// output: map[J:[jack jones janes] M:[mike] T:[tom terry]]
}

func ExampleGroupByP() {
	a := []string{`Jack`, `Mike`, `Jones`, `Janes`, `Tom`, `Terry`}
	m := GroupByP(a, func(s *string) string { return string((*s)[0:1]) }, func(s *string) string { return strings.ToLower(*s) })
	fmt.Print(m)
	// output: map[J:[jack jones janes] M:[mike] T:[tom terry]]
}
