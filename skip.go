package glinq

type (
	SkipIterator[T any] struct {
		iterator[T]
		source     IEnumerable[T]
		skip       int
		enumerator IEnumerator[T]
	}
)

func NewSkipIterator[T any](source IEnumerable[T], n int) *SkipIterator[T] {
	result := &SkipIterator[T]{
		source: source,
		skip:   n,
	}
	result.child = result
	return result
}

func Skip[T any](source IEnumerable[T], n int) IEnumerable[T] {
	return NewSkipIterator(source, n)
}

func (me *SkipIterator[T]) Clone() IEnumerator[T] {
	result := NewSkipIterator(me.source, me.skip)
	result.state = 1
	return result
}

func (me *SkipIterator[T]) MoveNext() error {
	switch me.state {
	case 1:
		enumerator := me.source.GetEnumerator()
		me.enumerator = enumerator
		if randomAcceesor, ok := enumerator.(IRangeEnumerator[T]); ok {
			err := randomAcceesor.SeekOnce(me.skip)
			me.current = enumerator.Current()
			me.state = 2
			return err
		}
		for i := 0; i < me.skip; i++ {
			if err := enumerator.MoveNext(); err != nil {
				return err
			}
		}
		me.current = enumerator.Current()
		me.state = 2
		fallthrough
	case 2:
		err := me.enumerator.MoveNext()
		me.current = me.enumerator.Current()
		return err
	}
	return ErrInvalidState
}

func (me *SkipIterator[T]) Any() bool {
	return me.GetEnumerator().MoveNext() == nil
}

func (me *SkipIterator[T]) Count() int {
	iter := me.source.GetEnumerator()
	if counter, ok := iter.(ICount[T]); ok {
		result := counter.Count() - me.skip
		if result >= 0 {
			return result
		}
		return 0
	}

	result := 0
	Foreach(me.source, func(T) bool {
		result++
		return true
	})
	result -= me.skip
	if result < 0 {
		result = 0
	}
	return result
}
