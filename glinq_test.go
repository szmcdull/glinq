package glinq

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
	m, err2 := ToMap(sl, func(i int) float64 { return float64(i) }, func(i int) string { return fmt.Sprintf(`%d`, i) })
	if err2 != nil || len(m) != 6 {
		t.Errorf(`err=%v or len=%d expected 6`, err2, len(m))
	}
}

func TestWhere(t *testing.T) {
	sl := FromSlice([]int{1, 2, 3, 4, 5, 6})
	q := Where(sl, func(x int) bool { return x%2 == 0 }) // 2 4 6
	q = Where(q, func(x int) bool { return x > 2 })      // 4 6
	sl2, err := ToSlice(q)
	if err != nil || len(sl2) != 2 || sl2[0] != 4 || sl2[1] != 6 {
		t.Errorf(`err=%v result=%+v`, err, sl2)
	}
}

func TestSelect(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	q := Select(sl, func(x int) int { return x * 2 })
	q2 := Select(q, func(x int) string { return fmt.Sprintf(`%d`, x) })
	sl2, err := ToSlice(q2)
	if err != nil || len(sl2) != 6 {
		t.Errorf(`err=%v len=%d expected 6`, err, len(sl2))
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

func TestSelectWhere(t *testing.T) {
	sl := FromSlice([]int{0, 1, 2, 3, 4, 5})
	q2 := Select(sl, func(x int) string { return fmt.Sprintf(`%v`, x*2) }) // 0 2 4 6 8 10
	q := Where(q2, func(x string) bool { return len(x) == 1 })             // 0 2 4 6 8
	sl2, err := ToSlice(q)
	if err != nil || len(sl2) != 5 {
		t.Errorf(`err=%v len=%d expected 5`, err, len(sl2))
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

func TestRange(t *testing.T) {
	q := Range(0, 10)
	sl, err := ToSlice(q)
	if err != nil || q.Count() != 10 || len(sl) != 10 {
		t.Errorf(`err=%v count=%d len=%d expected 10`, err, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[9] != 9 {
		t.Errorf(`[0]=%v [9]=%v`, sl[0], sl[9])
	}

	q = RangeStep(0, 10, 2)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 5 || len(sl) != 5 {
		t.Errorf(`err=%v count=%d len=%d expected 5`, err, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[4] != 8 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}
}

func TestRangeFloat(t *testing.T) {
	q := Range(0.0, 10.0)
	sl, err := ToSlice(q)
	if err != nil || q.Count() != 10 || len(sl) != 10 {
		t.Errorf(`err=%v count=%d len=%d expected 10`, err, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[9] != 9 {
		t.Errorf(`[0]=%v [9]=%v`, sl[0], sl[9])
	}

	q = RangeStep(0.0, 10.0, 2.0)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 5 || len(sl) != 5 {
		t.Errorf(`err=%v count=%d len=%d expected 5`, err, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[4] != 8 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}

	q = RangeStep(0.0, 0.5, 0.1)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 5 || len(sl) != 5 {
		t.Errorf(`err=%v count=%d len=%d expected 5`, err, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[4] != 0.4 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}

	q = RangeStep(0.0, 0.19, 0.1)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 2 || len(sl) != 2 {
		t.Errorf(`err=%v count=%d len=%d expected 2`, err, q.Count(), len(sl))
		return
	}
	if sl[0] != 0 || sl[1] != 0.1 {
		t.Errorf(`[0]=%v [4]=%v`, sl[0], sl[4])
	}

	q = RangeStep(0.0, 0.11, 0.1)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 2 || len(sl) != 2 {
		t.Errorf(`err=%v count=%d len=%d expected 2`, err, q.Count(), len(sl))
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
	sl, err := ToSlice(q)
	if err != nil || q.Count() != 7 || len(sl) != 7 {
		t.Errorf(`err=%v count=%d len=%d expected 7`, err, q.Count(), len(sl))
	}
	if sl[0] != 3 || sl[6] != 9 {
		t.Errorf(`[0]=%v expected 3 [6]=%v expected 9`, sl[0], sl[6])
	}

	q = Where(Range(0, 10), func(int) bool { return true }) // test non-seekable iter
	q = Skip(q, 3)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 7 || len(sl) != 7 {
		t.Errorf(`err=%v count=%d len=%d expected 7`, err, q.Count(), len(sl))
	}
	if sl[0] != 3 || sl[6] != 9 {
		t.Errorf(`[0]=%v expected 3 [6]=%v expected 9`, sl[0], sl[6])
	}
}

func TestTake(t *testing.T) {
	q := Range(0, 10)
	q = Take(q, 3)
	sl, err := ToSlice(q)
	if err != nil || q.Count() != 3 || len(sl) != 3 {
		t.Errorf(`err=%v count=%d len=%d expected 3`, err, q.Count(), len(sl))
	}
	if sl[0] != 0 || sl[2] != 2 {
		t.Errorf(`[0]=%v expected 0 [2]=%v expected 2`, sl[0], sl[2])
	}

	q = Where(Range(0, 10), func(int) bool { return true }) // test non-seekable iter
	q = Take(q, 3)
	sl, err = ToSlice(q)
	if err != nil || q.Count() != 3 || len(sl) != 3 {
		t.Errorf(`err=%v count=%d len=%d expected 3`, err, q.Count(), len(sl))
	}
	if sl[0] != 0 || sl[2] != 2 {
		t.Errorf(`[0]=%v expected 0 [2]=%v expected 2`, sl[0], sl[2])
	}
}

func TestMinMax(t *testing.T) {
	q := Range(1, 10)
	min, err := Min(q)
	max, err2 := Max(q)
	if err != nil || err2 != nil || min != 1 || max != 9 {
		t.Errorf(`err=%v err2=%v min=%d expected 1 max=%d expected 9`, err, err2, min, max)
	}

	q = FromSlice([]int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9})
	min, err = Min(q)
	max, err2 = Max(q)
	if err != nil || err2 != nil || min != 1 || max != 10 {
		t.Errorf(`err=%v err2=%v min=%d expected 1 max=%d expected 9`, err, err2, min, max)
	}
}

func TestMinMaxBy(t *testing.T) {
	q := FromSlice([]string{`the`, `silver`, `fox`, `jump`, `over`, `the`, `lazy`, `dog`})
	min, err := MinBy(q, func(s string) byte { return s[0] })
	max, err2 := MaxBy(q, func(s string) byte { return s[0] })
	if err != nil || err2 != nil || min != `dog` || max != `the` {
		t.Errorf(`err=%v err2=%v min=%s expected "dog" max=%s expected "the"`, err, err2, min, max)
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
	avg, err := Average(q)
	sum, err2 := Sum(q)
	if err != nil || avg != 5.5 {
		t.Errorf(`err=%v avg=%f expected 5.5`, err, avg)
	}
	if err2 != nil || sum != 55 {
		t.Errorf(`err=%v avg=%f expected 55`, err2, sum)
	}
}

func TestContains(t *testing.T) {
	q := Range(1, 10)
	if ok, err := Contains(q, -1); err != nil || ok {
		t.Fail()
	}
	if ok, err := Contains(q, 1); err != nil || !ok {
		t.Fail()
	}
}
