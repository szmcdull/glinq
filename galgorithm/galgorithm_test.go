package galgorithm

import "testing"

func TestAbs(t *testing.T) {
	r := Abs(-1)
	r2 := Abs(1)
	if r != 1 || r2 != 1 {
		t.Fail()
	}
}

func TestClamp(t *testing.T) {
	r := Clamp(-2, -1, 1)
	r2 := Clamp(2, -1, 1)
	r3 := Clamp(0, -1, 1)
	if r != -1 || r2 != 1 || r3 != 0 {
		t.Fail()
	}
}

func TestMinMax(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	if Min(values...) != 1 || Max(values...) != 5 {
		t.Fail()
	}
}
