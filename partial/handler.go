package partial

import (
	"sync"

	"github.com/szmcdull/glinq/garray"
)

type (
	Handler[F any] struct {
		subscribers []FuncOrMethod[F]
		lock        sync.Mutex
	}

	FuncOrMethod[F any] interface {
		Get() F
		Equals(other any) bool
	}

	// _MethodKey struct {
	// 	addr     uintptr
	// 	receiver uintptr
	// }
)

func NewHandler[F any]() *Handler[F] {
	return &Handler[F]{
		//subscribers: make(map[_MethodKey]FuncOrMethod[F]),
	}
}

// add a function to the handler (even if it is already added)
func (me *Handler[F]) AddFunc(f F) {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.subscribers = append(me.subscribers, NewFunc(f))
}

// add a method to the handler (even if it is already added)
func (me *Handler[F]) AddMethod(f F, receiver any) {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.subscribers = append(me.subscribers, NewMethod(f, receiver))
}

// remove a function or method from the handler. if funcOrMethod was added multiple times, all of it will be removed
func (me *Handler[F]) Remove(funcOrMethod any) {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.subscribers = garray.FilterI(me.subscribers, func(i int) bool {
		return !me.subscribers[i].Equals(funcOrMethod)
	})
}

// get all functions and methods
func (me *Handler[F]) Get() []F {
	me.lock.Lock()
	defer me.lock.Unlock()
	result := make([]F, len(me.subscribers))
	for i, v := range me.subscribers {
		result[i] = v.Get()
	}
	return result
}

// todo
// func (me *Handler[F]) Call(args ...any) ... {}
