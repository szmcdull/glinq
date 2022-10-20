package glinq

import (
	"io"
	"sync/atomic"
)

type (
	TakeWhileIterator[T any] struct {
		iterator[T]
		source     IEnumerable[T]
		pred       func(T) bool
		ended      bool
		enumerator IEnumerator[T]
	}
)

func NewTakeWhileIterator[T any](source IEnumerable[T], pred func(T) bool) *TakeWhileIterator[T] {
	result := &TakeWhileIterator[T]{
		source: source,
		pred:   pred,
	}
	result.child = result
	return result
}

func TakeWhile[T any](source IEnumerable[T], pred func(T) bool) IEnumerable[T] {
	return NewTakeWhileIterator(source, pred)
}

func (me *TakeWhileIterator[T]) Clone() IEnumerator[T] {
	result := NewTakeWhileIterator(me.source, me.pred)
	result.state = 1
	return result
}

func (me *TakeWhileIterator[T]) MoveNext() error {
	switch me.state {
	case 1:
		enumerator := me.source.GetEnumerator()
		me.enumerator = enumerator
		me.state = 2
		fallthrough
	case 2:
		err := me.enumerator.MoveNext()
		me.current = me.enumerator.Current()
		if !me.pred(me.current) {
			me.ended = true
			me.state = 100
			return io.EOF
		}
		return err
	case 100:
		return io.EOF
	}
	return ErrInvalidState
}

func (me *TakeWhileIterator[T]) Any() bool {
	return me.Clone().MoveNext() == nil
}

func (me *TakeWhileIterator[T]) Count() int {
	if atomic.CompareAndSwapInt32(&me.state, 0, 1) {
		enum := me.source.GetEnumerator()
		count := 0
		for err := enum.MoveNext(); err == nil; err = enum.MoveNext() {
			if me.pred(enum.Current()) {
				count++
			} else {
				break
			}
		}
		return count
	} else {
		result := NewTakeWhileIterator(me.source, me.pred)
		return result.Count()
	}
}
