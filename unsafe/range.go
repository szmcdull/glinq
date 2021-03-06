package unsafe

import (
	"io"
	"math"
	"sync/atomic"

	"golang.org/x/exp/constraints"
)

type (
	RangeIterator[T constraints.Integer | constraints.Float] struct {
		iterator[T]
		start, end, step T
	}
)

func NewRangeIterator[T constraints.Integer | constraints.Float](start, end, step T) *RangeIterator[T] {
	result := &RangeIterator[T]{
		start: start,
		end:   end,
		step:  step,
	}
	result.child = result
	return result
}

func Range[T constraints.Integer | constraints.Float](start, end T) IEnumerable[T] {
	return NewRangeIterator(start, end, T(1))
}

func RangeStep[T constraints.Integer | constraints.Float](start, end, step T) IEnumerable[T] {
	return NewRangeIterator(start, end, step)
}

func (me *RangeIterator[T]) MoveNext() bool {
	switch me.state {
	case 1:
		me.current = me.start
		me.state = 2
		return true
	case 2:
		me.current += me.step
		if me.current >= me.end {
			me.state = -1
			return false
		}
		return true
	}
	panic(ErrInvalidState)
}

func (me *RangeIterator[T]) Count() int {
	return int(math.Ceil(float64(me.end-me.start) / float64(me.step)))
}

func (me *RangeIterator[T]) Any() bool {
	return me.end > me.start
}

func (me *RangeIterator[T]) Clone() IEnumerator[T] {
	result := NewRangeIterator(me.start, me.end, me.step)
	result.state = 1
	return result
}

func (me *RangeIterator[T]) GetAt(pos int) (result T, err error) {
	current := me.start + me.step*T(pos)
	if current < me.end {
		return current, nil
	}
	return result, io.ErrUnexpectedEOF
}

func (me *RangeIterator[T]) SeekOnce(pos int) error {
	if atomic.CompareAndSwapInt32(&me.state, 1, 2) { // can only seek once
		current := me.start + me.step*T(pos)
		if current < me.end {
			me.current = current
			return nil
		}
		return io.ErrUnexpectedEOF
	}
	return ErrInvalidState
}
