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
	m.Range(func(k, v int) bool {
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
