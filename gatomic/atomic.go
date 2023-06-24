package gatomic

import (
	"sync"
)

type Pointer[T any] struct {
	p *T
	l sync.Mutex
}

func (me *Pointer[T]) LoadOrNew(creator func() (*T, error)) (result *T, loaded bool, err error) {
	me.l.Lock()
	defer me.l.Unlock()

	if me.p == nil {
		result, err = creator()
		if err != nil {
			return nil, false, err
		}
		me.p = result
		return result, false, nil
	} else {
		return me.p, true, nil
	}
}

func (me *Pointer[T]) Load() *T {
	return me.p
}

func (me *Pointer[T]) Store(p *T) {
	me.p = p
}

func (me *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool) {
	me.l.Lock()
	defer me.l.Unlock()

	if me.p == old {
		me.p = new
		return true
	}
	return false
}

func (me *Pointer[T]) Swap(new *T) (old *T) {
	me.l.Lock()
	defer me.l.Unlock()

	old = me.p
	me.p = new
	return
}
