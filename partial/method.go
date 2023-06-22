package partial

import "reflect"

type (
	Func[F any] struct {
		f    F
		addr uintptr
	}

	Method[F any, O any] struct {
		Func[F]
		receiver uintptr
	}
)

func NewFunc[F any](f F) *Func[F] {
	result := &Func[F]{f: f}
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		panic("not a func")
	}
	result.addr = v.Pointer()
	return result
}

func (me *Func[F]) Get() F {
	return me.f
}

// Equals compares the function addresses of the me and other
func (me *Func[F]) Equals(other any) bool {
	o, ok := other.(Func[F])
	if !ok {
		return false
	}
	return me.addr == o.addr
}

// NewMethod creates a new Method.
// f must be a method of receiver.
// It will not panic if it's not. But the receiver is used to check if 2 Methods are the same
func NewMethod[F any, O any](f F, receiver O) *Method[F, O] {
	return &Method[F, O]{
		Func: Func[F]{
			f:    f,
			addr: reflect.ValueOf(f).Pointer(),
		},
		receiver: reflect.ValueOf(receiver).Pointer(),
	}
}

// NewMethodStrict creates a new Method.
// This is a type checked version of NewMethod. But its use case is limited to 1 and only 1 arg+result.
func NewMethodStrict[A any, R any, O any](f func(O, A) R, receiver O) *Method[func(A) R, O] {
	return &Method[func(A) R, O]{
		Func: Func[func(A) R]{
			f:    func(a A) R { return f(receiver, a) },
			addr: reflect.ValueOf(f).Pointer(),
		},
		receiver: reflect.ValueOf(receiver).Pointer(),
	}
}

func (me *Method[F, O]) Equals(other any) bool {
	o, ok := other.(*Method[F, O])
	if !ok {
		return false
	}
	return me.addr == o.addr && me.receiver == o.receiver
}
