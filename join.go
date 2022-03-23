package glinq

import (
	"io"
	"sync/atomic"
)

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
	return me.Clone().MoveNext() == nil
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) runState() (result error, jump bool) {
	switch me.state {
	case 1:
		me.outerEnumerator = me.outer.GetEnumerator()
		me.innerEnumerator = me.inner.GetEnumerator()
		me.state = 2
		fallthrough
	case 2:
		err := me.outerEnumerator.MoveNext()
		if err != nil {
			return err, false
		}
		me.state = 4
		fallthrough
	case 4:
		for {
			err := me.innerEnumerator.MoveNext()
			if err != nil { // inner iteration finished, move on to the next outter
				if err != io.EOF {
					return err, false
				}
				me.innerEnumerator = me.inner.GetEnumerator()
				me.state = 2
				return nil, true // "jump" to case 2
			}
			o := me.outerEnumerator.Current()
			i := me.innerEnumerator.Current()
			if ok, r := me.selector(o, i); ok {
				me.current = r
				return nil, false
			}
		}
	}
	panic(ErrInvalidState)
}

func (me *JoinEnumerableIterator[Outer, Inner, Result]) MoveNext() error {
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
		for err := me.MoveNext(); err == nil; err = me.MoveNext() {
			count++
		}
		return count
	} else {
		result := NewJoinIterator(me.outer, me.inner, me.selector)
		return result.Count()
	}
}
