package unsafe

import (
	"fmt"
	"testing"
)

func TestFromSlice(t *testing.T) {
	sl := FromSlice([]int{1, 2, 3, 4, 5, 6})
	count := sl.Count()
	if count != 6 {
		t.Errorf(`len=%d expected 6`, count)
	}
}

func TestToMap(t *testing.T) {
	sl := FromSlice([]int{1, 2, 3, 4, 5, 6})
	m := ToMap(sl, func(i int) float64 { return float64(i) }, func(i int) string { return fmt.Sprintf(`%d`, i) })
	if len(m) != 6 {
		t.Errorf(`len=%d expected 6`, len(m))
	}
}

func TestWhere(t *testing.T) {
	sl := FromSlice([]int{1, 2, 3, 4, 5, 6})
	q := Where(sl, func(x int) bool { return x%2 == 0 }) // 2 4 6
	q = Where(q, func(x int) bool { return x > 2 })      // 4 6
	if !Any(q) {
		t.Fail()
	}
	sl2 := ToSlice(q)
	if len(sl2) != 2 || sl2[0] != 4 || sl2[1] != 6 {
		t.Errorf(`len=%d exected 2, [0]=%d expected 4, [1]=%d expected 6`, len(sl2), sl2[0], sl2[1])
	}
}

func TestSelect(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	q := Select(sl, func(x int) int { return x * 2 })
	q2 := Select(q, func(x int) string { return fmt.Sprintf(`%d`, x) })
	sl2 := ToSlice(q2)
	if len(sl2) != 6 {
		t.Errorf(`len=%d expected 6`, len(sl2))
		return
	}
	for i, v := range sl2 {
		if fmt.Sprintf(`%d`, i*2) != v {
			t.Errorf(`%d != %s`, i*2, v)
		}
	}
}

func TestWhereSelect(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	q := Where(sl, func(x int) bool { return x%2 == 0 })
	q2 := Select(q, func(x int) string { return fmt.Sprintf(`%v`, x) })
	sl2 := ToSlice(q2)
	if len(sl2) != 3 {
		t.Errorf(`len=%d expected 3`, len(sl2))
		return
	}
	for i, v := range sl2 {
		if fmt.Sprintf(`%d`, i*2) != v {
			t.Errorf(`%d != %s`, i*2, v)
		}
	}
}

func TestSelectWhere(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	q2 := Select(sl, func(x int) string { return fmt.Sprintf(`%v`, x*2) }) // 0 2 4 6 8 10
	q := Where(q2, func(x string) bool { return len(x) == 1 })             // 0 2 4 6 8
	sl2 := ToSlice(q)
	if len(sl2) != 5 {
		t.Errorf(`len=%d expected 5`, len(sl2))
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
	ForeachI(sl, func(i, x int) error {
		if i != x {
			t.Errorf(`%d expected %d`, x, i)
		}
		return nil
	})
}

func TestRange(t *testing.T) {
	q := Range(0, 10)
	sl := ToSlice(q)
	if q.Count() != 10 || len(sl) != 10 {
		t.Errorf(`count=%d len=%d expected 10`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[9] != 9 {
		t.Errorf(`[0]=%v [9]=%v`, sl[0], sl[9])
	}

	q = RangeStep(0, 10, 2)
	sl = ToSlice(q)
	if q.Count() != 5 || len(sl) != 5 {
		t.Errorf(`count=%d len=%d expected 5`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[4] != 8 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}
}

func TestRangeFloat(t *testing.T) {
	q := Range(0.0, 10.0)
	sl := ToSlice(q)
	if q.Count() != 10 || len(sl) != 10 {
		t.Errorf(`count=%d len=%d expected 10`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[9] != 9 {
		t.Errorf(`[0]=%v [9]=%v`, sl[0], sl[9])
	}

	q = RangeStep(0.0, 10.0, 2.0)
	sl = ToSlice(q)
	if q.Count() != 5 || len(sl) != 5 {
		t.Errorf(`count=%d len=%d expected 5`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[4] != 8 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}

	q = RangeStep(0.0, 0.5, 0.1)
	sl = ToSlice(q)
	if q.Count() != 5 || len(sl) != 5 {
		t.Errorf(`count=%d len=%d expected 5`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[4] != 0.4 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}

	q = RangeStep(0.0, 0.19, 0.1)
	sl = ToSlice(q)
	if q.Count() != 2 || len(sl) != 2 {
		t.Errorf(`count=%d len=%d expected 2`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[1] != 0.1 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}

	q = RangeStep(0.0, 0.11, 0.1)
	sl = ToSlice(q)
	if q.Count() != 2 || len(sl) != 2 {
		t.Errorf(`count=%d len=%d expected 2`, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[1] != 0.1 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}
}

func TestFirst(t *testing.T) {
	q := Range(0, 10)
	first := First(q, nil)
	if first != 0 {
		t.Errorf(`First=%d expected 0`, first)
	}
	first = FirstWhere(q, func(x int) bool {
		return x > 5
	}, nil)
	if first != 6 {
		t.Errorf(`First=%d expected 6`, first)
	}
}

func TestSkip(t *testing.T) {
	q := Range(0, 10)
	q = Skip(q, 3)
	sl := ToSlice(q)
	if q.Count() != 7 || len(sl) != 7 {
		t.Errorf(`count=%d len=%d expected 7`, q.Count(), len(sl))
	}
	if sl[0] != 3 || sl[6] != 9 {
		t.Errorf(`[0]=%v expected 3 [6]=%v expected 9`, sl[0], sl[6])
	}

	q = Where(Range(0, 10), func(int) bool { return true }) // test non-seekable iter
	q = Skip(q, 3)
	sl = ToSlice(q)
	if q.Count() != 7 || len(sl) != 7 {
		t.Errorf(`count=%d len=%d expected 7`, q.Count(), len(sl))
	}
	if sl[0] != 3 || sl[6] != 9 {
		t.Errorf(`[0]=%v expected 3 [6]=%v expected 9`, sl[0], sl[6])
	}
}

func TestSkip2(t *testing.T) {
	q := []int{0, 1}
	s := FromSlice(q)
	result := ToSlice(Skip(s, 0))
	fmt.Println(result)
	if len(result) != 2 {
		t.Fail()
	}
}

func TestSkip3(t *testing.T) {
	q1 := []int{1}
	s1 := Take(FromSlice(q1), 3)
	if len(ToSlice(s1)) != 1 {
		t.Fail()
	}
	var q2 []int
	s2 := Take(FromSlice(q2), 3)
	if len(ToSlice(s2)) != 0 {
		t.Fail()
	}
}

func TestTake(t *testing.T) {
	q := Range(0, 10)
	q = Take(q, 3)
	sl := ToSlice(q)
	if q.Count() != 3 || len(sl) != 3 {
		t.Errorf(`count=%d len=%d expected 3`, q.Count(), len(sl))
	}
	if sl[0] != 0 || sl[2] != 2 {
		t.Errorf(`[0]=%v expected 0 [2]=%v expected 2`, sl[0], sl[2])
	}

	q = Where(Range(0, 10), func(int) bool { return true }) // test non-seekable iter
	q = Take(q, 3)
	sl = ToSlice(q)
	if q.Count() != 3 || len(sl) != 3 {
		t.Errorf(`count=%d len=%d expected 3`, q.Count(), len(sl))
	}
	if sl[0] != 0 || sl[2] != 2 {
		t.Errorf(`[0]=%v expected 0 [2]=%v expected 2`, sl[0], sl[2])
	}
}

func TestMinMax(t *testing.T) {
	q := Range(1, 10)
	min := Min(q)
	max := Max(q)
	if min != 1 || max != 9 {
		t.Errorf(`min=%d expected 1 max=%d expected 9`, min, max)
	}

	q = FromSlice([]int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9})
	min = Min(q)
	max = Max(q)
	if min != 1 || max != 10 {
		t.Errorf(`min=%d expected 1 max=%d expected 9`, min, max)
	}
}

func TestMinMaxBy(t *testing.T) {
	q := FromSlice([]string{`the`, `silver`, `fox`, `jump`, `over`, `the`, `lazy`, `dog`})
	min := MinBy(q, func(s string) byte { return s[0] })
	max := MaxBy(q, func(s string) byte { return s[0] })
	if min != `dog` || max != `the` {
		t.Errorf(`min=%s expected "dog" max=%s expected "the"`, min, max)
	}
}

func TestCount(t *testing.T) {
	q := Range(0, 10)
	if q.Count() != 10 {
		t.Errorf(`count=%d expected 10`, q.Count())
	}
	q = Where(q, func(x int) bool { return x > 5 })
	if q.Count() != 4 {
		t.Errorf(`count=%d expected 4`, q.Count())
	}

	q2 := FromMapReflect(map[string]int{`a`: 1, `b`: 2})
	if q2.Count() != 2 {
		t.Errorf(`count=%d expected 2`, q2.Count())
	}
}

func TestAverageSum(t *testing.T) {
	q := Range(1.0, 11.0)
	avg := Average(q)
	sum := Sum(q)
	if avg != 5.5 {
		t.Errorf(`avg=%f expected 5.5`, avg)
	}
	if sum != 55 {
		t.Errorf(`avg=%f expected 55`, sum)
	}
}

func TestContains(t *testing.T) {
	q := Range(1, 10)
	if ok := Contains(q, -1); ok {
		t.Fail()
	}
	if ok := Contains(q, 1); !ok {
		t.Fail()
	}
}

func TestJoin(t *testing.T) {
	type country struct {
		id   int
		name string
	}
	type city struct {
		id        int
		name      string
		countryId int
	}
	type all struct {
		country
		city
	}

	outer := []country{{1, `USA`}, {2, `GB`}, {3, `China`}, {4, `Japan`}}
	inner := []city{{1, `Shenzhen`, 3}, {2, `London`, 2}, {3, `New York`, 1}, {4, `Guangzhou`, 3}, {5, `Bangkok`, 0}}
	result := ToSlice(
		Join(FromSlice(outer),
			FromSlice(inner),
			func(o country, i city) (ok bool, result all) {
				if o.id == i.countryId {
					return true, all{o, i}
				}
				return false, all{}
			}))
	s := fmt.Sprintf(`%+v`, result)
	if s != `[{country:{id:1 name:USA} city:{id:3 name:New York countryId:1}} {country:{id:2 name:GB} city:{id:2 name:London countryId:2}} {country:{id:3 name:China} city:{id:1 name:Shenzhen countryId:3}} {country:{id:3 name:China} city:{id:4 name:Guangzhou countryId:3}}]` {
		t.Fail()
	}
}
