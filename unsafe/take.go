package unsafe

import "io"

type (
	TakeIterator[T any] struct {
		iterator[T]
		source     IEnumerable[T]
		take, pos  int
		enumerator IEnumerator[T]
	}
)

func NewTakeIterator[T any](source IEnumerable[T], n int) *TakeIterator[T] {
	result := &TakeIterator[T]{
		source: source,
		take:   n,
	}
	result.child = result
	return result
}

func Take[T any](source IEnumerable[T], n int) IEnumerable[T] {
	return NewTakeIterator(source, n)
}

func (me *TakeIterator[T]) Clone() IEnumerator[T] {
	result := NewTakeIterator(me.source, me.take)
	result.state = 1
	return result
}

func (me *TakeIterator[T]) MoveNext() bool {
	if me.pos >= me.take {
		return false
	}
	me.pos++

	switch me.state {
	case 1:
		enumerator := me.source.GetEnumerator()
		me.enumerator = enumerator
		me.state = 2
		fallthrough
	case 2:
		ok := me.enumerator.MoveNext()
		me.current = me.enumerator.Current()
		return ok
	}
	panic(ErrInvalidState)
}

func (me *TakeIterator[T]) Any() bool {
	if me.take == 0 {
		return false
	}
	return me.source.Any()
}

func (me *TakeIterator[T]) Count() int {
	iter := me.source.GetEnumerator()
	if counter, ok := iter.(ICount[T]); ok {
		result := counter.Count()
		if result > me.take {
			result = me.take
		}
		return result
	}

	result := 0
	Foreach(me.source, func(T) error {
		result++
		if result < me.take {
			return nil
		}
		return io.EOF
	})
	return result
}
