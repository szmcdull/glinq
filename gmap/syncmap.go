package gmap

import (
	"encoding/json"
	"sync"
)

type (
	SyncMap[K comparable, V any] struct {
		m map[K]V
		l sync.RWMutex
	}
)

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func NewSyncMapLength[K comparable, V any](len int) *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V, len),
	}
}

func (me *SyncMap[K, V]) Delete(key K) {
	me.l.Lock()
	defer me.l.Unlock()
	delete(me.m, key)
}

func (me *SyncMap[K, V]) DeleteIf(pred func(k K, v V) bool) {
	me.l.Lock()
	defer me.l.Unlock()
	for k, v := range me.m {
		if pred(k, v) {
			delete(me.m, k)
		}
	}
}

func (me *SyncMap[K, V]) UpdateIf(pred func(k K, v V) (V, bool)) {
	me.l.Lock()
	defer me.l.Unlock()
	for k, v := range me.m {
		if newV, updated := pred(k, v); updated {
			me.m[k] = newV
		}
	}
}

func (me *SyncMap[K, V]) Store(key K, value V) {
	me.l.Lock()
	defer me.l.Unlock()
	me.m[key] = value
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (me *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	me.l.Lock()
	defer me.l.Unlock()
	value, loaded = me.m[key]
	if loaded {
		delete(me.m, key)
	}
	return
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
func (me *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	me.l.Lock()
	defer me.l.Unlock()
	actual, loaded = me.m[key]
	if !loaded {
		me.m[key] = value
		actual = value
	}
	return
}

// atomic load-or-new
func (me *SyncMap[K, V]) LoadOrNew(key K, newFunc func() V) (actual V, loaded bool) {
	me.l.Lock()
	defer me.l.Unlock()
	actual, loaded = me.m[key]
	if !loaded {
		actual = newFunc()
		me.m[key] = actual
	}
	return
}

// atomic load-or-new with error check
func (me *SyncMap[K, V]) LoadOrNewE(key K, newFunc func() (V, error)) (actual V, loaded bool, err error) {
	me.l.Lock()
	defer me.l.Unlock()
	actual, loaded = me.m[key]
	if !loaded {
		actual, err = newFunc()
		if err == nil {
			me.m[key] = actual
		}
	}
	return
}

// atomic load-and-update
func (me *SyncMap[K, V]) LoadAndUpdate(key K, updateFunc func(old V) (new V, updated bool)) {
	me.l.Lock()
	defer me.l.Unlock()
	old := me.m[key]
	new, updated := updateFunc(old)
	if updated {
		me.m[key] = new
	}
}

func (me *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	me.l.RLock()
	defer me.l.RUnlock()
	value, ok = me.m[key]
	return
}

// deprecated
// f must not call any methods of me.
// If you want to delete items using RangeNonReentrant, use DeleteIf instead.
func (me *SyncMap[K, V]) Range(f func(K, V) bool) {
	me.RangeNonReentrant(f)
}

// f must not call any methods of me.
// If you want to delete items using Range, use DeleteIf instead.
func (me *SyncMap[K, V]) RangeNonReentrant(f func(K, V) bool) {
	me.l.RLock()
	defer me.l.RUnlock()
	for k, v := range me.m {
		if !f(k, v) {
			break
		}
	}
}

// f must not call any methods of me.
// If you want to delete items using RangeE, use DeleteIf instead.
func (me *SyncMap[K, V]) RangeE(f func(K, V) error) error {
	me.l.RLock()
	defer me.l.RUnlock()
	for k, v := range me.m {
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (me *SyncMap[K, V]) ToSlice() []Pair[K, V] {
	me.l.RLock()
	defer me.l.RUnlock()
	result := make([]Pair[K, V], 0, len(me.m))
	for k, v := range me.m {
		result = append(result, Pair[K, V]{k, v})
	}
	return result
}

func (me *SyncMap[K, V]) Len() int {
	me.l.RLock()
	defer me.l.RUnlock()
	return len(me.m)
}

func (me *SyncMap[K, V]) Clear() {
	me.l.Lock()
	me.m = map[K]V{}
	me.l.Unlock()
}

func (me *SyncMap[K, V]) PopAll() map[K]V {
	me.l.Lock()
	result := me.m
	me.m = map[K]V{}
	me.l.Unlock()
	return result
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (me *SyncMap[K, V]) MarshalJSON() ([]byte, error) {
	me.l.RLock()
	defer me.l.RUnlock()
	return json.Marshal(me.m)
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (me *SyncMap[K, V]) UnmarshalJSON(b []byte) error {
	me.l.Lock()
	defer me.l.Unlock()
	//if err := UnmarshalUseNumber(b, &me.m); err != nil {
	if err := json.Unmarshal(b, &me.m); err != nil {
		return err
	}
	return nil
}

// func UnmarshalUseNumber(data []byte, v interface{}) (err error) {
// 	decoder := json.NewDecoder(bytes.NewReader(data))
// 	decoder.UseNumber()
// 	err = decoder.Decode(v)
// 	return
// }

// Replace the map with a new one, and returns the old one
func (me *SyncMap[K, V]) Replace(m map[K]V) map[K]V {
	me.l.Lock()
	result := me.m
	me.m = m
	me.l.Unlock()
	return result
}
