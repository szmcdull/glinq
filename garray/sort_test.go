package garray

import (
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
