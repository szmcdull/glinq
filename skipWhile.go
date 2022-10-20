package glinq

import "sync/atomic"

type (
	SkipWhileIterator[T any] struct {
		iterator[T]
		source     IEnumerable[T]
		pred       func(T) bool
		enumerator IEnumerator[T]
	}
)

func NewSkipWhileIterator[T any](source IEnumerable[T], pred func(T) bool) *SkipWhileIterator[T] {
	result := &SkipWhileIterator[T]{
		source: source,
		pred:   pred,
	}
	result.child = result
	return result
}

func SkipWhile[T any](source IEnumerable[T], pred func(T) bool) IEnumerable[T] {
	return NewSkipWhileIterator(source, pred)
}

func (me *SkipWhileIterator[T]) Clone() IEnumerator[T] {
	result := NewSkipWhileIterator(me.source, me.pred)
	result.state = 1
	return result
}

func (me *SkipWhileIterator[T]) MoveNext() error {
	switch me.state {
	case 1:
		enumerator := me.source.GetEnumerator()
		me.enumerator = enumerator
		for {
			if err := enumerator.MoveNext(); err != nil {
				return err
			}
			current := enumerator.Current()
			if !me.pred(current) {
				me.current = current
				break
			}
		}
		me.state = 2
		return nil
	case 2:
		err := me.enumerator.MoveNext()
		me.current = me.enumerator.Current()
		return err
	}
	return ErrInvalidState
}

func (me *SkipWhileIterator[T]) Any() bool {
	return me.GetEnumerator().MoveNext() == nil
}

func (me *SkipWhileIterator[T]) Count() int {
	if atomic.CompareAndSwapInt32(&me.state, 0, 1) {
		count := me.source.Count()
		if count < 0 {
			return count
		}
		enum := me.source.GetEnumerator()
		for err := enum.MoveNext(); err == nil; err = enum.MoveNext() {
			if me.pred(enum.Current()) {
				count--
			} else {
				break
			}
		}
		return count
	} else {
		result := NewSkipWhileIterator(me.source, me.pred)
		return result.Count()
	}
}
