package partial

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestBind(t *testing.T) {
	AddTo5 := Bind2_1r(Add, 5)
	if AddTo5(10) != 15 {
		t.Fail()
	}
}
