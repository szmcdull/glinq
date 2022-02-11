package glinq

import "sync/atomic"

type (
	iterator[T any] struct {
		child iteratorChild[T]
		// -1 = disposed, not usable anymore
		//  0 = unused
		//  1 = at first item
		//  2 = at other item
		//  3 = in use by Count()
		state   int32
		current T
	}
	IEnumerableAndEnumerator[T any] interface {
		IEnumerable[T]
		IEnumerator[T]
	}
	iteratorChild[T any] interface {
		IEnumerator[T]
		Clone() IEnumerator[T]
	}
)

func (me *iterator[T]) GetEnumerator() IEnumerator[T] {
	var result IEnumerator[T]
	if atomic.CompareAndSwapInt32(&me.state, 0, 1) {
		result = me.child
	} else {
		result = me.child.Clone()
	}
	return result
}

func (me *iterator[T]) Current() T {
	return me.current
}

func (me *iterator[T]) MoveNext() error {
	panic(`not implemented`)
}

// func (me *iterator[T]) Reset() error {
// 	panic(`not implemented`)
// }
