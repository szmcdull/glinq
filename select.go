package glinq

type (
	SelectEnumerableIterator[T any, T2 any] struct {
		iterator[T2]
		source     IEnumerable[T]
		selector   func(T) T2
		enumerator IEnumerator[T]
	}
)

func NewSelectIterator[T any, T2 any](source IEnumerable[T], selector func(T) T2) *SelectEnumerableIterator[T, T2] {
	result := &SelectEnumerableIterator[T, T2]{
		source:   source,
		selector: selector,
	}
	result.child = result
	return result
}

func Select[T any, T2 any](source IEnumerable[T], selector func(T) T2) IEnumerable[T2] {
	return NewSelectIterator(source, selector)
}

func (me *SelectEnumerableIterator[T, T2]) MoveNext() error {
	switch me.state {
	case 1:
		me.enumerator = me.source.GetEnumerator()
		me.state = 2
		fallthrough
	case 2:
		err := me.enumerator.MoveNext()
		if err != nil {
			return err
		}
		item := me.enumerator.Current()
		me.current = me.selector(item)
		return nil
	}
	return ErrInvalidState
}

func (me *SelectEnumerableIterator[T, T2]) Clone() IEnumerator[T2] {
	result := NewSelectIterator(me.source, me.selector)
	result.state = 1
	return result
}

func (me *SelectEnumerableIterator[T, T2]) Any() bool {
	return me.source.Any()
}

func (me *SelectEnumerableIterator[T, T2]) Count() int {
	return me.source.Count()
}
