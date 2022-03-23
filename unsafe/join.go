package unsafe

import "sync/atomic"

type JoinEnumerableIterator[Outer, Inner, Result any] struct {
	iterator[Result]
	outer           IEnumerable[Outer]
	inner           IEnumerable[Inner]
	selector        func(Outer, Inner) (selected bool, result Result)
	outerEnumerator IEnumerator[Outer]
	innerEnumerator IEnumerator[Inner]
}

// func (me *WhereEnumerator[T]) GetEnumerator() IEnumerator[T] {
// 	return me
// }

func NewJoinIterator[Outer, Inner, Result any](
	outer IEnumerable[Outer],
	inner IEnumerable[Inner],
	pred func(Outer, Inner) (selected bool, result Result)) *JoinEnumerableIterator[Outer, Inner, Result] {

	result := &JoinEnumerableIterator[Outer, Inner, Result]{
		outer:    outer,
		inner:    inner,
		selector: pred,
	}
	result.child = result
	return result
}

func Join[Outer, Inner, Result any](
	outer IEnumerable[Outer],
	inner IEnumerable[Inner],
	pred func(Outer, Inner) (selected bool, result Result)) IEnumerable[Result] {

	return NewJoinIterator(outer, inner, pred)
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) Clone() IEnumerator[Result] {
	result := NewJoinIterator(me.outer, me.inner, me.selector)
	result.state = 1
	return result
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) Any() bool {
	return me.Clone().MoveNext()
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) runState() (result bool, jump bool) {
	switch me.state {
	case 1:
		me.outerEnumerator = me.outer.GetEnumerator()
		me.innerEnumerator = me.inner.GetEnumerator()
		me.state = 2
		fallthrough
	case 2:
		ok := me.outerEnumerator.MoveNext()
		if !ok {
			return false, false
		}
		me.state = 4
		fallthrough
	case 4:
		for {
			ok := me.innerEnumerator.MoveNext()
			if !ok { // inner iteration finished, move on to the next outter
				me.innerEnumerator = me.inner.GetEnumerator()
				me.state = 2
				return false, true // "jump" to case 2
			}
			o := me.outerEnumerator.Current()
			i := me.innerEnumerator.Current()
			if ok, r := me.selector(o, i); ok {
				me.current = r
				return true, false
			}
		}
	}
	panic(ErrInvalidState)
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) MoveNext() bool {
	for {
		result, jump := me.runState()
		if !jump {
			return result
		}
	}
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) Count() int {
	if atomic.CompareAndSwapInt32(&me.state, 0, 1) {
		count := 0
		for ok := me.MoveNext(); ok; ok = me.MoveNext() {
			count++
		}
		return count
	} else {
		result := NewJoinIterator(me.outer, me.inner, me.selector)
		return result.Count()
	}
}
