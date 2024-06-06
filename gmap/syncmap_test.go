package gmap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Test(t *testing.T) {
	m := NewSyncMap[int, int]()
	m.Store(1, 4)
	m.Store(2, 5)
	m.Store(3, 6)

	count := 0
	m.RangeNonReentrant(func(k, v int) bool {
		if k < 1 || k > 3 {
			t.Errorf(`unexpected k %d (should be in range[1,3])`, k)
		}
		if v-k != 3 {
			t.Errorf(`v %d - k %d != 3`, v, k)
		}
		count++
		return true
	})
	if count != 3 {
		t.Errorf(`expected count 3, got %d`, count)
	}

	l := m.ToSlice()
	count = 0
	for _, p := range l {
		if p.Key < 1 || p.Key > 3 {
			t.Errorf(`ToSlice unexpected k %d (should be in range[1,3])`, p.Key)
		}
		if p.Value-p.Key != 3 {
			t.Errorf(`ToSlice v %d - k %d != 3`, p.Value, p.Key)
		}
		count++
	}
	if count != 3 {
		t.Errorf(`ToSlice expected count 3, got %d`, count)
	}

	len := m.Len()
	if len != 3 {
		t.Errorf(`Len expected 3 got %d`, len)
	}
}

func TestLoadAnd(t *testing.T) {
	m := NewSyncMap[int, int]()
	m.Store(1, 2)
	if v, ok := m.Load(1); v != 2 || !ok {
		t.Fail()
	}

	if v, ok := m.LoadOrStore(2, 4); v != 4 || ok {
		t.Fail()
	}
	if v, ok := m.Load(2); v != 4 || !ok {
		t.Fail()
	}

	if v, ok := m.LoadOrStore(2, 0); v != 4 || !ok {
		t.Fail()
	}
	if v, ok := m.Load(2); v != 4 || !ok {
		t.Fail()
	}

	if v, ok := m.LoadAndDelete(1); v != 2 || !ok {
		t.Fail()
	}
	if v, ok := m.LoadAndDelete(3); v != 0 || ok {
		t.Fail()
	}

	if v, ok := m.LoadOrNew(1, func() int { return 3 }); v != 3 || ok {
		t.Fail()
	}

	if v, ok, err := m.LoadOrNewE(11, func() (int, error) { return 3, nil }); v != 3 || ok || err != nil {
		t.Fail()
	}
	if v, ok := m.Load(11); v != 3 || !ok {
		t.Fail()
	}
}

func _Write(ch <-chan struct{}, m *SyncMap[int, int]) {
	count := 0
	for {
		count++
		select {
		case <-ch:
			fmt.Printf("%d Writes\n", count)
			return
		default:
			m.Store(rand.Intn(100), rand.Int())
		}
	}
}

func TestConcurrent(t *testing.T) {
	ch := make(chan struct{})
	m := NewSyncMap[int, int]()
	go _Write(ch, m)
	go _Write(ch, m)

	time.Sleep(time.Second * 3)
	close(ch)
}

// func TestRecursiveLock(t *testing.T) {
// 	done := false

// 	go func() {
// 		time.Sleep(time.Second)
// 		if !done {
// 			t.Fail()
// 		}
// 	}()

// 	m := NewSyncMap[int, int]()
// 	m.Store(1, 1)
// 	m.Store(2, 2)
// 	m.Range(func(k, v int) bool {
// 		if k == 1 {
// 			m.Delete(k)
// 		}
// 		return true
// 	})

// 	done = true
// }

func TestClear(t *testing.T) {
	m := NewSyncMap[int, int]()
	m.Store(1, 2)
	m.Clear()
	if m.Len() != 0 {
		t.Fail()
	}
	m.Store(1, 2)
}
