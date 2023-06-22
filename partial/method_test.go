package partial

import (
	"fmt"
	"testing"
)

type TT struct {
	A int
}

func (me *TT) Add(v any) bool {
	fmt.Println(v)
	return true
}

func TestHandler(t *testing.T) {
	t1 := TT{A: 1}
	t2 := TT{A: 2}

	method1 := NewMethod(t1.Add, &t1)
	method2 := NewMethod(t2.Add, &t2)
	method3 := NewMethod(t2.Add, &t1)
	if method1.Equals(method2) {
		t.Error(`method1.Equals(method2.Func)`)
	}
	if !method1.Equals(method3) {
		t.Error(`!method1.Equals(method3.Func)`)
	}
	method1.Get()(1)
	method2.Get()(2)
}
