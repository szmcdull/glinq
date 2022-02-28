package glinq

import "sync/atomic"

type WhereEnumerableIterator[T any] struct {
	iterator[T]
	source     IEnumerable[T]
	pred       func(T) bool
	enumerator IEnumerator[T]
}

// func (me *WhereEnumerator[T]) GetEnumerator() IEnumerator[T] {
// 	return me
// }

func NewWhereIterator[T any](source IEnumerable[T], pred func(T) bool) *WhereEnumerableIterator[T] {
	result := &WhereEnumerableIterator[T]{
		source: source,
		pred:   pred,
	}
	result.child = result
	return result
}

func Where[T any](source IEnumerable[T], pred func(T) bool) IEnumerable[T] {
	return NewWhereIterator(source, pred)
}

func (me *WhereEnumerableIterator[T]) MoveNext() error {
	switch me.state {
	case 1:
		me.enumerator = me.source.GetEnumerator()
		me.state = 2
		fallthrough
	case 2:
		for {
			err := me.enumerator.MoveNext()
			if err != nil {
				return err
			}
			item := me.enumerator.Current()
			if me.pred(item) {
				me.current = item
				return nil
			}
		}
	}
	return ErrInvalidState
}

func (me *WhereEnumerableIterator[T]) Clone() IEnumerator[T] {
	result := NewWhereIterator(me.source, me.pred)
	result.state = 1
	return result
}

func (me *WhereEnumerableIterator[T]) Any() bool {
	return me.Clone().MoveNext() == nil
}

func (me *WhereEnumerableIterator[T]) Count() int {
	if atomic.CompareAndSwapInt32(&me.state, 0, 1) {
		count := 0
		for err := me.MoveNext(); err == nil; err = me.MoveNext() {
			count++
		}
		return count
	} else {
		result := NewWhereIterator(me.source, me.pred)
		return result.Count()
	}
}
